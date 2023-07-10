package cmd

import (
	"context"

	"github.com/spf13/cobra"
)

var expirePersonCmd = &cobra.Command{
	Use:   "expire-person",
	Short: "auto expire person records",
	Run: func(cmd *cobra.Command, args []string) {
		nAffected, err := Repository().AutoExpirePeople(context.TODO())
		if err != nil {
			logger.Fatal(err)
		}
		logger.Infof("%d person records expired", nAffected)
	},
}

func init() {
	rootCmd.AddCommand(expirePersonCmd)
}
