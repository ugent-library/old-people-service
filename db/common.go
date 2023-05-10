package db

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ugent-library/people/ent"
)

var regexMultipleSpaces = regexp.MustCompile(`\s+`)
var regexNoBrackets = regexp.MustCompile(`[\[\]()\{\}]`)

func toTSQuery(column string, query string) *entsql.Predicate {

	// remove duplicate spaces
	query = regexMultipleSpaces.ReplaceAllString(query, " ")
	// trim
	query = strings.TrimSpace(query)

	queryParts := make([]string, 0)
	queryArgs := make([]any, 0)
	argCounter := 0

	for _, qp := range strings.Split(query, " ") {
		// remove terms that contain brackets
		if regexNoBrackets.MatchString(qp) {
			continue
		}
		argCounter++

		// $1 || ':*'
		queryParts = append(queryParts, fmt.Sprintf("$%d || ':*'", argCounter))
		queryArgs = append(queryArgs, qp)
	}

	// $1:* & $2:*
	tsQuery := fmt.Sprintf(
		"ts @@ to_tsquery('simple', %s)",
		strings.Join(queryParts, " || ' & ' || "),
	)

	return entsql.ExprP(
		tsQuery,
		queryArgs...)
}

func OpenClient(dsn string) (*ent.Client, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	driver := entsql.OpenDB(entdialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver)).Debug()

	/*
		run database migration in a transaction
		if two instances start, entgo might try
		to drop certain indexes more than once,
		depending on the number of instances,
		of which only one may succeed
	*/
	var migrationErr error
	func() {

		txClient, err := client.Tx(context.Background())
		if err != nil {
			migrationErr = err
			return
		}
		defer txClient.Rollback()

		/*
			TODO: advisory locks good choice for locking for migration?
			lock for ent migration
			otherwise two migrations might runs concurrently
			returning errors (like "table already created")
		*/
		txClient.Client().ExecContext(context.Background(), "SELECT pg_advisory_xact_lock(12345)")

		if err := txClient.Client().Schema.Create(context.Background(), schema.WithApplyHook(customSchemaChanges)); err != nil {
			migrationErr = err
			return
		}

		txClient.Commit()

	}()

	if migrationErr != nil {
		return nil, migrationErr
	}

	return client, nil
}

// cf. https://entgo.io/docs/migrate/#apply-hook-example
func customSchemaChanges(next schema.Applier) schema.Applier {
	return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {

		execQuery := `
		BEGIN;
		LOCK table organization IN EXCLUSIVE MODE;
		ALTER TABLE organization 
		ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS 
		(
			to_tsvector('simple', jsonb_path_query_array(other_id, '$[*].id')) || 
			to_tsvector('simple', public_id) || 
			to_tsvector('simple',name_dut) || 
			to_tsvector('simple', name_eng)
		) STORED;
		CREATE INDEX IF NOT EXISTS organization_ts_idx ON organization USING GIN(ts);
		LOCK table person IN EXCLUSIVE MODE;
		ALTER TABLE person 
			ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS (
				to_tsvector('simple',full_name)
			) STORED;
		CREATE INDEX IF NOT EXISTS person_ts_idx ON person USING GIN(ts);
		COMMIT;
		`
		return conn.Exec(context.Background(), execQuery, []any{}, nil)
	})
}
