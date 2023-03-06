// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people/ent/person"
)

// PersonCreate is the builder for creating a Person entity.
type PersonCreate struct {
	config
	mutation *PersonMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (pc *PersonCreate) SetDateCreated(t time.Time) *PersonCreate {
	pc.mutation.SetDateCreated(t)
	return pc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDateCreated(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetDateCreated(*t)
	}
	return pc
}

// SetDateUpdated sets the "date_updated" field.
func (pc *PersonCreate) SetDateUpdated(t time.Time) *PersonCreate {
	pc.mutation.SetDateUpdated(t)
	return pc
}

// SetNillableDateUpdated sets the "date_updated" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDateUpdated(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetDateUpdated(*t)
	}
	return pc
}

// SetObjectClass sets the "object_class" field.
func (pc *PersonCreate) SetObjectClass(s string) *PersonCreate {
	pc.mutation.SetObjectClass(s)
	return pc
}

// SetNillableObjectClass sets the "object_class" field if the given value is not nil.
func (pc *PersonCreate) SetNillableObjectClass(s *string) *PersonCreate {
	if s != nil {
		pc.SetObjectClass(*s)
	}
	return pc
}

// SetUgentUsername sets the "ugent_username" field.
func (pc *PersonCreate) SetUgentUsername(s string) *PersonCreate {
	pc.mutation.SetUgentUsername(s)
	return pc
}

// SetNillableUgentUsername sets the "ugent_username" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentUsername(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentUsername(*s)
	}
	return pc
}

// SetFirstName sets the "first_name" field.
func (pc *PersonCreate) SetFirstName(s string) *PersonCreate {
	pc.mutation.SetFirstName(s)
	return pc
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableFirstName(s *string) *PersonCreate {
	if s != nil {
		pc.SetFirstName(*s)
	}
	return pc
}

// SetMiddleName sets the "middle_name" field.
func (pc *PersonCreate) SetMiddleName(s string) *PersonCreate {
	pc.mutation.SetMiddleName(s)
	return pc
}

// SetNillableMiddleName sets the "middle_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableMiddleName(s *string) *PersonCreate {
	if s != nil {
		pc.SetMiddleName(*s)
	}
	return pc
}

// SetLastName sets the "last_name" field.
func (pc *PersonCreate) SetLastName(s string) *PersonCreate {
	pc.mutation.SetLastName(s)
	return pc
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableLastName(s *string) *PersonCreate {
	if s != nil {
		pc.SetLastName(*s)
	}
	return pc
}

// SetUgentID sets the "ugent_id" field.
func (pc *PersonCreate) SetUgentID(s []string) *PersonCreate {
	pc.mutation.SetUgentID(s)
	return pc
}

// SetBirthDate sets the "birth_date" field.
func (pc *PersonCreate) SetBirthDate(s string) *PersonCreate {
	pc.mutation.SetBirthDate(s)
	return pc
}

// SetNillableBirthDate sets the "birth_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableBirthDate(s *string) *PersonCreate {
	if s != nil {
		pc.SetBirthDate(*s)
	}
	return pc
}

// SetEmail sets the "email" field.
func (pc *PersonCreate) SetEmail(s string) *PersonCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetNillableEmail sets the "email" field if the given value is not nil.
func (pc *PersonCreate) SetNillableEmail(s *string) *PersonCreate {
	if s != nil {
		pc.SetEmail(*s)
	}
	return pc
}

// SetGender sets the "gender" field.
func (pc *PersonCreate) SetGender(pe person.Gender) *PersonCreate {
	pc.mutation.SetGender(pe)
	return pc
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (pc *PersonCreate) SetNillableGender(pe *person.Gender) *PersonCreate {
	if pe != nil {
		pc.SetGender(*pe)
	}
	return pc
}

// SetNationality sets the "nationality" field.
func (pc *PersonCreate) SetNationality(s string) *PersonCreate {
	pc.mutation.SetNationality(s)
	return pc
}

// SetNillableNationality sets the "nationality" field if the given value is not nil.
func (pc *PersonCreate) SetNillableNationality(s *string) *PersonCreate {
	if s != nil {
		pc.SetNationality(*s)
	}
	return pc
}

// SetUgentBarcode sets the "ugent_barcode" field.
func (pc *PersonCreate) SetUgentBarcode(s []string) *PersonCreate {
	pc.mutation.SetUgentBarcode(s)
	return pc
}

// SetUgentJobCategory sets the "ugent_job_category" field.
func (pc *PersonCreate) SetUgentJobCategory(s []string) *PersonCreate {
	pc.mutation.SetUgentJobCategory(s)
	return pc
}

// SetTitle sets the "title" field.
func (pc *PersonCreate) SetTitle(s string) *PersonCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (pc *PersonCreate) SetNillableTitle(s *string) *PersonCreate {
	if s != nil {
		pc.SetTitle(*s)
	}
	return pc
}

// SetUgentTel sets the "ugent_tel" field.
func (pc *PersonCreate) SetUgentTel(s string) *PersonCreate {
	pc.mutation.SetUgentTel(s)
	return pc
}

// SetNillableUgentTel sets the "ugent_tel" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentTel(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentTel(*s)
	}
	return pc
}

// SetUgentCampus sets the "ugent_campus" field.
func (pc *PersonCreate) SetUgentCampus(s []string) *PersonCreate {
	pc.mutation.SetUgentCampus(s)
	return pc
}

// SetUgentDepartmentID sets the "ugent_department_id" field.
func (pc *PersonCreate) SetUgentDepartmentID(s []string) *PersonCreate {
	pc.mutation.SetUgentDepartmentID(s)
	return pc
}

// SetUgentFacultyID sets the "ugent_faculty_id" field.
func (pc *PersonCreate) SetUgentFacultyID(s []string) *PersonCreate {
	pc.mutation.SetUgentFacultyID(s)
	return pc
}

// SetUgentJobTitle sets the "ugent_job_title" field.
func (pc *PersonCreate) SetUgentJobTitle(s []string) *PersonCreate {
	pc.mutation.SetUgentJobTitle(s)
	return pc
}

// SetUgentStreetAddress sets the "ugent_street_address" field.
func (pc *PersonCreate) SetUgentStreetAddress(s string) *PersonCreate {
	pc.mutation.SetUgentStreetAddress(s)
	return pc
}

// SetNillableUgentStreetAddress sets the "ugent_street_address" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentStreetAddress(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentStreetAddress(*s)
	}
	return pc
}

// SetUgentPostalCode sets the "ugent_postal_code" field.
func (pc *PersonCreate) SetUgentPostalCode(s string) *PersonCreate {
	pc.mutation.SetUgentPostalCode(s)
	return pc
}

// SetNillableUgentPostalCode sets the "ugent_postal_code" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentPostalCode(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentPostalCode(*s)
	}
	return pc
}

// SetUgentLocality sets the "ugent_locality" field.
func (pc *PersonCreate) SetUgentLocality(s string) *PersonCreate {
	pc.mutation.SetUgentLocality(s)
	return pc
}

// SetNillableUgentLocality sets the "ugent_locality" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentLocality(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentLocality(*s)
	}
	return pc
}

// SetUgentLastEnrolled sets the "ugent_last_enrolled" field.
func (pc *PersonCreate) SetUgentLastEnrolled(s string) *PersonCreate {
	pc.mutation.SetUgentLastEnrolled(s)
	return pc
}

// SetNillableUgentLastEnrolled sets the "ugent_last_enrolled" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentLastEnrolled(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentLastEnrolled(*s)
	}
	return pc
}

// SetHomeStreetAddress sets the "home_street_address" field.
func (pc *PersonCreate) SetHomeStreetAddress(s string) *PersonCreate {
	pc.mutation.SetHomeStreetAddress(s)
	return pc
}

// SetNillableHomeStreetAddress sets the "home_street_address" field if the given value is not nil.
func (pc *PersonCreate) SetNillableHomeStreetAddress(s *string) *PersonCreate {
	if s != nil {
		pc.SetHomeStreetAddress(*s)
	}
	return pc
}

// SetHomePostalCode sets the "home_postal_code" field.
func (pc *PersonCreate) SetHomePostalCode(s string) *PersonCreate {
	pc.mutation.SetHomePostalCode(s)
	return pc
}

// SetNillableHomePostalCode sets the "home_postal_code" field if the given value is not nil.
func (pc *PersonCreate) SetNillableHomePostalCode(s *string) *PersonCreate {
	if s != nil {
		pc.SetHomePostalCode(*s)
	}
	return pc
}

// SetHomeLocality sets the "home_locality" field.
func (pc *PersonCreate) SetHomeLocality(s string) *PersonCreate {
	pc.mutation.SetHomeLocality(s)
	return pc
}

// SetNillableHomeLocality sets the "home_locality" field if the given value is not nil.
func (pc *PersonCreate) SetNillableHomeLocality(s *string) *PersonCreate {
	if s != nil {
		pc.SetHomeLocality(*s)
	}
	return pc
}

// SetHomeCountry sets the "home_country" field.
func (pc *PersonCreate) SetHomeCountry(s string) *PersonCreate {
	pc.mutation.SetHomeCountry(s)
	return pc
}

// SetNillableHomeCountry sets the "home_country" field if the given value is not nil.
func (pc *PersonCreate) SetNillableHomeCountry(s *string) *PersonCreate {
	if s != nil {
		pc.SetHomeCountry(*s)
	}
	return pc
}

// SetHomeTel sets the "home_tel" field.
func (pc *PersonCreate) SetHomeTel(s string) *PersonCreate {
	pc.mutation.SetHomeTel(s)
	return pc
}

// SetNillableHomeTel sets the "home_tel" field if the given value is not nil.
func (pc *PersonCreate) SetNillableHomeTel(s *string) *PersonCreate {
	if s != nil {
		pc.SetHomeTel(*s)
	}
	return pc
}

// SetDormStreetAddress sets the "dorm_street_address" field.
func (pc *PersonCreate) SetDormStreetAddress(s string) *PersonCreate {
	pc.mutation.SetDormStreetAddress(s)
	return pc
}

// SetNillableDormStreetAddress sets the "dorm_street_address" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDormStreetAddress(s *string) *PersonCreate {
	if s != nil {
		pc.SetDormStreetAddress(*s)
	}
	return pc
}

// SetDormPostalCode sets the "dorm_postal_code" field.
func (pc *PersonCreate) SetDormPostalCode(s string) *PersonCreate {
	pc.mutation.SetDormPostalCode(s)
	return pc
}

// SetNillableDormPostalCode sets the "dorm_postal_code" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDormPostalCode(s *string) *PersonCreate {
	if s != nil {
		pc.SetDormPostalCode(*s)
	}
	return pc
}

// SetDormLocality sets the "dorm_locality" field.
func (pc *PersonCreate) SetDormLocality(s string) *PersonCreate {
	pc.mutation.SetDormLocality(s)
	return pc
}

// SetNillableDormLocality sets the "dorm_locality" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDormLocality(s *string) *PersonCreate {
	if s != nil {
		pc.SetDormLocality(*s)
	}
	return pc
}

// SetDormCountry sets the "dorm_country" field.
func (pc *PersonCreate) SetDormCountry(s string) *PersonCreate {
	pc.mutation.SetDormCountry(s)
	return pc
}

// SetNillableDormCountry sets the "dorm_country" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDormCountry(s *string) *PersonCreate {
	if s != nil {
		pc.SetDormCountry(*s)
	}
	return pc
}

// SetResearchDiscipline sets the "research_discipline" field.
func (pc *PersonCreate) SetResearchDiscipline(s []string) *PersonCreate {
	pc.mutation.SetResearchDiscipline(s)
	return pc
}

// SetResearchDisciplineCode sets the "research_discipline_code" field.
func (pc *PersonCreate) SetResearchDisciplineCode(s []string) *PersonCreate {
	pc.mutation.SetResearchDisciplineCode(s)
	return pc
}

// SetUgentExpirationDate sets the "ugent_expiration_date" field.
func (pc *PersonCreate) SetUgentExpirationDate(s string) *PersonCreate {
	pc.mutation.SetUgentExpirationDate(s)
	return pc
}

// SetNillableUgentExpirationDate sets the "ugent_expiration_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentExpirationDate(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentExpirationDate(*s)
	}
	return pc
}

// SetUzgentJobTitle sets the "uzgent_job_title" field.
func (pc *PersonCreate) SetUzgentJobTitle(s []string) *PersonCreate {
	pc.mutation.SetUzgentJobTitle(s)
	return pc
}

// SetUzgentDepartmentName sets the "uzgent_department_name" field.
func (pc *PersonCreate) SetUzgentDepartmentName(s []string) *PersonCreate {
	pc.mutation.SetUzgentDepartmentName(s)
	return pc
}

// SetUzgentID sets the "uzgent_id" field.
func (pc *PersonCreate) SetUzgentID(s []string) *PersonCreate {
	pc.mutation.SetUzgentID(s)
	return pc
}

// SetUgentExtCategory sets the "ugent_ext_category" field.
func (pc *PersonCreate) SetUgentExtCategory(s []string) *PersonCreate {
	pc.mutation.SetUgentExtCategory(s)
	return pc
}

// SetUgentAppointmentDate sets the "ugent_appointment_date" field.
func (pc *PersonCreate) SetUgentAppointmentDate(s string) *PersonCreate {
	pc.mutation.SetUgentAppointmentDate(s)
	return pc
}

// SetNillableUgentAppointmentDate sets the "ugent_appointment_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentAppointmentDate(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentAppointmentDate(*s)
	}
	return pc
}

// SetUgentDepartmentName sets the "ugent_department_name" field.
func (pc *PersonCreate) SetUgentDepartmentName(s []string) *PersonCreate {
	pc.mutation.SetUgentDepartmentName(s)
	return pc
}

// SetOrcidBio sets the "orcid_bio" field.
func (pc *PersonCreate) SetOrcidBio(s string) *PersonCreate {
	pc.mutation.SetOrcidBio(s)
	return pc
}

// SetNillableOrcidBio sets the "orcid_bio" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOrcidBio(s *string) *PersonCreate {
	if s != nil {
		pc.SetOrcidBio(*s)
	}
	return pc
}

// SetOrcidID sets the "orcid_id" field.
func (pc *PersonCreate) SetOrcidID(s string) *PersonCreate {
	pc.mutation.SetOrcidID(s)
	return pc
}

// SetNillableOrcidID sets the "orcid_id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOrcidID(s *string) *PersonCreate {
	if s != nil {
		pc.SetOrcidID(*s)
	}
	return pc
}

// SetOrcidSettings sets the "orcid_settings" field.
func (pc *PersonCreate) SetOrcidSettings(m map[string]interface{}) *PersonCreate {
	pc.mutation.SetOrcidSettings(m)
	return pc
}

// SetOrcidToken sets the "orcid_token" field.
func (pc *PersonCreate) SetOrcidToken(s string) *PersonCreate {
	pc.mutation.SetOrcidToken(s)
	return pc
}

// SetNillableOrcidToken sets the "orcid_token" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOrcidToken(s *string) *PersonCreate {
	if s != nil {
		pc.SetOrcidToken(*s)
	}
	return pc
}

// SetOrcidVerify sets the "orcid_verify" field.
func (pc *PersonCreate) SetOrcidVerify(s string) *PersonCreate {
	pc.mutation.SetOrcidVerify(s)
	return pc
}

// SetNillableOrcidVerify sets the "orcid_verify" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOrcidVerify(s *string) *PersonCreate {
	if s != nil {
		pc.SetOrcidVerify(*s)
	}
	return pc
}

// SetActive sets the "active" field.
func (pc *PersonCreate) SetActive(b bool) *PersonCreate {
	pc.mutation.SetActive(b)
	return pc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (pc *PersonCreate) SetNillableActive(b *bool) *PersonCreate {
	if b != nil {
		pc.SetActive(*b)
	}
	return pc
}

// SetDeleted sets the "deleted" field.
func (pc *PersonCreate) SetDeleted(b bool) *PersonCreate {
	pc.mutation.SetDeleted(b)
	return pc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDeleted(b *bool) *PersonCreate {
	if b != nil {
		pc.SetDeleted(*b)
	}
	return pc
}

// SetSettings sets the "settings" field.
func (pc *PersonCreate) SetSettings(m map[string]interface{}) *PersonCreate {
	pc.mutation.SetSettings(m)
	return pc
}

// SetRoles sets the "roles" field.
func (pc *PersonCreate) SetRoles(s []string) *PersonCreate {
	pc.mutation.SetRoles(s)
	return pc
}

// SetPublicationCount sets the "publication_count" field.
func (pc *PersonCreate) SetPublicationCount(i int) *PersonCreate {
	pc.mutation.SetPublicationCount(i)
	return pc
}

// SetNillablePublicationCount sets the "publication_count" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePublicationCount(i *int) *PersonCreate {
	if i != nil {
		pc.SetPublicationCount(*i)
	}
	return pc
}

// SetUgentMemorialisID sets the "ugent_memorialis_id" field.
func (pc *PersonCreate) SetUgentMemorialisID(s string) *PersonCreate {
	pc.mutation.SetUgentMemorialisID(s)
	return pc
}

// SetNillableUgentMemorialisID sets the "ugent_memorialis_id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableUgentMemorialisID(s *string) *PersonCreate {
	if s != nil {
		pc.SetUgentMemorialisID(*s)
	}
	return pc
}

// SetPreferredFirstName sets the "preferred_first_name" field.
func (pc *PersonCreate) SetPreferredFirstName(s string) *PersonCreate {
	pc.mutation.SetPreferredFirstName(s)
	return pc
}

// SetNillablePreferredFirstName sets the "preferred_first_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePreferredFirstName(s *string) *PersonCreate {
	if s != nil {
		pc.SetPreferredFirstName(*s)
	}
	return pc
}

// SetPreferredLastName sets the "preferred_last_name" field.
func (pc *PersonCreate) SetPreferredLastName(s string) *PersonCreate {
	pc.mutation.SetPreferredLastName(s)
	return pc
}

// SetNillablePreferredLastName sets the "preferred_last_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePreferredLastName(s *string) *PersonCreate {
	if s != nil {
		pc.SetPreferredLastName(*s)
	}
	return pc
}

// SetReplaces sets the "replaces" field.
func (pc *PersonCreate) SetReplaces(m []map[string]string) *PersonCreate {
	pc.mutation.SetReplaces(m)
	return pc
}

// SetReplacedBy sets the "replaced_by" field.
func (pc *PersonCreate) SetReplacedBy(m []map[string]string) *PersonCreate {
	pc.mutation.SetReplacedBy(m)
	return pc
}

// SetDateLastLogin sets the "date_last_login" field.
func (pc *PersonCreate) SetDateLastLogin(t time.Time) *PersonCreate {
	pc.mutation.SetDateLastLogin(t)
	return pc
}

// SetNillableDateLastLogin sets the "date_last_login" field if the given value is not nil.
func (pc *PersonCreate) SetNillableDateLastLogin(t *time.Time) *PersonCreate {
	if t != nil {
		pc.SetDateLastLogin(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PersonCreate) SetID(s string) *PersonCreate {
	pc.mutation.SetID(s)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableID(s *string) *PersonCreate {
	if s != nil {
		pc.SetID(*s)
	}
	return pc
}

// Mutation returns the PersonMutation object of the builder.
func (pc *PersonCreate) Mutation() *PersonMutation {
	return pc.mutation
}

// Save creates the Person in the database.
func (pc *PersonCreate) Save(ctx context.Context) (*Person, error) {
	pc.defaults()
	return withHooks[*Person, PersonMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PersonCreate) SaveX(ctx context.Context) *Person {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PersonCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PersonCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PersonCreate) defaults() {
	if _, ok := pc.mutation.DateCreated(); !ok {
		v := person.DefaultDateCreated()
		pc.mutation.SetDateCreated(v)
	}
	if _, ok := pc.mutation.DateUpdated(); !ok {
		v := person.DefaultDateUpdated()
		pc.mutation.SetDateUpdated(v)
	}
	if _, ok := pc.mutation.Active(); !ok {
		v := person.DefaultActive
		pc.mutation.SetActive(v)
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		v := person.DefaultDeleted
		pc.mutation.SetDeleted(v)
	}
	if _, ok := pc.mutation.PublicationCount(); !ok {
		v := person.DefaultPublicationCount
		pc.mutation.SetPublicationCount(v)
	}
	if _, ok := pc.mutation.ID(); !ok {
		v := person.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PersonCreate) check() error {
	if _, ok := pc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "Person.date_created"`)}
	}
	if _, ok := pc.mutation.DateUpdated(); !ok {
		return &ValidationError{Name: "date_updated", err: errors.New(`ent: missing required field "Person.date_updated"`)}
	}
	if v, ok := pc.mutation.Gender(); ok {
		if err := person.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "Person.gender": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Person.active"`)}
	}
	if _, ok := pc.mutation.Deleted(); !ok {
		return &ValidationError{Name: "deleted", err: errors.New(`ent: missing required field "Person.deleted"`)}
	}
	return nil
}

func (pc *PersonCreate) sqlSave(ctx context.Context) (*Person, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Person.ID type: %T", _spec.ID.Value)
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(person.Table, sqlgraph.NewFieldSpec(person.FieldID, field.TypeString))
	)
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.DateCreated(); ok {
		_spec.SetField(person.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := pc.mutation.DateUpdated(); ok {
		_spec.SetField(person.FieldDateUpdated, field.TypeTime, value)
		_node.DateUpdated = value
	}
	if value, ok := pc.mutation.ObjectClass(); ok {
		_spec.SetField(person.FieldObjectClass, field.TypeString, value)
		_node.ObjectClass = value
	}
	if value, ok := pc.mutation.UgentUsername(); ok {
		_spec.SetField(person.FieldUgentUsername, field.TypeString, value)
		_node.UgentUsername = value
	}
	if value, ok := pc.mutation.FirstName(); ok {
		_spec.SetField(person.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := pc.mutation.MiddleName(); ok {
		_spec.SetField(person.FieldMiddleName, field.TypeString, value)
		_node.MiddleName = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.SetField(person.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := pc.mutation.UgentID(); ok {
		_spec.SetField(person.FieldUgentID, field.TypeJSON, value)
		_node.UgentID = value
	}
	if value, ok := pc.mutation.BirthDate(); ok {
		_spec.SetField(person.FieldBirthDate, field.TypeString, value)
		_node.BirthDate = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := pc.mutation.Gender(); ok {
		_spec.SetField(person.FieldGender, field.TypeEnum, value)
		_node.Gender = value
	}
	if value, ok := pc.mutation.Nationality(); ok {
		_spec.SetField(person.FieldNationality, field.TypeString, value)
		_node.Nationality = value
	}
	if value, ok := pc.mutation.UgentBarcode(); ok {
		_spec.SetField(person.FieldUgentBarcode, field.TypeJSON, value)
		_node.UgentBarcode = value
	}
	if value, ok := pc.mutation.UgentJobCategory(); ok {
		_spec.SetField(person.FieldUgentJobCategory, field.TypeJSON, value)
		_node.UgentJobCategory = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(person.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.UgentTel(); ok {
		_spec.SetField(person.FieldUgentTel, field.TypeString, value)
		_node.UgentTel = value
	}
	if value, ok := pc.mutation.UgentCampus(); ok {
		_spec.SetField(person.FieldUgentCampus, field.TypeJSON, value)
		_node.UgentCampus = value
	}
	if value, ok := pc.mutation.UgentDepartmentID(); ok {
		_spec.SetField(person.FieldUgentDepartmentID, field.TypeJSON, value)
		_node.UgentDepartmentID = value
	}
	if value, ok := pc.mutation.UgentFacultyID(); ok {
		_spec.SetField(person.FieldUgentFacultyID, field.TypeJSON, value)
		_node.UgentFacultyID = value
	}
	if value, ok := pc.mutation.UgentJobTitle(); ok {
		_spec.SetField(person.FieldUgentJobTitle, field.TypeJSON, value)
		_node.UgentJobTitle = value
	}
	if value, ok := pc.mutation.UgentStreetAddress(); ok {
		_spec.SetField(person.FieldUgentStreetAddress, field.TypeString, value)
		_node.UgentStreetAddress = value
	}
	if value, ok := pc.mutation.UgentPostalCode(); ok {
		_spec.SetField(person.FieldUgentPostalCode, field.TypeString, value)
		_node.UgentPostalCode = value
	}
	if value, ok := pc.mutation.UgentLocality(); ok {
		_spec.SetField(person.FieldUgentLocality, field.TypeString, value)
		_node.UgentLocality = value
	}
	if value, ok := pc.mutation.UgentLastEnrolled(); ok {
		_spec.SetField(person.FieldUgentLastEnrolled, field.TypeString, value)
		_node.UgentLastEnrolled = value
	}
	if value, ok := pc.mutation.HomeStreetAddress(); ok {
		_spec.SetField(person.FieldHomeStreetAddress, field.TypeString, value)
		_node.HomeStreetAddress = value
	}
	if value, ok := pc.mutation.HomePostalCode(); ok {
		_spec.SetField(person.FieldHomePostalCode, field.TypeString, value)
		_node.HomePostalCode = value
	}
	if value, ok := pc.mutation.HomeLocality(); ok {
		_spec.SetField(person.FieldHomeLocality, field.TypeString, value)
		_node.HomeLocality = value
	}
	if value, ok := pc.mutation.HomeCountry(); ok {
		_spec.SetField(person.FieldHomeCountry, field.TypeString, value)
		_node.HomeCountry = value
	}
	if value, ok := pc.mutation.HomeTel(); ok {
		_spec.SetField(person.FieldHomeTel, field.TypeString, value)
		_node.HomeTel = value
	}
	if value, ok := pc.mutation.DormStreetAddress(); ok {
		_spec.SetField(person.FieldDormStreetAddress, field.TypeString, value)
		_node.DormStreetAddress = value
	}
	if value, ok := pc.mutation.DormPostalCode(); ok {
		_spec.SetField(person.FieldDormPostalCode, field.TypeString, value)
		_node.DormPostalCode = value
	}
	if value, ok := pc.mutation.DormLocality(); ok {
		_spec.SetField(person.FieldDormLocality, field.TypeString, value)
		_node.DormLocality = value
	}
	if value, ok := pc.mutation.DormCountry(); ok {
		_spec.SetField(person.FieldDormCountry, field.TypeString, value)
		_node.DormCountry = value
	}
	if value, ok := pc.mutation.ResearchDiscipline(); ok {
		_spec.SetField(person.FieldResearchDiscipline, field.TypeJSON, value)
		_node.ResearchDiscipline = value
	}
	if value, ok := pc.mutation.ResearchDisciplineCode(); ok {
		_spec.SetField(person.FieldResearchDisciplineCode, field.TypeJSON, value)
		_node.ResearchDisciplineCode = value
	}
	if value, ok := pc.mutation.UgentExpirationDate(); ok {
		_spec.SetField(person.FieldUgentExpirationDate, field.TypeString, value)
		_node.UgentExpirationDate = value
	}
	if value, ok := pc.mutation.UzgentJobTitle(); ok {
		_spec.SetField(person.FieldUzgentJobTitle, field.TypeJSON, value)
		_node.UzgentJobTitle = value
	}
	if value, ok := pc.mutation.UzgentDepartmentName(); ok {
		_spec.SetField(person.FieldUzgentDepartmentName, field.TypeJSON, value)
		_node.UzgentDepartmentName = value
	}
	if value, ok := pc.mutation.UzgentID(); ok {
		_spec.SetField(person.FieldUzgentID, field.TypeJSON, value)
		_node.UzgentID = value
	}
	if value, ok := pc.mutation.UgentExtCategory(); ok {
		_spec.SetField(person.FieldUgentExtCategory, field.TypeJSON, value)
		_node.UgentExtCategory = value
	}
	if value, ok := pc.mutation.UgentAppointmentDate(); ok {
		_spec.SetField(person.FieldUgentAppointmentDate, field.TypeString, value)
		_node.UgentAppointmentDate = value
	}
	if value, ok := pc.mutation.UgentDepartmentName(); ok {
		_spec.SetField(person.FieldUgentDepartmentName, field.TypeJSON, value)
		_node.UgentDepartmentName = value
	}
	if value, ok := pc.mutation.OrcidBio(); ok {
		_spec.SetField(person.FieldOrcidBio, field.TypeString, value)
		_node.OrcidBio = value
	}
	if value, ok := pc.mutation.OrcidID(); ok {
		_spec.SetField(person.FieldOrcidID, field.TypeString, value)
		_node.OrcidID = value
	}
	if value, ok := pc.mutation.OrcidSettings(); ok {
		_spec.SetField(person.FieldOrcidSettings, field.TypeJSON, value)
		_node.OrcidSettings = value
	}
	if value, ok := pc.mutation.OrcidToken(); ok {
		_spec.SetField(person.FieldOrcidToken, field.TypeString, value)
		_node.OrcidToken = value
	}
	if value, ok := pc.mutation.OrcidVerify(); ok {
		_spec.SetField(person.FieldOrcidVerify, field.TypeString, value)
		_node.OrcidVerify = value
	}
	if value, ok := pc.mutation.Active(); ok {
		_spec.SetField(person.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := pc.mutation.Deleted(); ok {
		_spec.SetField(person.FieldDeleted, field.TypeBool, value)
		_node.Deleted = value
	}
	if value, ok := pc.mutation.Settings(); ok {
		_spec.SetField(person.FieldSettings, field.TypeJSON, value)
		_node.Settings = value
	}
	if value, ok := pc.mutation.Roles(); ok {
		_spec.SetField(person.FieldRoles, field.TypeJSON, value)
		_node.Roles = value
	}
	if value, ok := pc.mutation.PublicationCount(); ok {
		_spec.SetField(person.FieldPublicationCount, field.TypeInt, value)
		_node.PublicationCount = value
	}
	if value, ok := pc.mutation.UgentMemorialisID(); ok {
		_spec.SetField(person.FieldUgentMemorialisID, field.TypeString, value)
		_node.UgentMemorialisID = value
	}
	if value, ok := pc.mutation.PreferredFirstName(); ok {
		_spec.SetField(person.FieldPreferredFirstName, field.TypeString, value)
		_node.PreferredFirstName = value
	}
	if value, ok := pc.mutation.PreferredLastName(); ok {
		_spec.SetField(person.FieldPreferredLastName, field.TypeString, value)
		_node.PreferredLastName = value
	}
	if value, ok := pc.mutation.Replaces(); ok {
		_spec.SetField(person.FieldReplaces, field.TypeJSON, value)
		_node.Replaces = value
	}
	if value, ok := pc.mutation.ReplacedBy(); ok {
		_spec.SetField(person.FieldReplacedBy, field.TypeJSON, value)
		_node.ReplacedBy = value
	}
	if value, ok := pc.mutation.DateLastLogin(); ok {
		_spec.SetField(person.FieldDateLastLogin, field.TypeTime, value)
		_node.DateLastLogin = value
	}
	return _node, _spec
}

// PersonCreateBulk is the builder for creating many Person entities in bulk.
type PersonCreateBulk struct {
	config
	builders []*PersonCreate
}

// Save creates the Person entities in the database.
func (pcb *PersonCreateBulk) Save(ctx context.Context) ([]*Person, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Person, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PersonMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PersonCreateBulk) SaveX(ctx context.Context) []*Person {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PersonCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PersonCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}
