// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// OrganizationColumns holds the columns for the "organization" table.
	OrganizationColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_updated", Type: field.TypeTime},
		{Name: "public_id", Type: field.TypeString, Unique: true},
		{Name: "identifier", Type: field.TypeJSON, Nullable: true},
		{Name: "identifier_values", Type: field.TypeJSON, Nullable: true},
		{Name: "type", Type: field.TypeString, Default: "organization"},
		{Name: "acronym", Type: field.TypeString, Nullable: true},
		{Name: "name_dut", Type: field.TypeString, Nullable: true},
		{Name: "name_eng", Type: field.TypeString, Nullable: true},
	}
	// OrganizationTable holds the schema information for the "organization" table.
	OrganizationTable = &schema.Table{
		Name:       "organization",
		Columns:    OrganizationColumns,
		PrimaryKey: []*schema.Column{OrganizationColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "organization_type",
				Unique:  false,
				Columns: []*schema.Column{OrganizationColumns[6]},
			},
			{
				Name:    "organization_acronym",
				Unique:  false,
				Columns: []*schema.Column{OrganizationColumns[7]},
			},
		},
	}
	// OrganizationParentColumns holds the columns for the "organization_parent" table.
	OrganizationParentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_updated", Type: field.TypeTime},
		{Name: "parent_organization_id", Type: field.TypeInt},
		{Name: "organization_id", Type: field.TypeInt},
		{Name: "from", Type: field.TypeTime},
		{Name: "until", Type: field.TypeTime, Nullable: true},
	}
	// OrganizationParentTable holds the schema information for the "organization_parent" table.
	OrganizationParentTable = &schema.Table{
		Name:       "organization_parent",
		Columns:    OrganizationParentColumns,
		PrimaryKey: []*schema.Column{OrganizationParentColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "organizationparent_parent_organization_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationParentColumns[3]},
			},
			{
				Name:    "organizationparent_organization_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationParentColumns[4]},
			},
			{
				Name:    "organizationparent_parent_organization_id_organization_id_from",
				Unique:  true,
				Columns: []*schema.Column{OrganizationParentColumns[3], OrganizationParentColumns[4], OrganizationParentColumns[5]},
			},
		},
	}
	// OrganizationPersonColumns holds the columns for the "organization_person" table.
	OrganizationPersonColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_updated", Type: field.TypeTime},
		{Name: "organization_id", Type: field.TypeInt},
		{Name: "person_id", Type: field.TypeInt},
		{Name: "from", Type: field.TypeTime},
		{Name: "until", Type: field.TypeTime, Nullable: true},
	}
	// OrganizationPersonTable holds the schema information for the "organization_person" table.
	OrganizationPersonTable = &schema.Table{
		Name:       "organization_person",
		Columns:    OrganizationPersonColumns,
		PrimaryKey: []*schema.Column{OrganizationPersonColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "organizationperson_person_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationPersonColumns[4]},
			},
			{
				Name:    "organizationperson_organization_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationPersonColumns[3]},
			},
			{
				Name:    "organizationperson_person_id_organization_id_from",
				Unique:  true,
				Columns: []*schema.Column{OrganizationPersonColumns[4], OrganizationPersonColumns[3], OrganizationPersonColumns[5]},
			},
		},
	}
	// PersonColumns holds the columns for the "person" table.
	PersonColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_updated", Type: field.TypeTime},
		{Name: "public_id", Type: field.TypeString, Unique: true},
		{Name: "identifier", Type: field.TypeJSON, Nullable: true},
		{Name: "identifier_values", Type: field.TypeJSON, Nullable: true},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "birth_date", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "given_name", Type: field.TypeString, Nullable: true},
		{Name: "name", Type: field.TypeString, Nullable: true},
		{Name: "family_name", Type: field.TypeString, Nullable: true},
		{Name: "job_category", Type: field.TypeJSON, Nullable: true},
		{Name: "preferred_given_name", Type: field.TypeString, Nullable: true},
		{Name: "preferred_family_name", Type: field.TypeString, Nullable: true},
		{Name: "honorific_prefix", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeJSON, Nullable: true},
		{Name: "settings", Type: field.TypeJSON, Nullable: true},
		{Name: "object_class", Type: field.TypeJSON, Nullable: true},
		{Name: "expiration_date", Type: field.TypeString, Nullable: true},
		{Name: "token", Type: field.TypeJSON, Nullable: true},
	}
	// PersonTable holds the schema information for the "person" table.
	PersonTable = &schema.Table{
		Name:       "person",
		Columns:    PersonColumns,
		PrimaryKey: []*schema.Column{PersonColumns[0]},
		Indexes: []*schema.Index{
			{
				Name:    "person_active",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[6]},
			},
			{
				Name:    "person_email",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[8]},
			},
			{
				Name:    "person_given_name",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[9]},
			},
			{
				Name:    "person_family_name",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[11]},
			},
			{
				Name:    "person_name",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[10]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationTable,
		OrganizationParentTable,
		OrganizationPersonTable,
		PersonTable,
	}
)

func init() {
	OrganizationTable.Annotation = &entsql.Annotation{
		Table: "organization",
	}
	OrganizationParentTable.Annotation = &entsql.Annotation{
		Table: "organization_parent",
	}
	OrganizationPersonTable.Annotation = &entsql.Annotation{
		Table: "organization_person",
	}
	PersonTable.Annotation = &entsql.Annotation{
		Table: "person",
	}
}
