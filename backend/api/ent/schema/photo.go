package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Photo holds the schema definition for the Photo entity.
type Photo struct {
	ent.Schema
}

// Fields of the Photo.
func (Photo) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("url"),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
