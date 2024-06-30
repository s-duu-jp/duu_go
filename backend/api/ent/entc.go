//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	entGqlEx, err := entgql.NewExtension(
		entgql.WithWhereInputs(true),
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("../graph/schema/query.graphqls"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	if err := entc.Generate("./schema", &gen.Config{Features: []gen.Feature{gen.FeatureVersionedMigration}}, entc.Extensions(entGqlEx)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
