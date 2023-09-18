package subscribers

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/people-service/gismo"
	"github.com/ugent-library/people-service/models"
)

type GismoOrganizationSubscriber struct {
	BaseSubscriber
	gismoProcessor *gismo.Processor
}

type GismoOrganizationConfig struct {
	BaseConfig
	GismoProcessor *gismo.Processor
}

func NewGismoOrganizationSubscriber(config GismoOrganizationConfig) *GismoOrganizationSubscriber {
	bs := newBaseSubscriber(config.Subject)
	os := &GismoOrganizationSubscriber{
		BaseSubscriber: bs,
		gismoProcessor: config.GismoProcessor,
	}
	os.subOpts = append(os.subOpts, config.SubOpts...)
	return os
}

func (oSub *GismoOrganizationSubscriber) Process(msg *nats.Msg) (string, error) {
	iMsg, err := oSub.gismoProcessor.ImportOrganizationByMessage(msg.Data)
	if err != nil {
		return "", fmt.Errorf("%w: %w", models.ErrSkipped, err)
	}
	return iMsg.ID, nil
}
