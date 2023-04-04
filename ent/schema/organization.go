package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Organization struct {
	ent.Schema
}

func (Organization) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization"},
	}
}

func (Organization) Fields() []ent.Field {
	// field "id" is implied
	return []ent.Field{
		field.String("primary_id"),
		field.String("name"),
	}
}

func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("people", Person.Type).
			StorageKey(edge.Table("organization_person")),
	}
}
