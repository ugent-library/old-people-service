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
		{Name: "gismo_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "type", Type: field.TypeString, Default: "organization"},
		{Name: "name_dut", Type: field.TypeString, Nullable: true},
		{Name: "name_eng", Type: field.TypeString, Nullable: true},
		{Name: "other_id", Type: field.TypeJSON, Nullable: true},
		{Name: "parent_id", Type: field.TypeInt, Nullable: true},
	}
	// OrganizationTable holds the schema information for the "organization" table.
	OrganizationTable = &schema.Table{
		Name:       "organization",
		Columns:    OrganizationColumns,
		PrimaryKey: []*schema.Column{OrganizationColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_organization_children",
				Columns:    []*schema.Column{OrganizationColumns[9]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "organization_type",
				Unique:  false,
				Columns: []*schema.Column{OrganizationColumns[5]},
			},
			{
				Name:    "organization_parent_id",
				Unique:  false,
				Columns: []*schema.Column{OrganizationColumns[9]},
			},
		},
	}
	// OrganizationPersonColumns holds the columns for the "organization_person" table.
	OrganizationPersonColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_updated", Type: field.TypeTime},
		{Name: "person_id", Type: field.TypeInt},
		{Name: "organization_id", Type: field.TypeInt},
	}
	// OrganizationPersonTable holds the schema information for the "organization_person" table.
	OrganizationPersonTable = &schema.Table{
		Name:       "organization_person",
		Columns:    OrganizationPersonColumns,
		PrimaryKey: []*schema.Column{OrganizationPersonColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "organization_person_person_people",
				Columns:    []*schema.Column{OrganizationPersonColumns[3]},
				RefColumns: []*schema.Column{PersonColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "organization_person_organization_organizations",
				Columns:    []*schema.Column{OrganizationPersonColumns[4]},
				RefColumns: []*schema.Column{OrganizationColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "organizationperson_person_id_organization_id",
				Unique:  true,
				Columns: []*schema.Column{OrganizationPersonColumns[3], OrganizationPersonColumns[4]},
			},
		},
	}
	// PersonColumns holds the columns for the "person" table.
	PersonColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_created", Type: field.TypeTime},
		{Name: "date_updated", Type: field.TypeTime},
		{Name: "public_id", Type: field.TypeString, Unique: true},
		{Name: "gismo_id", Type: field.TypeString, Unique: true, Nullable: true},
		{Name: "active", Type: field.TypeBool, Default: false},
		{Name: "birth_date", Type: field.TypeString, Nullable: true},
		{Name: "email", Type: field.TypeString, Nullable: true},
		{Name: "other_id", Type: field.TypeJSON, Nullable: true},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "full_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "job_category", Type: field.TypeJSON, Nullable: true},
		{Name: "orcid", Type: field.TypeString, Nullable: true},
		{Name: "orcid_token", Type: field.TypeString, Nullable: true},
		{Name: "preferred_first_name", Type: field.TypeString, Nullable: true},
		{Name: "preferred_last_name", Type: field.TypeString, Nullable: true},
		{Name: "title", Type: field.TypeString, Nullable: true},
		{Name: "role", Type: field.TypeJSON, Nullable: true},
		{Name: "settings", Type: field.TypeJSON, Nullable: true},
		{Name: "object_class", Type: field.TypeJSON, Nullable: true},
		{Name: "expiration_date", Type: field.TypeString, Nullable: true},
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
				Columns: []*schema.Column{PersonColumns[5]},
			},
			{
				Name:    "person_orcid",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[13]},
			},
			{
				Name:    "person_email",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[7]},
			},
			{
				Name:    "person_first_name",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[9]},
			},
			{
				Name:    "person_last_name",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[11]},
			},
			{
				Name:    "person_full_name",
				Unique:  false,
				Columns: []*schema.Column{PersonColumns[10]},
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		OrganizationTable,
		OrganizationPersonTable,
		PersonTable,
	}
)

func init() {
	OrganizationTable.ForeignKeys[0].RefTable = OrganizationTable
	OrganizationTable.Annotation = &entsql.Annotation{
		Table: "organization",
	}
	OrganizationPersonTable.ForeignKeys[0].RefTable = PersonTable
	OrganizationPersonTable.ForeignKeys[1].RefTable = OrganizationTable
	OrganizationPersonTable.Annotation = &entsql.Annotation{
		Table: "organization_person",
	}
	PersonTable.Annotation = &entsql.Annotation{
		Table: "person",
	}
}
