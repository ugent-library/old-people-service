// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/people-service/ent/organizationperson"
)

// OrganizationPerson is the model entity for the OrganizationPerson schema.
type OrganizationPerson struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// DateUpdated holds the value of the "date_updated" field.
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// OrganizationID holds the value of the "organization_id" field.
	OrganizationID int `json:"organization_id,omitempty"`
	// PersonID holds the value of the "person_id" field.
	PersonID int `json:"person_id,omitempty"`
	// From holds the value of the "from" field.
	From time.Time `json:"from,omitempty"`
	// Until holds the value of the "until" field.
	Until        *time.Time `json:"until,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*OrganizationPerson) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organizationperson.FieldID, organizationperson.FieldOrganizationID, organizationperson.FieldPersonID:
			values[i] = new(sql.NullInt64)
		case organizationperson.FieldDateCreated, organizationperson.FieldDateUpdated, organizationperson.FieldFrom, organizationperson.FieldUntil:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the OrganizationPerson fields.
func (op *OrganizationPerson) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organizationperson.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			op.ID = int(value.Int64)
		case organizationperson.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				op.DateCreated = value.Time
			}
		case organizationperson.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_updated", values[i])
			} else if value.Valid {
				op.DateUpdated = value.Time
			}
		case organizationperson.FieldOrganizationID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field organization_id", values[i])
			} else if value.Valid {
				op.OrganizationID = int(value.Int64)
			}
		case organizationperson.FieldPersonID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field person_id", values[i])
			} else if value.Valid {
				op.PersonID = int(value.Int64)
			}
		case organizationperson.FieldFrom:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field from", values[i])
			} else if value.Valid {
				op.From = value.Time
			}
		case organizationperson.FieldUntil:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field until", values[i])
			} else if value.Valid {
				op.Until = new(time.Time)
				*op.Until = value.Time
			}
		default:
			op.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the OrganizationPerson.
// This includes values selected through modifiers, order, etc.
func (op *OrganizationPerson) Value(name string) (ent.Value, error) {
	return op.selectValues.Get(name)
}

// Update returns a builder for updating this OrganizationPerson.
// Note that you need to call OrganizationPerson.Unwrap() before calling this method if this OrganizationPerson
// was returned from a transaction, and the transaction was committed or rolled back.
func (op *OrganizationPerson) Update() *OrganizationPersonUpdateOne {
	return NewOrganizationPersonClient(op.config).UpdateOne(op)
}

// Unwrap unwraps the OrganizationPerson entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (op *OrganizationPerson) Unwrap() *OrganizationPerson {
	_tx, ok := op.config.driver.(*txDriver)
	if !ok {
		panic("ent: OrganizationPerson is not a transactional entity")
	}
	op.config.driver = _tx.drv
	return op
}

// String implements the fmt.Stringer.
func (op *OrganizationPerson) String() string {
	var builder strings.Builder
	builder.WriteString("OrganizationPerson(")
	builder.WriteString(fmt.Sprintf("id=%v, ", op.ID))
	builder.WriteString("date_created=")
	builder.WriteString(op.DateCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date_updated=")
	builder.WriteString(op.DateUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("organization_id=")
	builder.WriteString(fmt.Sprintf("%v", op.OrganizationID))
	builder.WriteString(", ")
	builder.WriteString("person_id=")
	builder.WriteString(fmt.Sprintf("%v", op.PersonID))
	builder.WriteString(", ")
	builder.WriteString("from=")
	builder.WriteString(op.From.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := op.Until; v != nil {
		builder.WriteString("until=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteByte(')')
	return builder.String()
}

// OrganizationPersons is a parsable slice of OrganizationPerson.
type OrganizationPersons []*OrganizationPerson
