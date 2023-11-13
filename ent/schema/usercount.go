package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserCount holds the schema definition for the UserCount entity.
type UserCount struct {
	ent.Schema
}

func (UserCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_account"},
	}
}

// Fields of the UserCount.
func (UserCount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("follower_count").Default(0),
		field.Int("followings_count").Default(0),
	}
}

// Edges of the UserCount.
func (UserCount) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", UserAccount.Type).Ref("user_count_info").Required(),
	}
}
