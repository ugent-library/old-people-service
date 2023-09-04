package cli

import (
	"github.com/spf13/cobra"
	"github.com/ugent-library/people-service/models"
	"github.com/ugent-library/people-service/student"
)

var importStudentsCmd = &cobra.Command{
	Use:   "students",
	Short: "import student records from UGent LDAP",
	RunE: func(cmd *cobra.Command, args []string) error {
		ugentLdapClient := newUgentLdapClient()
		repo, err := newRepository()
		if err != nil {
			return err
		}

		importer := student.NewImporter(repo, ugentLdapClient)
		return importer.ImportAll(func(person *models.Person) {
			logger.Infof("successfully imported person %s", person.Id)
		})
	},
}

func init() {
	importCmd.AddCommand(importStudentsCmd)
}
