package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Thread holds the schema definition for the Thread entity.
type Thread struct {
	ent.Schema
}

// Fields of the Thread.
func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content"),
		field.Bool("is_comment_disabled").Default(false),
		field.Enum("visibility").Values("PRIVATE", "PUBLIC"),
		field.Enum("status").Values("PUBLISHED", "DRAFT"),
		field.Bool("is_deleted").Default(false),
	}
}

// Edges of the Thread.
func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child_threads", Thread.Type).From("parent_thread").Unique(),
		edge.To("images", Media.Type),
	}
}
