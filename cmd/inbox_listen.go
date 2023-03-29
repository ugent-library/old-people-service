package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people/inbox"
	"github.com/ugent-library/people/models"
)

var inboxStreamName string = "PEOPLE"
var inboxStreamSubjects []string = []string{"person.update"}
var inboxConsumerName string = "inbox"
var inboxDeliverSubject = "inboxDeliverSubject"

var outboxStreamName string = "PEOPLE_SERVICE"
var outboxStreamSubjects []string = []string{
	"person.updated",
	"inbox.rejected",
	"inbox.person.rejected",
}

func initInboxStream(js nats.JetStreamContext) error {

	stream, _ := js.StreamInfo(inboxStreamName)

	if stream == nil {
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     inboxStreamName,
			Subjects: inboxStreamSubjects,
			Storage:  nats.FileStorage,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func initOutboxStream(js nats.JetStreamContext) error {

	stream, _ := js.StreamInfo(outboxStreamName)

	if stream == nil {
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     outboxStreamName,
			Subjects: outboxStreamSubjects,
			Storage:  nats.FileStorage,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func initInboxConsumer(js nats.JetStreamContext) error {

	if err := initInboxStream(js); err != nil {
		return err
	}

	consumerInfo, _ := js.ConsumerInfo(
		inboxStreamName, inboxConsumerName,
	)
	consumerConfig := &nats.ConsumerConfig{
		AckPolicy: nats.AckExplicitPolicy,
		Durable:   inboxConsumerName,
		//DeliverSubject makes it a push based consumer
		//this must be different per consumer
		//reason: messages are republished by consumer using this subject
		DeliverSubject: inboxDeliverSubject, // makes it is a push based consumer
		AckWait:        time.Second * 10,
		/*
			when more than 1, messages can be delivered out of order
			when they need to be redelivered
		*/
		MaxAckPending: 1,
	}

	if consumerInfo == nil {
		_, err := js.AddConsumer(inboxStreamName, consumerConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func ensureAck(msg *nats.Msg) {
	if err := msg.Ack(); err != nil {
		logger.Fatal(fmt.Errorf("unable to acknowledge nats message: %w", err))
	}
}

var inboxListenCmd = &cobra.Command{
	Use: "listen",
	Run: func(cmd *cobra.Command, args []string) {

		personService, err := models.NewPersonService(&models.PersonConfig{
			DB: config.Db.Url,
		})

		if err != nil {
			logger.Fatal(fmt.Errorf("unable to connect to database: %w", err))
		}

		nc, err := nats.Connect(config.Nats.Url)

		if err != nil {
			logger.Fatal(fmt.Errorf("unable to connect to nats: %w", err))
		}

		// "drain" introduces weird behaviour (messages only sent after restart)
		//defer nc.Drain()

		js, _ := nc.JetStream()

		if err := initOutboxStream(js); err != nil {
			logger.Fatal(fmt.Errorf("unable to create nats outbox stream: %w", err))
		}

		if err := initInboxConsumer(js); err != nil {
			logger.Fatal(fmt.Errorf("unable to create nats consumer: %w", err))
		}

		// subscribe to person.update
		_, subErr := js.Subscribe("person.update", func(msg *nats.Msg) {

			iMsg := &inbox.InboxMessage{}
			iMsg.Subject = msg.Subject
			iMsg.Message = &inbox.Message{}

			// remove malformed message
			// leave no trace
			if err := json.Unmarshal(msg.Data, iMsg.Message); err != nil {
				logger.Errorf("unable to decode nats message data as json: %w", err)
				ensureAck(msg)
				return
			}

			// remove malformed message
			// push validation errors to subject inbox.rejected
			if vErrs := iMsg.Validate(); vErrs != nil {
				iErrMsg := &inbox.InboxErrorMessage{
					InboxMessage: iMsg,
					Errors:       vErrs,
				}
				outboxBytes, _ := json.Marshal(iErrMsg)
				if err := nc.Publish("inbox.rejected", outboxBytes); err != nil {
					logger.Fatal(fmt.Errorf("unable to publish message to subject %s: %w", "inbox.rejected", err))
				}
				ensureAck(msg)
				return
			}

			person, err := personService.Get(context.Background(), iMsg.Message.ID)

			if err != nil && err == models.ErrNotFound {
				person = &models.Person{}
			} else if err != nil {
				// something is seriously wrong. We should stop processing records
				log.Fatal(fmt.Errorf("unable to fetch person record '%s': %w", iMsg.Message.ID, err))
			}

			oldPerson := person.Dup()

			iMsg.UpdatePersonAttr(person)

			// TODO: how to deactive people?
			person.Active = true

			// report invalid changes to subject inbox.person.rejected
			if vErrs := person.Validate(); vErrs != nil {
				pChangeErr := &inbox.PersonChangeError{
					OldPerson: oldPerson,
					NewPerson: person,
					Errors:    vErrs,
				}
				outboxBytes, _ := json.Marshal(pChangeErr)
				if err := nc.Publish("inbox.person.rejected", outboxBytes); err != nil {
					logger.Fatal(fmt.Errorf("unable to publish message to subject %s: %w", "inbox.person.rejected", err))
				}
				ensureAck(msg)
				return
			}

			var updateErr error
			if person.IsStored() {
				person, updateErr = personService.Update(context.Background(), person)
			} else {
				person, updateErr = personService.Create(context.Background(), person)
			}

			// create/update failed: stop processing records
			if updateErr != nil {
				logger.Fatal(fmt.Errorf("unable to store person %s: %w", person.ID, updateErr))
			}

			// republish updated record to subject person.update
			logger.Infof("updated person %s via subject person.update", person.ID)
			outboxBytes, _ := json.Marshal(person)

			// failed to contact nats: stop processing records
			if err := nc.Publish("person.updated", outboxBytes); err != nil {
				logger.Fatal(fmt.Errorf("unable to publish message to subject %s: %w", "person.updated", err))
			}

			logger.Infof("published person %s to subject person.updated", person.ID)

			// acknowledge msg or die
			ensureAck(msg)

		},
			// second and next subscription will fail when using Bind
			nats.Bind(inboxStreamName, inboxConsumerName),
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			// subscription specific option
			// without this message is acknowledged automatically (see "mack" in nats.SubOpt)
			nats.ManualAck(), // don't ask: why needed if AckExplicit is set?
		)

		if subErr != nil {
			logger.Fatal(fmt.Errorf("unable to subscribe to nats subject: %w", subErr))
		}

		runtime.Goexit() // wait for go routine to end (never)

	},
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}
