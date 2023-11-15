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
	"github.com/0xfzz/tuwitt/ent/blockedusersrelationship"
	"github.com/0xfzz/tuwitt/ent/predicate"
	"github.com/0xfzz/tuwitt/ent/thread"
	"github.com/0xfzz/tuwitt/ent/useraccount"
	"github.com/0xfzz/tuwitt/ent/usercount"
	"github.com/0xfzz/tuwitt/ent/userfollowerrelationship"
	"github.com/0xfzz/tuwitt/ent/userprofile"
)

// UserAccountQuery is the builder for querying UserAccount entities.
type UserAccountQuery struct {
	config
	ctx              *QueryContext
	order            []useraccount.OrderOption
	inters           []Interceptor
	predicates       []predicate.UserAccount
	withProfile      *UserProfileQuery
	withFollowers    *UserFollowerRelationshipQuery
	withFollowings   *UserFollowerRelationshipQuery
	withBlockedBy    *BlockedUsersRelationshipQuery
	withBlockedUsers *BlockedUsersRelationshipQuery
	withUserCount    *UserCountQuery
	withThreads      *ThreadQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserAccountQuery builder.
func (uaq *UserAccountQuery) Where(ps ...predicate.UserAccount) *UserAccountQuery {
	uaq.predicates = append(uaq.predicates, ps...)
	return uaq
}

// Limit the number of records to be returned by this query.
func (uaq *UserAccountQuery) Limit(limit int) *UserAccountQuery {
	uaq.ctx.Limit = &limit
	return uaq
}

// Offset to start from.
func (uaq *UserAccountQuery) Offset(offset int) *UserAccountQuery {
	uaq.ctx.Offset = &offset
	return uaq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uaq *UserAccountQuery) Unique(unique bool) *UserAccountQuery {
	uaq.ctx.Unique = &unique
	return uaq
}

// Order specifies how the records should be ordered.
func (uaq *UserAccountQuery) Order(o ...useraccount.OrderOption) *UserAccountQuery {
	uaq.order = append(uaq.order, o...)
	return uaq
}

// QueryProfile chains the current query on the "profile" edge.
func (uaq *UserAccountQuery) QueryProfile() *UserProfileQuery {
	query := (&UserProfileClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(userprofile.Table, userprofile.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, useraccount.ProfileTable, useraccount.ProfileColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFollowers chains the current query on the "followers" edge.
func (uaq *UserAccountQuery) QueryFollowers() *UserFollowerRelationshipQuery {
	query := (&UserFollowerRelationshipClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(userfollowerrelationship.Table, userfollowerrelationship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, useraccount.FollowersTable, useraccount.FollowersColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFollowings chains the current query on the "followings" edge.
func (uaq *UserAccountQuery) QueryFollowings() *UserFollowerRelationshipQuery {
	query := (&UserFollowerRelationshipClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(userfollowerrelationship.Table, userfollowerrelationship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, useraccount.FollowingsTable, useraccount.FollowingsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBlockedBy chains the current query on the "blocked_by" edge.
func (uaq *UserAccountQuery) QueryBlockedBy() *BlockedUsersRelationshipQuery {
	query := (&BlockedUsersRelationshipClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(blockedusersrelationship.Table, blockedusersrelationship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, useraccount.BlockedByTable, useraccount.BlockedByColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryBlockedUsers chains the current query on the "blocked_users" edge.
func (uaq *UserAccountQuery) QueryBlockedUsers() *BlockedUsersRelationshipQuery {
	query := (&BlockedUsersRelationshipClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(blockedusersrelationship.Table, blockedusersrelationship.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, useraccount.BlockedUsersTable, useraccount.BlockedUsersColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUserCount chains the current query on the "user_count" edge.
func (uaq *UserAccountQuery) QueryUserCount() *UserCountQuery {
	query := (&UserCountClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(usercount.Table, usercount.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, useraccount.UserCountTable, useraccount.UserCountColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryThreads chains the current query on the "threads" edge.
func (uaq *UserAccountQuery) QueryThreads() *ThreadQuery {
	query := (&ThreadClient{config: uaq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := uaq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := uaq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(useraccount.Table, useraccount.FieldID, selector),
			sqlgraph.To(thread.Table, thread.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, useraccount.ThreadsTable, useraccount.ThreadsColumn),
		)
		fromU = sqlgraph.SetNeighbors(uaq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UserAccount entity from the query.
// Returns a *NotFoundError when no UserAccount was found.
func (uaq *UserAccountQuery) First(ctx context.Context) (*UserAccount, error) {
	nodes, err := uaq.Limit(1).All(setContextOp(ctx, uaq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{useraccount.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uaq *UserAccountQuery) FirstX(ctx context.Context) *UserAccount {
	node, err := uaq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserAccount ID from the query.
// Returns a *NotFoundError when no UserAccount ID was found.
func (uaq *UserAccountQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(1).IDs(setContextOp(ctx, uaq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{useraccount.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uaq *UserAccountQuery) FirstIDX(ctx context.Context) int {
	id, err := uaq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserAccount entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserAccount entity is found.
// Returns a *NotFoundError when no UserAccount entities are found.
func (uaq *UserAccountQuery) Only(ctx context.Context) (*UserAccount, error) {
	nodes, err := uaq.Limit(2).All(setContextOp(ctx, uaq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{useraccount.Label}
	default:
		return nil, &NotSingularError{useraccount.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uaq *UserAccountQuery) OnlyX(ctx context.Context) *UserAccount {
	node, err := uaq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserAccount ID in the query.
// Returns a *NotSingularError when more than one UserAccount ID is found.
// Returns a *NotFoundError when no entities are found.
func (uaq *UserAccountQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = uaq.Limit(2).IDs(setContextOp(ctx, uaq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{useraccount.Label}
	default:
		err = &NotSingularError{useraccount.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uaq *UserAccountQuery) OnlyIDX(ctx context.Context) int {
	id, err := uaq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserAccounts.
func (uaq *UserAccountQuery) All(ctx context.Context) ([]*UserAccount, error) {
	ctx = setContextOp(ctx, uaq.ctx, "All")
	if err := uaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UserAccount, *UserAccountQuery]()
	return withInterceptors[[]*UserAccount](ctx, uaq, qr, uaq.inters)
}

// AllX is like All, but panics if an error occurs.
func (uaq *UserAccountQuery) AllX(ctx context.Context) []*UserAccount {
	nodes, err := uaq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserAccount IDs.
func (uaq *UserAccountQuery) IDs(ctx context.Context) (ids []int, err error) {
	if uaq.ctx.Unique == nil && uaq.path != nil {
		uaq.Unique(true)
	}
	ctx = setContextOp(ctx, uaq.ctx, "IDs")
	if err = uaq.Select(useraccount.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uaq *UserAccountQuery) IDsX(ctx context.Context) []int {
	ids, err := uaq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uaq *UserAccountQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, uaq.ctx, "Count")
	if err := uaq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, uaq, querierCount[*UserAccountQuery](), uaq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (uaq *UserAccountQuery) CountX(ctx context.Context) int {
	count, err := uaq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uaq *UserAccountQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, uaq.ctx, "Exist")
	switch _, err := uaq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (uaq *UserAccountQuery) ExistX(ctx context.Context) bool {
	exist, err := uaq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserAccountQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uaq *UserAccountQuery) Clone() *UserAccountQuery {
	if uaq == nil {
		return nil
	}
	return &UserAccountQuery{
		config:           uaq.config,
		ctx:              uaq.ctx.Clone(),
		order:            append([]useraccount.OrderOption{}, uaq.order...),
		inters:           append([]Interceptor{}, uaq.inters...),
		predicates:       append([]predicate.UserAccount{}, uaq.predicates...),
		withProfile:      uaq.withProfile.Clone(),
		withFollowers:    uaq.withFollowers.Clone(),
		withFollowings:   uaq.withFollowings.Clone(),
		withBlockedBy:    uaq.withBlockedBy.Clone(),
		withBlockedUsers: uaq.withBlockedUsers.Clone(),
		withUserCount:    uaq.withUserCount.Clone(),
		withThreads:      uaq.withThreads.Clone(),
		// clone intermediate query.
		sql:  uaq.sql.Clone(),
		path: uaq.path,
	}
}

// WithProfile tells the query-builder to eager-load the nodes that are connected to
// the "profile" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithProfile(opts ...func(*UserProfileQuery)) *UserAccountQuery {
	query := (&UserProfileClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withProfile = query
	return uaq
}

// WithFollowers tells the query-builder to eager-load the nodes that are connected to
// the "followers" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithFollowers(opts ...func(*UserFollowerRelationshipQuery)) *UserAccountQuery {
	query := (&UserFollowerRelationshipClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withFollowers = query
	return uaq
}

// WithFollowings tells the query-builder to eager-load the nodes that are connected to
// the "followings" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithFollowings(opts ...func(*UserFollowerRelationshipQuery)) *UserAccountQuery {
	query := (&UserFollowerRelationshipClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withFollowings = query
	return uaq
}

// WithBlockedBy tells the query-builder to eager-load the nodes that are connected to
// the "blocked_by" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithBlockedBy(opts ...func(*BlockedUsersRelationshipQuery)) *UserAccountQuery {
	query := (&BlockedUsersRelationshipClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withBlockedBy = query
	return uaq
}

// WithBlockedUsers tells the query-builder to eager-load the nodes that are connected to
// the "blocked_users" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithBlockedUsers(opts ...func(*BlockedUsersRelationshipQuery)) *UserAccountQuery {
	query := (&BlockedUsersRelationshipClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withBlockedUsers = query
	return uaq
}

// WithUserCount tells the query-builder to eager-load the nodes that are connected to
// the "user_count" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithUserCount(opts ...func(*UserCountQuery)) *UserAccountQuery {
	query := (&UserCountClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withUserCount = query
	return uaq
}

// WithThreads tells the query-builder to eager-load the nodes that are connected to
// the "threads" edge. The optional arguments are used to configure the query builder of the edge.
func (uaq *UserAccountQuery) WithThreads(opts ...func(*ThreadQuery)) *UserAccountQuery {
	query := (&ThreadClient{config: uaq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	uaq.withThreads = query
	return uaq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserAccount.Query().
//		GroupBy(useraccount.FieldEmail).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (uaq *UserAccountQuery) GroupBy(field string, fields ...string) *UserAccountGroupBy {
	uaq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UserAccountGroupBy{build: uaq}
	grbuild.flds = &uaq.ctx.Fields
	grbuild.label = useraccount.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Email string `json:"email,omitempty"`
//	}
//
//	client.UserAccount.Query().
//		Select(useraccount.FieldEmail).
//		Scan(ctx, &v)
func (uaq *UserAccountQuery) Select(fields ...string) *UserAccountSelect {
	uaq.ctx.Fields = append(uaq.ctx.Fields, fields...)
	sbuild := &UserAccountSelect{UserAccountQuery: uaq}
	sbuild.label = useraccount.Label
	sbuild.flds, sbuild.scan = &uaq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UserAccountSelect configured with the given aggregations.
func (uaq *UserAccountQuery) Aggregate(fns ...AggregateFunc) *UserAccountSelect {
	return uaq.Select().Aggregate(fns...)
}

func (uaq *UserAccountQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range uaq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, uaq); err != nil {
				return err
			}
		}
	}
	for _, f := range uaq.ctx.Fields {
		if !useraccount.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if uaq.path != nil {
		prev, err := uaq.path(ctx)
		if err != nil {
			return err
		}
		uaq.sql = prev
	}
	return nil
}

func (uaq *UserAccountQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UserAccount, error) {
	var (
		nodes       = []*UserAccount{}
		_spec       = uaq.querySpec()
		loadedTypes = [7]bool{
			uaq.withProfile != nil,
			uaq.withFollowers != nil,
			uaq.withFollowings != nil,
			uaq.withBlockedBy != nil,
			uaq.withBlockedUsers != nil,
			uaq.withUserCount != nil,
			uaq.withThreads != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UserAccount).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UserAccount{config: uaq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, uaq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := uaq.withProfile; query != nil {
		if err := uaq.loadProfile(ctx, query, nodes, nil,
			func(n *UserAccount, e *UserProfile) { n.Edges.Profile = e }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withFollowers; query != nil {
		if err := uaq.loadFollowers(ctx, query, nodes,
			func(n *UserAccount) { n.Edges.Followers = []*UserFollowerRelationship{} },
			func(n *UserAccount, e *UserFollowerRelationship) { n.Edges.Followers = append(n.Edges.Followers, e) }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withFollowings; query != nil {
		if err := uaq.loadFollowings(ctx, query, nodes,
			func(n *UserAccount) { n.Edges.Followings = []*UserFollowerRelationship{} },
			func(n *UserAccount, e *UserFollowerRelationship) { n.Edges.Followings = append(n.Edges.Followings, e) }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withBlockedBy; query != nil {
		if err := uaq.loadBlockedBy(ctx, query, nodes,
			func(n *UserAccount) { n.Edges.BlockedBy = []*BlockedUsersRelationship{} },
			func(n *UserAccount, e *BlockedUsersRelationship) { n.Edges.BlockedBy = append(n.Edges.BlockedBy, e) }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withBlockedUsers; query != nil {
		if err := uaq.loadBlockedUsers(ctx, query, nodes,
			func(n *UserAccount) { n.Edges.BlockedUsers = []*BlockedUsersRelationship{} },
			func(n *UserAccount, e *BlockedUsersRelationship) {
				n.Edges.BlockedUsers = append(n.Edges.BlockedUsers, e)
			}); err != nil {
			return nil, err
		}
	}
	if query := uaq.withUserCount; query != nil {
		if err := uaq.loadUserCount(ctx, query, nodes, nil,
			func(n *UserAccount, e *UserCount) { n.Edges.UserCount = e }); err != nil {
			return nil, err
		}
	}
	if query := uaq.withThreads; query != nil {
		if err := uaq.loadThreads(ctx, query, nodes,
			func(n *UserAccount) { n.Edges.Threads = []*Thread{} },
			func(n *UserAccount, e *Thread) { n.Edges.Threads = append(n.Edges.Threads, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (uaq *UserAccountQuery) loadProfile(ctx context.Context, query *UserProfileQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *UserProfile)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.UserProfile(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.ProfileColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_account_profile
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_account_profile" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_account_profile" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uaq *UserAccountQuery) loadFollowers(ctx context.Context, query *UserFollowerRelationshipQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *UserFollowerRelationship)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userfollowerrelationship.FieldUserID)
	}
	query.Where(predicate.UserFollowerRelationship(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.FollowersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uaq *UserAccountQuery) loadFollowings(ctx context.Context, query *UserFollowerRelationshipQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *UserFollowerRelationship)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(userfollowerrelationship.FieldFollowerID)
	}
	query.Where(predicate.UserFollowerRelationship(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.FollowingsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.FollowerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "follower_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uaq *UserAccountQuery) loadBlockedBy(ctx context.Context, query *BlockedUsersRelationshipQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *BlockedUsersRelationship)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(blockedusersrelationship.FieldUserID)
	}
	query.Where(predicate.BlockedUsersRelationship(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.BlockedByColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UserID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uaq *UserAccountQuery) loadBlockedUsers(ctx context.Context, query *BlockedUsersRelationshipQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *BlockedUsersRelationship)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(blockedusersrelationship.FieldBlockerID)
	}
	query.Where(predicate.BlockedUsersRelationship(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.BlockedUsersColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.BlockerID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "blocker_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uaq *UserAccountQuery) loadUserCount(ctx context.Context, query *UserCountQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *UserCount)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
	}
	query.withFKs = true
	query.Where(predicate.UserCount(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.UserCountColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.user_account_user_count
		if fk == nil {
			return fmt.Errorf(`foreign-key "user_account_user_count" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "user_account_user_count" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (uaq *UserAccountQuery) loadThreads(ctx context.Context, query *ThreadQuery, nodes []*UserAccount, init func(*UserAccount), assign func(*UserAccount, *Thread)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UserAccount)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	if len(query.ctx.Fields) > 0 {
		query.ctx.AppendFieldOnce(thread.FieldAuthorID)
	}
	query.Where(predicate.Thread(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(useraccount.ThreadsColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.AuthorID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "author_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (uaq *UserAccountQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uaq.querySpec()
	_spec.Node.Columns = uaq.ctx.Fields
	if len(uaq.ctx.Fields) > 0 {
		_spec.Unique = uaq.ctx.Unique != nil && *uaq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, uaq.driver, _spec)
}

func (uaq *UserAccountQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(useraccount.Table, useraccount.Columns, sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt))
	_spec.From = uaq.sql
	if unique := uaq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if uaq.path != nil {
		_spec.Unique = true
	}
	if fields := uaq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, useraccount.FieldID)
		for i := range fields {
			if fields[i] != useraccount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uaq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uaq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uaq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uaq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uaq *UserAccountQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uaq.driver.Dialect())
	t1 := builder.Table(useraccount.Table)
	columns := uaq.ctx.Fields
	if len(columns) == 0 {
		columns = useraccount.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uaq.sql != nil {
		selector = uaq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uaq.ctx.Unique != nil && *uaq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range uaq.predicates {
		p(selector)
	}
	for _, p := range uaq.order {
		p(selector)
	}
	if offset := uaq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uaq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserAccountGroupBy is the group-by builder for UserAccount entities.
type UserAccountGroupBy struct {
	selector
	build *UserAccountQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uagb *UserAccountGroupBy) Aggregate(fns ...AggregateFunc) *UserAccountGroupBy {
	uagb.fns = append(uagb.fns, fns...)
	return uagb
}

// Scan applies the selector query and scans the result into the given value.
func (uagb *UserAccountGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uagb.build.ctx, "GroupBy")
	if err := uagb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAccountQuery, *UserAccountGroupBy](ctx, uagb.build, uagb, uagb.build.inters, v)
}

func (uagb *UserAccountGroupBy) sqlScan(ctx context.Context, root *UserAccountQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(uagb.fns))
	for _, fn := range uagb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*uagb.flds)+len(uagb.fns))
		for _, f := range *uagb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*uagb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uagb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UserAccountSelect is the builder for selecting fields of UserAccount entities.
type UserAccountSelect struct {
	*UserAccountQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (uas *UserAccountSelect) Aggregate(fns ...AggregateFunc) *UserAccountSelect {
	uas.fns = append(uas.fns, fns...)
	return uas
}

// Scan applies the selector query and scans the result into the given value.
func (uas *UserAccountSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, uas.ctx, "Select")
	if err := uas.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UserAccountQuery, *UserAccountSelect](ctx, uas.UserAccountQuery, uas, uas.inters, v)
}

func (uas *UserAccountSelect) sqlScan(ctx context.Context, root *UserAccountQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(uas.fns))
	for _, fn := range uas.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*uas.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uas.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
