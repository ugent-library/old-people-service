package gismo

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ugent-library/people-service/models"
)

type OrganizationProcessor struct {
	repository models.Repository
}

func NewOrganizationProcessor(repo models.Repository) *OrganizationProcessor {
	return &OrganizationProcessor{
		repository: repo,
	}
}

func (op *OrganizationProcessor) Process(buf []byte) (*models.Message, error) {
	msg, err := parseOrganizationMessage(buf)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	org, err := op.repository.GetOrganizationByGismoId(ctx, msg.ID)
	if errors.Is(err, models.ErrNotFound) {
		org = models.NewOrganization()
	} else if err != nil {
		return nil, fmt.Errorf("%w: unable to fetch organization record: %s", models.ErrFatal, err)
	}

	if msg.Source == "gismo.organization.update" {
		now := time.Now()
		org.NameDut = ""
		org.NameEng = ""
		org.OtherId.Clear()
		org.Type = "organization"
		org.ParentId = ""
		org.GismoId = msg.ID

		// only recent values needed: name_dut, name_eng, type
		// all values needed: ugent_memorialis_id, code, biblio_code
		for _, attr := range msg.Attributes {
			withinDateRange := attr.ValidAt(now)
			switch attr.Name {
			case "parent_id":
				if withinDateRange {
					orgParentByGismo, err := op.repository.GetOrganizationByGismoId(ctx, attr.Value)
					if errors.Is(err, models.ErrNotFound) {
						orgParentByGismo := models.NewOrganization()
						orgParentByGismo.GismoId = attr.Value
						orgParentByGismo, err = op.repository.CreateOrganization(ctx, orgParentByGismo)
						if err != nil {
							return nil, fmt.Errorf("%w: unable to create parent organization: %s", models.ErrFatal, err)
						}
						org.ParentId = orgParentByGismo.Id
					} else if err != nil {
						return nil, fmt.Errorf("%w: unable to query database: %s", models.ErrFatal, err)
					} else {
						org.ParentId = orgParentByGismo.Id
					}
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
				org.Type = attr.Value
			case "ugent_memorialis_id":
				org.OtherId.Add("ugent_memorialis_id", attr.Value)
			case "code":
				org.OtherId.Add("ugent_id", attr.Value)
			case "biblio_code":
				org.OtherId.Add("biblio_id", attr.Value)
			}
		}

		if _, err := op.repository.SaveOrganization(ctx, org); err != nil {
			return nil, fmt.Errorf("%w: unable to save organization record: %s", models.ErrFatal, err)
		}
	} else if msg.Source == "gismo.organization.delete" {
		if org.IsStored() {
			if err := op.repository.DeleteOrganization(ctx, org.Id); err != nil {
				return nil, fmt.Errorf("%w: unable to delete organization record: %s", models.ErrFatal, err)
			}
		}
	}
	return msg, nil
}
