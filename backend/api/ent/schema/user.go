package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("sid").
			DefaultFunc(func() string {
				return uuid.New().String()
			}),
		field.String("uid").Unique(),
		field.String("name"),
		field.String("email").Unique(),
		field.String("password").Optional().Sensitive(),
		field.String("role_type"),
		field.String("status_type"),
		field.String("oauth_type"),
		field.String("sub").Optional(),
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
