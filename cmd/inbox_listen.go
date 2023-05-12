package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people/inbox"
	"github.com/ugent-library/people/models"
	"github.com/ugent-library/people/validation"
)

/*
TODO: on delete of organizations push all related people to person.updated?
*/

var inboxListenCmd = &cobra.Command{
	Use: "listen",
	Run: func(cmd *cobra.Command, args []string) {

		services := Services()
		personService := services.PersonService
		organizationService := services.OrganizationService

		nc, err := natsConnect(config.Nats)

		if err != nil {
			logger.Fatal(fmt.Errorf("unable to connect to nats: %w", err))
		}

		// "drain" introduces weird behaviour (messages only sent after restart)
		//defer nc.Drain()

		js, _ := nc.JetStream()

		if err := initConsumer(js, natsStreamConfig.Name, &natsPersonConsumerConfig); err != nil {
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
			nats.Bind(natsStreamConfig.Name, natsPersonConsumerConfig.Durable),
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

		if err := initConsumer(js, natsStreamConfig.Name, &natsOrganizationConsumerConfig); err != nil {
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
			nats.Bind(natsStreamConfig.Name, natsOrganizationConsumerConfig.Durable),
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
