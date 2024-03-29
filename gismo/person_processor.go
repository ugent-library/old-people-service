package gismo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/samber/lo"
	"github.com/ugent-library/old-people-service/models"
	"github.com/ugent-library/old-people-service/ugentldap"
)

type PersonProcessor struct {
	repository        models.Repository
	ugentLdapSearcher ugentldap.Searcher
}

func NewPersonProcessor(repo models.Repository, ugentLdapSearcher ugentldap.Searcher) *PersonProcessor {
	return &PersonProcessor{
		repository:        repo,
		ugentLdapSearcher: ugentLdapSearcher,
	}
}

func (pp *PersonProcessor) Process(buf []byte) (*models.Message, error) {
	msg, err := ParsePersonMessage(buf)
	if err != nil {
		return nil, err
	}

	// retrieve old person by matching on attributes in gismo message
	// returns new person when no match is found
	person, err := pp.getPersonByMessage(msg)
	if err != nil {
		return nil, err
	}

	pp.clearOldPersonValues(person)

	// enrich person with incoming gismo attributes
	person, err = pp.enrichPersonWithMessage(person, msg)
	if err != nil {
		return nil, err
	}

	// enrich person with ugent ldap attributes -> TODO: very slow
	person, err = pp.enrichPersonWithLdap(person)
	if err != nil {
		return nil, err
	}

	pp.cleanupPerson(person)

	// create/update record
	if _, err = pp.repository.SavePerson(context.TODO(), person); err != nil {
		return nil, fmt.Errorf("%w: unable to save person record: %s", models.ErrFatal, err)
	}

	return msg, nil
}

func (pp *PersonProcessor) cleanupPerson(person *models.Person) {
	// non ldap users receive no full name
	if person.Name == "" && person.FamilyName != "" && person.GivenName != "" {
		person.Name = person.GivenName + " " + person.FamilyName
	}
}

func (pp *PersonProcessor) clearOldPersonValues(person *models.Person) {
	person.ClearIdentifier()
	person.Email = ""
	person.GivenName = ""
	person.FamilyName = ""
	person.Organization = nil
	person.PreferredGivenName = ""
	person.PreferredFamilyName = ""
	person.HonorificPrefix = ""
	person.Active = false
	person.JobCategory = nil
	person.ObjectClass = nil
	person.BirthDate = ""
}

func (pp *PersonProcessor) enrichPersonWithMessage(person *models.Person, msg *models.Message) (*models.Person, error) {
	now := time.Now()
	ctx := context.TODO()

	person.AddIdentifier(models.NewURN("gismo_id", msg.ID))
	var gismoOrganizationMembers []models.OrganizationMember

	// TODO: evaluate what to keep always
	// add attributes from GISMO
	for _, attr := range msg.Attributes {
		withinDateRange := attr.ValidAt(now)
		switch attr.Name {
		case "email":
			if withinDateRange {
				person.SetEmail(attr.Value)
			}
		case "given_name":
			if withinDateRange {
				person.GivenName = attr.Value
			}
		case "family_name":
			if withinDateRange {
				person.FamilyName = attr.Value
			}
		case "ugent_id":
			person.AddIdentifier(models.NewURN("ugent_id", attr.Value))
			person.AddIdentifier(models.NewURN("historic_ugent_id", attr.Value))
		case "ugent_memorialis_id":
			person.AddIdentifier(models.NewURN("ugent_memorialis_id", attr.Value))
		case "honorific_prefix":
			if withinDateRange {
				person.HonorificPrefix = attr.Value
			}
		case "organization_id":
			// duplicates expected (historical)
			orgMember := models.NewOrganizationMember(attr.Value)
			orgMember.From = attr.StartDate
			orgMember.Until = attr.EndDate
			gismoOrganizationMembers = append(gismoOrganizationMembers, orgMember)
		case "preferred_given_name":
			if withinDateRange {
				person.PreferredGivenName = attr.Value
			}
		case "preferred_family_name":
			if withinDateRange {
				person.PreferredFamilyName = attr.Value
			}
		case "orcid":
			if withinDateRange {
				person.AddIdentifier(models.NewURN("orcid", attr.Value))
			}
		}
	}

	if len(gismoOrganizationMembers) > 0 {
		var gismoOrganizationIds []string
		for _, orgMember := range gismoOrganizationMembers {
			gismoOrganizationIds = append(gismoOrganizationIds, orgMember.ID)
		}
		gismoOrganizationIds = lo.Uniq(gismoOrganizationIds)
		qGismoOrganizationIds := lo.Map(gismoOrganizationIds, func(id string, idx int) models.URN {
			return models.NewURN("gismo_id", id)
		})
		orgsByGismo, err := pp.repository.GetOrganizationsByIdentifier(ctx, qGismoOrganizationIds...)
		if err != nil {
			return nil, fmt.Errorf("%w: unable to fetch affiliated organization records: %s", models.ErrFatal, err)
		}

		// create dummy organizations when organization is not yet known
		for _, gismoOrganizationId := range gismoOrganizationIds {
			var gismoOrg *models.Organization
			for _, org := range orgsByGismo {
				if org.GetIdentifierValueByNS("gismo_id") == gismoOrganizationId {
					gismoOrg = org
					break
				}
			}
			if gismoOrg == nil {
				gismoOrg = models.NewOrganization()
				gismoOrg.AddIdentifier(models.NewURN("gismo_id", gismoOrganizationId))
				gismoOrg, err = pp.repository.SaveOrganization(ctx, gismoOrg)
				if err != nil {
					return nil, fmt.Errorf("%w: unable to save affiliated organization for person: %s", models.ErrFatal, err)
				}
				orgsByGismo = append(orgsByGismo, gismoOrg)
			}
		}

		var organizationMembers []models.OrganizationMember
		for _, gismoOrgMember := range gismoOrganizationMembers {
			for _, org := range orgsByGismo {
				if gismoOrgMember.ID == org.GetIdentifierValueByNS("gismo_id") {
					oMember := models.NewOrganizationMember(org.ID)
					oMember.From = gismoOrgMember.From
					oMember.Until = gismoOrgMember.Until
					organizationMembers = append(organizationMembers, oMember)
					break
				}
			}
		}
		person.SetOrganizationMember(organizationMembers...)
	}

	return person, nil
}

func (pp *PersonProcessor) getPersonByMessage(msg *models.Message) (*models.Person, error) {
	ctx := context.TODO()
	now := time.Now()

	// Without ugentId no linking possible
	ugentIds := lo.Uniq(msg.GetAttributesAt("ugent_id", now))
	if len(ugentIds) == 0 {
		return nil, fmt.Errorf("%w: missing ugent_id in message %s", models.ErrSkipped, msg.ID)
	}

	// trial 1: retrieve old person by gismo-id
	person, err := pp.repository.GetPersonByIdentifier(ctx, models.NewURN("gismo_id", msg.ID))
	if err == nil {
		return person, nil
	}
	if !errors.Is(err, models.ErrNotFound) {
		return nil, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	}

	// trial 2: retrieve old person by ugent-id
	qHistoricUgentIds := lo.Map(ugentIds, func(ugentId string, idx int) models.URN {
		return models.NewURN("historic_ugent_id", ugentId)
	})
	person, err = pp.repository.GetPersonByIdentifier(ctx, qHistoricUgentIds...)
	if err == nil {
		return person, nil
	}
	if !errors.Is(err, models.ErrNotFound) {
		return nil, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	}

	// trial 3: new person record
	return models.NewPerson(), nil
}

func (pp *PersonProcessor) enrichPersonWithLdap(person *models.Person) (*models.Person, error) {
	historicUgentIds := person.GetIdentifierValuesByNS("historic_ugent_id")
	ldapQueryParts := make([]string, 0, len(historicUgentIds))
	for _, ugentId := range historicUgentIds {
		ldapQueryParts = append(ldapQueryParts, fmt.Sprintf("(ugentHistoricIDs=%s)", ugentId))
	}
	ldapQuery := "(&" + strings.Join(ldapQueryParts, "") + ")"
	ldapEntries := make([]*ldap.Entry, 0)
	err := pp.ugentLdapSearcher.SearchPeople(ldapQuery, func(ldapEntry *ldap.Entry) error {
		ldapEntries = append(ldapEntries, ldapEntry)
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("%w: unable to query ugent ldap: %s", models.ErrFatal, err)
	}

	person.Active = len(ldapEntries) > 0

	if len(ldapEntries) >= 1 {
		ldapEntry := ldapEntries[0]

		for _, attr := range ldapEntry.Attributes {
			for _, val := range attr.Values {
				switch attr.Name {
				case "mail":
					person.Email = strings.ToLower(val)
				case "ugentPreferredGivenName":
					person.GivenName = val
				case "ugentPreferredSn":
					person.FamilyName = val
				case "displayName":
					person.Name = val
				case "ugentBarcode":
					person.AddIdentifier(models.NewURN("ugent_barcode", val))
				case "uid":
					person.AddIdentifier(models.NewURN("ugent_username", val))
				case "ugentExpirationDate":
					person.ExpirationDate = val
				case "objectClass":
					person.AddObjectClass(val)
				case "ugentJobCategory":
					person.AddJobCategory(val)
				case "ugentBirthDate":
					person.BirthDate = val
				}
			}
		}
	}

	return person, nil
}
