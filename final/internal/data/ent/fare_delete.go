// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"golang_weekend/final/internal/data/ent/fare"
	"golang_weekend/final/internal/data/ent/predicate"
)

// FareDelete is the builder for deleting a Fare entity.
type FareDelete struct {
	config
	hooks    []Hook
	mutation *FareMutation
}

// Where adds a new predicate to the FareDelete builder.
func (fd *FareDelete) Where(ps ...predicate.Fare) *FareDelete {
	fd.mutation.predicates = append(fd.mutation.predicates, ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FareDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(fd.hooks) == 0 {
		affected, err = fd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FareMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			fd.mutation = mutation
			affected, err = fd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(fd.hooks) - 1; i >= 0; i-- {
			mut = fd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FareDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FareDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: fare.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: fare.FieldID,
			},
		},
	}
	if ps := fd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, fd.driver, _spec)
}

// FareDeleteOne is the builder for deleting a single Fare entity.
type FareDeleteOne struct {
	fd *FareDelete
}

// Exec executes the deletion query.
func (fdo *FareDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fare.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FareDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
