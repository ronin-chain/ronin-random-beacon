package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// RandomRequest holds the schema definition for the RandomRequest entity.
type RandomRequest struct {
	ent.Schema
}

// Fields of the RandomRequest.
func (RandomRequest) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").NotEmpty().Unique(),
		field.Uint64("blockNumber"),
		field.Uint("logIndex"),
		field.String("raw").NotEmpty(),
		field.Uint64("period").Immutable(),
	}
}

// Edges of the RandomRequest.
func (RandomRequest) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("randomrequest").Unique()}
}

func (RandomRequest) Index() []ent.Index {
	return []ent.Index{
		index.Fields("blockNumber", "logIndex").Unique(),
	}
}
