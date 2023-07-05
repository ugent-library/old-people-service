// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/person-service/ent/organization"
	"github.com/ugent-library/person-service/ent/organizationperson"
	"github.com/ugent-library/person-service/ent/person"
	"github.com/ugent-library/person-service/ent/schema"
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

// SetPublicID sets the "public_id" field.
func (pc *PersonCreate) SetPublicID(s string) *PersonCreate {
	pc.mutation.SetPublicID(s)
	return pc
}

// SetNillablePublicID sets the "public_id" field if the given value is not nil.
func (pc *PersonCreate) SetNillablePublicID(s *string) *PersonCreate {
	if s != nil {
		pc.SetPublicID(*s)
	}
	return pc
}

// SetGismoID sets the "gismo_id" field.
func (pc *PersonCreate) SetGismoID(s string) *PersonCreate {
	pc.mutation.SetGismoID(s)
	return pc
}

// SetNillableGismoID sets the "gismo_id" field if the given value is not nil.
func (pc *PersonCreate) SetNillableGismoID(s *string) *PersonCreate {
	if s != nil {
		pc.SetGismoID(*s)
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

// SetOtherID sets the "other_id" field.
func (pc *PersonCreate) SetOtherID(sr []schema.IdRef) *PersonCreate {
	pc.mutation.SetOtherID(sr)
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

// SetFullName sets the "full_name" field.
func (pc *PersonCreate) SetFullName(s string) *PersonCreate {
	pc.mutation.SetFullName(s)
	return pc
}

// SetNillableFullName sets the "full_name" field if the given value is not nil.
func (pc *PersonCreate) SetNillableFullName(s *string) *PersonCreate {
	if s != nil {
		pc.SetFullName(*s)
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

// SetJobCategory sets the "job_category" field.
func (pc *PersonCreate) SetJobCategory(s []string) *PersonCreate {
	pc.mutation.SetJobCategory(s)
	return pc
}

// SetOrcid sets the "orcid" field.
func (pc *PersonCreate) SetOrcid(s string) *PersonCreate {
	pc.mutation.SetOrcid(s)
	return pc
}

// SetNillableOrcid sets the "orcid" field if the given value is not nil.
func (pc *PersonCreate) SetNillableOrcid(s *string) *PersonCreate {
	if s != nil {
		pc.SetOrcid(*s)
	}
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

// SetRole sets the "role" field.
func (pc *PersonCreate) SetRole(s []string) *PersonCreate {
	pc.mutation.SetRole(s)
	return pc
}

// SetSettings sets the "settings" field.
func (pc *PersonCreate) SetSettings(m map[string]string) *PersonCreate {
	pc.mutation.SetSettings(m)
	return pc
}

// SetObjectClass sets the "object_class" field.
func (pc *PersonCreate) SetObjectClass(s []string) *PersonCreate {
	pc.mutation.SetObjectClass(s)
	return pc
}

// SetExpirationDate sets the "expiration_date" field.
func (pc *PersonCreate) SetExpirationDate(s string) *PersonCreate {
	pc.mutation.SetExpirationDate(s)
	return pc
}

// SetNillableExpirationDate sets the "expiration_date" field if the given value is not nil.
func (pc *PersonCreate) SetNillableExpirationDate(s *string) *PersonCreate {
	if s != nil {
		pc.SetExpirationDate(*s)
	}
	return pc
}

// AddOrganizationIDs adds the "organizations" edge to the Organization entity by IDs.
func (pc *PersonCreate) AddOrganizationIDs(ids ...int) *PersonCreate {
	pc.mutation.AddOrganizationIDs(ids...)
	return pc
}

// AddOrganizations adds the "organizations" edges to the Organization entity.
func (pc *PersonCreate) AddOrganizations(o ...*Organization) *PersonCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrganizationIDs(ids...)
}

// AddOrganizationPersonIDs adds the "organization_person" edge to the OrganizationPerson entity by IDs.
func (pc *PersonCreate) AddOrganizationPersonIDs(ids ...int) *PersonCreate {
	pc.mutation.AddOrganizationPersonIDs(ids...)
	return pc
}

// AddOrganizationPerson adds the "organization_person" edges to the OrganizationPerson entity.
func (pc *PersonCreate) AddOrganizationPerson(o ...*OrganizationPerson) *PersonCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrganizationPersonIDs(ids...)
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
	if _, ok := pc.mutation.PublicID(); !ok {
		v := person.DefaultPublicID()
		pc.mutation.SetPublicID(v)
	}
	if _, ok := pc.mutation.Active(); !ok {
		v := person.DefaultActive
		pc.mutation.SetActive(v)
	}
	if _, ok := pc.mutation.OtherID(); !ok {
		v := person.DefaultOtherID
		pc.mutation.SetOtherID(v)
	}
	if _, ok := pc.mutation.JobCategory(); !ok {
		v := person.DefaultJobCategory
		pc.mutation.SetJobCategory(v)
	}
	if _, ok := pc.mutation.Role(); !ok {
		v := person.DefaultRole
		pc.mutation.SetRole(v)
	}
	if _, ok := pc.mutation.Settings(); !ok {
		v := person.DefaultSettings
		pc.mutation.SetSettings(v)
	}
	if _, ok := pc.mutation.ObjectClass(); !ok {
		v := person.DefaultObjectClass
		pc.mutation.SetObjectClass(v)
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
	if _, ok := pc.mutation.PublicID(); !ok {
		return &ValidationError{Name: "public_id", err: errors.New(`ent: missing required field "Person.public_id"`)}
	}
	if _, ok := pc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Person.active"`)}
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
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PersonCreate) createSpec() (*Person, *sqlgraph.CreateSpec) {
	var (
		_node = &Person{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(person.Table, sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt))
	)
	if value, ok := pc.mutation.DateCreated(); ok {
		_spec.SetField(person.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := pc.mutation.DateUpdated(); ok {
		_spec.SetField(person.FieldDateUpdated, field.TypeTime, value)
		_node.DateUpdated = value
	}
	if value, ok := pc.mutation.PublicID(); ok {
		_spec.SetField(person.FieldPublicID, field.TypeString, value)
		_node.PublicID = value
	}
	if value, ok := pc.mutation.GismoID(); ok {
		_spec.SetField(person.FieldGismoID, field.TypeString, value)
		_node.GismoID = &value
	}
	if value, ok := pc.mutation.Active(); ok {
		_spec.SetField(person.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := pc.mutation.BirthDate(); ok {
		_spec.SetField(person.FieldBirthDate, field.TypeString, value)
		_node.BirthDate = value
	}
	if value, ok := pc.mutation.Email(); ok {
		_spec.SetField(person.FieldEmail, field.TypeString, value)
		_node.Email = value
	}
	if value, ok := pc.mutation.OtherID(); ok {
		_spec.SetField(person.FieldOtherID, field.TypeJSON, value)
		_node.OtherID = value
	}
	if value, ok := pc.mutation.FirstName(); ok {
		_spec.SetField(person.FieldFirstName, field.TypeString, value)
		_node.FirstName = value
	}
	if value, ok := pc.mutation.FullName(); ok {
		_spec.SetField(person.FieldFullName, field.TypeString, value)
		_node.FullName = value
	}
	if value, ok := pc.mutation.LastName(); ok {
		_spec.SetField(person.FieldLastName, field.TypeString, value)
		_node.LastName = value
	}
	if value, ok := pc.mutation.JobCategory(); ok {
		_spec.SetField(person.FieldJobCategory, field.TypeJSON, value)
		_node.JobCategory = value
	}
	if value, ok := pc.mutation.Orcid(); ok {
		_spec.SetField(person.FieldOrcid, field.TypeString, value)
		_node.Orcid = value
	}
	if value, ok := pc.mutation.OrcidToken(); ok {
		_spec.SetField(person.FieldOrcidToken, field.TypeString, value)
		_node.OrcidToken = value
	}
	if value, ok := pc.mutation.PreferredFirstName(); ok {
		_spec.SetField(person.FieldPreferredFirstName, field.TypeString, value)
		_node.PreferredFirstName = value
	}
	if value, ok := pc.mutation.PreferredLastName(); ok {
		_spec.SetField(person.FieldPreferredLastName, field.TypeString, value)
		_node.PreferredLastName = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(person.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Role(); ok {
		_spec.SetField(person.FieldRole, field.TypeJSON, value)
		_node.Role = value
	}
	if value, ok := pc.mutation.Settings(); ok {
		_spec.SetField(person.FieldSettings, field.TypeJSON, value)
		_node.Settings = value
	}
	if value, ok := pc.mutation.ObjectClass(); ok {
		_spec.SetField(person.FieldObjectClass, field.TypeJSON, value)
		_node.ObjectClass = value
	}
	if value, ok := pc.mutation.ExpirationDate(); ok {
		_spec.SetField(person.FieldExpirationDate, field.TypeString, value)
		_node.ExpirationDate = value
	}
	if nodes := pc.mutation.OrganizationsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   person.OrganizationsTable,
			Columns: person.OrganizationsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrganizationPersonCreate{config: pc.config, mutation: newOrganizationPersonMutation(pc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OrganizationPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   person.OrganizationPersonTable,
			Columns: []string{person.OrganizationPersonColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(organizationperson.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
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
				var err error
				nodes[i], specs[i] = builder.createSpec()
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
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
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
