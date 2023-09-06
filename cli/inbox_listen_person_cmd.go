package cli

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/models"
)

var inboxListenPersonCmd = &cobra.Command{
	Use: "person",
	RunE: func(cmd *cobra.Command, args []string) error {
		return backOffRetry(cmd.Context(), listenPersonFn)
	},
}

func listenPersonFn() error {
	nc, err := natsConnect(config.Nats)

	if err != nil {
		return fmt.Errorf("%w: unable to connect to nats: %w", models.ErrFatal, err)
	}

	js, _ := nc.JetStream()

	sub, err := newPersonSubscriber(js)

	if err != nil {
		return fmt.Errorf("%w: unable to build person subscriber: %w", models.ErrFatal, err)
	}

	_, err = js.Subscribe(sub.Subject(), func(msg *nats.Msg) {
		id, lErr := sub.Process(msg)

		if errors.Is(lErr, models.ErrFatal) {
			logger.Fatal(lErr) // escape loop
		} else if errors.Is(lErr, models.ErrNonFatal) {
			logger.Errorf("subject %s: caught non fatal error: %s", sub.Subject(), lErr)
		} else if errors.Is(lErr, models.ErrSkipped) {
			logger.Errorf("subject %s: message was skipped: %s", sub.Subject(), lErr)
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
	inboxListenCmd.AddCommand(inboxListenPersonCmd)
}
