package cmd

import (
	"github.com/spf13/cobra"
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
}

func init() {
	inboxCmd.AddCommand(inboxListenCmd)
}
