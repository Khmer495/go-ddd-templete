package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type UlidMixin struct {
	mixin.Schema
}

func (UlidMixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("ulid").
			Unique().
			Immutable(),
	}
}
