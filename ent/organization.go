// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ugent-library/person-service/ent/organization"
	"github.com/ugent-library/person-service/ent/schema"
)

// Organization is the model entity for the Organization schema.
type Organization struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// DateCreated holds the value of the "date_created" field.
	DateCreated time.Time `json:"date_created,omitempty"`
	// DateUpdated holds the value of the "date_updated" field.
	DateUpdated time.Time `json:"date_updated,omitempty"`
	// PublicID holds the value of the "public_id" field.
	PublicID string `json:"public_id,omitempty"`
	// Type holds the value of the "type" field.
	Type string `json:"type,omitempty"`
	// NameDut holds the value of the "name_dut" field.
	NameDut string `json:"name_dut,omitempty"`
	// NameEng holds the value of the "name_eng" field.
	NameEng string `json:"name_eng,omitempty"`
	// OtherID holds the value of the "other_id" field.
	OtherID []schema.IdRef `json:"other_id,omitempty"`
	// OtherParentID holds the value of the "other_parent_id" field.
	OtherParentID string `json:"other_parent_id,omitempty"`
	// ParentID holds the value of the "parent_id" field.
	ParentID int `json:"parent_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrganizationQuery when eager-loading is set.
	Edges        OrganizationEdges `json:"edges"`
	selectValues sql.SelectValues
}

// OrganizationEdges holds the relations/edges for other nodes in the graph.
type OrganizationEdges struct {
	// People holds the value of the people edge.
	People []*Person `json:"people,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *Organization `json:"parent,omitempty"`
	// Children holds the value of the children edge.
	Children []*Organization `json:"children,omitempty"`
	// OrganizationPerson holds the value of the organization_person edge.
	OrganizationPerson []*OrganizationPerson `json:"organization_person,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PeopleOrErr returns the People value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) PeopleOrErr() ([]*Person, error) {
	if e.loadedTypes[0] {
		return e.People, nil
	}
	return nil, &NotLoadedError{edge: "people"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrganizationEdges) ParentOrErr() (*Organization, error) {
	if e.loadedTypes[1] {
		if e.Parent == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: organization.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) ChildrenOrErr() ([]*Organization, error) {
	if e.loadedTypes[2] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// OrganizationPersonOrErr returns the OrganizationPerson value or an error if the edge
// was not loaded in eager-loading.
func (e OrganizationEdges) OrganizationPersonOrErr() ([]*OrganizationPerson, error) {
	if e.loadedTypes[3] {
		return e.OrganizationPerson, nil
	}
	return nil, &NotLoadedError{edge: "organization_person"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Organization) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case organization.FieldOtherID:
			values[i] = new([]byte)
		case organization.FieldID, organization.FieldParentID:
			values[i] = new(sql.NullInt64)
		case organization.FieldPublicID, organization.FieldType, organization.FieldNameDut, organization.FieldNameEng, organization.FieldOtherParentID:
			values[i] = new(sql.NullString)
		case organization.FieldDateCreated, organization.FieldDateUpdated:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Organization fields.
func (o *Organization) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case organization.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case organization.FieldDateCreated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_created", values[i])
			} else if value.Valid {
				o.DateCreated = value.Time
			}
		case organization.FieldDateUpdated:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field date_updated", values[i])
			} else if value.Valid {
				o.DateUpdated = value.Time
			}
		case organization.FieldPublicID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field public_id", values[i])
			} else if value.Valid {
				o.PublicID = value.String
			}
		case organization.FieldType:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field type", values[i])
			} else if value.Valid {
				o.Type = value.String
			}
		case organization.FieldNameDut:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_dut", values[i])
			} else if value.Valid {
				o.NameDut = value.String
			}
		case organization.FieldNameEng:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name_eng", values[i])
			} else if value.Valid {
				o.NameEng = value.String
			}
		case organization.FieldOtherID:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field other_id", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &o.OtherID); err != nil {
					return fmt.Errorf("unmarshal field other_id: %w", err)
				}
			}
		case organization.FieldOtherParentID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field other_parent_id", values[i])
			} else if value.Valid {
				o.OtherParentID = value.String
			}
		case organization.FieldParentID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field parent_id", values[i])
			} else if value.Valid {
				o.ParentID = int(value.Int64)
			}
		default:
			o.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Organization.
// This includes values selected through modifiers, order, etc.
func (o *Organization) Value(name string) (ent.Value, error) {
	return o.selectValues.Get(name)
}

// QueryPeople queries the "people" edge of the Organization entity.
func (o *Organization) QueryPeople() *PersonQuery {
	return NewOrganizationClient(o.config).QueryPeople(o)
}

// QueryParent queries the "parent" edge of the Organization entity.
func (o *Organization) QueryParent() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryParent(o)
}

// QueryChildren queries the "children" edge of the Organization entity.
func (o *Organization) QueryChildren() *OrganizationQuery {
	return NewOrganizationClient(o.config).QueryChildren(o)
}

// QueryOrganizationPerson queries the "organization_person" edge of the Organization entity.
func (o *Organization) QueryOrganizationPerson() *OrganizationPersonQuery {
	return NewOrganizationClient(o.config).QueryOrganizationPerson(o)
}

// Update returns a builder for updating this Organization.
// Note that you need to call Organization.Unwrap() before calling this method if this Organization
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Organization) Update() *OrganizationUpdateOne {
	return NewOrganizationClient(o.config).UpdateOne(o)
}

// Unwrap unwraps the Organization entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Organization) Unwrap() *Organization {
	_tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Organization is not a transactional entity")
	}
	o.config.driver = _tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Organization) String() string {
	var builder strings.Builder
	builder.WriteString("Organization(")
	builder.WriteString(fmt.Sprintf("id=%v, ", o.ID))
	builder.WriteString("date_created=")
	builder.WriteString(o.DateCreated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("date_updated=")
	builder.WriteString(o.DateUpdated.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("public_id=")
	builder.WriteString(o.PublicID)
	builder.WriteString(", ")
	builder.WriteString("type=")
	builder.WriteString(o.Type)
	builder.WriteString(", ")
	builder.WriteString("name_dut=")
	builder.WriteString(o.NameDut)
	builder.WriteString(", ")
	builder.WriteString("name_eng=")
	builder.WriteString(o.NameEng)
	builder.WriteString(", ")
	builder.WriteString("other_id=")
	builder.WriteString(fmt.Sprintf("%v", o.OtherID))
	builder.WriteString(", ")
	builder.WriteString("other_parent_id=")
	builder.WriteString(o.OtherParentID)
	builder.WriteString(", ")
	builder.WriteString("parent_id=")
	builder.WriteString(fmt.Sprintf("%v", o.ParentID))
	builder.WriteByte(')')
	return builder.String()
}

// Organizations is a parsable slice of Organization.
type Organizations []*Organization
