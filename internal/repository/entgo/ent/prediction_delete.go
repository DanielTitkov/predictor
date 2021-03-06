// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/prediction"
)

// PredictionDelete is the builder for deleting a Prediction entity.
type PredictionDelete struct {
	config
	hooks    []Hook
	mutation *PredictionMutation
}

// Where appends a list predicates to the PredictionDelete builder.
func (pd *PredictionDelete) Where(ps ...predicate.Prediction) *PredictionDelete {
	pd.mutation.Where(ps...)
	return pd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pd *PredictionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pd.hooks) == 0 {
		affected, err = pd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PredictionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pd.mutation = mutation
			affected, err = pd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pd.hooks) - 1; i >= 0; i-- {
			if pd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PredictionDelete) ExecX(ctx context.Context) int {
	n, err := pd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pd *PredictionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: prediction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: prediction.FieldID,
			},
		},
	}
	if ps := pd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pd.driver, _spec)
}

// PredictionDeleteOne is the builder for deleting a single Prediction entity.
type PredictionDeleteOne struct {
	pd *PredictionDelete
}

// Exec executes the deletion query.
func (pdo *PredictionDeleteOne) Exec(ctx context.Context) error {
	n, err := pdo.pd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{prediction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PredictionDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
