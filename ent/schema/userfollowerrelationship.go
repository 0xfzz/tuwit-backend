package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserFollowerRelationship holds the schema definition for the UserFollowerRelationship entity.
type UserFollowerRelationship struct {
	ent.Schema
}

func (UserFollowerRelationship) Mixin() []ent.Mixin {
	return []ent.Mixin{TimeMixin{}}
}

// Fields of the UserFollowerRelationship.
func (UserFollowerRelationship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Int("follower_id").Optional(),
	}
}

// Edges of the UserFollowerRelationship.
func (UserFollowerRelationship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("follower", UserAccount.Type).Ref("followings").Unique().Field("follower_id"),
		edge.From("following", UserAccount.Type).Ref("followers").Unique().Field("user_id"),
	}
}
