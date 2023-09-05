package student

import (
	"context"
	"errors"
	"fmt"
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

func (si *Importer) ImportAll(cb func(*models.Person)) error {
	ctx := context.TODO()
	err := si.ugentLdapClient.SearchPeople("(objectClass=ugentStudent)", func(ldapEntry *ldap.Entry) error {
		newPerson, err := si.ldapEntryToPerson(ldapEntry)
		if err != nil {
			return err
		}

		var oldPerson *models.Person
		for _, otherId := range newPerson.OtherId {
			if otherId.Type != "historic_ugent_id" {
				continue
			}
			op, err := si.repository.GetPersonByOtherId(ctx, "historic_ugent_id", otherId.Id)
			if errors.Is(err, models.ErrNotFound) {
				continue
			} else if err != nil {
				return err
			}
			oldPerson = op
			break
		}

		if oldPerson == nil {
			np, err := si.repository.CreatePerson(ctx, newPerson)
			if err != nil {
				return fmt.Errorf("unable to create person: %w", err)
			}
			cb(np)
		} else {
			oldPerson.Active = true
			oldPerson.BirthDate = newPerson.BirthDate
			oldPerson.Email = newPerson.Email
			oldPerson.FirstName = newPerson.FirstName
			oldPerson.LastName = newPerson.LastName
			oldPerson.FullName = newPerson.FullName
			oldPerson.JobCategory = newPerson.JobCategory
			oldPerson.Title = newPerson.Title
			oldPerson.OtherId = newPerson.OtherId
			oldPerson.ObjectClass = newPerson.ObjectClass
			oldPerson.ExpirationDate = newPerson.ExpirationDate
			oldPerson.Organization = newPerson.Organization

			oldPerson, err := si.repository.UpdatePerson(ctx, oldPerson)
			if err != nil {
				return fmt.Errorf("unable to update person: %w", err)
			}
			cb(oldPerson)
		}

		return nil
	})

	return err
}

func (si *Importer) ldapEntryToPerson(ldapEntry *ldap.Entry) (*models.Person, error) {
	newPerson := models.NewPerson()
	newPerson.Active = true
	ctx := context.TODO()

	for _, attr := range ldapEntry.Attributes {
		for _, val := range attr.Values {
			switch attr.Name {
			case "uid":
				newPerson.OtherId = append(newPerson.OtherId, &models.IdRef{
					Type: "ugent_username",
					Id:   val,
				})
			// contains current active ugentID
			case "ugentID":
				newPerson.OtherId = append(newPerson.OtherId, &models.IdRef{
					Type: "ugent_id",
					Id:   val,
				})
			// contains ugentID also (at the end)
			case "ugentHistoricIDs":
				newPerson.OtherId = append(newPerson.OtherId, &models.IdRef{
					Type: "historic_ugent_id",
					Id:   val,
				})
			case "ugentBarcode":
				newPerson.OtherId = append(newPerson.OtherId, &models.IdRef{
					Type: "ugent_barcode",
					Id:   val,
				})
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
				realOrg, err := si.repository.GetOrganizationByOtherId(ctx, "ugent_id", val)
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