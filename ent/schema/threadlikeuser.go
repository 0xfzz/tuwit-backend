package schema

import "entgo.io/ent"

// ThreadLikeUser holds the schema definition for the ThreadLikeUser entity.
type ThreadLikeUser struct {
	ent.Schema
}

// Fields of the ThreadLikeUser.
func (ThreadLikeUser) Fields() []ent.Field {
	return nil
}

// Edges of the ThreadLikeUser.
func (ThreadLikeUser) Edges() []ent.Edge {
	return nil
}
