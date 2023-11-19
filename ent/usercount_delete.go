// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwit-backend/ent/predicate"
	"github.com/0xfzz/tuwit-backend/ent/usercount"
)

// UserCountDelete is the builder for deleting a UserCount entity.
type UserCountDelete struct {
	config
	hooks    []Hook
	mutation *UserCountMutation
}

// Where appends a list predicates to the UserCountDelete builder.
func (ucd *UserCountDelete) Where(ps ...predicate.UserCount) *UserCountDelete {
	ucd.mutation.Where(ps...)
	return ucd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ucd *UserCountDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ucd.sqlExec, ucd.mutation, ucd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ucd *UserCountDelete) ExecX(ctx context.Context) int {
	n, err := ucd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ucd *UserCountDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(usercount.Table, sqlgraph.NewFieldSpec(usercount.FieldID, field.TypeInt))
	if ps := ucd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ucd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ucd.mutation.done = true
	return affected, err
}

// UserCountDeleteOne is the builder for deleting a single UserCount entity.
type UserCountDeleteOne struct {
	ucd *UserCountDelete
}

// Where appends a list predicates to the UserCountDelete builder.
func (ucdo *UserCountDeleteOne) Where(ps ...predicate.UserCount) *UserCountDeleteOne {
	ucdo.ucd.mutation.Where(ps...)
	return ucdo
}

// Exec executes the deletion query.
func (ucdo *UserCountDeleteOne) Exec(ctx context.Context) error {
	n, err := ucdo.ucd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{usercount.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ucdo *UserCountDeleteOne) ExecX(ctx context.Context) {
	if err := ucdo.Exec(ctx); err != nil {
		panic(err)
	}
}
