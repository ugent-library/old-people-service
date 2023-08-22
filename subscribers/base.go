package subscribers

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/people-service/inbox"
	"go.uber.org/zap"
)

type Subcriber interface {
	Subject() string
	SubOpts() []nats.SubOpt
	Process(*nats.Msg) (*inbox.Message, error)
	EnsureAck(*nats.Msg)
}

type BaseSubscriber struct {
	subject string
	subOpts []nats.SubOpt
	logger  *zap.SugaredLogger
}

type BaseConfig struct {
	Subject string
	SubOpts []nats.SubOpt
	Logger  *zap.SugaredLogger
}

func (bs *BaseSubscriber) Subject() string {
	return bs.subject
}

func (bs *BaseSubscriber) SubOpts() []nats.SubOpt {
	return bs.subOpts
}

func (bs *BaseSubscriber) EnsureAck(msg *nats.Msg) {
	if err := msg.Ack(); err != nil {
		bs.logger.Fatal(fmt.Errorf("unable to acknowledge nats message: %w", err))
	}
}

func NewBaseSubscriber(subject string) BaseSubscriber {
	return BaseSubscriber{
		subject: subject,
		subOpts: []nats.SubOpt{
			nats.AckExplicit(),
			nats.MaxAckPending(1),
			nats.ManualAck(),
		},
	}
}
