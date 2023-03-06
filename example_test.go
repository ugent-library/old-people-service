package main

import (
	"context"
	"log"
	"testing"

	"github.com/ugent-library/people/ent"

	"entgo.io/ent/dialect"
	_ "github.com/lib/pq"
)

func TestPeople(t *testing.T) {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.Postgres, "host=localhost port=5432 user=biblio dbname=authority password=biblio sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:
}
