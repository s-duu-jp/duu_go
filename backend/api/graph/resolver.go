package graph

import (
	"api/ent"
	mysql "api/services/db"
	"log"

	"github.com/99designs/gqlgen/graphql"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema() graphql.ExecutableSchema {
	client, _, err := mysql.GetDatabaseClient()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
	})
}
