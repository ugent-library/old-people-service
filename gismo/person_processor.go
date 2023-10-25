package gismo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/go-ldap/ldap/v3"
	"github.com/samber/lo"
	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/ugentldap"
)

type PersonProcessor struct {
	repository      models.Repository
	ugentLdapClient *ugentldap.Client
}

func NewPersonProcessor(repo models.Repository, ugentLdapClient *ugentldap.Client) *PersonProcessor {
	return &PersonProcessor{
		repository:      repo,
		ugentLdapClient: ugentLdapClient,
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

	// create/update record
	if _, err = pp.repository.SavePerson(context.TODO(), person); err != nil {
		return nil, fmt.Errorf("%w: unable to save person record: %s", models.ErrFatal, err)
	}

	return msg, nil
}

func (pp *PersonProcessor) enrichPersonWithMessage(person *models.Person, msg *models.Message) (*models.Person, error) {
	now := time.Now()
	ctx := context.TODO()

	// clear old values
	person.ClearIdentifier()
	person.AddIdentifier("gismo_id", msg.ID)
	person.Email = ""
	person.GivenName = ""
	person.FamilyName = ""
	person.Organization = nil
	person.PreferredGivenName = ""
	person.PreferredFamilyName = ""
	person.HonorificPrefix = ""
	var gismoOrganizationRefs []*models.OrganizationMember

	// add attributes from GISMO
	for _, attr := range msg.Attributes {
		withinDateRange := attr.ValidAt(now)
		switch attr.Name {
		case "email":
			if withinDateRange {
				person.Email = strings.ToLower(attr.Value)
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
			person.AddIdentifier("ugent_id", attr.Value)
			person.AddIdentifier("historic_ugent_id", attr.Value)
		case "ugent_memorialis_id":
			person.AddIdentifier("ugent_memorialis_id", attr.Value)
		case "honorific_prefix":
			if withinDateRange {
				person.HonorificPrefix = attr.Value
			}
		case "organization_id":
			// sometimes double entries
			found := false
			for _, orgRef := range gismoOrganizationRefs {
				if orgRef.Id == attr.Value {
					found = true
				}
			}
			if !found {
				orgRef := models.NewOrganizationMember(attr.Value)
				orgRef.From = attr.StartDate
				orgRef.Until = attr.EndDate
				gismoOrganizationRefs = append(gismoOrganizationRefs, orgRef)
			}
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
				person.AddIdentifier("orcid", attr.Value)
			}
		}
	}

	if len(gismoOrganizationRefs) > 0 {
		var gismoOrganizationIds []string
		for _, orgRef := range gismoOrganizationRefs {
			gismoOrganizationIds = append(gismoOrganizationIds, orgRef.Id)
		}
		orgsByGismo, err := pp.repository.GetOrganizationsByIdentifier(ctx, "gismo_id", gismoOrganizationIds...)
		if err != nil {
			return nil, fmt.Errorf("%w: unable to fetch affiliated organization records: %s", models.ErrFatal, err)
		}

		// create dummy organizations when organization is not yet known
		for _, gismoOrganizationId := range gismoOrganizationIds {
			var gismoOrg *models.Organization
			for _, org := range orgsByGismo {
				if org.GetIdentifierValue("gismo_id") == gismoOrganizationId {
					gismoOrg = org
					break
				}
			}
			if gismoOrg == nil {
				gismoOrg = models.NewOrganization()
				gismoOrg.AddIdentifier("gismo_id", gismoOrganizationId)
				gismoOrg, err = pp.repository.SaveOrganization(ctx, gismoOrg)
				if err != nil {
					return nil, fmt.Errorf("%w: unable to save affiliated organization for person: %s", models.ErrFatal, err)
				}
				orgsByGismo = append(orgsByGismo, gismoOrg)
			}
		}

		var orgRefs []*models.OrganizationMember
		for _, gismoOrgMember := range gismoOrganizationRefs {
			for _, org := range orgsByGismo {
				if gismoOrgMember.Id == org.GetIdentifierValue("gismo_id") {
					oMember := models.NewOrganizationMember(org.ID)
					oMember.From = gismoOrgMember.From
					oMember.Until = gismoOrgMember.Until
					orgRefs = append(orgRefs, oMember)
					break
				}
			}
		}
		person.Organization = orgRefs
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
	person, err := pp.repository.GetPersonByIdentifier(ctx, "gismo_id", msg.ID)
	if err == nil {
		return person, nil
	}
	if !errors.Is(err, models.ErrNotFound) {
		return nil, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	}

	// trial 2: retrieve old person by ugent-id
	person, err = pp.repository.GetPersonByIdentifier(ctx, "historic_ugent_id", ugentIds...)
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
	historicUgentIds := person.GetIdentifierValues("historic_ugent_id")
	ldapQueryParts := make([]string, 0, len(historicUgentIds))
	for _, ugentId := range historicUgentIds {
		ldapQueryParts = append(ldapQueryParts, fmt.Sprintf("(ugentHistoricIDs=%s)", ugentId))
	}
	ldapQuery := "(&" + strings.Join(ldapQueryParts, "") + ")"
	ldapEntries := make([]*ldap.Entry, 0)
	err := pp.ugentLdapClient.SearchPeople(ldapQuery, func(ldapEntry *ldap.Entry) error {
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
				case "displayName":
					person.Name = val
				case "ugentBarcode":
					person.AddIdentifier("ugent_barcode", val)
				case "uid":
					person.AddIdentifier("ugent_username", val)
				case "ugentExpirationDate":
					person.ExpirationDate = val
				case "objectClass":
					person.ObjectClass = append(person.ObjectClass, val)
				}
			}
		}
	}

	return person, nil
}
