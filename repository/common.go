package repository

import (
	"context"
	"database/sql"
	"encoding/base64"
	"fmt"
	"regexp"
	"strings"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/ugent-library/crypt"
	"github.com/ugent-library/person-service/ent"
)

var regexMultipleSpaces = regexp.MustCompile(`\s+`)
var regexNoBrackets = regexp.MustCompile(`[\[\]()\{\}]`)

func toTSQuery(query string) (string, []any) {
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
		"to_tsquery('usimple', %s)",
		strings.Join(queryParts, " || ' & ' || "),
	)

	return tsQuery, queryArgs
}

func OpenClient(dsn string) (*ent.Client, error) {

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	driver := entsql.OpenDB(dialect.Postgres, db)
	client := ent.NewClient(ent.Driver(driver))

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

		if err := txClient.Client().Schema.Create(
			context.Background(),
			schema.WithApplyHook(customSchemaChanges),
		); err != nil {
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

		CREATE EXTENSION IF NOT EXISTS unaccent;
		DO
		$$BEGIN
			CREATE TEXT SEARCH CONFIGURATION usimple ( COPY = simple );
		EXCEPTION
			WHEN unique_violation THEN
			   NULL;
		END;$$;
		ALTER TEXT SEARCH CONFIGURATION usimple ALTER MAPPING FOR hword, hword_part, word WITH unaccent, simple;
		
		ALTER TABLE organization
		ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS
		(
			to_tsvector('simple', jsonb_path_query_array(other_id, '$[*].id')) ||
			to_tsvector('simple', public_id) ||
			to_tsvector('usimple',name_dut) ||
			to_tsvector('usimple', name_eng)
		) STORED;
		CREATE INDEX IF NOT EXISTS organization_ts_idx ON organization USING GIN(ts);
		LOCK table person IN EXCLUSIVE MODE;
		ALTER TABLE person
			ADD COLUMN IF NOT EXISTS ts tsvector GENERATED ALWAYS AS (
				to_tsvector('usimple',full_name)
			) STORED;
		CREATE INDEX IF NOT EXISTS person_ts_idx ON person USING GIN(ts);
		COMMIT;
		`

		plan.Changes = append(plan.Changes, &migrate.Change{
			Cmd: execQuery,
		})

		return next.Apply(ctx, conn, plan)

	})
}

func encryptMessage(key []byte, message string) (string, error) {
	cryptedMsgInBytes, err := crypt.Encrypt(key, []byte(message))
	if err != nil {
		return "", err
	}

	return base64.URLEncoding.EncodeToString(cryptedMsgInBytes), nil
}

func decryptMessage(key []byte, cryptedMsg string) (string, error) {
	cryptedMsgInBytes, err := base64.URLEncoding.DecodeString(cryptedMsg)
	if err != nil {
		return "", err
	}

	msgInBytes, err := crypt.Decrypt(key, cryptedMsgInBytes)
	if err != nil {
		return "", err
	}

	return string(msgInBytes), nil
}
