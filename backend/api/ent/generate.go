package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migration ./schema
//go:generate go run -mod=mod entc.go
//go:generate go run -mod=mod github.com/99designs/gqlgen generate
