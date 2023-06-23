package cmd

import (
	"errors"
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
	"github.com/ugent-library/person-service/subscribers"
)

/*
	notes:

	- "nc.Drain()" not used. It introduced weird behaviour in our case (messages only sent after restart)
	- subscribe options:
		- Bind(streamName, consumerName) reserves connection to exactly one worker
		- AckExplicit: true
		- MaxAckPending: 1
		- ManualAck: true (not sure why this is need if AckExplicit is true)

	todo:

	- how to report malformed nats messages, as they do not have any id to track?
*/

var inboxListenCmd = &cobra.Command{
	Use: "listen",
	Run: func(cmd *cobra.Command, args []string) {

		nc, err := natsConnect(config.Nats)

		if err != nil {
			logger.Fatal(fmt.Errorf("unable to connect to nats: %w", err))
		}

		js, _ := nc.JetStream()

		subs, err := buildSubscribers(js, Services())

		if err != nil {
			logger.Fatal(err)
		}

		for _, sub := range subs {
			_, err := js.Subscribe(sub.Subject(), func(msg *nats.Msg) {
				iMsg, lErr := sub.Listen(msg)

				if errors.Is(lErr, subscribers.ErrFatal) {
					logger.Fatal(lErr)
				} else if errors.Is(lErr, subscribers.ErrNonFatal) {
					logger.Errorf("caught non fatal error: %s", lErr)
				} else if lErr != nil {
					logger.Errorf("caught unexpected error: %s", lErr)
				} else {
					logger.Infof("subject %s: processed message %s", sub.Subject(), iMsg.ID)
				}

				ensureAck(msg)
			}, sub.SubOpts()...)
			if err != nil {
				logger.Fatal(fmt.Errorf("unable to subscribe to nats subject %s: %w", sub.Subject(), err))
			}
			logger.Infof("listening to messages on nats subject %s", sub.Subject())
		}

		runtime.Goexit()

	},
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}
