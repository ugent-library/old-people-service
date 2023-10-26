package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
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
		field.String("type").Default("organization"),
		field.String("acronym").Optional(),
		field.String("name_dut").Optional(),
		field.String("name_eng").Optional(),
		field.JSON("identifier", TypeVals{}).Optional().Default(TypeVals{}),
	}
}

func (Organization) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		PublicIdMixin{},
	}
}

func (Organization) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("type"),
		index.Fields("acronym"),
	}
}
