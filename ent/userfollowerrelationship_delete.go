// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwitt/ent/predicate"
	"github.com/0xfzz/tuwitt/ent/userfollowerrelationship"
)

// UserFollowerRelationshipDelete is the builder for deleting a UserFollowerRelationship entity.
type UserFollowerRelationshipDelete struct {
	config
	hooks    []Hook
	mutation *UserFollowerRelationshipMutation
}

// Where appends a list predicates to the UserFollowerRelationshipDelete builder.
func (ufrd *UserFollowerRelationshipDelete) Where(ps ...predicate.UserFollowerRelationship) *UserFollowerRelationshipDelete {
	ufrd.mutation.Where(ps...)
	return ufrd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ufrd *UserFollowerRelationshipDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ufrd.sqlExec, ufrd.mutation, ufrd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ufrd *UserFollowerRelationshipDelete) ExecX(ctx context.Context) int {
	n, err := ufrd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ufrd *UserFollowerRelationshipDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(userfollowerrelationship.Table, sqlgraph.NewFieldSpec(userfollowerrelationship.FieldID, field.TypeInt))
	if ps := ufrd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ufrd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ufrd.mutation.done = true
	return affected, err
}

// UserFollowerRelationshipDeleteOne is the builder for deleting a single UserFollowerRelationship entity.
type UserFollowerRelationshipDeleteOne struct {
	ufrd *UserFollowerRelationshipDelete
}

// Where appends a list predicates to the UserFollowerRelationshipDelete builder.
func (ufrdo *UserFollowerRelationshipDeleteOne) Where(ps ...predicate.UserFollowerRelationship) *UserFollowerRelationshipDeleteOne {
	ufrdo.ufrd.mutation.Where(ps...)
	return ufrdo
}

// Exec executes the deletion query.
func (ufrdo *UserFollowerRelationshipDeleteOne) Exec(ctx context.Context) error {
	n, err := ufrdo.ufrd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{userfollowerrelationship.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ufrdo *UserFollowerRelationshipDeleteOne) ExecX(ctx context.Context) {
	if err := ufrdo.Exec(ctx); err != nil {
		panic(err)
	}
}
