package cli

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/gismo"
	"github.com/ugent-library/people-service/models"
)

var inboxListenOrganizationCmd = &cobra.Command{
	Use: "organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		return backOffRetry(cmd.Context(), organizationListener)
	},
}

func organizationListener() error {
	repo, err := newRepository()
	if err != nil {
		return err
	}

	gismoProc := gismo.NewOrganizationProcessor(repo)

	jsClient, err := newJetstreamClient()
	if err != nil {
		return fmt.Errorf("%w: unable to connect to nats: %w", models.ErrFatal, err)
	}

	_, err = jsClient.SubscribeOrganization(func(msg *nats.Msg) {
		id, err := gismoProc.Process(msg.Data)

		if errors.Is(err, models.ErrFatal) {
			logger.Fatal(err)
		} else if errors.Is(err, models.ErrSkipped) {
			logger.Errorf("subject %s: message was skipped: %s", jsClient.OrganizationSubject(), err)
		} else if err != nil {
			logger.Errorf("subject %s: caught unexpected error: %s", jsClient.OrganizationSubject(), err)
		} else {
			logger.Infof("subject %s: processed message %s", jsClient.OrganizationSubject(), id)
		}

		if err = msg.Ack(); err != nil {
			logger.Fatal(err)
		}
	})

	if err != nil {
		return fmt.Errorf("%w: nats subscriber %s failed with error: %w", models.ErrFatal, jsClient.OrganizationSubject(), err)
	}
	logger.Infof("listening to messages on nats subject %s", jsClient.OrganizationSubject())

	runtime.Goexit()

	return nil
}

func init() {
	inboxListenCmd.AddCommand(inboxListenOrganizationCmd)
}
