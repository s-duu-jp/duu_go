package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"

	entHelper "api/helpers"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		entHelper.UuidField("sid"),
		entHelper.UniqueField("uid"),
		entHelper.OrderField("name"),
		entHelper.UniqueField("email"),
		field.String("password").Optional().Sensitive(),
		entHelper.OrderField("role_type"),
		entHelper.OrderField("status_type"),
		entHelper.OrderField("oauth_type"),
		entHelper.OptionalField("sub"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("photos", Photo.Type),
		edge.From("organization", Organization.Type).
			Ref("users").
			Unique(),
	}
}

// Userのアノテーション
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
