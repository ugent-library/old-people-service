// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people-service/ent/organizationperson"
	"github.com/ugent-library/people-service/ent/predicate"
)

// OrganizationPersonUpdate is the builder for updating OrganizationPerson entities.
type OrganizationPersonUpdate struct {
	config
	hooks     []Hook
	mutation  *OrganizationPersonMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the OrganizationPersonUpdate builder.
func (opu *OrganizationPersonUpdate) Where(ps ...predicate.OrganizationPerson) *OrganizationPersonUpdate {
	opu.mutation.Where(ps...)
	return opu
}

// SetDateUpdated sets the "date_updated" field.
func (opu *OrganizationPersonUpdate) SetDateUpdated(t time.Time) *OrganizationPersonUpdate {
	opu.mutation.SetDateUpdated(t)
	return opu
}

// SetOrganizationID sets the "organization_id" field.
func (opu *OrganizationPersonUpdate) SetOrganizationID(i int) *OrganizationPersonUpdate {
	opu.mutation.ResetOrganizationID()
	opu.mutation.SetOrganizationID(i)
	return opu
}

// AddOrganizationID adds i to the "organization_id" field.
func (opu *OrganizationPersonUpdate) AddOrganizationID(i int) *OrganizationPersonUpdate {
	opu.mutation.AddOrganizationID(i)
	return opu
}

// SetPersonID sets the "person_id" field.
func (opu *OrganizationPersonUpdate) SetPersonID(i int) *OrganizationPersonUpdate {
	opu.mutation.ResetPersonID()
	opu.mutation.SetPersonID(i)
	return opu
}

// AddPersonID adds i to the "person_id" field.
func (opu *OrganizationPersonUpdate) AddPersonID(i int) *OrganizationPersonUpdate {
	opu.mutation.AddPersonID(i)
	return opu
}

// SetFrom sets the "from" field.
func (opu *OrganizationPersonUpdate) SetFrom(t time.Time) *OrganizationPersonUpdate {
	opu.mutation.SetFrom(t)
	return opu
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (opu *OrganizationPersonUpdate) SetNillableFrom(t *time.Time) *OrganizationPersonUpdate {
	if t != nil {
		opu.SetFrom(*t)
	}
	return opu
}

// SetUntil sets the "until" field.
func (opu *OrganizationPersonUpdate) SetUntil(t time.Time) *OrganizationPersonUpdate {
	opu.mutation.SetUntil(t)
	return opu
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (opu *OrganizationPersonUpdate) SetNillableUntil(t *time.Time) *OrganizationPersonUpdate {
	if t != nil {
		opu.SetUntil(*t)
	}
	return opu
}

// ClearUntil clears the value of the "until" field.
func (opu *OrganizationPersonUpdate) ClearUntil() *OrganizationPersonUpdate {
	opu.mutation.ClearUntil()
	return opu
}

// Mutation returns the OrganizationPersonMutation object of the builder.
func (opu *OrganizationPersonUpdate) Mutation() *OrganizationPersonMutation {
	return opu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (opu *OrganizationPersonUpdate) Save(ctx context.Context) (int, error) {
	opu.defaults()
	return withHooks[int, OrganizationPersonMutation](ctx, opu.sqlSave, opu.mutation, opu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opu *OrganizationPersonUpdate) SaveX(ctx context.Context) int {
	affected, err := opu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (opu *OrganizationPersonUpdate) Exec(ctx context.Context) error {
	_, err := opu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opu *OrganizationPersonUpdate) ExecX(ctx context.Context) {
	if err := opu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opu *OrganizationPersonUpdate) defaults() {
	if _, ok := opu.mutation.DateUpdated(); !ok {
		v := organizationperson.UpdateDefaultDateUpdated()
		opu.mutation.SetDateUpdated(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opu *OrganizationPersonUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrganizationPersonUpdate {
	opu.modifiers = append(opu.modifiers, modifiers...)
	return opu
}

func (opu *OrganizationPersonUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationperson.Table, organizationperson.Columns, sqlgraph.NewFieldSpec(organizationperson.FieldID, field.TypeInt))
	if ps := opu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opu.mutation.DateUpdated(); ok {
		_spec.SetField(organizationperson.FieldDateUpdated, field.TypeTime, value)
	}
	if value, ok := opu.mutation.OrganizationID(); ok {
		_spec.SetField(organizationperson.FieldOrganizationID, field.TypeInt, value)
	}
	if value, ok := opu.mutation.AddedOrganizationID(); ok {
		_spec.AddField(organizationperson.FieldOrganizationID, field.TypeInt, value)
	}
	if value, ok := opu.mutation.PersonID(); ok {
		_spec.SetField(organizationperson.FieldPersonID, field.TypeInt, value)
	}
	if value, ok := opu.mutation.AddedPersonID(); ok {
		_spec.AddField(organizationperson.FieldPersonID, field.TypeInt, value)
	}
	if value, ok := opu.mutation.From(); ok {
		_spec.SetField(organizationperson.FieldFrom, field.TypeTime, value)
	}
	if value, ok := opu.mutation.Until(); ok {
		_spec.SetField(organizationperson.FieldUntil, field.TypeTime, value)
	}
	if opu.mutation.UntilCleared() {
		_spec.ClearField(organizationperson.FieldUntil, field.TypeTime)
	}
	_spec.AddModifiers(opu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, opu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationperson.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	opu.mutation.done = true
	return n, nil
}

// OrganizationPersonUpdateOne is the builder for updating a single OrganizationPerson entity.
type OrganizationPersonUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *OrganizationPersonMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetDateUpdated sets the "date_updated" field.
func (opuo *OrganizationPersonUpdateOne) SetDateUpdated(t time.Time) *OrganizationPersonUpdateOne {
	opuo.mutation.SetDateUpdated(t)
	return opuo
}

// SetOrganizationID sets the "organization_id" field.
func (opuo *OrganizationPersonUpdateOne) SetOrganizationID(i int) *OrganizationPersonUpdateOne {
	opuo.mutation.ResetOrganizationID()
	opuo.mutation.SetOrganizationID(i)
	return opuo
}

// AddOrganizationID adds i to the "organization_id" field.
func (opuo *OrganizationPersonUpdateOne) AddOrganizationID(i int) *OrganizationPersonUpdateOne {
	opuo.mutation.AddOrganizationID(i)
	return opuo
}

// SetPersonID sets the "person_id" field.
func (opuo *OrganizationPersonUpdateOne) SetPersonID(i int) *OrganizationPersonUpdateOne {
	opuo.mutation.ResetPersonID()
	opuo.mutation.SetPersonID(i)
	return opuo
}

// AddPersonID adds i to the "person_id" field.
func (opuo *OrganizationPersonUpdateOne) AddPersonID(i int) *OrganizationPersonUpdateOne {
	opuo.mutation.AddPersonID(i)
	return opuo
}

// SetFrom sets the "from" field.
func (opuo *OrganizationPersonUpdateOne) SetFrom(t time.Time) *OrganizationPersonUpdateOne {
	opuo.mutation.SetFrom(t)
	return opuo
}

// SetNillableFrom sets the "from" field if the given value is not nil.
func (opuo *OrganizationPersonUpdateOne) SetNillableFrom(t *time.Time) *OrganizationPersonUpdateOne {
	if t != nil {
		opuo.SetFrom(*t)
	}
	return opuo
}

// SetUntil sets the "until" field.
func (opuo *OrganizationPersonUpdateOne) SetUntil(t time.Time) *OrganizationPersonUpdateOne {
	opuo.mutation.SetUntil(t)
	return opuo
}

// SetNillableUntil sets the "until" field if the given value is not nil.
func (opuo *OrganizationPersonUpdateOne) SetNillableUntil(t *time.Time) *OrganizationPersonUpdateOne {
	if t != nil {
		opuo.SetUntil(*t)
	}
	return opuo
}

// ClearUntil clears the value of the "until" field.
func (opuo *OrganizationPersonUpdateOne) ClearUntil() *OrganizationPersonUpdateOne {
	opuo.mutation.ClearUntil()
	return opuo
}

// Mutation returns the OrganizationPersonMutation object of the builder.
func (opuo *OrganizationPersonUpdateOne) Mutation() *OrganizationPersonMutation {
	return opuo.mutation
}

// Where appends a list predicates to the OrganizationPersonUpdate builder.
func (opuo *OrganizationPersonUpdateOne) Where(ps ...predicate.OrganizationPerson) *OrganizationPersonUpdateOne {
	opuo.mutation.Where(ps...)
	return opuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (opuo *OrganizationPersonUpdateOne) Select(field string, fields ...string) *OrganizationPersonUpdateOne {
	opuo.fields = append([]string{field}, fields...)
	return opuo
}

// Save executes the query and returns the updated OrganizationPerson entity.
func (opuo *OrganizationPersonUpdateOne) Save(ctx context.Context) (*OrganizationPerson, error) {
	opuo.defaults()
	return withHooks[*OrganizationPerson, OrganizationPersonMutation](ctx, opuo.sqlSave, opuo.mutation, opuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (opuo *OrganizationPersonUpdateOne) SaveX(ctx context.Context) *OrganizationPerson {
	node, err := opuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (opuo *OrganizationPersonUpdateOne) Exec(ctx context.Context) error {
	_, err := opuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (opuo *OrganizationPersonUpdateOne) ExecX(ctx context.Context) {
	if err := opuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (opuo *OrganizationPersonUpdateOne) defaults() {
	if _, ok := opuo.mutation.DateUpdated(); !ok {
		v := organizationperson.UpdateDefaultDateUpdated()
		opuo.mutation.SetDateUpdated(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (opuo *OrganizationPersonUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *OrganizationPersonUpdateOne {
	opuo.modifiers = append(opuo.modifiers, modifiers...)
	return opuo
}

func (opuo *OrganizationPersonUpdateOne) sqlSave(ctx context.Context) (_node *OrganizationPerson, err error) {
	_spec := sqlgraph.NewUpdateSpec(organizationperson.Table, organizationperson.Columns, sqlgraph.NewFieldSpec(organizationperson.FieldID, field.TypeInt))
	id, ok := opuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "OrganizationPerson.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := opuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationperson.FieldID)
		for _, f := range fields {
			if !organizationperson.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != organizationperson.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := opuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := opuo.mutation.DateUpdated(); ok {
		_spec.SetField(organizationperson.FieldDateUpdated, field.TypeTime, value)
	}
	if value, ok := opuo.mutation.OrganizationID(); ok {
		_spec.SetField(organizationperson.FieldOrganizationID, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.AddedOrganizationID(); ok {
		_spec.AddField(organizationperson.FieldOrganizationID, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.PersonID(); ok {
		_spec.SetField(organizationperson.FieldPersonID, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.AddedPersonID(); ok {
		_spec.AddField(organizationperson.FieldPersonID, field.TypeInt, value)
	}
	if value, ok := opuo.mutation.From(); ok {
		_spec.SetField(organizationperson.FieldFrom, field.TypeTime, value)
	}
	if value, ok := opuo.mutation.Until(); ok {
		_spec.SetField(organizationperson.FieldUntil, field.TypeTime, value)
	}
	if opuo.mutation.UntilCleared() {
		_spec.ClearField(organizationperson.FieldUntil, field.TypeTime)
	}
	_spec.AddModifiers(opuo.modifiers...)
	_node = &OrganizationPerson{config: opuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, opuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{organizationperson.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	opuo.mutation.done = true
	return _node, nil
}
