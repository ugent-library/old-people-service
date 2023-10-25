package student

import (
	"context"
	"errors"
	"strings"

	"github.com/go-ldap/ldap/v3"
	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/ugentldap"
)

type Importer struct {
	repository      models.Repository
	ugentLdapClient *ugentldap.Client
}

func NewImporter(repo models.Repository, ugentLdapClient *ugentldap.Client) *Importer {
	return &Importer{
		repository:      repo,
		ugentLdapClient: ugentLdapClient,
	}
}

// Each calls callback function with valid models.Person to save
func (si *Importer) Each(cb func(*models.Person) error) error {
	ctx := context.TODO()
	err := si.ugentLdapClient.SearchPeople("(objectClass=ugentStudent)", func(ldapEntry *ldap.Entry) error {
		newPerson, err := si.ldapEntryToPerson(ldapEntry)
		if err != nil {
			return err
		}

		oldPerson, err := si.repository.GetPersonByIdentifier(ctx, "historic_ugent_id", newPerson.GetIdentifierValues("historic_ugent_id")...)
		if err != nil && !errors.Is(err, models.ErrNotFound) {
			return err
		}

		if oldPerson == nil {
			if err := cb(newPerson); err != nil {
				return err
			}
		} else {
			var gismoId string
			var orcid string
			for _, id := range oldPerson.Identifier {
				switch id.PropertyID {
				case "orcid":
					orcid = id.Value
				case "gismo_id":
					gismoId = id.Value
				}
			}
			oldPerson.ClearIdentifier()
			for _, id := range newPerson.Identifier {
				oldPerson.AddIdentifier(id.PropertyID, id.Value)
			}
			if gismoId != "" {
				oldPerson.AddIdentifier("gismo_id", gismoId)
			}
			if orcid != "" {
				oldPerson.AddIdentifier("orcid", orcid)
			}
			oldPerson.Active = true
			oldPerson.BirthDate = newPerson.BirthDate
			oldPerson.Email = newPerson.Email
			oldPerson.GivenName = newPerson.GivenName
			oldPerson.FamilyName = newPerson.FamilyName
			oldPerson.Name = newPerson.Name
			oldPerson.JobCategory = newPerson.JobCategory
			oldPerson.HonorificPrefix = newPerson.HonorificPrefix
			oldPerson.ObjectClass = newPerson.ObjectClass
			oldPerson.ExpirationDate = newPerson.ExpirationDate
			oldPerson.Organization = newPerson.Organization

			if err := cb(oldPerson); err != nil {
				return err
			}
		}

		return nil
	})

	return err
}

// ldapEntryToPerson maps ldap entry to new Person
func (si *Importer) ldapEntryToPerson(ldapEntry *ldap.Entry) (*models.Person, error) {
	newPerson := models.NewPerson()
	newPerson.Active = true
	ctx := context.TODO()

	for _, attr := range ldapEntry.Attributes {
		for _, val := range attr.Values {
			switch attr.Name {
			case "uid":
				newPerson.AddIdentifier("ugent_username", val)
			case "ugentID":
				newPerson.AddIdentifier("ugent_id", val)
			case "ugentHistoricIDs":
				newPerson.AddIdentifier("historic_ugent_id", val)
			case "ugentBarcode":
				newPerson.AddIdentifier("ugent_barcode", val)
			case "ugentPreferredGivenName":
				newPerson.GivenName = val
			case "ugentPreferredSn":
				newPerson.FamilyName = val
			case "displayName":
				newPerson.Name = val
			case "ugentBirthDate":
				newPerson.BirthDate = val
			case "mail":
				newPerson.Email = strings.ToLower(val)
			case "ugentJobCategory":
				newPerson.JobCategory = append(newPerson.JobCategory, val)
			case "ugentAddressingTitle":
				newPerson.HonorificPrefix = val
			case "objectClass":
				newPerson.ObjectClass = append(newPerson.ObjectClass, val)
			case "ugentExpirationDate":
				newPerson.ExpirationDate = val
			case "departmentNumber":
				realOrg, err := si.repository.GetOrganizationByIdentifier(ctx, "ugent_id", val)
				// ignore for now. Maybe tomorrow on the next run
				if errors.Is(err, models.ErrNotFound) {
					continue
				} else if err != nil {
					return nil, err
				}
				newOrgMember := models.NewOrganizationMember(realOrg.ID)
				newPerson.Organization = append(newPerson.Organization, newOrgMember)
			}
		}
	}

	return newPerson, nil
}
