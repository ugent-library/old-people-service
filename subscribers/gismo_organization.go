package subscribers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/nats-io/nats.go"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/gismo"
	"github.com/ugent-library/person-service/inbox"
	"github.com/ugent-library/person-service/models"
)

type GismoOrganizationSubscriber struct {
	BaseSubscriber
	repository models.Repository
}

type GismoOrganizationConfig struct {
	BaseConfig
	Repository models.Repository
}

func NewGismoOrganizationSubscriber(config GismoOrganizationConfig) *GismoOrganizationSubscriber {
	bs := NewBaseSubscriber(config.Subject)
	bs.logger = config.Logger
	os := &GismoOrganizationSubscriber{
		BaseSubscriber: bs,
		repository:     config.Repository,
	}
	os.subOpts = append(os.subOpts, config.SubOpts...)
	return os
}

func (oSub *GismoOrganizationSubscriber) Process(msg *nats.Msg) (*inbox.Message, error) {
	ctx := context.Background()
	now := time.Now()

	// parse soap xml message into json inbox message
	iMsg, err := gismo.ParseOrganizationMessage(msg.Data)
	if err != nil {
		return nil, fmt.Errorf("%w: unable to process malformed message: %s", models.ErrNonFatal, err)
	}

	jsonBytes, _ := json.Marshal(iMsg)
	oSub.logger.Infof("converted soap message %s into json: %s", iMsg.ID, string(jsonBytes))

	// fetch organization by gismo_id
	org, err := oSub.repository.GetOrganizationByGismoId(ctx, iMsg.ID)
	if err != nil && err == models.ErrNotFound {
		org = models.NewOrganization()
	} else if err != nil {
		return iMsg, fmt.Errorf("%w: unable to fetch organization record: %s", models.ErrFatal, err)
	}

	if iMsg.Source == "gismo.organization.update" {
		org.NameDut = ""
		org.NameEng = ""
		org.OtherId = make([]*v1.IdRef, 0)
		org.Type = "organization"
		org.ParentId = ""
		org.GismoId = iMsg.ID

		// only recent values needed: name_dut, name_eng, type
		// all values needed: ugent_memorialis_id, code, biblio_code
		for _, attr := range iMsg.Attributes {
			withinDateRange := attr.ValidAt(now)
			switch attr.Name {
			case "parent_id":
				if withinDateRange {
					orgParentByGismo, err := oSub.repository.GetOrganizationByGismoId(ctx, attr.Value)
					if err != nil {
						return nil, fmt.Errorf("%w: unable to find parent organization with gismo id %s", models.ErrNotFound, attr.Value)
					}
					org.ParentId = orgParentByGismo.Id
				}
			case "name_dut":
				if withinDateRange {
					org.NameDut = attr.Value
				}
			case "name_eng":
				if withinDateRange {
					org.NameEng = attr.Value
				}
			case "type":
				if withinDateRange {
					org.Type = attr.Value
				}
			case "ugent_memorialis_id":
				org.OtherId = append(org.OtherId, &v1.IdRef{
					Type: "ugent_memorialis_id",
					Id:   attr.Value,
				})
			case "code":
				org.OtherId = append(org.OtherId, &v1.IdRef{
					Type: "ugent_id",
					Id:   attr.Value,
				})
			case "biblio_code":
				org.OtherId = append(org.OtherId, &v1.IdRef{
					Type: "biblio_id",
					Id:   attr.Value,
				})
			}
		}

		if org.IsStored() {
			o, err := oSub.repository.UpdateOrganization(ctx, org)
			if err == nil {
				oSub.logger.Infof("updated organization %s", o.Id)
			}
		} else {
			o, err := oSub.repository.CreateOrganization(ctx, org)
			if err == nil {
				oSub.logger.Infof("created organization %s", o.Id)
			}
		}
		if err != nil {
			return iMsg, fmt.Errorf("%w: unable to store organization record: %s", models.ErrFatal, err)
		}
	} else if iMsg.Source == "gismo.organization.delete" {
		if org.IsStored() {
			if err := oSub.repository.DeleteOrganization(ctx, org.Id); err != nil {
				return iMsg, fmt.Errorf("%w: unable to delete organization record: %s", models.ErrFatal, err)
			}
		}
	}

	return iMsg, nil
}
