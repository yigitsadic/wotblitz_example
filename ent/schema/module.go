package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"time"
)

// Module holds the schema definition for the Module entity.
type Module struct {
	ent.Schema
}

// Fields of the Module.
func (Module) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("moduleType").NotEmpty(),
		field.Time("createdAt").Default(time.Now),
	}
}

// Edges of the Module.
func (Module) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tanks", Tank.Type).Ref("modules"),
	}
}
