// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people-service/ent/organization"
	"github.com/ugent-library/people-service/ent/organizationperson"
	"github.com/ugent-library/people-service/ent/person"
	"github.com/ugent-library/people-service/ent/schema"
)

// OrganizationCreate is the builder for creating a Organization entity.
type OrganizationCreate struct {
	config
	mutation *OrganizationMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (oc *OrganizationCreate) SetDateCreated(t time.Time) *OrganizationCreate {
	oc.mutation.SetDateCreated(t)
	return oc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDateCreated(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetDateCreated(*t)
	}
	return oc
}

// SetDateUpdated sets the "date_updated" field.
func (oc *OrganizationCreate) SetDateUpdated(t time.Time) *OrganizationCreate {
	oc.mutation.SetDateUpdated(t)
	return oc
}

// SetNillableDateUpdated sets the "date_updated" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableDateUpdated(t *time.Time) *OrganizationCreate {
	if t != nil {
		oc.SetDateUpdated(*t)
	}
	return oc
}

// SetPublicID sets the "public_id" field.
func (oc *OrganizationCreate) SetPublicID(s string) *OrganizationCreate {
	oc.mutation.SetPublicID(s)
	return oc
}

// SetNillablePublicID sets the "public_id" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillablePublicID(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetPublicID(*s)
	}
	return oc
}

// SetType sets the "type" field.
func (oc *OrganizationCreate) SetType(s string) *OrganizationCreate {
	oc.mutation.SetType(s)
	return oc
}

// SetNillableType sets the "type" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableType(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetType(*s)
	}
	return oc
}

// SetNameDut sets the "name_dut" field.
func (oc *OrganizationCreate) SetNameDut(s string) *OrganizationCreate {
	oc.mutation.SetNameDut(s)
	return oc
}

// SetNillableNameDut sets the "name_dut" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableNameDut(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetNameDut(*s)
	}
	return oc
}

// SetNameEng sets the "name_eng" field.
func (oc *OrganizationCreate) SetNameEng(s string) *OrganizationCreate {
	oc.mutation.SetNameEng(s)
	return oc
}

// SetNillableNameEng sets the "name_eng" field if the given value is not nil.
func (oc *OrganizationCreate) SetNillableNameEng(s *string) *OrganizationCreate {
	if s != nil {
		oc.SetNameEng(*s)
	}
	return oc
}

// SetIdentifier sets the "identifier" field.
func (oc *OrganizationCreate) SetIdentifier(sv schema.TypeVals) *OrganizationCreate {
	oc.mutation.SetIdentifier(sv)
	return oc
}

// AddPersonIDs adds the "people" edge to the Person entity by IDs.
func (oc *OrganizationCreate) AddPersonIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddPersonIDs(ids...)
	return oc
}

// AddPeople adds the "people" edges to the Person entity.
func (oc *OrganizationCreate) AddPeople(p ...*Person) *OrganizationCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return oc.AddPersonIDs(ids...)
}

// AddOrganizationPersonIDs adds the "organization_person" edge to the OrganizationPerson entity by IDs.
func (oc *OrganizationCreate) AddOrganizationPersonIDs(ids ...int) *OrganizationCreate {
	oc.mutation.AddOrganizationPersonIDs(ids...)
	return oc
}

// AddOrganizationPerson adds the "organization_person" edges to the OrganizationPerson entity.
func (oc *OrganizationCreate) AddOrganizationPerson(o ...*OrganizationPerson) *OrganizationCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return oc.AddOrganizationPersonIDs(ids...)
}

// Mutation returns the OrganizationMutation object of the builder.
func (oc *OrganizationCreate) Mutation() *OrganizationMutation {
	return oc.mutation
}

// Save creates the Organization in the database.
func (oc *OrganizationCreate) Save(ctx context.Context) (*Organization, error) {
	oc.defaults()
	return withHooks[*Organization, OrganizationMutation](ctx, oc.sqlSave, oc.mutation, oc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrganizationCreate) SaveX(ctx context.Context) *Organization {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (oc *OrganizationCreate) Exec(ctx context.Context) error {
	_, err := oc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (oc *OrganizationCreate) ExecX(ctx context.Context) {
	if err := oc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (oc *OrganizationCreate) defaults() {
	if _, ok := oc.mutation.DateCreated(); !ok {
		v := organization.DefaultDateCreated()
		oc.mutation.SetDateCreated(v)
	}
	if _, ok := oc.mutation.DateUpdated(); !ok {
		v := organization.DefaultDateUpdated()
		oc.mutation.SetDateUpdated(v)
	}
	if _, ok := oc.mutation.PublicID(); !ok {
		v := organization.DefaultPublicID()
		oc.mutation.SetPublicID(v)
	}
	if _, ok := oc.mutation.GetType(); !ok {
		v := organization.DefaultType
		oc.mutation.SetType(v)
	}
	if _, ok := oc.mutation.Identifier(); !ok {
		v := organization.DefaultIdentifier
		oc.mutation.SetIdentifier(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (oc *OrganizationCreate) check() error {
	if _, ok := oc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "Organization.date_created"`)}
	}
	if _, ok := oc.mutation.DateUpdated(); !ok {
		return &ValidationError{Name: "date_updated", err: errors.New(`ent: missing required field "Organization.date_updated"`)}
	}
	if _, ok := oc.mutation.PublicID(); !ok {
		return &ValidationError{Name: "public_id", err: errors.New(`ent: missing required field "Organization.public_id"`)}
	}
	if _, ok := oc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "Organization.type"`)}
	}
	return nil
}

func (oc *OrganizationCreate) sqlSave(ctx context.Context) (*Organization, error) {
	if err := oc.check(); err != nil {
		return nil, err
	}
	_node, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	oc.mutation.id = &_node.ID
	oc.mutation.done = true
	return _node, nil
}

func (oc *OrganizationCreate) createSpec() (*Organization, *sqlgraph.CreateSpec) {
	var (
		_node = &Organization{config: oc.config}
		_spec = sqlgraph.NewCreateSpec(organization.Table, sqlgraph.NewFieldSpec(organization.FieldID, field.TypeInt))
	)
	if value, ok := oc.mutation.DateCreated(); ok {
		_spec.SetField(organization.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := oc.mutation.DateUpdated(); ok {
		_spec.SetField(organization.FieldDateUpdated, field.TypeTime, value)
		_node.DateUpdated = value
	}
	if value, ok := oc.mutation.PublicID(); ok {
		_spec.SetField(organization.FieldPublicID, field.TypeString, value)
		_node.PublicID = value
	}
	if value, ok := oc.mutation.GetType(); ok {
		_spec.SetField(organization.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := oc.mutation.NameDut(); ok {
		_spec.SetField(organization.FieldNameDut, field.TypeString, value)
		_node.NameDut = value
	}
	if value, ok := oc.mutation.NameEng(); ok {
		_spec.SetField(organization.FieldNameEng, field.TypeString, value)
		_node.NameEng = value
	}
	if value, ok := oc.mutation.Identifier(); ok {
		_spec.SetField(organization.FieldIdentifier, field.TypeJSON, value)
		_node.Identifier = value
	}
	if nodes := oc.mutation.PeopleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   organization.PeopleTable,
			Columns: organization.PeoplePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(person.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &OrganizationPersonCreate{config: oc.config, mutation: newOrganizationPersonMutation(oc.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.OrganizationPersonIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   organization.OrganizationPersonTable,
			Columns: []string{organization.OrganizationPersonColumn},
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

// OrganizationCreateBulk is the builder for creating many Organization entities in bulk.
type OrganizationCreateBulk struct {
	config
	builders []*OrganizationCreate
}

// Save creates the Organization entities in the database.
func (ocb *OrganizationCreateBulk) Save(ctx context.Context) ([]*Organization, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ocb.builders))
	nodes := make([]*Organization, len(ocb.builders))
	mutators := make([]Mutator, len(ocb.builders))
	for i := range ocb.builders {
		func(i int, root context.Context) {
			builder := ocb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationMutation)
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
					_, err = mutators[i+1].Mutate(root, ocb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ocb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ocb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) SaveX(ctx context.Context) []*Organization {
	v, err := ocb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ocb *OrganizationCreateBulk) Exec(ctx context.Context) error {
	_, err := ocb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ocb *OrganizationCreateBulk) ExecX(ctx context.Context) {
	if err := ocb.Exec(ctx); err != nil {
		panic(err)
	}
}
