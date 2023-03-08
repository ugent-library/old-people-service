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

type OrcidSettings struct {
	SendEmails bool `json:"send-emails"`
}

type Settings struct {
	ShowBio         bool `json:"show_bio"`
	ShowEmail       bool `json:"show_email"`
	ShowGravatar    bool `json:"show_gravatar"`
	ShowWorkAddress bool `json:"show_work_address"`
}

type OrcidVerify struct {
	OrcidID       string         `json:"orcid_id,omitempty"`
	OrcidToken    string         `json:"orcid_token,omitempty"`
	OrcidSettings *OrcidSettings `json:"orcid_settings,omitempty"`
}

func (Person) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "person"},
	}
}

// Fields of the Person.
func (Person) Fields() []ent.Field {
	return []ent.Field{
		// attributes mapped and imported from LDAP
		field.Strings("object_class").Optional(),
		field.String("ugent_username").Optional(),
		field.String("first_name").Optional(),
		field.String("middle_name").Optional(),
		field.String("last_name").Optional(),
		field.Strings("ugent_id").Optional(),
		field.String("birth_date").Optional(),
		field.String("email").Optional(),
		field.String("nationality").Optional(),
		field.Strings("ugent_barcode").Optional(),
		field.Strings("ugent_job_category").Optional(),
		field.String("title").Optional(),
		field.String("ugent_tel").Optional(),
		field.Strings("ugent_campus").Optional(),
		field.Strings("ugent_department_id").Optional(),
		field.Strings("ugent_faculty_id").Optional(),
		field.Strings("ugent_job_title").Optional(),
		field.String("ugent_street_address").Optional(),
		field.String("ugent_postal_code").Optional(),
		field.String("ugent_locality").Optional(),
		field.String("ugent_last_enrolled").Optional(),
		field.String("home_street_address").Optional(),
		field.String("home_postal_code").Optional(),
		field.String("home_locality").Optional(),
		field.String("home_country").Optional(),
		field.String("home_tel").Optional(),
		field.String("dorm_street_address").Optional(),
		field.String("dorm_postal_code").Optional(),
		field.String("dorm_locality").Optional(),
		field.String("dorm_country").Optional(),
		field.Strings("research_discipline").Optional(),
		field.Strings("research_discipline_code").Optional(),
		field.String("ugent_expiration_date").Optional(),
		field.Strings("uzgent_job_title").Optional(),
		field.Strings("uzgent_department_name").Optional(),
		field.Strings("uzgent_id").Optional(),
		field.Strings("ugent_ext_category").Optional(),

		// old ldap attributes?
		// 'YYYYmmdd'
		field.String("ugent_appointment_date").Optional(),
		field.Strings("ugent_department_name").Optional(),

		// orcid
		field.String("orcid_bio").Optional(),
		field.String("orcid_id").Optional(),
		field.JSON("orcid_settings", OrcidSettings{}).Optional(),
		field.String("orcid_token").Optional(),
		field.JSON("orcid_verify", OrcidVerify{}).Optional(),

		// internal attributes
		field.Bool("active").Default(false),
		field.Bool("deleted").Default(false),
		field.JSON("settings", Settings{}).Optional(),
		field.Strings("roles").Optional(),
		field.Int("publication_count").Default(0).Optional(),
		field.String("ugent_memorialis_id").Optional(),

		// source?
		field.String("preferred_first_name").Optional(),
		field.String("preferred_last_name").Optional(),
		// [{ "_id": "<id>" }]
		field.JSON("replaces", []map[string]string{}).Optional(),
		field.JSON("replaced_by", []map[string]string{}).Optional(),
		field.Time("date_last_login").Optional(),
	}
}

func (Person) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UUIDMixin{},
		TimeMixin{},
	}
}

// Edges of the Person.
func (Person) Edges() []ent.Edge {
	return nil
}

func (Person) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("active"),
		index.Fields("orcid_id"),
		index.Fields("ugent_id"),
		index.Fields("ugent_username"),
		index.Fields("email"),
		index.Fields("first_name"),
		index.Fields("last_name"),
		index.Fields("publication_count"),
	}
}
