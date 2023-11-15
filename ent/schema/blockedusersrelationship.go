package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BlockedUsersRelationship holds the schema definition for the BlockedUsersRelationship entity.
type BlockedUsersRelationship struct {
	ent.Schema
}

func (BlockedUsersRelationship) Mixin() []ent.Mixin {
	return []ent.Mixin{TimeMixin{}}
}

// Fields of the BlockedUsersRelationship.
func (BlockedUsersRelationship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Int("blocker_id").Optional(),
	}
}

// Edges of the BlockedUsersRelationship.
func (BlockedUsersRelationship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("blocker", UserAccount.Type).Ref("blocked_users").Unique().Field("blocker_id"),
		edge.From("blocked_user", UserAccount.Type).Ref("blocked_by").Unique().Field("user_id"),
	}
}
