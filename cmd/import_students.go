package cmd

import (
	"context"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/ugent-library/person-service/models"
)

var importStudentsCmd = &cobra.Command{
	Use:   "students",
	Short: "import student records from UGent LDAP",
	Run: func(cmd *cobra.Command, args []string) {

		ctx := context.TODO()
		repo := Services().Repository

		//TODO: require organisation "Universiteit Gent"?
		org, err := repo.GetOrganizationByOtherId(ctx, "ugent_id", "UGent")
		if errors.Is(err, models.ErrNotFound) {
			logger.Fatal(errors.New("unable to find parent organization UGent"))
		}
		if err != nil {
			logger.Fatal(err)
		}

		err = LDAPClient().SearchPeople("(objectClass=ugentStudent)", func(np *models.Person) error {

			//TODO: assign to 'Universiteit Gent' automatically?
			//TODO: loop by cursor instead of fetching all data at once?
			//TODO: use ugentmodifytimestamp to fetch records incrementally?

			var op *models.Person
			for _, otherId := range np.OtherId {
				if otherId.Type != "historic_ugent_id" {
					continue
				}
				oldPerson, err := repo.GetPersonByOtherId(ctx, "historic_ugent_id", otherId.Id)
				if errors.Is(err, models.ErrNotFound) {
					continue
				} else if err != nil {
					return err
				}
				op = oldPerson
				break
			}

			if op == nil {
				np.Organization = append(np.Organization, models.NewOrganizationRef(org.Id))
				np, err := repo.CreatePerson(ctx, np)
				if err != nil {
					return fmt.Errorf("unable to create person: %w", err)
				}
				logger.Infof("successfully inserted person record %s", np.Id)
			} else {
				orgFound := false
				for _, orgRef := range op.Organization {
					if orgRef.Id == org.Id {
						orgFound = true
						break
					}
				}
				if !orgFound {
					op.Organization = append(op.Organization, models.NewOrganizationRef(org.Id))
				}
				op.Active = true
				op.BirthDate = np.BirthDate
				op.Email = np.Email
				op.FirstName = np.FirstName
				op.LastName = np.LastName
				op.FullName = np.FullName
				op.JobCategory = np.JobCategory
				op.Title = np.Title
				op.OtherId = np.OtherId // TODO: change this if non ldap identifiers are added
				op.ObjectClass = np.ObjectClass
				op.ExpirationDate = np.ExpirationDate

				op, err := repo.UpdatePerson(ctx, op)
				if err != nil {
					return fmt.Errorf("unable to update person: %w", err)
				}
				logger.Infof("success updated person record %s", op.Id)
			}

			return nil
		})

		if err != nil {
			logger.Fatal(err)
		}

		// TODO: make person records that where not found inactive

	},
}

func init() {
	importCmd.AddCommand(importStudentsCmd)
}
