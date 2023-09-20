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

		oldPerson, err := si.repository.GetPersonByAnyOtherId(ctx, "historic_ugent_id", newPerson.OtherId["historic_ugent_id"]...)
		if err != nil && !errors.Is(err, models.ErrNotFound) {
			return err
		}

		if oldPerson == nil {
			if err := cb(newPerson); err != nil {
				return err
			}
		} else {
			oldPerson.Active = true
			oldPerson.BirthDate = newPerson.BirthDate
			oldPerson.Email = newPerson.Email
			oldPerson.FirstName = newPerson.FirstName
			oldPerson.LastName = newPerson.LastName
			oldPerson.FullName = newPerson.FullName
			oldPerson.JobCategory = newPerson.JobCategory
			oldPerson.Title = newPerson.Title
			oldPerson.OtherId = newPerson.OtherId.Dup()
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

// ldapEntryToPerson maps ldap entry to new Person, and creates Organization's if necessary (side effect)
func (si *Importer) ldapEntryToPerson(ldapEntry *ldap.Entry) (*models.Person, error) {
	newPerson := models.NewPerson()
	newPerson.Active = true
	ctx := context.TODO()

	for _, attr := range ldapEntry.Attributes {
		for _, val := range attr.Values {
			switch attr.Name {
			case "uid":
				newPerson.OtherId.Add("ugent_username", val)
			// contains current active ugentID
			case "ugentID":
				newPerson.OtherId.Add("ugent_id", val)
			// contains ugentID also (at the end)
			case "ugentHistoricIDs":
				newPerson.OtherId.Add("historic_ugent_id", val)
			case "ugentBarcode":
				newPerson.OtherId.Add("ugent_barcode", val)
			case "ugentPreferredGivenName":
				newPerson.FirstName = val
			case "ugentPreferredSn":
				newPerson.LastName = val
			case "displayName":
				newPerson.FullName = val
			case "ugentBirthDate":
				newPerson.BirthDate = val
			case "mail":
				newPerson.Email = strings.ToLower(val)
			case "ugentJobCategory":
				newPerson.JobCategory = append(newPerson.JobCategory, val)
			case "ugentAddressingTitle":
				newPerson.Title = val
			case "objectClass":
				newPerson.ObjectClass = append(newPerson.ObjectClass, val)
			case "ugentExpirationDate":
				newPerson.ExpirationDate = val
			case "departmentNumber":
				realOrg, err := si.repository.GetOrganizationByAnyOtherId(ctx, "ugent_id", val)
				if errors.Is(err, models.ErrNotFound) {
					continue
				} else if err != nil {
					return nil, err
				}
				newOrgRef := models.NewOrganizationRef(realOrg.Id)
				newPerson.Organization = append(newPerson.Organization, newOrgRef)
			}
		}
	}

	return newPerson, nil
}
