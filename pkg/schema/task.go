package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("txHash").Optional(),
		field.String("data").NotEmpty(),
		field.Enum("status").Values("pending", "sent", "success", "failed"),
		field.String("lastError").Optional(),
		field.Bool("reOrg").Default(false),
		field.Int("attempts").Default(0),
		field.Int("lastSent").Optional(),
		field.Bool("isFinalized").Default(false),
		field.Int("sentBlock").Default(0),
		field.Int("createdAt").NonNegative(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("randomrequest", RandomRequest.Type),
	}
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("status"),
		index.Fields("isFinalized"),
	}
}
