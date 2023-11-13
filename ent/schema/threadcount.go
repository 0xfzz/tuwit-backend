package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

func (ThreadCount) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "thread_account"},
	}
}

// ThreadCount holds the schema definition for the ThreadCount entity.
type ThreadCount struct {
	ent.Schema
}

// Fields of the ThreadCount.
func (ThreadCount) Fields() []ent.Field {
	return []ent.Field{
		field.Int("reply_count").Default(0),
		field.Int("like_count").Default(0),
	}
}

// Edges of the ThreadCount.
func (ThreadCount) Edges() []ent.Edge {
	return nil
}
