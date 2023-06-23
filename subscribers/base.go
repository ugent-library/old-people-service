package subscribers

import (
	"errors"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/person-service/inbox"
)

var ErrFatal = errors.New("fatal error")
var ErrNonFatal = errors.New("non fatal error")

type Subcriber interface {
	Subject() string
	SubOpts() []nats.SubOpt
	Listen(*nats.Msg) (*inbox.Message, error)
}

type BaseSubscriber struct {
	subject string
	subOpts []nats.SubOpt
}

func (bs *BaseSubscriber) Subject() string {
	return bs.subject
}

func (bs *BaseSubscriber) SubOpts() []nats.SubOpt {
	return bs.subOpts
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
