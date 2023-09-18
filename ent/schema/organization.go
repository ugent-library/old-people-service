package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Organization struct {
	ent.Schema
}

var OrganizationTypes = []string{
	"organization",
	"department",
}

var OrganizationIdTypes = []string{
	// vb. WE03, WE03V
	// (gismo: org-code)
	"ugent_id",
	// vb. WE03* (biblio uses "*" to mark historic organizations)
	// (gismo: biblio-code)
	"biblio_id",
	// vb. 000006045
	// (gismo: memorialis)
	"ugent_memorialis_id",
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
		field.JSON("other_id", IdRefs{}).Optional().Default(IdRefs{}),
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
