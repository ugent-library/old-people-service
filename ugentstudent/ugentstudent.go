package ugentstudent

import (
	"context"
	"errors"
	"fmt"

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
	err := si.ugentLdapClient.SearchPeople("(objectClass=ugentStudent)", func(newPerson *models.Person) error {

		/*
			"newPerson" is a "dummy" person record as returned by SearchPeople

			Notes:

			* newPerson.Id is empty
			* newPerson.Active is always true
			* newPerson.Organization contains "dummy" *models.OrganizationRef where Id is an ugent identifier (e.g CA20).
			  We try to match the ugent identifier against a stored organization.
			  Every provided models.OrganizationRef requires a match.
			  Make sure the organizations are already stored.
		*/

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

		// provided OrganizationRef#Id is an ugent_id. Match with stored organization
		var newOrgRefs []*models.OrganizationRef
		for _, oRef := range newPerson.Organization {
			realOrg, err := si.repository.GetOrganizationByOtherId(ctx, "ugent_id", oRef.Id)
			if errors.Is(err, models.ErrNotFound) {
				continue
			} else if err != nil {
				return err
			}
			newOrgRef := models.NewOrganizationRef(realOrg.Id)
			newOrgRef.From = oRef.From
			newOrgRef.Until = oRef.Until
			newOrgRefs = append(newOrgRefs, newOrgRef)
		}
		newPerson.Organization = newOrgRefs

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
