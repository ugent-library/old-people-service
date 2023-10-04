// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people-service/ent/person"
	"github.com/ugent-library/people-service/ent/schema"
)

// Person is the model entity for the Person schema.
type Person struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// DateUpdated holds the value of the "date_updated" field.
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// PublicID holds the value of the "public_id" field.
	PublicID string `json:"public_id,omitempty"`
	// GismoID holds the value of the "gismo_id" field.
	GismoID *string `json:"gismo_id,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// BirthDate holds the value of the "birth_date" field.
	BirthDate string `json:"birth_date,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// OtherID holds the value of the "other_id" field.
	OtherID schema.IdRefs `json:"other_id,omitempty"`
	// FirstName holds the value of the "first_name" field.
	FirstName string `json:"first_name,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName string `json:"full_name,omitempty"`
	// LastName holds the value of the "last_name" field.
	LastName string `json:"last_name,omitempty"`
	// JobCategory holds the value of the "job_category" field.
	JobCategory []string `json:"job_category,omitempty"`
	// Orcid holds the value of the "orcid" field.
	Orcid string `json:"orcid,omitempty"`
	// OrcidToken holds the value of the "orcid_token" field.
	OrcidToken string `json:"orcid_token,omitempty"`
	// PreferredFirstName holds the value of the "preferred_first_name" field.
	PreferredFirstName string `json:"preferred_first_name,omitempty"`
	// PreferredLastName holds the value of the "preferred_last_name" field.
	PreferredLastName string `json:"preferred_last_name,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Role holds the value of the "role" field.
	Role []string `json:"role,omitempty"`
	// Settings holds the value of the "settings" field.
	Settings map[string]string `json:"settings,omitempty"`
	// ObjectClass holds the value of the "object_class" field.
	ObjectClass []string `json:"object_class,omitempty"`
	// ExpirationDate holds the value of the "expiration_date" field.
	ExpirationDate string `json:"expiration_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PersonQuery when eager-loading is set.
	Edges        PersonEdges `json:"edges"`
	selectValues sql.SelectValues
}

// PersonEdges holds the relations/edges for other nodes in the graph.
type PersonEdges struct {
	// Organizations holds the value of the organizations edge.
	Organizations []*Organization `json:"organizations,omitempty"`
	// OrganizationPerson holds the value of the organization_person edge.
	OrganizationPerson []*OrganizationPerson `json:"organization_person,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OrganizationsOrErr returns the Organizations value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) OrganizationsOrErr() ([]*Organization, error) {
	if e.loadedTypes[0] {
		return e.Organizations, nil
	}
	return nil, &NotLoadedError{edge: "organizations"}
}

// OrganizationPersonOrErr returns the OrganizationPerson value or an error if the edge
// was not loaded in eager-loading.
func (e PersonEdges) OrganizationPersonOrErr() ([]*OrganizationPerson, error) {
	if e.loadedTypes[1] {
		return e.OrganizationPerson, nil
	}
	return nil, &NotLoadedError{edge: "organization_person"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Person) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case person.FieldOtherID, person.FieldJobCategory, person.FieldRole, person.FieldSettings, person.FieldObjectClass:
			values[i] = new([]byte)
		case person.FieldActive:
			values[i] = new(sql.NullBool)
		case person.FieldID:
			values[i] = new(sql.NullInt64)
		case person.FieldPublicID, person.FieldGismoID, person.FieldBirthDate, person.FieldEmail, person.FieldFirstName, person.FieldFullName, person.FieldLastName, person.FieldOrcid, person.FieldOrcidToken, person.FieldPreferredFirstName, person.FieldPreferredLastName, person.FieldTitle, person.FieldExpirationDate:
			values[i] = new(sql.NullString)
		case person.FieldDateCreated, person.FieldDateUpdated:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Person fields.
func (pe *Person) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case person.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pe.ID = int(value.Int64)
		case person.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				pe.DateCreated = value.Time
			}
		case person.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_updated", values[i])
			} else if value.Valid {
				pe.DateUpdated = value.Time
			}
		case person.FieldPublicID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_id", values[i])
			} else if value.Valid {
				pe.PublicID = value.String
			}
		case person.FieldGismoID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gismo_id", values[i])
			} else if value.Valid {
				pe.GismoID = new(string)
				*pe.GismoID = value.String
			}
		case person.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				pe.Active = value.Bool
			}
		case person.FieldBirthDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field birth_date", values[i])
			} else if value.Valid {
				pe.BirthDate = value.String
			}
		case person.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				pe.Email = value.String
			}
		case person.FieldOtherID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field other_id", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.OtherID); err != nil {
					return fmt.Errorf("unmarshal field other_id: %w", err)
				}
			}
		case person.FieldFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field first_name", values[i])
			} else if value.Valid {
				pe.FirstName = value.String
			}
		case person.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				pe.FullName = value.String
			}
		case person.FieldLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field last_name", values[i])
			} else if value.Valid {
				pe.LastName = value.String
			}
		case person.FieldJobCategory:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field job_category", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.JobCategory); err != nil {
					return fmt.Errorf("unmarshal field job_category: %w", err)
				}
			}
		case person.FieldOrcid:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orcid", values[i])
			} else if value.Valid {
				pe.Orcid = value.String
			}
		case person.FieldOrcidToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field orcid_token", values[i])
			} else if value.Valid {
				pe.OrcidToken = value.String
			}
		case person.FieldPreferredFirstName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preferred_first_name", values[i])
			} else if value.Valid {
				pe.PreferredFirstName = value.String
			}
		case person.FieldPreferredLastName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field preferred_last_name", values[i])
			} else if value.Valid {
				pe.PreferredLastName = value.String
			}
		case person.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				pe.Title = value.String
			}
		case person.FieldRole:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Role); err != nil {
					return fmt.Errorf("unmarshal field role: %w", err)
				}
			}
		case person.FieldSettings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field settings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.Settings); err != nil {
					return fmt.Errorf("unmarshal field settings: %w", err)
				}
			}
		case person.FieldObjectClass:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field object_class", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &pe.ObjectClass); err != nil {
					return fmt.Errorf("unmarshal field object_class: %w", err)
				}
			}
		case person.FieldExpirationDate:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field expiration_date", values[i])
			} else if value.Valid {
				pe.ExpirationDate = value.String
			}
		default:
			pe.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Person.
// This includes values selected through modifiers, order, etc.
func (pe *Person) Value(name string) (ent.Value, error) {
	return pe.selectValues.Get(name)
}

// QueryOrganizations queries the "organizations" edge of the Person entity.
func (pe *Person) QueryOrganizations() *OrganizationQuery {
	return NewPersonClient(pe.config).QueryOrganizations(pe)
}

// QueryOrganizationPerson queries the "organization_person" edge of the Person entity.
func (pe *Person) QueryOrganizationPerson() *OrganizationPersonQuery {
	return NewPersonClient(pe.config).QueryOrganizationPerson(pe)
}

// Update returns a builder for updating this Person.
// Note that you need to call Person.Unwrap() before calling this method if this Person
// was returned from a transaction, and the transaction was committed or rolled back.
func (pe *Person) Update() *PersonUpdateOne {
	return NewPersonClient(pe.config).UpdateOne(pe)
}

// Unwrap unwraps the Person entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pe *Person) Unwrap() *Person {
	_tx, ok := pe.config.driver.(*txDriver)
	if !ok {
		panic("ent: Person is not a transactional entity")
	}
	pe.config.driver = _tx.drv
	return pe
}

// String implements the fmt.Stringer.
func (pe *Person) String() string {
	var builder strings.Builder
	builder.WriteString("Person(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pe.ID))
	builder.WriteString("date_created=")
	builder.WriteString(pe.DateCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date_updated=")
	builder.WriteString(pe.DateUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("public_id=")
	builder.WriteString(pe.PublicID)
	builder.WriteString(", ")
	if v := pe.GismoID; v != nil {
		builder.WriteString("gismo_id=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", pe.Active))
	builder.WriteString(", ")
	builder.WriteString("birth_date=")
	builder.WriteString(pe.BirthDate)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(pe.Email)
	builder.WriteString(", ")
	builder.WriteString("other_id=")
	builder.WriteString(fmt.Sprintf("%v", pe.OtherID))
	builder.WriteString(", ")
	builder.WriteString("first_name=")
	builder.WriteString(pe.FirstName)
	builder.WriteString(", ")
	builder.WriteString("full_name=")
	builder.WriteString(pe.FullName)
	builder.WriteString(", ")
	builder.WriteString("last_name=")
	builder.WriteString(pe.LastName)
	builder.WriteString(", ")
	builder.WriteString("job_category=")
	builder.WriteString(fmt.Sprintf("%v", pe.JobCategory))
	builder.WriteString(", ")
	builder.WriteString("orcid=")
	builder.WriteString(pe.Orcid)
	builder.WriteString(", ")
	builder.WriteString("orcid_token=")
	builder.WriteString(pe.OrcidToken)
	builder.WriteString(", ")
	builder.WriteString("preferred_first_name=")
	builder.WriteString(pe.PreferredFirstName)
	builder.WriteString(", ")
	builder.WriteString("preferred_last_name=")
	builder.WriteString(pe.PreferredLastName)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(pe.Title)
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", pe.Role))
	builder.WriteString(", ")
	builder.WriteString("settings=")
	builder.WriteString(fmt.Sprintf("%v", pe.Settings))
	builder.WriteString(", ")
	builder.WriteString("object_class=")
	builder.WriteString(fmt.Sprintf("%v", pe.ObjectClass))
	builder.WriteString(", ")
	builder.WriteString("expiration_date=")
	builder.WriteString(pe.ExpirationDate)
	builder.WriteByte(')')
	return builder.String()
}

// Persons is a parsable slice of Person.
type Persons []*Person
