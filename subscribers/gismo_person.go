package subscribers

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/people-service/gismo"
	"github.com/ugent-library/people-service/models"
)

type GismoPersonSubscriber struct {
	BaseSubscriber
	gismoProcessor *gismo.Processor
}

type GismoPersonConfig struct {
	BaseConfig
	GismoProcessor *gismo.Processor
}

func NewGismoPersonSubscriber(config GismoPersonConfig) *GismoPersonSubscriber {
	bs := newBaseSubscriber(config.Subject)
	sub := &GismoPersonSubscriber{
		BaseSubscriber: bs,
		gismoProcessor: config.GismoProcessor,
	}
	sub.subOpts = append(sub.subOpts, config.SubOpts...)
	return sub
}

func (ps *GismoPersonSubscriber) Process(msg *nats.Msg) (string, error) {
	iMsg, err := ps.gismoProcessor.ImportPersonByMessage(msg.Data)
	if err != nil {
		return "", fmt.Errorf("%w: %w", models.ErrSkipped, err)
	}
	return iMsg.ID, nil
}
