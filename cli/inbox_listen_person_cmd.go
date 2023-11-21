package cli

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/old-people-service/cache"
	"github.com/ugent-library/old-people-service/gismo"
	"github.com/ugent-library/old-people-service/models"
)

var inboxListenPersonCmd = &cobra.Command{
	Use: "person",
	RunE: func(cmd *cobra.Command, args []string) error {
		return backOffRetry(cmd.Context(), personListener)
	},
}

func personListener() error {
	repo, err := newRepository()
	if err != nil {
		return err
	}

	ugentLdapSearcher := cache.NewUgentLdapSearcher(newUgentLdapClient(), 100, 10*time.Minute)
	gismoProc := gismo.NewPersonProcessor(repo, ugentLdapSearcher)

	jsClient, err := newJetstreamClient()
	if err != nil {
		return fmt.Errorf("%w: unable to connect to nats: %w", models.ErrFatal, err)
	}

	_, err = jsClient.SubscribePerson(func(msg *nats.Msg) {
		id, err := gismoProc.Process(msg.Data)

		if errors.Is(err, models.ErrFatal) {
			logger.Fatal(err)
		} else if errors.Is(err, models.ErrSkipped) {
			logger.Errorf("subject %s: message was skipped: %s", jsClient.PersonSubject(), err)
		} else if err != nil {
			logger.Errorf("subject %s: caught unexpected error: %s", jsClient.PersonSubject(), err)
		} else {
			logger.Infof("subject %s: processed message %s", jsClient.PersonSubject(), id)
		}

		if err = msg.Ack(); err != nil {
			logger.Fatal(err)
		}
	})

	if err != nil {
		return fmt.Errorf("%w: nats subscriber %s failed with error: %w", models.ErrFatal, jsClient.PersonSubject(), err)
	}
	logger.Infof("listening to messages on nats subject %s", jsClient.PersonSubject())

	runtime.Goexit()

	return nil
}

func init() {
	inboxListenCmd.AddCommand(inboxListenPersonCmd)
}
