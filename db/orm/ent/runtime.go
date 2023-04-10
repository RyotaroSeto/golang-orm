// Code generated by ent, DO NOT EDIT.

package ent

import (
	"golang-orm/db/orm/ent/schema"
	"golang-orm/db/orm/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescUsername is the schema descriptor for username field.
	userDescUsername := userFields[0].Descriptor()
	// user.DefaultUsername holds the default value on creation for the username field.
	user.DefaultUsername = userDescUsername.Default.(string)
	// userDescFullName is the schema descriptor for full_name field.
	userDescFullName := userFields[1].Descriptor()
	// user.DefaultFullName holds the default value on creation for the full_name field.
	user.DefaultFullName = userDescFullName.Default.(string)
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[2].Descriptor()
	// user.DefaultEmail holds the default value on creation for the email field.
	user.DefaultEmail = userDescEmail.Default.(string)
}
