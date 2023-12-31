// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/0xfzz/tuwit-backend/ent/thread"
	"github.com/0xfzz/tuwit-backend/ent/threadcount"
)

// ThreadCountCreate is the builder for creating a ThreadCount entity.
type ThreadCountCreate struct {
	config
	mutation *ThreadCountMutation
	hooks    []Hook
}

// SetReplyCount sets the "reply_count" field.
func (tcc *ThreadCountCreate) SetReplyCount(i int) *ThreadCountCreate {
	tcc.mutation.SetReplyCount(i)
	return tcc
}

// SetNillableReplyCount sets the "reply_count" field if the given value is not nil.
func (tcc *ThreadCountCreate) SetNillableReplyCount(i *int) *ThreadCountCreate {
	if i != nil {
		tcc.SetReplyCount(*i)
	}
	return tcc
}

// SetLikeCount sets the "like_count" field.
func (tcc *ThreadCountCreate) SetLikeCount(i int) *ThreadCountCreate {
	tcc.mutation.SetLikeCount(i)
	return tcc
}

// SetNillableLikeCount sets the "like_count" field if the given value is not nil.
func (tcc *ThreadCountCreate) SetNillableLikeCount(i *int) *ThreadCountCreate {
	if i != nil {
		tcc.SetLikeCount(*i)
	}
	return tcc
}

// SetThreadID sets the "thread" edge to the Thread entity by ID.
func (tcc *ThreadCountCreate) SetThreadID(id int) *ThreadCountCreate {
	tcc.mutation.SetThreadID(id)
	return tcc
}

// SetThread sets the "thread" edge to the Thread entity.
func (tcc *ThreadCountCreate) SetThread(t *Thread) *ThreadCountCreate {
	return tcc.SetThreadID(t.ID)
}

// Mutation returns the ThreadCountMutation object of the builder.
func (tcc *ThreadCountCreate) Mutation() *ThreadCountMutation {
	return tcc.mutation
}

// Save creates the ThreadCount in the database.
func (tcc *ThreadCountCreate) Save(ctx context.Context) (*ThreadCount, error) {
	tcc.defaults()
	return withHooks(ctx, tcc.sqlSave, tcc.mutation, tcc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tcc *ThreadCountCreate) SaveX(ctx context.Context) *ThreadCount {
	v, err := tcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcc *ThreadCountCreate) Exec(ctx context.Context) error {
	_, err := tcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcc *ThreadCountCreate) ExecX(ctx context.Context) {
	if err := tcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tcc *ThreadCountCreate) defaults() {
	if _, ok := tcc.mutation.ReplyCount(); !ok {
		v := threadcount.DefaultReplyCount
		tcc.mutation.SetReplyCount(v)
	}
	if _, ok := tcc.mutation.LikeCount(); !ok {
		v := threadcount.DefaultLikeCount
		tcc.mutation.SetLikeCount(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tcc *ThreadCountCreate) check() error {
	if _, ok := tcc.mutation.ReplyCount(); !ok {
		return &ValidationError{Name: "reply_count", err: errors.New(`ent: missing required field "ThreadCount.reply_count"`)}
	}
	if _, ok := tcc.mutation.LikeCount(); !ok {
		return &ValidationError{Name: "like_count", err: errors.New(`ent: missing required field "ThreadCount.like_count"`)}
	}
	if _, ok := tcc.mutation.ThreadID(); !ok {
		return &ValidationError{Name: "thread", err: errors.New(`ent: missing required edge "ThreadCount.thread"`)}
	}
	return nil
}

func (tcc *ThreadCountCreate) sqlSave(ctx context.Context) (*ThreadCount, error) {
	if err := tcc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	tcc.mutation.id = &_node.ID
	tcc.mutation.done = true
	return _node, nil
}

func (tcc *ThreadCountCreate) createSpec() (*ThreadCount, *sqlgraph.CreateSpec) {
	var (
		_node = &ThreadCount{config: tcc.config}
		_spec = sqlgraph.NewCreateSpec(threadcount.Table, sqlgraph.NewFieldSpec(threadcount.FieldID, field.TypeInt))
	)
	if value, ok := tcc.mutation.ReplyCount(); ok {
		_spec.SetField(threadcount.FieldReplyCount, field.TypeInt, value)
		_node.ReplyCount = value
	}
	if value, ok := tcc.mutation.LikeCount(); ok {
		_spec.SetField(threadcount.FieldLikeCount, field.TypeInt, value)
		_node.LikeCount = value
	}
	if nodes := tcc.mutation.ThreadIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   threadcount.ThreadTable,
			Columns: []string{threadcount.ThreadColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(thread.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.thread_thread_count = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ThreadCountCreateBulk is the builder for creating many ThreadCount entities in bulk.
type ThreadCountCreateBulk struct {
	config
	err      error
	builders []*ThreadCountCreate
}

// Save creates the ThreadCount entities in the database.
func (tccb *ThreadCountCreateBulk) Save(ctx context.Context) ([]*ThreadCount, error) {
	if tccb.err != nil {
		return nil, tccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tccb.builders))
	nodes := make([]*ThreadCount, len(tccb.builders))
	mutators := make([]Mutator, len(tccb.builders))
	for i := range tccb.builders {
		func(i int, root context.Context) {
			builder := tccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ThreadCountMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tccb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, tccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tccb *ThreadCountCreateBulk) SaveX(ctx context.Context) []*ThreadCount {
	v, err := tccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tccb *ThreadCountCreateBulk) Exec(ctx context.Context) error {
	_, err := tccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tccb *ThreadCountCreateBulk) ExecX(ctx context.Context) {
	if err := tccb.Exec(ctx); err != nil {
		panic(err)
	}
}
