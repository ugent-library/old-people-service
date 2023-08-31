package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	v1 "github.com/ugent-library/people-service/api/v1"
	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/repository"
	"github.com/ugent-library/people-service/ugentldap"
)

var importStudentsCmd = &cobra.Command{
	Use:   "students",
	Short: "import student records from UGent LDAP",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.TODO()
		ldapClient := ugentldap.NewClient(ugentldap.Config{
			Url:      config.Ldap.Url,
			Username: config.Ldap.Username,
			Password: config.Ldap.Password,
		})
		repo, err := repository.NewRepository(&repository.Config{
			DbUrl:  config.Db.Url,
			AesKey: config.Db.AesKey,
		})
		if err != nil {
			logger.Fatal(err)
		}

		err = ldapClient.SearchPeople("(objectClass=ugentStudent)", func(np *models.Person) error {

			/*
				np = "dummy" person record as returned by SearchPeople

				Notes:

				* np.Id is empty
				* np.Active is always true
				* np.Organization contains "dummy" *v1.OrganizationRef where Id is an ugent identifier (e.g CA20).
				  We try to match the ugent identifier against a stored organization.
				  Every provided v1.OrganizationRef requires a match.
				  Make sure the organizations are already stored.


			*/

			//TODO: use modifytimestamp to fetch records incrementally?

			var oldPerson *models.Person
			for _, otherId := range np.OtherId {
				if otherId.Type != "historic_ugent_id" {
					continue
				}

				oldPerson, err = repo.GetPersonByOtherId(ctx, "historic_ugent_id", otherId.Id)
				if errors.Is(err, models.ErrNotFound) {
					continue
				} else if err != nil {
					return err
				}

				logger.Infof("found existing person by matching historic_ugent_id %s", otherId.Id)
				break
			}

			// provided OrganizationRef#Id is an ugent_id. Match with stored organization
			var newOrgRefs []*v1.OrganizationRef
			for _, oRef := range np.Organization {
				realOrg, err := repo.GetOrganizationByOtherId(ctx, "ugent_id", oRef.Id)
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
			np.Organization = newOrgRefs

			if oldPerson == nil {
				np, err := repo.CreatePerson(ctx, np)
				if err != nil {
					return fmt.Errorf("unable to create person: %w", err)
				}
				logger.Infof("successfully inserted person record %s", np.Id)
			} else {
				oldPerson.Active = true
				oldPerson.BirthDate = np.BirthDate
				oldPerson.Email = np.Email
				oldPerson.FirstName = np.FirstName
				oldPerson.LastName = np.LastName
				oldPerson.FullName = np.FullName
				oldPerson.JobCategory = np.JobCategory
				oldPerson.Title = np.Title
				oldPerson.OtherId = np.OtherId
				oldPerson.ObjectClass = np.ObjectClass
				oldPerson.ExpirationDate = np.ExpirationDate
				oldPerson.Organization = np.Organization

				oldPerson, err = repo.UpdatePerson(ctx, oldPerson)
				if err != nil {
					return fmt.Errorf("unable to update person: %w", err)
				}
				logger.Infof("successfully updated person record %s", oldPerson.Id)
			}

			return nil
		})

		if err != nil {
			logger.Fatal(err)
		}

	},
}

func init() {
	importCmd.AddCommand(importStudentsCmd)
}
