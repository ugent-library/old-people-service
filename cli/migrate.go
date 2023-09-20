package cli

import (
	"context"
	"database/sql"
	"errors"

	"ariga.io/atlas/sql/migrate"
	"ariga.io/atlas/sql/postgres"
	"github.com/spf13/cobra"
)

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "migrate database",
	RunE: func(cmd *cobra.Command, args []string) error {

		db, err := sql.Open("pgx", config.Db.Url)
		if err != nil {
			return err
		}

		dir, err := migrate.NewLocalDir("migrations")
		if err != nil {
			return err
		}

		driver, err := postgres.Open(db)
		if err != nil {
			return err
		}

		executor, err := migrate.NewExecutor(
			driver,
			dir,
			migrate.NopRevisionReadWriter{},
			// why is this needed? "atlas migrate" runs fine without this!
			migrate.WithBaselineVersion("20230919120247"),
		)
		if err != nil {
			return err
		}

		if err = executor.ExecuteN(context.TODO(), 0); err != nil {
			if !errors.Is(err, migrate.ErrNoPendingFiles) {
				return err
			}
			logger.Infof("all is fine")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(migrateCmd)
}
