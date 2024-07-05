package entHelper

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Normal field
func OrderField(name string) ent.Field {
	return field.String(name).Annotations(orderFieldAnnotation(name))
}

// Unique field
func UniqueField(name string) ent.Field {
	return field.String(name).Unique().Annotations(orderFieldAnnotation(name))
}

// Optional field
func OptionalField(name string) ent.Field {
	return field.String(name).Optional().Annotations(orderFieldAnnotation(name))
}

// UUID field
func UuidField(name string) ent.Field {
	return field.String(name).
		DefaultFunc(func() string {
			return uuid.New().String()
		}).Annotations(orderFieldAnnotation(name))
}

// 共通アノテーションをまとめる
var orderFieldAnnotation = func(name string) entgql.Annotation {
	return entgql.OrderField(name)
}
