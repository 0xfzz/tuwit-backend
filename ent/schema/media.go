package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("path"),
		field.Enum("category").Values("IMAGE", "DOCUMENT", "OTHER"), // IMAGE, DOCUMENT
	}
}

// Edges of the Media.
func (Media) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("threads", Thread.Type).Ref("images"),
		edge.From("owner_profile_picture", UserProfile.Type).Ref("profile_picture"),
		edge.From("owner_banner", UserProfile.Type).Ref("banner"),
	}
}
