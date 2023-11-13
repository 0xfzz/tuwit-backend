// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/0xfzz/tuwitt/ent/schema"
	"github.com/0xfzz/tuwitt/ent/thread"
	"github.com/0xfzz/tuwitt/ent/useraccount"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	threadFields := schema.Thread{}.Fields()
	_ = threadFields
	// threadDescIsCommentDisabled is the schema descriptor for is_comment_disabled field.
	threadDescIsCommentDisabled := threadFields[1].Descriptor()
	// thread.DefaultIsCommentDisabled holds the default value on creation for the is_comment_disabled field.
	thread.DefaultIsCommentDisabled = threadDescIsCommentDisabled.Default.(bool)
	// threadDescIsDeleted is the schema descriptor for is_deleted field.
	threadDescIsDeleted := threadFields[4].Descriptor()
	// thread.DefaultIsDeleted holds the default value on creation for the is_deleted field.
	thread.DefaultIsDeleted = threadDescIsDeleted.Default.(bool)
	useraccountFields := schema.UserAccount{}.Fields()
	_ = useraccountFields
	// useraccountDescIsVerified is the schema descriptor for is_verified field.
	useraccountDescIsVerified := useraccountFields[3].Descriptor()
	// useraccount.DefaultIsVerified holds the default value on creation for the is_verified field.
	useraccount.DefaultIsVerified = useraccountDescIsVerified.Default.(bool)
	// useraccountDescIsPrivate is the schema descriptor for is_private field.
	useraccountDescIsPrivate := useraccountFields[4].Descriptor()
	// useraccount.DefaultIsPrivate holds the default value on creation for the is_private field.
	useraccount.DefaultIsPrivate = useraccountDescIsPrivate.Default.(bool)
	// useraccountDescIsEmailVerified is the schema descriptor for is_email_verified field.
	useraccountDescIsEmailVerified := useraccountFields[5].Descriptor()
	// useraccount.DefaultIsEmailVerified holds the default value on creation for the is_email_verified field.
	useraccount.DefaultIsEmailVerified = useraccountDescIsEmailVerified.Default.(bool)
}