package cmd

import (
	"errors"
	"fmt"
	"runtime"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
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
	- nats subscribers (js.Subscribe) use go routines, and therefore need runtime.Goexit() to wait for
	- backoff wrapper retries main function a 100 times until it finally stops, using a exponential wait time
	  between takes, and a maximum wait time of one minute

	todo:

	- how to report malformed nats messages, as they do not have any id to track?
*/

var inboxListenCmd = &cobra.Command{
	Use: "listen",
	Run: func(cmd *cobra.Command, args []string) {

		expB := backoff.NewExponentialBackOff()
		expB.MaxInterval = time.Minute
		b := backoff.WithMaxRetries(expB, 100)
		b = backoff.WithContext(b, cmd.Context())

		err := backoff.Retry(func() error {
			err := listenFn()
			if err != nil {
				logger.Error(err)
			}
			return err
		}, b)

		if err != nil {
			logger.Errorf("fatal error: %s", err)
		}

	},
}

func listenFn() error {
	nc, err := natsConnect(config.Nats)

	if err != nil {
		return fmt.Errorf("%w: unable to connect to nats: %w", models.ErrFatal, err)
	}

	js, _ := nc.JetStream()

	subs, err := buildSubscribers(js, Services())

	if err != nil {
		return fmt.Errorf("%w: unable to build subscribers: %w", models.ErrFatal, err)
	}

	for _, sub := range subs {
		_, err := js.Subscribe(sub.Subject(), func(msg *nats.Msg) {
			iMsg, lErr := sub.Listen(msg)

			if errors.Is(lErr, models.ErrFatal) {
				logger.Fatal(lErr) // escape loop
			} else if errors.Is(lErr, models.ErrNonFatal) {
				logger.Errorf("caught non fatal error: %s", lErr)
			} else if lErr != nil {
				logger.Errorf("caught unexpected error: %s", lErr)
			} else {
				logger.Infof("subject %s: processed message %s", sub.Subject(), iMsg.ID)
			}

			ensureAck(msg)
		}, sub.SubOpts()...)
		if err != nil {
			return fmt.Errorf("%w: nats subscriber %s failed with error: %w", models.ErrFatal, sub.Subject(), err)
		}
		logger.Infof("listening to messages on nats subject %s", sub.Subject())
	}

	runtime.Goexit()

	return nil
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}
