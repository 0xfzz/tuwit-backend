package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserAccount holds the schema definition for the UserAccount entity.
type UserAccount struct {
	ent.Schema
}

func (UserAccount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_account"},
	}
}

// Fields of the UserAccount.
func (UserAccount) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").Unique(),
		field.String("username").Unique(),
		field.Bytes("password"),
		field.Bool("is_verified").Default(false),
		field.Bool("is_private").Default(false),
		field.Bool("is_email_verified").Default(false),
	}
}

// Edges of the UserAccount.
func (UserAccount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("profile", UserProfile.Type).Unique(),
		edge.To("followers", UserFollowerRelationship.Type),
		edge.To("followings", UserFollowerRelationship.Type),
		edge.To("blocked_by", BlockedUsersRelationship.Type),
		edge.To("blocked_users", BlockedUsersRelationship.Type),
		edge.To("user_count", UserCount.Type).Unique(),
		edge.To("threads", Thread.Type),
	}
}
