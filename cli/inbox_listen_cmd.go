package cli

import (
	"context"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/jetstreamclient"
)

var inboxListenCmd = &cobra.Command{
	Use: "listen",
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}

// TODO: make this configurable
func backOffRetry(ctx context.Context, fn func() error) error {
	expB := backoff.NewExponentialBackOff()
	expB.MaxInterval = time.Minute
	b := backoff.WithMaxRetries(expB, 100)
	b = backoff.WithContext(b, ctx)

	return backoff.Retry(func() error {
		err := fn()
		if err != nil {
			logger.Error(err)
		}
		return err
	}, b)
}

func newJetstreamClient() (*jetstreamclient.Client, error) {
	return jetstreamclient.New(&jetstreamclient.Config{
		NatsUrl:    config.Nats.Url,
		StreamName: config.Nats.StreamName,
		Nkey:       config.Nats.Nkey,
		NkeySeed:   config.Nats.NkeySeed,
		Logger:     logger,
	})
}
