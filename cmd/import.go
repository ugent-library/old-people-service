package cmd

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"
	"github.com/ugent-library/people/models"
)

var importCmd = &cobra.Command{
	Use: "import",
	Run: func(cmd *cobra.Command, args []string) {

		ps, err := models.NewPersonService(&models.PersonConfig{
			DB: config.Db.Url,
		})

		if err != nil {
			logger.Fatal(err)
		}

		decoder := json.NewDecoder(os.Stdin)

		// TODO: use transaction?
		for {
			person := &models.Person{}
			err := decoder.Decode(person)
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				logger.Fatal(err)
			}
			person, err = ps.Create(context.Background(), person)
			if err != nil {
				logger.Fatal(err)
			}
			fmt.Fprintf(os.Stderr, "added person %s\n", person.ID)
		}
	},
}

func init() {
	rootCmd.AddCommand(importCmd)
}
