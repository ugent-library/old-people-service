// Code generated by ent, DO NOT EDIT.

package person

import (
	"time"
)

const (
	// Label holds the string label denoting the person type in the database.
	Label = "person"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDateCreated holds the string denoting the date_created field in the database.
	FieldDateCreated = "date_created"
	// FieldDateUpdated holds the string denoting the date_updated field in the database.
	FieldDateUpdated = "date_updated"
	// FieldActive holds the string denoting the active field in the database.
	FieldActive = "active"
	// FieldBirthDate holds the string denoting the birth_date field in the database.
	FieldBirthDate = "birth_date"
	// FieldEmail holds the string denoting the email field in the database.
	FieldEmail = "email"
	// FieldOtherID holds the string denoting the other_id field in the database.
	FieldOtherID = "other_id"
	// FieldOrganizationID holds the string denoting the organization_id field in the database.
	FieldOrganizationID = "organization_id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldJobCategory holds the string denoting the job_category field in the database.
	FieldJobCategory = "job_category"
	// FieldOrcid holds the string denoting the orcid field in the database.
	FieldOrcid = "orcid"
	// FieldOrcidToken holds the string denoting the orcid_token field in the database.
	FieldOrcidToken = "orcid_token"
	// FieldPreferredFirstName holds the string denoting the preferred_first_name field in the database.
	FieldPreferredFirstName = "preferred_first_name"
	// FieldPreferredLastName holds the string denoting the preferred_last_name field in the database.
	FieldPreferredLastName = "preferred_last_name"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// Table holds the table name of the person in the database.
	Table = "person"
)

// Columns holds all SQL columns for person fields.
var Columns = []string{
	FieldID,
	FieldDateCreated,
	FieldDateUpdated,
	FieldActive,
	FieldBirthDate,
	FieldEmail,
	FieldOtherID,
	FieldOrganizationID,
	FieldFirstName,
	FieldFullName,
	FieldLastName,
	FieldJobCategory,
	FieldOrcid,
	FieldOrcidToken,
	FieldPreferredFirstName,
	FieldPreferredLastName,
	FieldTitle,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultDateCreated holds the default value on creation for the "date_created" field.
	DefaultDateCreated func() time.Time
	// DefaultDateUpdated holds the default value on creation for the "date_updated" field.
	DefaultDateUpdated func() time.Time
	// UpdateDefaultDateUpdated holds the default value on update for the "date_updated" field.
	UpdateDefaultDateUpdated func() time.Time
	// DefaultActive holds the default value on creation for the "active" field.
	DefaultActive bool
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() string
)
