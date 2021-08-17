package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname"),
		field.String("brand"),
		field.Int("model_year"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return nil
}
