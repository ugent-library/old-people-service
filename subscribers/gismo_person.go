package subscribers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
	v1 "github.com/ugent-library/person-service/api/v1"
	"github.com/ugent-library/person-service/gismo"
	"github.com/ugent-library/person-service/inbox"
	"github.com/ugent-library/person-service/models"
	"github.com/ugent-library/person-service/ugentldap"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GismoPersonSubscriber struct {
	BaseSubscriber
	repository models.Repository
	ldapClient *ugentldap.UgentLdap
}

type GismoPersonConfig struct {
	BaseConfig
	Repository models.Repository
	LdapClient *ugentldap.UgentLdap
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
	bs := NewBaseSubscriber(config.Subject)
	bs.logger = config.Logger
	sub := &GismoPersonSubscriber{
		BaseSubscriber: bs,
		repository:     config.Repository,
		ldapClient:     config.LdapClient,
	}
	sub.subOpts = append(sub.subOpts, config.SubOpts...)
	return sub
}

func (ps *GismoPersonSubscriber) Process(msg *nats.Msg) (*inbox.Message, error) {
	ctx := context.Background()
	now := time.Now()

	// parse soap xml into json inbox message
	iMsg, err := gismo.ParsePersonMessage(msg.Data)
	if err != nil {
		return nil, fmt.Errorf("%w: unable to process malformed message: %s", models.ErrSkipped, err)
	}

	jsonBytes, _ := json.Marshal(iMsg)
	ps.logger.Infof("converted soap message %s into json: %s", iMsg.ID, string(jsonBytes))

	// Without ugentId no linking possible
	ugentIds := iMsg.GetAttributesAt("ugent_id", now)
	if len(ugentIds) == 0 {
		return nil, fmt.Errorf("%w: missing ugent_id in message %s", models.ErrSkipped, iMsg.ID)
	}

	// trial 1: retrieve old person by gismo-id
	person, err := ps.repository.GetPersonByGismoId(ctx, iMsg.ID)

	// trial 2: retrieve old person by ugent-id
	if errors.Is(err, models.ErrNotFound) {
		for _, ugentId := range ugentIds {
			person, err = ps.repository.GetPersonByOtherId(ctx, "historic_ugent_id", ugentId)
			if errors.Is(err, models.ErrNotFound) {
				continue
			}
			if err != nil {
				return iMsg, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
			}
			ps.logger.Infof("found match in table person on other_id.historic_ugent_id = %s", ugentId)
		}
	} else if err != nil {
		return iMsg, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	} else {
		ps.logger.Infof("found match in table person on gismo_id = %s", iMsg.ID)
	}

	// trial 3: new person record
	if person == nil {
		ps.logger.Infof("found no match in table person")
		person = models.NewPerson()
	}

	// clear old values
	person.GismoId = iMsg.ID
	person.OtherId = nil
	person.Email = ""
	person.FirstName = ""
	person.LastName = ""
	person.Orcid = ""
	person.Organization = nil
	person.PreferredFirstName = ""
	person.PreferredLastName = ""
	person.Title = ""
	person.Organization = nil
	var gismoOrganizationRefs []*v1.OrganizationRef

	// add attributes from GISMO
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
			person.OtherId = append(person.OtherId, &v1.IdRef{
				Type: "historic_ugent_id",
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
			orgRef := models.NewOrganizationRef(attr.Value)
			orgRef.From = timestamppb.New(*attr.StartDate)
			orgRef.Until = timestamppb.New(*attr.EndDate)
			gismoOrganizationRefs = append(gismoOrganizationRefs, orgRef)
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

	if len(gismoOrganizationRefs) > 0 {
		var gismoOrganizationIds []string
		for _, orgRef := range gismoOrganizationRefs {
			gismoOrganizationIds = append(gismoOrganizationIds, orgRef.Id)
		}
		orgsByGismo, err := ps.repository.GetOrganizationsByGismoId(ctx, gismoOrganizationIds...)
		if err != nil {
			return nil, err
		}
		// return fatal error when person arrives with organization that we do not know yet
		if len(orgsByGismo) != len(gismoOrganizationIds) {
			return nil, fmt.Errorf("%w: person.organization_id contains invalid gismo organization identifiers", models.ErrFatal)
		}

		var orgRefs []*v1.OrganizationRef
		for _, gismoOrgRef := range gismoOrganizationRefs {
			for _, org := range orgsByGismo {
				if gismoOrgRef.Id == org.GismoId {
					oRef := models.NewOrganizationRef(org.Id)
					oRef.From = gismoOrgRef.From
					oRef.Until = gismoOrgRef.Until
					orgRefs = append(orgRefs, oRef)
					break
				}
			}
		}
		person.Organization = orgRefs
	}

	// enrich with ugent ldap attributes
	ldapQueryParts := make([]string, 0, len(ugentIds))
	for _, ugentId := range ugentIds {
		ldapQueryParts = append(ldapQueryParts, fmt.Sprintf("(ugentHistoricIDs=%s)", ugentId))
	}
	ldapQuery := "(&" + strings.Join(ldapQueryParts, "") + ")"
	ldapPersons := make([]*models.Person, 0)
	err = ps.ldapClient.SearchPeople(ldapQuery, func(p *models.Person) error {
		ldapPersons = append(ldapPersons, p)
		return nil
	})
	if err != nil {
		ps.logger.Errorf("failed to query ugent ldap: %s", err)
		return nil, err
	}
	ps.logger.Infof("found %d matches for ugent id in ugent ldap", len(ldapPersons))

	// TODO: what if there are multiple matches?
	// TODO: what if we match the wrong user (ugent id is reused)

	// TODO: better check: ugentStudent or ugentEmployee also
	person.Active = len(ldapPersons) > 0

	if len(ldapPersons) >= 1 {
		ldapPerson := ldapPersons[0]
		person.FullName = ldapPerson.FullName
		for _, otherId := range ldapPerson.OtherId {
			if otherId.Type == "ugent_username" || otherId.Type == "ugent_barcode" {
				person.OtherId = append(person.OtherId, otherId)
			}
		}
		person.ExpirationDate = ldapPerson.ExpirationDate
		person.ObjectClass = ldapPerson.ObjectClass
	}

	// create/update record
	p, err := ps.repository.SavePerson(ctx, person)
	if err == nil {
		ps.logger.Infof("saved person %s", p.Id)
	} else {
		return iMsg, fmt.Errorf("%w: unable to save person record: %s", models.ErrFatal, err)
	}

	return iMsg, nil
}
