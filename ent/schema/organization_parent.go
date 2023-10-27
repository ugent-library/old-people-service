package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// TODO: give constraints and indexes more sane names
type OrganizationParent struct {
	ent.Schema
}

func (OrganizationParent) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "organization_parent"},
	}
}

func (OrganizationParent) Fields() []ent.Field {
	return []ent.Field{
		field.Int("parent_organization_id"),
		field.Int("organization_id"),
		field.Time("from").Default(genBeginningOfTime),
		field.Time("until").Nillable().Optional(),
	}
}

func (OrganizationParent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Note: cannot model M2M with duplicate entries in entgo
// see https://github.com/ent/ent/issues/2964
// see migration file for additional checks
func (OrganizationParent) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("parent_organization_id"),
		index.Fields("organization_id"),
		index.Fields("parent_organization_id", "organization_id", "from").Unique(),
	}
}
