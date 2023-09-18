package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/ugent-library/people-service/models"
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
		field.String("gismo_id").Optional().Unique().Nillable(),
		field.String("type").Default("organization"),
		field.String("name_dut").Optional(),
		field.String("name_eng").Optional(),
		field.JSON("other_id", models.IdRefs{}).Optional().Default(models.IdRefs{}),
		field.Int("parent_id").Optional(),
	}
}

func (Organization) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("people", Person.Type).
			Ref("organizations").
			Through("organization_person", OrganizationPerson.Type),
		edge.To("children", Organization.Type).
			From("parent").
			Field("parent_id").
			Unique(),
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
		index.Fields("parent_id"),
	}
}
