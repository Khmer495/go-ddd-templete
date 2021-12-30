package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Team holds the schema definition for the Team entity.
type Team struct {
	ent.Schema
}

func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UlidMixin{},
		TimeMixin{},
	}
}

// Fields of the Team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.Int("create_user_id"),
		field.String("name"),
		field.String("description"),
	}
}

func (Team) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name"),
	}
}

// Edges of the Team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("create_user", User.Type).
			Unique().
			Required().
			Field("create_user_id"),
		edge.From("team_users", TeamUser.Type).
			Ref("team"),
	}
}
