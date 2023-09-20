package cli

import (
	"context"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/spf13/cobra"
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
