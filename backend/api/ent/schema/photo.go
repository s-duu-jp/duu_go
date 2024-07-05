package schema

import (
	entHelper "api/helpers"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
)

// Photo holds the schema definition for the Photo entity.
type Photo struct {
	ent.Schema
}

// Fields of the Photo.
func (Photo) Fields() []ent.Field {
	return []ent.Field{
		entHelper.OrderField("name"),
		entHelper.OrderField("url"),
	}
}

// Edges of the Photo.
func (Photo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("photos").
			Unique(),
	}
}

// アノテーション
func (Photo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
