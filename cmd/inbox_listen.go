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
	"github.com/ugent-library/people/validation"
)

var inboxStreamName = "PEOPLE"
var inboxPersonStreamSubjects = []string{"person.update", "person.delete"}
var inboxPersonConsumerName = "inboxPerson"
var inboxPersonDeliverSubject = "inboxPersonDeliverSubject"

var inboxOrganizationStreamSubjects = []string{"organization.update", "organization.delete"}
var inboxOrganizationConsumerName = "inboxOrganization"
var inboxOrganizationDeliverSubject = "inboxOrganizationDeliverSubject"

var inboxAllStreamSubjects = []string{
	"person.update",
	"person.delete",
	"organization.update",
	"organization.delete",
}

/*
TODO: on delete of organizations push all related people to person.updated?
*/

func initInboxStream(js nats.JetStreamContext) error {

	stream, _ := js.StreamInfo(inboxStreamName)

	if stream == nil {
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     inboxStreamName,
			Subjects: inboxAllStreamSubjects,
			Storage:  nats.FileStorage,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func initPersonInboxConsumer(js nats.JetStreamContext) error {

	if err := initInboxStream(js); err != nil {
		return err
	}

	consumerInfo, _ := js.ConsumerInfo(
		inboxStreamName, inboxPersonConsumerName,
	)
	consumerConfig := &nats.ConsumerConfig{
		AckPolicy: nats.AckExplicitPolicy,
		Durable:   inboxPersonConsumerName,
		//DeliverSubject makes it a push based consumer
		//this must be different per consumer
		//reason: messages are republished by consumer using this subject
		DeliverSubject: inboxPersonDeliverSubject, // makes it is a push based consumer
		AckWait:        time.Second * 10,
		/*
			when more than 1, messages can be delivered out of order
			when they need to be redelivered
		*/
		MaxAckPending: 1,
		FilterSubject: "person.*",
	}

	if consumerInfo == nil {
		_, err := js.AddConsumer(inboxStreamName, consumerConfig)
		if err != nil {
			return err
		}
	}

	return nil
}

func initOrganizationInboxConsumer(js nats.JetStreamContext) error {

	if err := initInboxStream(js); err != nil {
		return err
	}

	consumerInfo, _ := js.ConsumerInfo(
		inboxStreamName, inboxOrganizationConsumerName,
	)
	consumerConfig := &nats.ConsumerConfig{
		AckPolicy: nats.AckExplicitPolicy,
		Durable:   inboxOrganizationConsumerName,
		//DeliverSubject makes it a push based consumer
		//this must be different per consumer
		//reason: messages are republished by consumer using this subject
		DeliverSubject: inboxOrganizationDeliverSubject, // makes it is a push based consumer
		AckWait:        time.Second * 10,
		/*
			when more than 1, messages can be delivered out of order
			when they need to be redelivered
		*/
		MaxAckPending: 1,
		FilterSubject: "organization.*",
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

		services := Services()
		personService := services.PersonService
		//personSearchService := services.PersonSearchService
		organizationService := services.OrganizationService
		//organizationSearchService := services.OrganizationSearchService

		nc, err := nats.Connect(config.Nats.Url)

		if err != nil {
			logger.Fatal(fmt.Errorf("unable to connect to nats: %w", err))
		}

		// "drain" introduces weird behaviour (messages only sent after restart)
		//defer nc.Drain()

		js, _ := nc.JetStream()

		if err := initPersonInboxConsumer(js); err != nil {
			logger.Fatal(fmt.Errorf("unable to create nats consumer for person: %w", err))
		}

		// subscribe to person.*
		_, subPersonErr := js.Subscribe("person.*", func(msg *nats.Msg) {

			iMsg := &inbox.InboxMessage{}
			iMsg.Subject = msg.Subject
			iMsg.Message = &inbox.Message{}

			// remove message on invalid subject
			if !validation.InArray(inboxPersonStreamSubjects, iMsg.Subject) {
				logger.Errorf("subscriber person: removed message with invalid subject %s", iMsg.Subject)
				ensureAck(msg)
				return
			}

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
				logger.Errorf("unable to validate message with id %s: %s", iMsg.Message.ID, vErrs.Error())
				ensureAck(msg)
				return
			}

			person, err := personService.GetPerson(context.Background(), iMsg.Message.ID)

			if err != nil && err == models.ErrNotFound {
				person = &models.Person{}
			} else if err != nil {
				// something is seriously wrong. We should stop processing records
				log.Fatal(fmt.Errorf("unable to fetch person record '%s': %w", iMsg.Message.ID, err))
			}

			// TODO update attributes during a delete?
			if iMsg.Subject == "person.update" {
				iMsg.UpdatePersonAttr(person)
				person.Active = true
			} else if iMsg.Subject == "person.delete" {
				person.Active = false
			}

			// report invalid changes to subject inbox.person.rejected
			if vErrs := person.Validate(); vErrs != nil {
				logger.Errorf("message with %s lead to invalid person record: %s", iMsg.Message.ID, vErrs)
				ensureAck(msg)
				return
			}

			var updateErr error
			var updatedPerson *models.Person
			if person.IsStored() {
				updatedPerson, updateErr = personService.UpdatePerson(context.Background(), person)
			} else {
				updatedPerson, updateErr = personService.CreatePerson(context.Background(), person)
			}

			// create/update failed: stop processing records
			if updateErr != nil {
				logger.Fatal(fmt.Errorf("unable to store person %s: %w", person.Id, updateErr))
			}

			person = updatedPerson

			logger.Infof("updated person %s", person.Id)

			// TODO: post serialized model.Person back to subject person.updated

			// acknowledge msg or die
			ensureAck(msg)

		},
			// second and next subscription will fail when using Bind
			nats.Bind(inboxStreamName, inboxPersonConsumerName),
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			// subscription specific option
			// without this message is acknowledged automatically (see "mack" in nats.SubOpt)
			nats.ManualAck(), // don't ask: why needed if AckExplicit is set?
		)

		if subPersonErr != nil {
			logger.Fatal(fmt.Errorf("unable to subscribe to nats subject person.*: %w", subPersonErr))
		}

		logger.Info("started to listen to messages at subject person.*")

		if err := initOrganizationInboxConsumer(js); err != nil {
			logger.Fatal(fmt.Errorf("unable to create nats consumer for organization: %w", err))
		}

		// subscribe to organization.*
		_, subOrgErr := js.Subscribe("organization.*", func(msg *nats.Msg) {

			ctx := context.Background()
			iMsg := &inbox.InboxMessage{}
			iMsg.Subject = msg.Subject
			iMsg.Message = &inbox.Message{}

			// remove message on invalid subject
			if !validation.InArray(inboxOrganizationStreamSubjects, iMsg.Subject) {
				logger.Errorf("subscriber org:removed message with invalid subject %s", iMsg.Subject)
				ensureAck(msg)
				return
			}

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
				logger.Errorf("unable to validate message with id %s: %s", iMsg.Message.ID, vErrs.Error())
				ensureAck(msg)
				return
			}

			org, err := organizationService.GetOrganization(ctx, iMsg.Message.ID)

			if err != nil && err == models.ErrNotFound {
				org = &models.Organization{}
			} else if err != nil {
				// something is seriously wrong. We should stop processing records
				log.Fatal(fmt.Errorf("unable to fetch organization record '%s': %w", iMsg.Message.ID, err))
			}

			// TODO update attributes during a delete?
			if iMsg.Subject == "organization.update" {

				iMsg.UpdateOrganizationAttr(org)

				// report invalid changes to subject inbox.organization.rejected
				if vErrs := org.Validate(); vErrs != nil {
					logger.Errorf("message with %s lead to invalid organization record: %s", iMsg.Message.ID, vErrs)
					ensureAck(msg)
					return
				}

				var updateErr error
				var updatedOrg *models.Organization
				if org.IsStored() {
					updatedOrg, updateErr = organizationService.UpdateOrganization(ctx, org)
				} else {
					updatedOrg, updateErr = organizationService.CreateOrganization(ctx, org)
				}

				// create/update failed: stop processing records
				if updateErr != nil {
					logger.Fatal(fmt.Errorf("unable to store organization %s: %w", org.Id, updateErr))
				}

				org = updatedOrg

			} else if iMsg.Subject == "organization.delete" {

				if org.IsStored() {
					if err := organizationService.DeleteOrganization(ctx, org.Id); err != nil {
						logger.Fatalf("unable to delete organization %s: %w", org.Id, err)
					}
				}

			}

			logger.Infof("updated organization %s from message %s", org.Id, iMsg.Message.ID)
			// TODO: republish updated record to subject organization.updated/organization.deleted

			// acknowledge msg or die
			ensureAck(msg)

		},
			// second and next subscription will fail when using Bind
			nats.Bind(inboxStreamName, inboxOrganizationConsumerName),
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			// subscription specific option
			// without this message is acknowledged automatically (see "mack" in nats.SubOpt)
			nats.ManualAck(), // don't ask: why needed if AckExplicit is set?
		)

		if subOrgErr != nil {
			logger.Fatal(fmt.Errorf("unable to subscribe to nats subject organization.*: %w", subOrgErr))
		}

		logger.Info("started to listen to messages at subject organization.*")

		runtime.Goexit() // wait for go routine to end (never)

	},
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}
