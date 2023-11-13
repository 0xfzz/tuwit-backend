// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwitt/ent/predicate"
	"github.com/0xfzz/tuwitt/ent/useraccount"
	"github.com/0xfzz/tuwitt/ent/usercount"
)

// UserCountQuery is the builder for querying UserCount entities.
type UserCountQuery struct {
	config
	ctx        *QueryContext
	order      []usercount.OrderOption
	inters     []Interceptor
	predicates []predicate.UserCount
	withUser   *UserAccountQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserCountQuery builder.
func (ucq *UserCountQuery) Where(ps ...predicate.UserCount) *UserCountQuery {
	ucq.predicates = append(ucq.predicates, ps...)
	return ucq
}

// Limit the number of records to be returned by this query.
func (ucq *UserCountQuery) Limit(limit int) *UserCountQuery {
	ucq.ctx.Limit = &limit
	return ucq
}

// Offset to start from.
func (ucq *UserCountQuery) Offset(offset int) *UserCountQuery {
	ucq.ctx.Offset = &offset
	return ucq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ucq *UserCountQuery) Unique(unique bool) *UserCountQuery {
	ucq.ctx.Unique = &unique
	return ucq
}

// Order specifies how the records should be ordered.
func (ucq *UserCountQuery) Order(o ...usercount.OrderOption) *UserCountQuery {
	ucq.order = append(ucq.order, o...)
	return ucq
}

// QueryUser chains the current query on the "user" edge.
func (ucq *UserCountQuery) QueryUser() *UserAccountQuery {
	query := (&UserAccountClient{config: ucq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ucq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ucq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(usercount.Table, usercount.FieldID, selector),
			sqlgraph.To(useraccount.Table, useraccount.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, usercount.UserTable, usercount.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(ucq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserCount entity from the query.
// Returns a *NotFoundError when no UserCount was found.
func (ucq *UserCountQuery) First(ctx context.Context) (*UserCount, error) {
	nodes, err := ucq.Limit(1).All(setContextOp(ctx, ucq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{usercount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ucq *UserCountQuery) FirstX(ctx context.Context) *UserCount {
	node, err := ucq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserCount ID from the query.
// Returns a *NotFoundError when no UserCount ID was found.
func (ucq *UserCountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(1).IDs(setContextOp(ctx, ucq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{usercount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ucq *UserCountQuery) FirstIDX(ctx context.Context) int {
	id, err := ucq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserCount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserCount entity is found.
// Returns a *NotFoundError when no UserCount entities are found.
func (ucq *UserCountQuery) Only(ctx context.Context) (*UserCount, error) {
	nodes, err := ucq.Limit(2).All(setContextOp(ctx, ucq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{usercount.Label}
	default:
		return nil, &NotSingularError{usercount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ucq *UserCountQuery) OnlyX(ctx context.Context) *UserCount {
	node, err := ucq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserCount ID in the query.
// Returns a *NotSingularError when more than one UserCount ID is found.
// Returns a *NotFoundError when no entities are found.
func (ucq *UserCountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ucq.Limit(2).IDs(setContextOp(ctx, ucq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{usercount.Label}
	default:
		err = &NotSingularError{usercount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ucq *UserCountQuery) OnlyIDX(ctx context.Context) int {
	id, err := ucq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserCounts.
func (ucq *UserCountQuery) All(ctx context.Context) ([]*UserCount, error) {
	ctx = setContextOp(ctx, ucq.ctx, "All")
	if err := ucq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserCount, *UserCountQuery]()
	return withInterceptors[[]*UserCount](ctx, ucq, qr, ucq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ucq *UserCountQuery) AllX(ctx context.Context) []*UserCount {
	nodes, err := ucq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserCount IDs.
func (ucq *UserCountQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ucq.ctx.Unique == nil && ucq.path != nil {
		ucq.Unique(true)
	}
	ctx = setContextOp(ctx, ucq.ctx, "IDs")
	if err = ucq.Select(usercount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ucq *UserCountQuery) IDsX(ctx context.Context) []int {
	ids, err := ucq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ucq *UserCountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ucq.ctx, "Count")
	if err := ucq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ucq, querierCount[*UserCountQuery](), ucq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ucq *UserCountQuery) CountX(ctx context.Context) int {
	count, err := ucq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ucq *UserCountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ucq.ctx, "Exist")
	switch _, err := ucq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ucq *UserCountQuery) ExistX(ctx context.Context) bool {
	exist, err := ucq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserCountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ucq *UserCountQuery) Clone() *UserCountQuery {
	if ucq == nil {
		return nil
	}
	return &UserCountQuery{
		config:     ucq.config,
		ctx:        ucq.ctx.Clone(),
		order:      append([]usercount.OrderOption{}, ucq.order...),
		inters:     append([]Interceptor{}, ucq.inters...),
		predicates: append([]predicate.UserCount{}, ucq.predicates...),
		withUser:   ucq.withUser.Clone(),
		// clone intermediate query.
		sql:  ucq.sql.Clone(),
		path: ucq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (ucq *UserCountQuery) WithUser(opts ...func(*UserAccountQuery)) *UserCountQuery {
	query := (&UserAccountClient{config: ucq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ucq.withUser = query
	return ucq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FollowerCount int `json:"follower_count,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserCount.Query().
//		GroupBy(usercount.FieldFollowerCount).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ucq *UserCountQuery) GroupBy(field string, fields ...string) *UserCountGroupBy {
	ucq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserCountGroupBy{build: ucq}
	grbuild.flds = &ucq.ctx.Fields
	grbuild.label = usercount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FollowerCount int `json:"follower_count,omitempty"`
//	}
//
//	client.UserCount.Query().
//		Select(usercount.FieldFollowerCount).
//		Scan(ctx, &v)
func (ucq *UserCountQuery) Select(fields ...string) *UserCountSelect {
	ucq.ctx.Fields = append(ucq.ctx.Fields, fields...)
	sbuild := &UserCountSelect{UserCountQuery: ucq}
	sbuild.label = usercount.Label
	sbuild.flds, sbuild.scan = &ucq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserCountSelect configured with the given aggregations.
func (ucq *UserCountQuery) Aggregate(fns ...AggregateFunc) *UserCountSelect {
	return ucq.Select().Aggregate(fns...)
}

func (ucq *UserCountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ucq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ucq); err != nil {
				return err
			}
		}
	}
	for _, f := range ucq.ctx.Fields {
		if !usercount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ucq.path != nil {
		prev, err := ucq.path(ctx)
		if err != nil {
			return err
		}
		ucq.sql = prev
	}
	return nil
}

func (ucq *UserCountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserCount, error) {
	var (
		nodes       = []*UserCount{}
		_spec       = ucq.querySpec()
		loadedTypes = [1]bool{
			ucq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserCount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserCount{config: ucq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ucq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ucq.withUser; query != nil {
		if err := ucq.loadUser(ctx, query, nodes,
			func(n *UserCount) { n.Edges.User = []*UserAccount{} },
			func(n *UserCount, e *UserAccount) { n.Edges.User = append(n.Edges.User, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ucq *UserCountQuery) loadUser(ctx context.Context, query *UserAccountQuery, nodes []*UserCount, init func(*UserCount), assign func(*UserCount, *UserAccount)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserCount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.UserAccount(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(usercount.UserColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_account_user_count_info
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_account_user_count_info" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_account_user_count_info" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ucq *UserCountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ucq.querySpec()
	_spec.Node.Columns = ucq.ctx.Fields
	if len(ucq.ctx.Fields) > 0 {
		_spec.Unique = ucq.ctx.Unique != nil && *ucq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ucq.driver, _spec)
}

func (ucq *UserCountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(usercount.Table, usercount.Columns, sqlgraph.NewFieldSpec(usercount.FieldID, field.TypeInt))
	_spec.From = ucq.sql
	if unique := ucq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ucq.path != nil {
		_spec.Unique = true
	}
	if fields := ucq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercount.FieldID)
		for i := range fields {
			if fields[i] != usercount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ucq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ucq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ucq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ucq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ucq *UserCountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ucq.driver.Dialect())
	t1 := builder.Table(usercount.Table)
	columns := ucq.ctx.Fields
	if len(columns) == 0 {
		columns = usercount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ucq.sql != nil {
		selector = ucq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ucq.ctx.Unique != nil && *ucq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ucq.predicates {
		p(selector)
	}
	for _, p := range ucq.order {
		p(selector)
	}
	if offset := ucq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ucq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserCountGroupBy is the group-by builder for UserCount entities.
type UserCountGroupBy struct {
	selector
	build *UserCountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ucgb *UserCountGroupBy) Aggregate(fns ...AggregateFunc) *UserCountGroupBy {
	ucgb.fns = append(ucgb.fns, fns...)
	return ucgb
}

// Scan applies the selector query and scans the result into the given value.
func (ucgb *UserCountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucgb.build.ctx, "GroupBy")
	if err := ucgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCountQuery, *UserCountGroupBy](ctx, ucgb.build, ucgb, ucgb.build.inters, v)
}

func (ucgb *UserCountGroupBy) sqlScan(ctx context.Context, root *UserCountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ucgb.fns))
	for _, fn := range ucgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ucgb.flds)+len(ucgb.fns))
		for _, f := range *ucgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ucgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserCountSelect is the builder for selecting fields of UserCount entities.
type UserCountSelect struct {
	*UserCountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ucs *UserCountSelect) Aggregate(fns ...AggregateFunc) *UserCountSelect {
	ucs.fns = append(ucs.fns, fns...)
	return ucs
}

// Scan applies the selector query and scans the result into the given value.
func (ucs *UserCountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ucs.ctx, "Select")
	if err := ucs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserCountQuery, *UserCountSelect](ctx, ucs.UserCountQuery, ucs, ucs.inters, v)
}

func (ucs *UserCountSelect) sqlScan(ctx context.Context, root *UserCountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ucs.fns))
	for _, fn := range ucs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ucs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ucs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}