// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwit-backend/ent/predicate"
	"github.com/0xfzz/tuwit-backend/ent/thread"
	"github.com/0xfzz/tuwit-backend/ent/threadcount"
)

// ThreadCountQuery is the builder for querying ThreadCount entities.
type ThreadCountQuery struct {
	config
	ctx        *QueryContext
	order      []threadcount.OrderOption
	inters     []Interceptor
	predicates []predicate.ThreadCount
	withThread *ThreadQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ThreadCountQuery builder.
func (tcq *ThreadCountQuery) Where(ps ...predicate.ThreadCount) *ThreadCountQuery {
	tcq.predicates = append(tcq.predicates, ps...)
	return tcq
}

// Limit the number of records to be returned by this query.
func (tcq *ThreadCountQuery) Limit(limit int) *ThreadCountQuery {
	tcq.ctx.Limit = &limit
	return tcq
}

// Offset to start from.
func (tcq *ThreadCountQuery) Offset(offset int) *ThreadCountQuery {
	tcq.ctx.Offset = &offset
	return tcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tcq *ThreadCountQuery) Unique(unique bool) *ThreadCountQuery {
	tcq.ctx.Unique = &unique
	return tcq
}

// Order specifies how the records should be ordered.
func (tcq *ThreadCountQuery) Order(o ...threadcount.OrderOption) *ThreadCountQuery {
	tcq.order = append(tcq.order, o...)
	return tcq
}

// QueryThread chains the current query on the "thread" edge.
func (tcq *ThreadCountQuery) QueryThread() *ThreadQuery {
	query := (&ThreadClient{config: tcq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tcq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(threadcount.Table, threadcount.FieldID, selector),
			sqlgraph.To(thread.Table, thread.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, threadcount.ThreadTable, threadcount.ThreadColumn),
		)
		fromU = sqlgraph.SetNeighbors(tcq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ThreadCount entity from the query.
// Returns a *NotFoundError when no ThreadCount was found.
func (tcq *ThreadCountQuery) First(ctx context.Context) (*ThreadCount, error) {
	nodes, err := tcq.Limit(1).All(setContextOp(ctx, tcq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{threadcount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tcq *ThreadCountQuery) FirstX(ctx context.Context) *ThreadCount {
	node, err := tcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ThreadCount ID from the query.
// Returns a *NotFoundError when no ThreadCount ID was found.
func (tcq *ThreadCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(1).IDs(setContextOp(ctx, tcq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{threadcount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tcq *ThreadCountQuery) FirstIDX(ctx context.Context) int {
	id, err := tcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ThreadCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ThreadCount entity is found.
// Returns a *NotFoundError when no ThreadCount entities are found.
func (tcq *ThreadCountQuery) Only(ctx context.Context) (*ThreadCount, error) {
	nodes, err := tcq.Limit(2).All(setContextOp(ctx, tcq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{threadcount.Label}
	default:
		return nil, &NotSingularError{threadcount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tcq *ThreadCountQuery) OnlyX(ctx context.Context) *ThreadCount {
	node, err := tcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ThreadCount ID in the query.
// Returns a *NotSingularError when more than one ThreadCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (tcq *ThreadCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tcq.Limit(2).IDs(setContextOp(ctx, tcq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{threadcount.Label}
	default:
		err = &NotSingularError{threadcount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tcq *ThreadCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := tcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ThreadCounts.
func (tcq *ThreadCountQuery) All(ctx context.Context) ([]*ThreadCount, error) {
	ctx = setContextOp(ctx, tcq.ctx, "All")
	if err := tcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ThreadCount, *ThreadCountQuery]()
	return withInterceptors[[]*ThreadCount](ctx, tcq, qr, tcq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tcq *ThreadCountQuery) AllX(ctx context.Context) []*ThreadCount {
	nodes, err := tcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ThreadCount IDs.
func (tcq *ThreadCountQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tcq.ctx.Unique == nil && tcq.path != nil {
		tcq.Unique(true)
	}
	ctx = setContextOp(ctx, tcq.ctx, "IDs")
	if err = tcq.Select(threadcount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tcq *ThreadCountQuery) IDsX(ctx context.Context) []int {
	ids, err := tcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tcq *ThreadCountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Count")
	if err := tcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tcq, querierCount[*ThreadCountQuery](), tcq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tcq *ThreadCountQuery) CountX(ctx context.Context) int {
	count, err := tcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tcq *ThreadCountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tcq.ctx, "Exist")
	switch _, err := tcq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tcq *ThreadCountQuery) ExistX(ctx context.Context) bool {
	exist, err := tcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ThreadCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tcq *ThreadCountQuery) Clone() *ThreadCountQuery {
	if tcq == nil {
		return nil
	}
	return &ThreadCountQuery{
		config:     tcq.config,
		ctx:        tcq.ctx.Clone(),
		order:      append([]threadcount.OrderOption{}, tcq.order...),
		inters:     append([]Interceptor{}, tcq.inters...),
		predicates: append([]predicate.ThreadCount{}, tcq.predicates...),
		withThread: tcq.withThread.Clone(),
		// clone intermediate query.
		sql:  tcq.sql.Clone(),
		path: tcq.path,
	}
}

// WithThread tells the query-builder to eager-load the nodes that are connected to
// the "thread" edge. The optional arguments are used to configure the query builder of the edge.
func (tcq *ThreadCountQuery) WithThread(opts ...func(*ThreadQuery)) *ThreadCountQuery {
	query := (&ThreadClient{config: tcq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tcq.withThread = query
	return tcq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ReplyCount int `json:"reply_count,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ThreadCount.Query().
//		GroupBy(threadcount.FieldReplyCount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tcq *ThreadCountQuery) GroupBy(field string, fields ...string) *ThreadCountGroupBy {
	tcq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ThreadCountGroupBy{build: tcq}
	grbuild.flds = &tcq.ctx.Fields
	grbuild.label = threadcount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ReplyCount int `json:"reply_count,omitempty"`
//	}
//
//	client.ThreadCount.Query().
//		Select(threadcount.FieldReplyCount).
//		Scan(ctx, &v)
func (tcq *ThreadCountQuery) Select(fields ...string) *ThreadCountSelect {
	tcq.ctx.Fields = append(tcq.ctx.Fields, fields...)
	sbuild := &ThreadCountSelect{ThreadCountQuery: tcq}
	sbuild.label = threadcount.Label
	sbuild.flds, sbuild.scan = &tcq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ThreadCountSelect configured with the given aggregations.
func (tcq *ThreadCountQuery) Aggregate(fns ...AggregateFunc) *ThreadCountSelect {
	return tcq.Select().Aggregate(fns...)
}

func (tcq *ThreadCountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tcq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tcq); err != nil {
				return err
			}
		}
	}
	for _, f := range tcq.ctx.Fields {
		if !threadcount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tcq.path != nil {
		prev, err := tcq.path(ctx)
		if err != nil {
			return err
		}
		tcq.sql = prev
	}
	return nil
}

func (tcq *ThreadCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ThreadCount, error) {
	var (
		nodes       = []*ThreadCount{}
		withFKs     = tcq.withFKs
		_spec       = tcq.querySpec()
		loadedTypes = [1]bool{
			tcq.withThread != nil,
		}
	)
	if tcq.withThread != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, threadcount.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ThreadCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ThreadCount{config: tcq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tcq.withThread; query != nil {
		if err := tcq.loadThread(ctx, query, nodes, nil,
			func(n *ThreadCount, e *Thread) { n.Edges.Thread = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tcq *ThreadCountQuery) loadThread(ctx context.Context, query *ThreadQuery, nodes []*ThreadCount, init func(*ThreadCount), assign func(*ThreadCount, *Thread)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ThreadCount)
	for i := range nodes {
		if nodes[i].thread_thread_count == nil {
			continue
		}
		fk := *nodes[i].thread_thread_count
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(thread.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "thread_thread_count" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (tcq *ThreadCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tcq.querySpec()
	_spec.Node.Columns = tcq.ctx.Fields
	if len(tcq.ctx.Fields) > 0 {
		_spec.Unique = tcq.ctx.Unique != nil && *tcq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tcq.driver, _spec)
}

func (tcq *ThreadCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(threadcount.Table, threadcount.Columns, sqlgraph.NewFieldSpec(threadcount.FieldID, field.TypeInt))
	_spec.From = tcq.sql
	if unique := tcq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tcq.path != nil {
		_spec.Unique = true
	}
	if fields := tcq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, threadcount.FieldID)
		for i := range fields {
			if fields[i] != threadcount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tcq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tcq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tcq *ThreadCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tcq.driver.Dialect())
	t1 := builder.Table(threadcount.Table)
	columns := tcq.ctx.Fields
	if len(columns) == 0 {
		columns = threadcount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tcq.sql != nil {
		selector = tcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tcq.ctx.Unique != nil && *tcq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tcq.predicates {
		p(selector)
	}
	for _, p := range tcq.order {
		p(selector)
	}
	if offset := tcq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tcq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ThreadCountGroupBy is the group-by builder for ThreadCount entities.
type ThreadCountGroupBy struct {
	selector
	build *ThreadCountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tcgb *ThreadCountGroupBy) Aggregate(fns ...AggregateFunc) *ThreadCountGroupBy {
	tcgb.fns = append(tcgb.fns, fns...)
	return tcgb
}

// Scan applies the selector query and scans the result into the given value.
func (tcgb *ThreadCountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcgb.build.ctx, "GroupBy")
	if err := tcgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThreadCountQuery, *ThreadCountGroupBy](ctx, tcgb.build, tcgb, tcgb.build.inters, v)
}

func (tcgb *ThreadCountGroupBy) sqlScan(ctx context.Context, root *ThreadCountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tcgb.fns))
	for _, fn := range tcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tcgb.flds)+len(tcgb.fns))
		for _, f := range *tcgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tcgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ThreadCountSelect is the builder for selecting fields of ThreadCount entities.
type ThreadCountSelect struct {
	*ThreadCountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tcs *ThreadCountSelect) Aggregate(fns ...AggregateFunc) *ThreadCountSelect {
	tcs.fns = append(tcs.fns, fns...)
	return tcs
}

// Scan applies the selector query and scans the result into the given value.
func (tcs *ThreadCountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tcs.ctx, "Select")
	if err := tcs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ThreadCountQuery, *ThreadCountSelect](ctx, tcs.ThreadCountQuery, tcs, tcs.inters, v)
}

func (tcs *ThreadCountSelect) sqlScan(ctx context.Context, root *ThreadCountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tcs.fns))
	for _, fn := range tcs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tcs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
