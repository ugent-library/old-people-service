package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Person holds the schema definition for the Person entity.
type Person struct {
	ent.Schema
}

func (Person) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "person"},
	}
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active").Default(false),
		field.String("birth_date").Optional(),
		field.String("email").Optional(),
		field.JSON("identifier", TypeVals{}).Optional().Default(TypeVals{}),
		field.String("given_name").Optional(),
		field.String("name").Optional(),
		field.String("family_name").Optional(),
		field.Strings("job_category").Optional().Default([]string{}),
		field.String("preferred_given_name").Optional(),
		field.String("preferred_family_name").Optional(),
		field.String("honorific_prefix").Optional(),
		field.Strings("role").Optional().Default([]string{}),
		field.JSON("settings", map[string]string{}).Optional().Default(map[string]string{}),
		field.Strings("object_class").Optional().Default([]string{}),
		field.String("expiration_date").Optional(),
		field.JSON("token", TypeVals{}).Optional().Default(TypeVals{}),
	}
}

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		PublicIdMixin{},
	}
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("active"),
		index.Fields("email"),
		index.Fields("given_name"),
		index.Fields("family_name"),
		index.Fields("name"),
	}
}
