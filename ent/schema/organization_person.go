package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type OrganizationPerson struct {
	ent.Schema
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

// Note: cannot model M2M with duplicate entries in entgo
// see https://github.com/ent/ent/issues/2964
// see migration file for additional checks
func (OrganizationPerson) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("person_id"),
		index.Fields("organization_id"),
		index.Fields("person_id", "organization_id", "from").Unique(),
	}
}
