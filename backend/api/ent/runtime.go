// Code generated by ent, DO NOT EDIT.

package ent

import (
	"api/ent/schema"
	"api/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescSid is the schema descriptor for sid field.
	userDescSid := userFields[0].Descriptor()
	// user.DefaultSid holds the default value on creation for the sid field.
	user.DefaultSid = userDescSid.Default.(func() string)
}