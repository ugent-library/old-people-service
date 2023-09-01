package cmd

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/models"
)

var inboxListenOrganizationCmd = &cobra.Command{
	Use: "organization",
	RunE: func(cmd *cobra.Command, args []string) error {

		expB := backoff.NewExponentialBackOff()
		expB.MaxInterval = time.Minute
		b := backoff.WithMaxRetries(expB, 100)
		b = backoff.WithContext(b, cmd.Context())

		return backoff.Retry(func() error {
			err := listenOrganizationFn()
			if err != nil {
				logger.Error(err)
			}
			return err
		}, b)
	},
}

func listenOrganizationFn() error {
	nc, err := natsConnect(config.Nats)

	if err != nil {
		return fmt.Errorf("%w: unable to connect to nats: %w", models.ErrFatal, err)
	}

	js, _ := nc.JetStream()

	sub, err := newOrganizationSubscriber(js)

	if err != nil {
		return fmt.Errorf("%w: unable to build organization subscriber: %w", models.ErrFatal, err)
	}

	_, err = js.Subscribe(sub.Subject(), func(msg *nats.Msg) {
		id, lErr := sub.Process(msg)

		if errors.Is(lErr, models.ErrFatal) {
			logger.Fatal(lErr) // escape loop
		} else if errors.Is(lErr, models.ErrNonFatal) {
			logger.Errorf("subject %s: caught non fatal error: %s", sub.Subject(), lErr)
		} else if errors.Is(lErr, models.ErrSkipped) {
			logger.Errorf("subject %s: message was skipped: %s", lErr)
		} else if lErr != nil {
			logger.Errorf("subject %s: caught unexpected error: %s", sub.Subject(), lErr)
		} else {
			logger.Infof("subject %s: processed message %s", sub.Subject(), id)
		}

		if err := sub.EnsureAck(msg); err != nil {
			logger.Fatal(err)
		}
	}, sub.SubOpts()...)

	if err != nil {
		return fmt.Errorf("%w: nats subscriber %s failed with error: %w", models.ErrFatal, sub.Subject(), err)
	}
	logger.Infof("listening to messages on nats subject %s", sub.Subject())

	runtime.Goexit()

	return nil
}

func init() {
	inboxListenCmd.AddCommand(inboxListenOrganizationCmd)
}
