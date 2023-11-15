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

func (Thread) Mixin() []ent.Mixin {
	return []ent.Mixin{TimeMixin{}}
}

// Fields of the Thread.
func (Thread) Fields() []ent.Field {
	return []ent.Field{
		field.Text("content"),
		field.Bool("is_comment_disabled").Default(false),
		field.Enum("visibility").Values("PRIVATE", "PUBLIC").Default("PUBLIC"),
		field.Enum("status").Values("PUBLISHED", "DRAFT").Default("PUBLISHED"),
		field.Bool("is_deleted").Default(false),
		field.Int("author_id").Optional(),
		field.Int("parent_id").Optional(),
		field.Int("repost_thread_id").Optional(),
	}
}

// Edges of the Thread.
func (Thread) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", UserAccount.Type).Ref("threads").Unique().Field("author_id"),
		edge.To("children", Thread.Type).From("parent").Unique().Field("parent_id"),
		edge.To("thread_count", ThreadCount.Type),
		edge.To("repost", Thread.Type).Unique().From("reposted").Unique().Field("repost_thread_id"),
		edge.To("images", Media.Type),
	}
}
