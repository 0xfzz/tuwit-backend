// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwit-backend/ent/predicate"
	"github.com/0xfzz/tuwit-backend/ent/useraccount"
	"github.com/0xfzz/tuwit-backend/ent/usercount"
)

// UserCountUpdate is the builder for updating UserCount entities.
type UserCountUpdate struct {
	config
	hooks    []Hook
	mutation *UserCountMutation
}

// Where appends a list predicates to the UserCountUpdate builder.
func (ucu *UserCountUpdate) Where(ps ...predicate.UserCount) *UserCountUpdate {
	ucu.mutation.Where(ps...)
	return ucu
}

// SetFollowerCount sets the "follower_count" field.
func (ucu *UserCountUpdate) SetFollowerCount(i int) *UserCountUpdate {
	ucu.mutation.ResetFollowerCount()
	ucu.mutation.SetFollowerCount(i)
	return ucu
}

// SetNillableFollowerCount sets the "follower_count" field if the given value is not nil.
func (ucu *UserCountUpdate) SetNillableFollowerCount(i *int) *UserCountUpdate {
	if i != nil {
		ucu.SetFollowerCount(*i)
	}
	return ucu
}

// AddFollowerCount adds i to the "follower_count" field.
func (ucu *UserCountUpdate) AddFollowerCount(i int) *UserCountUpdate {
	ucu.mutation.AddFollowerCount(i)
	return ucu
}

// SetFollowingsCount sets the "followings_count" field.
func (ucu *UserCountUpdate) SetFollowingsCount(i int) *UserCountUpdate {
	ucu.mutation.ResetFollowingsCount()
	ucu.mutation.SetFollowingsCount(i)
	return ucu
}

// SetNillableFollowingsCount sets the "followings_count" field if the given value is not nil.
func (ucu *UserCountUpdate) SetNillableFollowingsCount(i *int) *UserCountUpdate {
	if i != nil {
		ucu.SetFollowingsCount(*i)
	}
	return ucu
}

// AddFollowingsCount adds i to the "followings_count" field.
func (ucu *UserCountUpdate) AddFollowingsCount(i int) *UserCountUpdate {
	ucu.mutation.AddFollowingsCount(i)
	return ucu
}

// SetUserID sets the "user" edge to the UserAccount entity by ID.
func (ucu *UserCountUpdate) SetUserID(id int) *UserCountUpdate {
	ucu.mutation.SetUserID(id)
	return ucu
}

// SetUser sets the "user" edge to the UserAccount entity.
func (ucu *UserCountUpdate) SetUser(u *UserAccount) *UserCountUpdate {
	return ucu.SetUserID(u.ID)
}

// Mutation returns the UserCountMutation object of the builder.
func (ucu *UserCountUpdate) Mutation() *UserCountMutation {
	return ucu.mutation
}

// ClearUser clears the "user" edge to the UserAccount entity.
func (ucu *UserCountUpdate) ClearUser() *UserCountUpdate {
	ucu.mutation.ClearUser()
	return ucu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ucu *UserCountUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ucu.sqlSave, ucu.mutation, ucu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucu *UserCountUpdate) SaveX(ctx context.Context) int {
	affected, err := ucu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ucu *UserCountUpdate) Exec(ctx context.Context) error {
	_, err := ucu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucu *UserCountUpdate) ExecX(ctx context.Context) {
	if err := ucu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucu *UserCountUpdate) check() error {
	if _, ok := ucu.mutation.UserID(); ucu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCount.user"`)
	}
	return nil
}

func (ucu *UserCountUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ucu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercount.Table, usercount.Columns, sqlgraph.NewFieldSpec(usercount.FieldID, field.TypeInt))
	if ps := ucu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucu.mutation.FollowerCount(); ok {
		_spec.SetField(usercount.FieldFollowerCount, field.TypeInt, value)
	}
	if value, ok := ucu.mutation.AddedFollowerCount(); ok {
		_spec.AddField(usercount.FieldFollowerCount, field.TypeInt, value)
	}
	if value, ok := ucu.mutation.FollowingsCount(); ok {
		_spec.SetField(usercount.FieldFollowingsCount, field.TypeInt, value)
	}
	if value, ok := ucu.mutation.AddedFollowingsCount(); ok {
		_spec.AddField(usercount.FieldFollowingsCount, field.TypeInt, value)
	}
	if ucu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usercount.UserTable,
			Columns: []string{usercount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usercount.UserTable,
			Columns: []string{usercount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ucu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ucu.mutation.done = true
	return n, nil
}

// UserCountUpdateOne is the builder for updating a single UserCount entity.
type UserCountUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserCountMutation
}

// SetFollowerCount sets the "follower_count" field.
func (ucuo *UserCountUpdateOne) SetFollowerCount(i int) *UserCountUpdateOne {
	ucuo.mutation.ResetFollowerCount()
	ucuo.mutation.SetFollowerCount(i)
	return ucuo
}

// SetNillableFollowerCount sets the "follower_count" field if the given value is not nil.
func (ucuo *UserCountUpdateOne) SetNillableFollowerCount(i *int) *UserCountUpdateOne {
	if i != nil {
		ucuo.SetFollowerCount(*i)
	}
	return ucuo
}

// AddFollowerCount adds i to the "follower_count" field.
func (ucuo *UserCountUpdateOne) AddFollowerCount(i int) *UserCountUpdateOne {
	ucuo.mutation.AddFollowerCount(i)
	return ucuo
}

// SetFollowingsCount sets the "followings_count" field.
func (ucuo *UserCountUpdateOne) SetFollowingsCount(i int) *UserCountUpdateOne {
	ucuo.mutation.ResetFollowingsCount()
	ucuo.mutation.SetFollowingsCount(i)
	return ucuo
}

// SetNillableFollowingsCount sets the "followings_count" field if the given value is not nil.
func (ucuo *UserCountUpdateOne) SetNillableFollowingsCount(i *int) *UserCountUpdateOne {
	if i != nil {
		ucuo.SetFollowingsCount(*i)
	}
	return ucuo
}

// AddFollowingsCount adds i to the "followings_count" field.
func (ucuo *UserCountUpdateOne) AddFollowingsCount(i int) *UserCountUpdateOne {
	ucuo.mutation.AddFollowingsCount(i)
	return ucuo
}

// SetUserID sets the "user" edge to the UserAccount entity by ID.
func (ucuo *UserCountUpdateOne) SetUserID(id int) *UserCountUpdateOne {
	ucuo.mutation.SetUserID(id)
	return ucuo
}

// SetUser sets the "user" edge to the UserAccount entity.
func (ucuo *UserCountUpdateOne) SetUser(u *UserAccount) *UserCountUpdateOne {
	return ucuo.SetUserID(u.ID)
}

// Mutation returns the UserCountMutation object of the builder.
func (ucuo *UserCountUpdateOne) Mutation() *UserCountMutation {
	return ucuo.mutation
}

// ClearUser clears the "user" edge to the UserAccount entity.
func (ucuo *UserCountUpdateOne) ClearUser() *UserCountUpdateOne {
	ucuo.mutation.ClearUser()
	return ucuo
}

// Where appends a list predicates to the UserCountUpdate builder.
func (ucuo *UserCountUpdateOne) Where(ps ...predicate.UserCount) *UserCountUpdateOne {
	ucuo.mutation.Where(ps...)
	return ucuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ucuo *UserCountUpdateOne) Select(field string, fields ...string) *UserCountUpdateOne {
	ucuo.fields = append([]string{field}, fields...)
	return ucuo
}

// Save executes the query and returns the updated UserCount entity.
func (ucuo *UserCountUpdateOne) Save(ctx context.Context) (*UserCount, error) {
	return withHooks(ctx, ucuo.sqlSave, ucuo.mutation, ucuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ucuo *UserCountUpdateOne) SaveX(ctx context.Context) *UserCount {
	node, err := ucuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ucuo *UserCountUpdateOne) Exec(ctx context.Context) error {
	_, err := ucuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ucuo *UserCountUpdateOne) ExecX(ctx context.Context) {
	if err := ucuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ucuo *UserCountUpdateOne) check() error {
	if _, ok := ucuo.mutation.UserID(); ucuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "UserCount.user"`)
	}
	return nil
}

func (ucuo *UserCountUpdateOne) sqlSave(ctx context.Context) (_node *UserCount, err error) {
	if err := ucuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(usercount.Table, usercount.Columns, sqlgraph.NewFieldSpec(usercount.FieldID, field.TypeInt))
	id, ok := ucuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserCount.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ucuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usercount.FieldID)
		for _, f := range fields {
			if !usercount.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usercount.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ucuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ucuo.mutation.FollowerCount(); ok {
		_spec.SetField(usercount.FieldFollowerCount, field.TypeInt, value)
	}
	if value, ok := ucuo.mutation.AddedFollowerCount(); ok {
		_spec.AddField(usercount.FieldFollowerCount, field.TypeInt, value)
	}
	if value, ok := ucuo.mutation.FollowingsCount(); ok {
		_spec.SetField(usercount.FieldFollowingsCount, field.TypeInt, value)
	}
	if value, ok := ucuo.mutation.AddedFollowingsCount(); ok {
		_spec.AddField(usercount.FieldFollowingsCount, field.TypeInt, value)
	}
	if ucuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usercount.UserTable,
			Columns: []string{usercount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ucuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   usercount.UserTable,
			Columns: []string{usercount.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(useraccount.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &UserCount{config: ucuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ucuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usercount.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ucuo.mutation.done = true
	return _node, nil
}
