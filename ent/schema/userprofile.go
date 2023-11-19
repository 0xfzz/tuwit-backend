package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UserProfile holds the schema definition for the UserProfile entity.
type UserProfile struct {
	ent.Schema
}

func (UserProfile) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "user_profile"},
	}
}

// Fields of the UserProfile.
func (UserProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name"),
		field.Text("bio").Optional(),
		field.Int("profile_picture_id").Optional(),
		field.Int("banner_id").Optional(),
	}
}

// Edges of the UserProfile.
func (UserProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("account", UserAccount.Type).Ref("profile").Unique().Required(),
		edge.To("profile_picture", Media.Type).Unique().Field("profile_picture_id"),
		edge.To("banner", Media.Type).Unique().Field("banner_id"),
	}
}
