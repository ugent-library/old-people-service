package subscribers

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type Subcriber interface {
	Subject() string
	SubOpts() []nats.SubOpt
	Process(*nats.Msg) (string, error)
	EnsureAck(*nats.Msg) error
}

type BaseSubscriber struct {
	subject string
	subOpts []nats.SubOpt
}

type BaseConfig struct {
	Subject string
	SubOpts []nats.SubOpt
}

func (bs *BaseSubscriber) Subject() string {
	return bs.subject
}

func (bs *BaseSubscriber) SubOpts() []nats.SubOpt {
	return bs.subOpts
}

func (bs *BaseSubscriber) EnsureAck(msg *nats.Msg) error {
	if err := msg.Ack(); err != nil {
		return fmt.Errorf("unable to acknowledge nats message: %w", err)
	}
	return nil
}

func newBaseSubscriber(subject string) BaseSubscriber {
	return BaseSubscriber{
		subject: subject,
		subOpts: []nats.SubOpt{
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			nats.ManualAck(),
		},
	}
}
