package db

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"strings"

	entdialect "entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
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

		if err := txClient.Client().Schema.Create(context.Background()); err != nil {
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
