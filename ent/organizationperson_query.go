// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ugent-library/people-service/ent/organizationperson"
	"github.com/ugent-library/people-service/ent/predicate"
)

// OrganizationPersonQuery is the builder for querying OrganizationPerson entities.
type OrganizationPersonQuery struct {
	config
	ctx        *QueryContext
	order      []organizationperson.OrderOption
	inters     []Interceptor
	predicates []predicate.OrganizationPerson
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the OrganizationPersonQuery builder.
func (opq *OrganizationPersonQuery) Where(ps ...predicate.OrganizationPerson) *OrganizationPersonQuery {
	opq.predicates = append(opq.predicates, ps...)
	return opq
}

// Limit the number of records to be returned by this query.
func (opq *OrganizationPersonQuery) Limit(limit int) *OrganizationPersonQuery {
	opq.ctx.Limit = &limit
	return opq
}

// Offset to start from.
func (opq *OrganizationPersonQuery) Offset(offset int) *OrganizationPersonQuery {
	opq.ctx.Offset = &offset
	return opq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (opq *OrganizationPersonQuery) Unique(unique bool) *OrganizationPersonQuery {
	opq.ctx.Unique = &unique
	return opq
}

// Order specifies how the records should be ordered.
func (opq *OrganizationPersonQuery) Order(o ...organizationperson.OrderOption) *OrganizationPersonQuery {
	opq.order = append(opq.order, o...)
	return opq
}

// First returns the first OrganizationPerson entity from the query.
// Returns a *NotFoundError when no OrganizationPerson was found.
func (opq *OrganizationPersonQuery) First(ctx context.Context) (*OrganizationPerson, error) {
	nodes, err := opq.Limit(1).All(setContextOp(ctx, opq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{organizationperson.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (opq *OrganizationPersonQuery) FirstX(ctx context.Context) *OrganizationPerson {
	node, err := opq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first OrganizationPerson ID from the query.
// Returns a *NotFoundError when no OrganizationPerson ID was found.
func (opq *OrganizationPersonQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = opq.Limit(1).IDs(setContextOp(ctx, opq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{organizationperson.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (opq *OrganizationPersonQuery) FirstIDX(ctx context.Context) int {
	id, err := opq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single OrganizationPerson entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one OrganizationPerson entity is found.
// Returns a *NotFoundError when no OrganizationPerson entities are found.
func (opq *OrganizationPersonQuery) Only(ctx context.Context) (*OrganizationPerson, error) {
	nodes, err := opq.Limit(2).All(setContextOp(ctx, opq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{organizationperson.Label}
	default:
		return nil, &NotSingularError{organizationperson.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (opq *OrganizationPersonQuery) OnlyX(ctx context.Context) *OrganizationPerson {
	node, err := opq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only OrganizationPerson ID in the query.
// Returns a *NotSingularError when more than one OrganizationPerson ID is found.
// Returns a *NotFoundError when no entities are found.
func (opq *OrganizationPersonQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = opq.Limit(2).IDs(setContextOp(ctx, opq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{organizationperson.Label}
	default:
		err = &NotSingularError{organizationperson.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (opq *OrganizationPersonQuery) OnlyIDX(ctx context.Context) int {
	id, err := opq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of OrganizationPersons.
func (opq *OrganizationPersonQuery) All(ctx context.Context) ([]*OrganizationPerson, error) {
	ctx = setContextOp(ctx, opq.ctx, "All")
	if err := opq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*OrganizationPerson, *OrganizationPersonQuery]()
	return withInterceptors[[]*OrganizationPerson](ctx, opq, qr, opq.inters)
}

// AllX is like All, but panics if an error occurs.
func (opq *OrganizationPersonQuery) AllX(ctx context.Context) []*OrganizationPerson {
	nodes, err := opq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of OrganizationPerson IDs.
func (opq *OrganizationPersonQuery) IDs(ctx context.Context) (ids []int, err error) {
	if opq.ctx.Unique == nil && opq.path != nil {
		opq.Unique(true)
	}
	ctx = setContextOp(ctx, opq.ctx, "IDs")
	if err = opq.Select(organizationperson.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (opq *OrganizationPersonQuery) IDsX(ctx context.Context) []int {
	ids, err := opq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (opq *OrganizationPersonQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, opq.ctx, "Count")
	if err := opq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, opq, querierCount[*OrganizationPersonQuery](), opq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (opq *OrganizationPersonQuery) CountX(ctx context.Context) int {
	count, err := opq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (opq *OrganizationPersonQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, opq.ctx, "Exist")
	switch _, err := opq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (opq *OrganizationPersonQuery) ExistX(ctx context.Context) bool {
	exist, err := opq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the OrganizationPersonQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (opq *OrganizationPersonQuery) Clone() *OrganizationPersonQuery {
	if opq == nil {
		return nil
	}
	return &OrganizationPersonQuery{
		config:     opq.config,
		ctx:        opq.ctx.Clone(),
		order:      append([]organizationperson.OrderOption{}, opq.order...),
		inters:     append([]Interceptor{}, opq.inters...),
		predicates: append([]predicate.OrganizationPerson{}, opq.predicates...),
		// clone intermediate query.
		sql:  opq.sql.Clone(),
		path: opq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		DateCreated time.Time `json:"date_created,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.OrganizationPerson.Query().
//		GroupBy(organizationperson.FieldDateCreated).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (opq *OrganizationPersonQuery) GroupBy(field string, fields ...string) *OrganizationPersonGroupBy {
	opq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &OrganizationPersonGroupBy{build: opq}
	grbuild.flds = &opq.ctx.Fields
	grbuild.label = organizationperson.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		DateCreated time.Time `json:"date_created,omitempty"`
//	}
//
//	client.OrganizationPerson.Query().
//		Select(organizationperson.FieldDateCreated).
//		Scan(ctx, &v)
func (opq *OrganizationPersonQuery) Select(fields ...string) *OrganizationPersonSelect {
	opq.ctx.Fields = append(opq.ctx.Fields, fields...)
	sbuild := &OrganizationPersonSelect{OrganizationPersonQuery: opq}
	sbuild.label = organizationperson.Label
	sbuild.flds, sbuild.scan = &opq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a OrganizationPersonSelect configured with the given aggregations.
func (opq *OrganizationPersonQuery) Aggregate(fns ...AggregateFunc) *OrganizationPersonSelect {
	return opq.Select().Aggregate(fns...)
}

func (opq *OrganizationPersonQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range opq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, opq); err != nil {
				return err
			}
		}
	}
	for _, f := range opq.ctx.Fields {
		if !organizationperson.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if opq.path != nil {
		prev, err := opq.path(ctx)
		if err != nil {
			return err
		}
		opq.sql = prev
	}
	return nil
}

func (opq *OrganizationPersonQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*OrganizationPerson, error) {
	var (
		nodes = []*OrganizationPerson{}
		_spec = opq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*OrganizationPerson).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &OrganizationPerson{config: opq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, opq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (opq *OrganizationPersonQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := opq.querySpec()
	if len(opq.modifiers) > 0 {
		_spec.Modifiers = opq.modifiers
	}
	_spec.Node.Columns = opq.ctx.Fields
	if len(opq.ctx.Fields) > 0 {
		_spec.Unique = opq.ctx.Unique != nil && *opq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, opq.driver, _spec)
}

func (opq *OrganizationPersonQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(organizationperson.Table, organizationperson.Columns, sqlgraph.NewFieldSpec(organizationperson.FieldID, field.TypeInt))
	_spec.From = opq.sql
	if unique := opq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if opq.path != nil {
		_spec.Unique = true
	}
	if fields := opq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, organizationperson.FieldID)
		for i := range fields {
			if fields[i] != organizationperson.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := opq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := opq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := opq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := opq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (opq *OrganizationPersonQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(opq.driver.Dialect())
	t1 := builder.Table(organizationperson.Table)
	columns := opq.ctx.Fields
	if len(columns) == 0 {
		columns = organizationperson.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if opq.sql != nil {
		selector = opq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if opq.ctx.Unique != nil && *opq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range opq.modifiers {
		m(selector)
	}
	for _, p := range opq.predicates {
		p(selector)
	}
	for _, p := range opq.order {
		p(selector)
	}
	if offset := opq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := opq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (opq *OrganizationPersonQuery) ForUpdate(opts ...sql.LockOption) *OrganizationPersonQuery {
	if opq.driver.Dialect() == dialect.Postgres {
		opq.Unique(false)
	}
	opq.modifiers = append(opq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return opq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (opq *OrganizationPersonQuery) ForShare(opts ...sql.LockOption) *OrganizationPersonQuery {
	if opq.driver.Dialect() == dialect.Postgres {
		opq.Unique(false)
	}
	opq.modifiers = append(opq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return opq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (opq *OrganizationPersonQuery) Modify(modifiers ...func(s *sql.Selector)) *OrganizationPersonSelect {
	opq.modifiers = append(opq.modifiers, modifiers...)
	return opq.Select()
}

// OrganizationPersonGroupBy is the group-by builder for OrganizationPerson entities.
type OrganizationPersonGroupBy struct {
	selector
	build *OrganizationPersonQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (opgb *OrganizationPersonGroupBy) Aggregate(fns ...AggregateFunc) *OrganizationPersonGroupBy {
	opgb.fns = append(opgb.fns, fns...)
	return opgb
}

// Scan applies the selector query and scans the result into the given value.
func (opgb *OrganizationPersonGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, opgb.build.ctx, "GroupBy")
	if err := opgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationPersonQuery, *OrganizationPersonGroupBy](ctx, opgb.build, opgb, opgb.build.inters, v)
}

func (opgb *OrganizationPersonGroupBy) sqlScan(ctx context.Context, root *OrganizationPersonQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(opgb.fns))
	for _, fn := range opgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*opgb.flds)+len(opgb.fns))
		for _, f := range *opgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*opgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := opgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// OrganizationPersonSelect is the builder for selecting fields of OrganizationPerson entities.
type OrganizationPersonSelect struct {
	*OrganizationPersonQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ops *OrganizationPersonSelect) Aggregate(fns ...AggregateFunc) *OrganizationPersonSelect {
	ops.fns = append(ops.fns, fns...)
	return ops
}

// Scan applies the selector query and scans the result into the given value.
func (ops *OrganizationPersonSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ops.ctx, "Select")
	if err := ops.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*OrganizationPersonQuery, *OrganizationPersonSelect](ctx, ops.OrganizationPersonQuery, ops, ops.inters, v)
}

func (ops *OrganizationPersonSelect) sqlScan(ctx context.Context, root *OrganizationPersonQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ops.fns))
	for _, fn := range ops.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ops.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ops.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ops *OrganizationPersonSelect) Modify(modifiers ...func(s *sql.Selector)) *OrganizationPersonSelect {
	ops.modifiers = append(ops.modifiers, modifiers...)
	return ops
}
