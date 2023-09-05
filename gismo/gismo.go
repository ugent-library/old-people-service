package gismo

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/antchfx/xmlquery"
	"github.com/go-ldap/ldap/v3"
	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/ugentldap"
)

type Importer struct {
	repository      models.Repository
	ugentLdapClient *ugentldap.Client
}

/*
	ImportPerson and ImportOrganization return the following
	wrapped errors

	* `models.ErrFatal` when message cannot be processed at the time and needs to be resent.
	   e.g. database is unavailabe
	* `models.ErrSkipped` when message cannot be processed, but can be skipped.
	   Mostly for message that cannot be mapped (e.g. missing ugent_id)
*/

func NewImporter(repo models.Repository, ugentLdapClient *ugentldap.Client) *Importer {
	return &Importer{
		repository:      repo,
		ugentLdapClient: ugentLdapClient,
	}
}

func (gi *Importer) parseOrganizationMessage(buf []byte) (*models.Message, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	removeNamespace(doc)

	node := xmlquery.FindOne(doc, "//cfOrgUnit")

	if node == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfOrgUnit", ErrNonCompliantXml)
	}

	idNode := xmlquery.FindOne(node, "cfOrgUnitId")

	if idNode == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfOrgUnit/cfOrgUnitId", ErrNonCompliantXml)
	}

	orgNode := xmlquery.FindOne(doc, "//Body/organizations")

	if orgNode == nil {
		return nil, fmt.Errorf("%w: unabel to find node //Body/organizations", ErrNonCompliantXml)
	}

	msg := &models.Message{
		ID:   strings.TrimSpace(idNode.InnerText()),
		Date: orgNode.SelectAttr("date"),
	}

	if node.SelectAttr("action") == "DELETE" {
		msg.Source = "gismo.organization.delete"
	} else {
		msg.Source = "gismo.organization.update"

		for _, n := range xmlquery.Find(node, "./cfName") {
			t1, err := time.Parse(time.RFC3339, n.SelectAttr("cfStartDate"))
			if err != nil {
				return nil, err
			}
			t2, err := time.Parse(time.RFC3339, n.SelectAttr("cfEndDate"))
			if err != nil {
				return nil, err
			}
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "name_" + n.SelectAttr("cfLangCode"),
				Value:     strings.TrimSpace(n.InnerText()),
				StartDate: &t1,
				EndDate:   &t2,
			})
		}

		for _, v := range cerifValuesByClassName(doc, "cfOrgUnit_Class", "/be.ugent/organisatie/type/vakgroep", "") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "type",
				Value:     "department",
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		for _, v := range cerifValuesByClassName(doc, "cfOrgUnit_OrgUnit", "/be.ugent/gismo/organisatie-organisatie/type/kind-van", "cfOrgUnitId1") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "parent_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		// e.g. 000006047
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/memorialis", "cfFedId") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "ugent_memorialis_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		// e.g. "WE03V"
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/org-code", "cfFedId") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "code",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		// e.g. WE03*
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/biblio-code", "cfFedId") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "biblio_code",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}

	}

	return msg, nil
}

func (gi *Importer) ImportOrganization(buf []byte) (*models.Message, error) {
	msg, err := gi.parseOrganizationMessage(buf)
	if err != nil {
		return nil, err
	}

	ctx := context.TODO()

	// fetch existing organization by gismo_id
	org, err := gi.repository.GetOrganizationByGismoId(ctx, msg.ID)
	if err != nil && err == models.ErrNotFound {
		org = models.NewOrganization()
	} else if err != nil {
		return nil, fmt.Errorf("%w: unable to fetch organization record: %s", models.ErrFatal, err)
	}

	if msg.Source == "gismo.organization.update" {
		now := time.Now()
		org.NameDut = ""
		org.NameEng = ""
		org.OtherId = make([]*models.IdRef, 0)
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
					orgParentByGismo, err := gi.repository.GetOrganizationByGismoId(ctx, attr.Value)
					if errors.Is(err, models.ErrNotFound) {
						return nil, fmt.Errorf("%w: unable to find parent organization with gismo id %s", models.ErrFatal, attr.Value)
					} else if err != nil {
						return nil, fmt.Errorf("%w", models.ErrFatal)
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
				org.OtherId = append(org.OtherId, &models.IdRef{
					Type: "ugent_memorialis_id",
					Id:   attr.Value,
				})
			case "code":
				org.OtherId = append(org.OtherId, &models.IdRef{
					Type: "ugent_id",
					Id:   attr.Value,
				})
			case "biblio_code":
				org.OtherId = append(org.OtherId, &models.IdRef{
					Type: "biblio_id",
					Id:   attr.Value,
				})
			}
		}

		if _, err := gi.repository.SaveOrganization(ctx, org); err != nil {
			return nil, fmt.Errorf("%w: unable to save organization record: %s", models.ErrFatal, err)
		}
	} else if msg.Source == "gismo.organization.delete" {
		if org.IsStored() {
			if err := gi.repository.DeleteOrganization(ctx, org.Id); err != nil {
				return nil, fmt.Errorf("%w: unable to delete organization record: %s", models.ErrFatal, err)
			}
		}
	}
	return msg, nil
}

func (gi *Importer) parsePersonMessage(buf []byte) (*models.Message, error) {
	doc, err := xmlquery.Parse(bytes.NewReader(buf))
	if err != nil {
		return nil, err
	}

	removeNamespace(doc)

	node := xmlquery.FindOne(doc, "//cfPers")

	if node == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfPers", ErrNonCompliantXml)
	}

	idNode := xmlquery.FindOne(node, "cfPersId")

	if idNode == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfPers/cfPersId", ErrNonCompliantXml)
	}

	langNode := xmlquery.FindOne(node, "cfPers_Lang/cfLangCode")

	if langNode == nil {
		return nil, fmt.Errorf("%w: unable to find xml node //cfPers/cfPers_Lang/cfLangCode", ErrNonCompliantXml)
	}

	persNode := xmlquery.FindOne(doc, "//Body/persons")

	if persNode == nil {
		return nil, fmt.Errorf("%w: unabel to find node //Body/persons", ErrNonCompliantXml)
	}

	msg := &models.Message{
		ID:       strings.TrimSpace(idNode.InnerText()),
		Language: strings.TrimSpace(langNode.InnerText()),
		Date:     persNode.SelectAttr("date"),
	}

	if node.SelectAttr("action") == "DELETE" {
		msg.Source = "gismo.person.delete"
	} else {
		msg.Source = "gismo.person.update"

		for _, nameNode := range cerifNodesByClassName(doc, "cfPersName_Pers", "/be.ugent/gismo/persoon/persoonsnaam/type/officiele-naam") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			if n := xmlquery.FindOne(nameNode, "cfFamilyNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "last_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "cfFirstNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "first_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "ugTitles"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "title",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
		}
		for _, nameNode := range cerifNodesByClassName(doc, "cfPersName_Pers", "/be.ugent/gismo/persoon/persoonsnaam/type/voorkeursweergave") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(nameNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			if n := xmlquery.FindOne(nameNode, "cfFamilyNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "preferred_last_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
			if n := xmlquery.FindOne(nameNode, "cfFirstNames"); n != nil {
				msg.Attributes = append(msg.Attributes, models.Attribute{
					Name:      "preferred_first_name",
					Value:     strings.TrimSpace(n.InnerText()),
					StartDate: &startDate,
					EndDate:   &endDate,
				})
			}
		}

		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/persoon/federated-id/ugent-id", "cfFedId") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "ugent_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/persoon/federated-id/orcid", "cfFedId") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "orcid",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}
		for _, v := range cerifValuesByClassName(doc, "cfFedId", "/be.ugent/gismo/organisatie/federated-id/memorialis", "cfFedId") {
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "ugent_memorialis_id",
				Value:     v.Value,
				StartDate: &v.StartDate,
				EndDate:   &v.EndDate,
			})
		}

		for _, persAddrNode := range xmlquery.Find(node, "cfPers_EAddr") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			addrID := strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfEAddrId").InnerText())
			addrNode := xmlquery.FindOne(doc, "//cfEAddr[contains(cfEAddrId, '"+addrID+"')]")
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "email",
				Value:     strings.TrimSpace(xmlquery.FindOne(addrNode, "cfURI").InnerText()),
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}

		for _, persAddrNode := range xmlquery.Find(node, "cfPers_OrgUnit") {
			startDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfStartDate").InnerText()))
			if err != nil {
				return nil, err
			}
			endDate, err := time.Parse(time.RFC3339, strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfEndDate").InnerText()))
			if err != nil {
				return nil, err
			}
			msg.Attributes = append(msg.Attributes, models.Attribute{
				Name:      "organization_id",
				Value:     strings.TrimSpace(xmlquery.FindOne(persAddrNode, "cfOrgUnitId").InnerText()),
				StartDate: &startDate,
				EndDate:   &endDate,
			})
		}
	}

	return msg, nil
}

func (gi *Importer) ImportPerson(buf []byte) (*models.Message, error) {
	msg, err := gi.parsePersonMessage(buf)
	if err != nil {
		return nil, err
	}

	// retrieve old person by matching on attributes in gismo message
	// returns new person when no match is found
	person, err := gi.getPersonByMessage(msg)
	if err != nil {
		return nil, err
	}

	// enrich person with incoming gismo attributes
	person, err = gi.enrichPersonWithMessage(person, msg)
	if err != nil {
		return nil, err
	}

	// enrich person with ugent ldap attributes
	person, err = gi.enrichPersonWithLdap(person)
	if err != nil {
		return nil, err
	}

	// create/update record
	if _, err = gi.repository.SavePerson(context.TODO(), person); err != nil {
		return nil, fmt.Errorf("%w: unable to save person record: %s", models.ErrFatal, err)
	}

	return msg, nil
}

func (gi *Importer) enrichPersonWithMessage(person *models.Person, msg *models.Message) (*models.Person, error) {
	now := time.Now()
	ctx := context.TODO()

	// clear old values
	person.GismoId = msg.ID
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
	var gismoOrganizationRefs []*models.OrganizationRef

	// add attributes from GISMO
	for _, attr := range msg.Attributes {
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
			person.OtherId = append(person.OtherId, &models.IdRef{
				Type: "ugent_id",
				Id:   attr.Value,
			})
			person.OtherId = append(person.OtherId, &models.IdRef{
				Type: "historic_ugent_id",
				Id:   attr.Value,
			})
		case "ugent_memorialis_id":
			person.OtherId = append(person.OtherId, &models.IdRef{
				Type: "ugent_memorialis_id",
				Id:   attr.Value,
			})
		case "title":
			if withinDateRange {
				person.Title = attr.Value
			}
		case "organization_id":
			orgRef := models.NewOrganizationRef(attr.Value)
			orgRef.From = attr.StartDate
			orgRef.Until = attr.EndDate
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
		orgsByGismo, err := gi.repository.GetOrganizationsByGismoId(ctx, gismoOrganizationIds...)
		if err != nil {
			return nil, err
		}
		// return fatal error when person arrives with organization that we do not know yet
		if len(orgsByGismo) != len(gismoOrganizationIds) {
			return nil, fmt.Errorf("%w: person.organization_id contains invalid gismo organization identifiers", models.ErrFatal)
		}

		var orgRefs []*models.OrganizationRef
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

	return person, nil
}

func (gi *Importer) getPersonByMessage(msg *models.Message) (*models.Person, error) {
	ctx := context.TODO()
	now := time.Now()

	// Without ugentId no linking possible
	ugentIds := msg.GetAttributesAt("ugent_id", now)
	if len(ugentIds) == 0 {
		return nil, fmt.Errorf("%w: missing ugent_id in message %s", models.ErrSkipped, msg.ID)
	}

	// trial 1: retrieve old person by gismo-id
	person, err := gi.repository.GetPersonByGismoId(ctx, msg.ID)
	if err == nil {
		return person, nil
	}
	if !errors.Is(err, models.ErrNotFound) {
		return nil, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
	}

	// trial 2: retrieve old person by ugent-id
	for _, ugentId := range ugentIds {
		person, err = gi.repository.GetPersonByOtherId(ctx, "historic_ugent_id", ugentId)
		if errors.Is(err, models.ErrNotFound) {
			continue
		}
		if err != nil {
			return nil, fmt.Errorf("%w: unable to fetch person record: %s", models.ErrFatal, err)
		}
		return person, nil
	}

	// trial 3: new person record
	return models.NewPerson(), nil
}

func (gi *Importer) enrichPersonWithLdap(person *models.Person) (*models.Person, error) {
	ugentIds := []string{}
	for _, otherId := range person.OtherId {
		if otherId.Type == "historic_ugent_id" {
			ugentIds = append(ugentIds, otherId.Id)
		}
	}

	ldapQueryParts := make([]string, 0, len(ugentIds))
	for _, ugentId := range ugentIds {
		ldapQueryParts = append(ldapQueryParts, fmt.Sprintf("(ugentHistoricIDs=%s)", ugentId))
	}
	ldapQuery := "(&" + strings.Join(ldapQueryParts, "") + ")"
	ldapEntries := make([]*ldap.Entry, 0)
	err := gi.ugentLdapClient.SearchPeople(ldapQuery, func(ldapEntry *ldap.Entry) error {
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
					person.FullName = val
				case "ugentBarcode":
					person.OtherId = append(person.OtherId, &models.IdRef{
						Type: "ugent_barcode",
						Id:   val,
					})
				case "uid":
					person.OtherId = append(person.OtherId, &models.IdRef{
						Type: "ugent_username",
						Id:   val,
					})
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
