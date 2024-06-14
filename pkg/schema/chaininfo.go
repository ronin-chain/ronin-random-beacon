package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type ChainInfo struct {
	ent.Schema
}

func (ChainInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.Uint64("processed_block").Default(0),
	}
}
