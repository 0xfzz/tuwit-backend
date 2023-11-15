// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwitt/ent/predicate"
	"github.com/0xfzz/tuwitt/ent/threadlikeuser"
)

// ThreadLikeUserQuery is the builder for querying ThreadLikeUser entities.
type ThreadLikeUserQuery struct {
	config
	ctx        *QueryContext
	order      []threadlikeuser.OrderOption
	inters     []Interceptor
	predicates []predicate.ThreadLikeUser
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ThreadLikeUserQuery builder.
func (tluq *ThreadLikeUserQuery) Where(ps ...predicate.ThreadLikeUser) *ThreadLikeUserQuery {
	tluq.predicates = append(tluq.predicates, ps...)
	return tluq
}

// Limit the number of records to be returned by this query.
func (tluq *ThreadLikeUserQuery) Limit(limit int) *ThreadLikeUserQuery {
	tluq.ctx.Limit = &limit
	return tluq
}

// Offset to start from.
func (tluq *ThreadLikeUserQuery) Offset(offset int) *ThreadLikeUserQuery {
	tluq.ctx.Offset = &offset
	return tluq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tluq *ThreadLikeUserQuery) Unique(unique bool) *ThreadLikeUserQuery {
	tluq.ctx.Unique = &unique
	return tluq
}

// Order specifies how the records should be ordered.
func (tluq *ThreadLikeUserQuery) Order(o ...threadlikeuser.OrderOption) *ThreadLikeUserQuery {
	tluq.order = append(tluq.order, o...)
	return tluq
}

// First returns the first ThreadLikeUser entity from the query.
// Returns a *NotFoundError when no ThreadLikeUser was found.
func (tluq *ThreadLikeUserQuery) First(ctx context.Context) (*ThreadLikeUser, error) {
	nodes, err := tluq.Limit(1).All(setContextOp(ctx, tluq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{threadlikeuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) FirstX(ctx context.Context) *ThreadLikeUser {
	node, err := tluq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ThreadLikeUser ID from the query.
// Returns a *NotFoundError when no ThreadLikeUser ID was found.
func (tluq *ThreadLikeUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tluq.Limit(1).IDs(setContextOp(ctx, tluq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{threadlikeuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) FirstIDX(ctx context.Context) int {
	id, err := tluq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ThreadLikeUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ThreadLikeUser entity is found.
// Returns a *NotFoundError when no ThreadLikeUser entities are found.
func (tluq *ThreadLikeUserQuery) Only(ctx context.Context) (*ThreadLikeUser, error) {
	nodes, err := tluq.Limit(2).All(setContextOp(ctx, tluq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{threadlikeuser.Label}
	default:
		return nil, &NotSingularError{threadlikeuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) OnlyX(ctx context.Context) *ThreadLikeUser {
	node, err := tluq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ThreadLikeUser ID in the query.
// Returns a *NotSingularError when more than one ThreadLikeUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (tluq *ThreadLikeUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tluq.Limit(2).IDs(setContextOp(ctx, tluq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{threadlikeuser.Label}
	default:
		err = &NotSingularError{threadlikeuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := tluq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ThreadLikeUsers.
func (tluq *ThreadLikeUserQuery) All(ctx context.Context) ([]*ThreadLikeUser, error) {
	ctx = setContextOp(ctx, tluq.ctx, "All")
	if err := tluq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ThreadLikeUser, *ThreadLikeUserQuery]()
	return withInterceptors[[]*ThreadLikeUser](ctx, tluq, qr, tluq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) AllX(ctx context.Context) []*ThreadLikeUser {
	nodes, err := tluq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ThreadLikeUser IDs.
func (tluq *ThreadLikeUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tluq.ctx.Unique == nil && tluq.path != nil {
		tluq.Unique(true)
	}
	ctx = setContextOp(ctx, tluq.ctx, "IDs")
	if err = tluq.Select(threadlikeuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) IDsX(ctx context.Context) []int {
	ids, err := tluq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tluq *ThreadLikeUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tluq.ctx, "Count")
	if err := tluq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tluq, querierCount[*ThreadLikeUserQuery](), tluq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) CountX(ctx context.Context) int {
	count, err := tluq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tluq *ThreadLikeUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tluq.ctx, "Exist")
	switch _, err := tluq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tluq *ThreadLikeUserQuery) ExistX(ctx context.Context) bool {
	exist, err := tluq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ThreadLikeUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tluq *ThreadLikeUserQuery) Clone() *ThreadLikeUserQuery {
	if tluq == nil {
		return nil
	}
	return &ThreadLikeUserQuery{
		config:     tluq.config,
		ctx:        tluq.ctx.Clone(),
		order:      append([]threadlikeuser.OrderOption{}, tluq.order...),
		inters:     append([]Interceptor{}, tluq.inters...),
		predicates: append([]predicate.ThreadLikeUser{}, tluq.predicates...),
		// clone intermediate query.
		sql:  tluq.sql.Clone(),
		path: tluq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (tluq *ThreadLikeUserQuery) GroupBy(field string, fields ...string) *ThreadLikeUserGroupBy {
	tluq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ThreadLikeUserGroupBy{build: tluq}
	grbuild.flds = &tluq.ctx.Fields
	grbuild.label = threadlikeuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (tluq *ThreadLikeUserQuery) Select(fields ...string) *ThreadLikeUserSelect {
	tluq.ctx.Fields = append(tluq.ctx.Fields, fields...)
	sbuild := &ThreadLikeUserSelect{ThreadLikeUserQuery: tluq}
	sbuild.label = threadlikeuser.Label
	sbuild.flds, sbuild.scan = &tluq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ThreadLikeUserSelect configured with the given aggregations.
func (tluq *ThreadLikeUserQuery) Aggregate(fns ...AggregateFunc) *ThreadLikeUserSelect {
	return tluq.Select().Aggregate(fns...)
}

func (tluq *ThreadLikeUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tluq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tluq); err != nil {
				return err
			}
		}
	}
	for _, f := range tluq.ctx.Fields {
		if !threadlikeuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tluq.path != nil {
		prev, err := tluq.path(ctx)
		if err != nil {
			return err
		}
		tluq.sql = prev
	}
	return nil
}

func (tluq *ThreadLikeUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ThreadLikeUser, error) {
	var (
		nodes = []*ThreadLikeUser{}
		_spec = tluq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ThreadLikeUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ThreadLikeUser{config: tluq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tluq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (tluq *ThreadLikeUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tluq.querySpec()
	_spec.Node.Columns = tluq.ctx.Fields
	if len(tluq.ctx.Fields) > 0 {
		_spec.Unique = tluq.ctx.Unique != nil && *tluq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tluq.driver, _spec)
}

func (tluq *ThreadLikeUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(threadlikeuser.Table, threadlikeuser.Columns, sqlgraph.NewFieldSpec(threadlikeuser.FieldID, field.TypeInt))
	_spec.From = tluq.sql
	if unique := tluq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tluq.path != nil {
		_spec.Unique = true
	}
	if fields := tluq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, threadlikeuser.FieldID)
		for i := range fields {
			if fields[i] != threadlikeuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tluq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tluq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tluq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tluq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tluq *ThreadLikeUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tluq.driver.Dialect())
	t1 := builder.Table(threadlikeuser.Table)
	columns := tluq.ctx.Fields
	if len(columns) == 0 {
		columns = threadlikeuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tluq.sql != nil {
		selector = tluq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tluq.ctx.Unique != nil && *tluq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tluq.predicates {
		p(selector)
	}
	for _, p := range tluq.order {
		p(selector)
	}
	if offset := tluq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tluq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ThreadLikeUserGroupBy is the group-by builder for ThreadLikeUser entities.
type ThreadLikeUserGroupBy struct {
	selector
	build *ThreadLikeUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tlugb *ThreadLikeUserGroupBy) Aggregate(fns ...AggregateFunc) *ThreadLikeUserGroupBy {
	tlugb.fns = append(tlugb.fns, fns...)
	return tlugb
}

// Scan applies the selector query and scans the result into the given value.
func (tlugb *ThreadLikeUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tlugb.build.ctx, "GroupBy")
	if err := tlugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThreadLikeUserQuery, *ThreadLikeUserGroupBy](ctx, tlugb.build, tlugb, tlugb.build.inters, v)
}

func (tlugb *ThreadLikeUserGroupBy) sqlScan(ctx context.Context, root *ThreadLikeUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tlugb.fns))
	for _, fn := range tlugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tlugb.flds)+len(tlugb.fns))
		for _, f := range *tlugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tlugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ThreadLikeUserSelect is the builder for selecting fields of ThreadLikeUser entities.
type ThreadLikeUserSelect struct {
	*ThreadLikeUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tlus *ThreadLikeUserSelect) Aggregate(fns ...AggregateFunc) *ThreadLikeUserSelect {
	tlus.fns = append(tlus.fns, fns...)
	return tlus
}

// Scan applies the selector query and scans the result into the given value.
func (tlus *ThreadLikeUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tlus.ctx, "Select")
	if err := tlus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThreadLikeUserQuery, *ThreadLikeUserSelect](ctx, tlus.ThreadLikeUserQuery, tlus, tlus.inters, v)
}

func (tlus *ThreadLikeUserSelect) sqlScan(ctx context.Context, root *ThreadLikeUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tlus.fns))
	for _, fn := range tlus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tlus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tlus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}