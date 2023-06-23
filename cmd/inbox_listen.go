package cmd

import (
	"context"
	"fmt"
	"log"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/person-service/gismo"
	"github.com/ugent-library/person-service/models"
)

/*
	notes:

	- "nc.Drain()" not used. It introduced weird behaviour in our case (messages only sent after restart)
	- subscribe options:
		- Bind(streamName, consumerName) reserves connection to exactly one worker
		- AckExplicit: true
		- MaxAckPending: 1
		- ManualAck: true (not sure why this is need if AckExplicit is true)
*/

var inboxListenCmd = &cobra.Command{
	Use: "listen",
	Run: func(cmd *cobra.Command, args []string) {
		personService := Services().PersonService
		organizationService := Services().OrganizationService

		nc, err := natsConnect(config.Nats)

		if err != nil {
			logger.Fatal(fmt.Errorf("unable to connect to nats: %w", err))
		}

		js, _ := nc.JetStream()

		if err := initConsumer(js, natsStreamConfig.Name, &natsPersonConsumerConfig); err != nil {
			logger.Fatal(fmt.Errorf("unable to create nats consumer for %s: %w", personSubject, err))
		}

		// subscribe to gismo.person xml messages
		// subject are translated into person.update and person.delete
		_, subPersonErr := js.Subscribe(personSubject, func(msg *nats.Msg) {
			ctx := context.Background()
			iMsg, err := gismo.ParsePersonMessage(msg.Data)

			if err != nil {
				logger.Errorf("subscriber %s: unable to process malformed message: %s", personSubject, err)
				ensureAck(msg)
				return
			}

			person, err := personService.GetPerson(ctx, iMsg.ID)

			if err != nil && err == models.ErrNotFound {
				person = models.NewPerson()
			} else if err != nil {
				log.Fatal(fmt.Errorf("subscriber %s: unable to fetch person record '%s': %w", personSubject, iMsg.ID, err))
			}

			if iMsg.Subject == "person.update" {
				iMsg.UpdatePersonAttr(person)
				person.Active = true
			} else if iMsg.Subject == "person.delete" {
				person.Active = false
			}

			if person.IsStored() {
				person, err = personService.UpdatePerson(ctx, person)
			} else {
				person, err = personService.CreatePerson(ctx, person)
			}

			if err != nil {
				logger.Fatal(fmt.Errorf("subscriber %s: unable to store person %s: %w", organizationSubject, iMsg.ID, err))
			}

			logger.Infof("subscriber %s: updated person %s", organizationSubject, person.Id)

			ensureAck(msg)

		},
			nats.Bind(natsStreamConfig.Name, natsPersonConsumerConfig.Durable),
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			nats.ManualAck(),
		)

		if subPersonErr != nil {
			logger.Fatal(fmt.Errorf("unable to subscribe to nats subject %s: %w", personSubject, subPersonErr))
		}

		logger.Infof("started to listen to messages at subject %s", personSubject)

		if err := initConsumer(js, natsStreamConfig.Name, &natsOrganizationConsumerConfig); err != nil {
			logger.Fatal(fmt.Errorf("unable to create nats consumer for %s: %w", organizationSubject, err))
		}

		// subscribe to gismo.organization
		// subject is translated into organization.update and organization.delete
		_, subOrgErr := js.Subscribe(organizationSubject, func(msg *nats.Msg) {
			ctx := context.Background()
			iMsg, err := gismo.ParseOrganizationMessage(msg.Data)

			if err != nil {
				logger.Errorf("subscriber %s: unable to process malformed message: %s", organizationSubject, err)
				ensureAck(msg)
				return
			}

			org, err := organizationService.GetOrganization(ctx, iMsg.ID)

			if err != nil && err == models.ErrNotFound {
				org = models.NewOrganization()
			} else if err != nil {
				log.Fatal(fmt.Errorf("subscriber %s: unable to fetch organization record '%s': %w", organizationSubject, iMsg.ID, err))
			}

			if iMsg.Subject == "organization.update" {
				iMsg.UpdateOrganizationAttr(org)

				if org.IsStored() {
					org, err = organizationService.UpdateOrganization(ctx, org)
				} else {
					org, err = organizationService.CreateOrganization(ctx, org)
				}

				if err != nil {
					logger.Fatal(fmt.Errorf("subscriber %s: unable to store organization %s: %w", organizationSubject, iMsg.ID, err))
				}
			} else if iMsg.Subject == "organization.delete" {
				if org.IsStored() {
					if err := organizationService.DeleteOrganization(ctx, org.Id); err != nil {
						logger.Fatalf("subscriber %s: unable to delete organization %s: %w", organizationSubject, org.Id, err)
					}
				}
			}

			logger.Infof("subscriber %s: updated organization %s", organizationSubject, org.Id)

			ensureAck(msg)

		},
			nats.Bind(natsStreamConfig.Name, natsOrganizationConsumerConfig.Durable),
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			nats.ManualAck(),
		)

		if subOrgErr != nil {
			logger.Fatal(fmt.Errorf("unable to subscribe to nats subject %s: %w", organizationSubject, subOrgErr))
		}

		logger.Infof("started to listen to messages at subject %s", organizationSubject)

		runtime.Goexit() // wait for go routine to end (never)

	},
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}
