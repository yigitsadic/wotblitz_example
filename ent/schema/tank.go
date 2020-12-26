package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Tank holds the schema definition for the Tank entity.
type Tank struct {
	ent.Schema
}

// Fields of the Tank.
func (Tank) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.Int("tier"),
		field.Bool("isPremium").Default(false),
		field.String("tankClass").NotEmpty(),
		field.String("country").NotEmpty(),
	}
}

// Edges of the Tank.
func (Tank) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("nextTanks", Tank.Type).From("fromTankId"),
		edge.To("previousTanks", Tank.Type).From("nextTankId"),
		edge.To("modules", Module.Type),
	}
}
