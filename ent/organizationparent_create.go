// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people-service/ent/organizationparent"
)

// OrganizationParentCreate is the builder for creating a OrganizationParent entity.
type OrganizationParentCreate struct {
	config
	mutation *OrganizationParentMutation
	hooks    []Hook
}

// SetDateCreated sets the "date_created" field.
func (opc *OrganizationParentCreate) SetDateCreated(t time.Time) *OrganizationParentCreate {
	opc.mutation.SetDateCreated(t)
	return opc
}

// SetNillableDateCreated sets the "date_created" field if the given value is not nil.
func (opc *OrganizationParentCreate) SetNillableDateCreated(t *time.Time) *OrganizationParentCreate {
	if t != nil {
		opc.SetDateCreated(*t)
	}
	return opc
}

// SetDateUpdated sets the "date_updated" field.
func (opc *OrganizationParentCreate) SetDateUpdated(t time.Time) *OrganizationParentCreate {
	opc.mutation.SetDateUpdated(t)
	return opc
}

// SetNillableDateUpdated sets the "date_updated" field if the given value is not nil.
func (opc *OrganizationParentCreate) SetNillableDateUpdated(t *time.Time) *OrganizationParentCreate {
	if t != nil {
		opc.SetDateUpdated(*t)
	}
	return opc
}

// SetParentOrganizationID sets the "parent_organization_id" field.
func (opc *OrganizationParentCreate) SetParentOrganizationID(i int) *OrganizationParentCreate {
	opc.mutation.SetParentOrganizationID(i)
	return opc
}

// SetOrganizationID sets the "organization_id" field.
func (opc *OrganizationParentCreate) SetOrganizationID(i int) *OrganizationParentCreate {
	opc.mutation.SetOrganizationID(i)
	return opc
}

// SetFrom sets the "from" field.
func (opc *OrganizationParentCreate) SetFrom(t time.Time) *OrganizationParentCreate {
	opc.mutation.SetFrom(t)
	return opc
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (opc *OrganizationParentCreate) SetNillableFrom(t *time.Time) *OrganizationParentCreate {
	if t != nil {
		opc.SetFrom(*t)
	}
	return opc
}

// SetUntil sets the "until" field.
func (opc *OrganizationParentCreate) SetUntil(t time.Time) *OrganizationParentCreate {
	opc.mutation.SetUntil(t)
	return opc
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (opc *OrganizationParentCreate) SetNillableUntil(t *time.Time) *OrganizationParentCreate {
	if t != nil {
		opc.SetUntil(*t)
	}
	return opc
}

// Mutation returns the OrganizationParentMutation object of the builder.
func (opc *OrganizationParentCreate) Mutation() *OrganizationParentMutation {
	return opc.mutation
}

// Save creates the OrganizationParent in the database.
func (opc *OrganizationParentCreate) Save(ctx context.Context) (*OrganizationParent, error) {
	opc.defaults()
	return withHooks[*OrganizationParent, OrganizationParentMutation](ctx, opc.sqlSave, opc.mutation, opc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (opc *OrganizationParentCreate) SaveX(ctx context.Context) *OrganizationParent {
	v, err := opc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opc *OrganizationParentCreate) Exec(ctx context.Context) error {
	_, err := opc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opc *OrganizationParentCreate) ExecX(ctx context.Context) {
	if err := opc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opc *OrganizationParentCreate) defaults() {
	if _, ok := opc.mutation.DateCreated(); !ok {
		v := organizationparent.DefaultDateCreated()
		opc.mutation.SetDateCreated(v)
	}
	if _, ok := opc.mutation.DateUpdated(); !ok {
		v := organizationparent.DefaultDateUpdated()
		opc.mutation.SetDateUpdated(v)
	}
	if _, ok := opc.mutation.From(); !ok {
		v := organizationparent.DefaultFrom()
		opc.mutation.SetFrom(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (opc *OrganizationParentCreate) check() error {
	if _, ok := opc.mutation.DateCreated(); !ok {
		return &ValidationError{Name: "date_created", err: errors.New(`ent: missing required field "OrganizationParent.date_created"`)}
	}
	if _, ok := opc.mutation.DateUpdated(); !ok {
		return &ValidationError{Name: "date_updated", err: errors.New(`ent: missing required field "OrganizationParent.date_updated"`)}
	}
	if _, ok := opc.mutation.ParentOrganizationID(); !ok {
		return &ValidationError{Name: "parent_organization_id", err: errors.New(`ent: missing required field "OrganizationParent.parent_organization_id"`)}
	}
	if _, ok := opc.mutation.OrganizationID(); !ok {
		return &ValidationError{Name: "organization_id", err: errors.New(`ent: missing required field "OrganizationParent.organization_id"`)}
	}
	if _, ok := opc.mutation.From(); !ok {
		return &ValidationError{Name: "from", err: errors.New(`ent: missing required field "OrganizationParent.from"`)}
	}
	return nil
}

func (opc *OrganizationParentCreate) sqlSave(ctx context.Context) (*OrganizationParent, error) {
	if err := opc.check(); err != nil {
		return nil, err
	}
	_node, _spec := opc.createSpec()
	if err := sqlgraph.CreateNode(ctx, opc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	opc.mutation.id = &_node.ID
	opc.mutation.done = true
	return _node, nil
}

func (opc *OrganizationParentCreate) createSpec() (*OrganizationParent, *sqlgraph.CreateSpec) {
	var (
		_node = &OrganizationParent{config: opc.config}
		_spec = sqlgraph.NewCreateSpec(organizationparent.Table, sqlgraph.NewFieldSpec(organizationparent.FieldID, field.TypeInt))
	)
	if value, ok := opc.mutation.DateCreated(); ok {
		_spec.SetField(organizationparent.FieldDateCreated, field.TypeTime, value)
		_node.DateCreated = value
	}
	if value, ok := opc.mutation.DateUpdated(); ok {
		_spec.SetField(organizationparent.FieldDateUpdated, field.TypeTime, value)
		_node.DateUpdated = value
	}
	if value, ok := opc.mutation.ParentOrganizationID(); ok {
		_spec.SetField(organizationparent.FieldParentOrganizationID, field.TypeInt, value)
		_node.ParentOrganizationID = value
	}
	if value, ok := opc.mutation.OrganizationID(); ok {
		_spec.SetField(organizationparent.FieldOrganizationID, field.TypeInt, value)
		_node.OrganizationID = value
	}
	if value, ok := opc.mutation.From(); ok {
		_spec.SetField(organizationparent.FieldFrom, field.TypeTime, value)
		_node.From = value
	}
	if value, ok := opc.mutation.Until(); ok {
		_spec.SetField(organizationparent.FieldUntil, field.TypeTime, value)
		_node.Until = &value
	}
	return _node, _spec
}

// OrganizationParentCreateBulk is the builder for creating many OrganizationParent entities in bulk.
type OrganizationParentCreateBulk struct {
	config
	builders []*OrganizationParentCreate
}

// Save creates the OrganizationParent entities in the database.
func (opcb *OrganizationParentCreateBulk) Save(ctx context.Context) ([]*OrganizationParent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(opcb.builders))
	nodes := make([]*OrganizationParent, len(opcb.builders))
	mutators := make([]Mutator, len(opcb.builders))
	for i := range opcb.builders {
		func(i int, root context.Context) {
			builder := opcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*OrganizationParentMutation)
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
					_, err = mutators[i+1].Mutate(root, opcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, opcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, opcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (opcb *OrganizationParentCreateBulk) SaveX(ctx context.Context) []*OrganizationParent {
	v, err := opcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (opcb *OrganizationParentCreateBulk) Exec(ctx context.Context) error {
	_, err := opcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opcb *OrganizationParentCreateBulk) ExecX(ctx context.Context) {
	if err := opcb.Exec(ctx); err != nil {
		panic(err)
	}
}
