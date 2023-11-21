// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/old-people-service/ent/organizationperson"
	"github.com/ugent-library/old-people-service/ent/predicate"
)

// OrganizationPersonDelete is the builder for deleting a OrganizationPerson entity.
type OrganizationPersonDelete struct {
	config
	hooks    []Hook
	mutation *OrganizationPersonMutation
}

// Where appends a list predicates to the OrganizationPersonDelete builder.
func (opd *OrganizationPersonDelete) Where(ps ...predicate.OrganizationPerson) *OrganizationPersonDelete {
	opd.mutation.Where(ps...)
	return opd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (opd *OrganizationPersonDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, OrganizationPersonMutation](ctx, opd.sqlExec, opd.mutation, opd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (opd *OrganizationPersonDelete) ExecX(ctx context.Context) int {
	n, err := opd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (opd *OrganizationPersonDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(organizationperson.Table, sqlgraph.NewFieldSpec(organizationperson.FieldID, field.TypeInt))
	if ps := opd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, opd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	opd.mutation.done = true
	return affected, err
}

// OrganizationPersonDeleteOne is the builder for deleting a single OrganizationPerson entity.
type OrganizationPersonDeleteOne struct {
	opd *OrganizationPersonDelete
}

// Where appends a list predicates to the OrganizationPersonDelete builder.
func (opdo *OrganizationPersonDeleteOne) Where(ps ...predicate.OrganizationPerson) *OrganizationPersonDeleteOne {
	opdo.opd.mutation.Where(ps...)
	return opdo
}

// Exec executes the deletion query.
func (opdo *OrganizationPersonDeleteOne) Exec(ctx context.Context) error {
	n, err := opdo.opd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{organizationperson.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (opdo *OrganizationPersonDeleteOne) ExecX(ctx context.Context) {
	if err := opdo.Exec(ctx); err != nil {
		panic(err)
	}
}
