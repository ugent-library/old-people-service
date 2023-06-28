package subscribers

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/gismo"
	"github.com/ugent-library/person-service/inbox"
	"github.com/ugent-library/person-service/models"
	"github.com/ugent-library/person-service/validation"
)

type GismoPersonSubscriber struct {
	BaseSubscriber
	repository models.Repository
}

type GismoPersonConfig struct {
	BaseConfig
	Repository models.Repository
}

/*
managed fields:
- gismo_id (set as id in inbox.Message)
- first_name
- last_name
- preferred_first_name
- preferred_last_name
- title
- ugent_id
- orcid
- email
- ugent_memorialis_id
- organization_id
*/

func NewGismoPersonSubscriber(config GismoPersonConfig) *GismoPersonSubscriber {
	sub := &GismoPersonSubscriber{
		BaseSubscriber: NewBaseSubscriber(config.Subject),
		repository:     config.Repository,
	}
	sub.subOpts = append(sub.subOpts, config.SubOpts...)
	return sub
}

func (ps *GismoPersonSubscriber) Listen(msg *nats.Msg) (*inbox.Message, error) {
	ctx := context.Background()
	now := time.Now()

	// parse soap xml into json inbox message
	iMsg, err := gismo.ParsePersonMessage(msg.Data)
	if err != nil {
		return nil, fmt.Errorf("%w: unable to process malformed message: %s", models.ErrNonFatal, err)
	}

	// trial 1: retrieve old person by gismo-id
	person, err := ps.repository.GetPersonByGismoId(ctx, iMsg.ID)

	// trial 2: retrieve old person by ugent-id
	// TODO: panic when no ugent_id is set??
	if err != nil && err == models.ErrNotFound {
		ugentId, err := iMsg.GetAttributeAt("ugent_id", now)
		if err == nil {
			person, err = ps.repository.GetPersonByUgentId(ctx, ugentId)
			if err != nil && err != models.ErrNotFound {
				return iMsg, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
			}
		}
		if person == nil {
			person = models.NewPerson()
		}
	} else if err != nil {
		return iMsg, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	}

	// clear old values
	person.Active = iMsg.Source == "gismo.person.update"
	person.GismoId = iMsg.ID
	otherId := make([]*v1.IdRef, 0)
	for _, oi := range person.OtherId {
		if oi.Type != "ugent_id" && oi.Type != "ugent_memorialis_id" {
			otherId = append(otherId, oi)
		}
	}
	person.OtherId = otherId
	person.Email = ""
	person.FirstName = ""
	person.LastName = ""
	person.Orcid = ""
	person.OrganizationId = make([]string, 0)
	person.PreferredFirstName = ""
	person.PreferredLastName = ""
	person.Title = ""
	gismoOrganizationId := make([]string, 0)

	for _, attr := range iMsg.Attributes {
		withinDateRange := attr.ValidAt(now)
		switch attr.Name {
		case "email":
			if withinDateRange {
				person.Email = strings.ToLower(attr.Value)
			}
		case "first_name":
			if withinDateRange {
				person.FirstName = attr.Value
			}
		case "last_name":
			if withinDateRange {
				person.LastName = attr.Value
			}
		case "ugent_id":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_id",
				Id:   attr.Value,
			})
		case "ugent_memorialis_id":
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "ugent_memorialis_id",
				Id:   attr.Value,
			})
		case "title":
			if withinDateRange {
				person.Title = attr.Value
			}
		case "organization_id":
			if withinDateRange {
				gismoOrganizationId = append(gismoOrganizationId, attr.Value)
			}
		case "preferred_first_name":
			if withinDateRange {
				person.PreferredFirstName = attr.Value
			}
		case "preferred_last_name":
			if withinDateRange {
				person.PreferredLastName = attr.Value
			}
		case "orcid":
			if withinDateRange {
				person.Orcid = attr.Value
			}
		}
	}

	gismoOrganizationId = validation.Uniq(gismoOrganizationId)

	if len(gismoOrganizationId) > 0 {
		orgIds := make([]string, 0)
		orgsByGismo, err := ps.repository.GetOrganizationsByGismoId(ctx, gismoOrganizationId...)
		if err != nil {
			return nil, err
		}
		if len(orgsByGismo) != len(gismoOrganizationId) {
			return nil, fmt.Errorf("person.organization_id contains invalid gismo organization identifiers")
		}
		for _, org := range orgsByGismo {
			orgIds = append(orgIds, org.Id)
		}
		person.OrganizationId = orgIds
	}

	if person.IsStored() {
		_, err = ps.repository.UpdatePerson(ctx, person)
	} else {
		_, err = ps.repository.CreatePerson(ctx, person)
	}

	if err != nil {
		return iMsg, fmt.Errorf("%w: unable to store person record: %s", models.ErrFatal, err)
	}

	return iMsg, nil
}
