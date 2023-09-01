package subscribers

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/ugent-library/people-service/gismo"
	"github.com/ugent-library/people-service/models"
)

type GismoPersonSubscriber struct {
	BaseSubscriber
	gismoImporter *gismo.Importer
}

type GismoPersonConfig struct {
	BaseConfig
	GismoImporter *gismo.Importer
}

func NewGismoPersonSubscriber(config GismoPersonConfig) *GismoPersonSubscriber {
	bs := newBaseSubscriber(config.Subject)
	sub := &GismoPersonSubscriber{
		BaseSubscriber: bs,
		gismoImporter:  config.GismoImporter,
	}
	sub.subOpts = append(sub.subOpts, config.SubOpts...)
	return sub
}

func (ps *GismoPersonSubscriber) Process(msg *nats.Msg) (string, error) {
	iMsg, err := ps.gismoImporter.ImportPerson(msg.Data)
	if err != nil {
		return "", fmt.Errorf("%w: %w", models.ErrSkipped, err)
	}
	return iMsg.ID, nil
}
