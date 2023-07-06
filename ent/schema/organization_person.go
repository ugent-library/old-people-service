package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/person-service/models"
)

type OrganizationPerson struct {
	ent.Schema
}

var genBeginningOfTime = func() time.Time {
	return models.BeginningOfTime
}
var genEndOfTime = func() time.Time {
	return models.EndOfTime
}

func (OrganizationPerson) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization_person"},
	}
}

func (OrganizationPerson) Fields() []ent.Field {
	return []ent.Field{
		field.Int("organization_id"),
		field.Int("person_id"),
		field.Time("from").Default(genBeginningOfTime),
		field.Time("until").Default(genEndOfTime).UpdateDefault(genEndOfTime),
	}
}

func (OrganizationPerson) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

func (OrganizationPerson) Edges() []ent.Edge {
	/*
		Not sure why this works

		cf. https://github.com/ent/ent/issues/2964

		without "Required" per field entgo claims that
		"person_id" is not holding a foreign key

		It will generate a unique index though on the combination
		of person_id and organization_id, not on each separately,
		from some reason.
	*/
	return []ent.Edge{
		edge.To("people", Person.Type).
			Unique().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Required().Field("person_id"),
		edge.To("organizations", Organization.Type).
			Unique().
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}).
			Required().Field("organization_id"),
	}
}
