package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
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
			}).Annotations(entgql.OrderField("SID")),
		field.String("uid").Unique().Annotations(entgql.OrderField("UID")),
		field.String("name").Annotations(entgql.OrderField("NAME")),
		field.String("email").Unique().Annotations(entgql.OrderField("EMAIL")),
		field.String("password").Optional().Sensitive(),
		field.String("role_type").Annotations(entgql.OrderField("ROLE_TYPE")),
		field.String("status_type").Annotations(entgql.OrderField("STATUS_TYPE")),
		field.String("oauth_type").Annotations(entgql.OrderField("OAUTH_TYPE")),
		field.String("sub").Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

// Userのアノテーション
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
