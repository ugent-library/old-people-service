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
var outboxStreamName = "PEOPLE_SERVICE"
var outboxStreamSubjects = []string{
	"person.updated",
	"organization.updated",
	"organization.deleted",
	"inbox.rejected",
	"inbox.person.rejected",
	"inbox.organization.rejected",
}

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
		personSearchService := services.PersonSearchService
		organizationService := services.OrganizationService
		organizationSearchService := services.OrganizationSearchService

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

			person, err := personService.GetPerson(context.Background(), iMsg.Message.ID)

			if err != nil && err == models.ErrNotFound {
				person = &models.Person{}
			} else if err != nil {
				// something is seriously wrong. We should stop processing records
				log.Fatal(fmt.Errorf("unable to fetch person record '%s': %w", iMsg.Message.ID, err))
			}

			oldPerson := person.Dup()

			// TODO update attributes during a delete?
			if iMsg.Subject == "person.update" {
				iMsg.UpdatePersonAttr(person)
				person.Active = true
			} else if iMsg.Subject == "person.delete" {
				person.Active = false
			}

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

			// index
			if err := personSearchService.Index(person); err != nil {
				logger.Errorf("unable to index person %s: %s", person.Id, err)
			}

			// outbox subject
			outboxSubject := "person.updated"

			// republish updated record to subject person.update
			logger.Infof("updated person %s via subject %s", person.Id, outboxSubject)
			outboxBytes, _ := json.Marshal(person)

			// failed to contact nats: stop processing records
			if err := nc.Publish(outboxSubject, outboxBytes); err != nil {
				logger.Fatal(fmt.Errorf("unable to publish message to subject %s: %w", outboxSubject, err))
			}

			logger.Infof("published person %s to subject %s", person.Id, outboxSubject)

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

			org, err := organizationService.GetOrganization(ctx, iMsg.Message.ID)

			if err != nil && err == models.ErrNotFound {
				org = &models.Organization{}
			} else if err != nil {
				// something is seriously wrong. We should stop processing records
				log.Fatal(fmt.Errorf("unable to fetch organization record '%s': %w", iMsg.Message.ID, err))
			}

			// TODO update attributes during a delete?
			if iMsg.Subject == "organization.update" {

				oldOrg := org.Dup()

				iMsg.UpdateOrganizationAttr(org)

				// report invalid changes to subject inbox.organization.rejected
				if vErrs := org.Validate(); vErrs != nil {
					pChangeErr := &inbox.OrganizationChangeError{
						OldOrganization: oldOrg,
						NewOrganization: org,
						Errors:          vErrs,
					}
					outboxBytes, _ := json.Marshal(pChangeErr)
					if err := nc.Publish("inbox.organization.rejected", outboxBytes); err != nil {
						logger.Fatal(fmt.Errorf("unable to publish message to subject %s: %w", "inbox.organization.rejected", err))
					}
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

				// index
				if err := organizationSearchService.Index(org); err != nil {
					logger.Errorf("unable to index organization %s: %s", org.Id, err)
				}

			} else if iMsg.Subject == "organization.delete" {

				if org.IsStored() {
					if err := organizationService.DeleteOrganization(ctx, org.Id); err != nil {
						logger.Fatalf("unable to delete organization %s: %w", org.Id, err)
					}
					if err := organizationSearchService.Delete(org.Id); err != nil {
						logger.Fatalf("unable to delete organization %s from index: %w", org.Id, err)
					}
				}

			}

			// outbox subject
			outboxSubject := "organization.updated"
			if msg.Subject == "organization.delete" {
				outboxSubject = "organization.deleted"
			}

			// republish updated record to subject organization.updated
			logger.Infof("updated organization %s via subject %s", org.Id, outboxSubject)
			outboxBytes, _ := json.Marshal(org)

			// failed to contact nats: stop processing records
			if err := nc.Publish(outboxSubject, outboxBytes); err != nil {
				logger.Fatal(fmt.Errorf("unable to publish message to subject %s: %w", outboxSubject, err))
			}

			logger.Infof("published organization %s to subject %s", org.Id, outboxSubject)

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
