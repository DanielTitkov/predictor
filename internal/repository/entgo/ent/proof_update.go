// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/challenge"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/predicate"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/proof"
	"github.com/google/uuid"
)

// ProofUpdate is the builder for updating Proof entities.
type ProofUpdate struct {
	config
	hooks    []Hook
	mutation *ProofMutation
}

// Where appends a list predicates to the ProofUpdate builder.
func (pu *ProofUpdate) Where(ps ...predicate.Proof) *ProofUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetUpdateTime sets the "update_time" field.
func (pu *ProofUpdate) SetUpdateTime(t time.Time) *ProofUpdate {
	pu.mutation.SetUpdateTime(t)
	return pu
}

// SetMeta sets the "meta" field.
func (pu *ProofUpdate) SetMeta(m map[string]interface{}) *ProofUpdate {
	pu.mutation.SetMeta(m)
	return pu
}

// ClearMeta clears the value of the "meta" field.
func (pu *ProofUpdate) ClearMeta() *ProofUpdate {
	pu.mutation.ClearMeta()
	return pu
}

// SetChallengeID sets the "challenge" edge to the Challenge entity by ID.
func (pu *ProofUpdate) SetChallengeID(id uuid.UUID) *ProofUpdate {
	pu.mutation.SetChallengeID(id)
	return pu
}

// SetChallenge sets the "challenge" edge to the Challenge entity.
func (pu *ProofUpdate) SetChallenge(c *Challenge) *ProofUpdate {
	return pu.SetChallengeID(c.ID)
}

// Mutation returns the ProofMutation object of the builder.
func (pu *ProofUpdate) Mutation() *ProofMutation {
	return pu.mutation
}

// ClearChallenge clears the "challenge" edge to the Challenge entity.
func (pu *ProofUpdate) ClearChallenge() *ProofUpdate {
	pu.mutation.ClearChallenge()
	return pu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *ProofUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	pu.defaults()
	if len(pu.hooks) == 0 {
		if err = pu.check(); err != nil {
			return 0, err
		}
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProofMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pu.check(); err != nil {
				return 0, err
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *ProofUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *ProofUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *ProofUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pu *ProofUpdate) defaults() {
	if _, ok := pu.mutation.UpdateTime(); !ok {
		v := proof.UpdateDefaultUpdateTime()
		pu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pu *ProofUpdate) check() error {
	if _, ok := pu.mutation.ChallengeID(); pu.mutation.ChallengeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Proof.challenge"`)
	}
	return nil
}

func (pu *ProofUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   proof.Table,
			Columns: proof.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: proof.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: proof.FieldUpdateTime,
		})
	}
	if value, ok := pu.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: proof.FieldMeta,
		})
	}
	if pu.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: proof.FieldMeta,
		})
	}
	if pu.mutation.ChallengeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proof.ChallengeTable,
			Columns: []string{proof.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.ChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proof.ChallengeTable,
			Columns: []string{proof.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{proof.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ProofUpdateOne is the builder for updating a single Proof entity.
type ProofUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProofMutation
}

// SetUpdateTime sets the "update_time" field.
func (puo *ProofUpdateOne) SetUpdateTime(t time.Time) *ProofUpdateOne {
	puo.mutation.SetUpdateTime(t)
	return puo
}

// SetMeta sets the "meta" field.
func (puo *ProofUpdateOne) SetMeta(m map[string]interface{}) *ProofUpdateOne {
	puo.mutation.SetMeta(m)
	return puo
}

// ClearMeta clears the value of the "meta" field.
func (puo *ProofUpdateOne) ClearMeta() *ProofUpdateOne {
	puo.mutation.ClearMeta()
	return puo
}

// SetChallengeID sets the "challenge" edge to the Challenge entity by ID.
func (puo *ProofUpdateOne) SetChallengeID(id uuid.UUID) *ProofUpdateOne {
	puo.mutation.SetChallengeID(id)
	return puo
}

// SetChallenge sets the "challenge" edge to the Challenge entity.
func (puo *ProofUpdateOne) SetChallenge(c *Challenge) *ProofUpdateOne {
	return puo.SetChallengeID(c.ID)
}

// Mutation returns the ProofMutation object of the builder.
func (puo *ProofUpdateOne) Mutation() *ProofMutation {
	return puo.mutation
}

// ClearChallenge clears the "challenge" edge to the Challenge entity.
func (puo *ProofUpdateOne) ClearChallenge() *ProofUpdateOne {
	puo.mutation.ClearChallenge()
	return puo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *ProofUpdateOne) Select(field string, fields ...string) *ProofUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Proof entity.
func (puo *ProofUpdateOne) Save(ctx context.Context) (*Proof, error) {
	var (
		err  error
		node *Proof
	)
	puo.defaults()
	if len(puo.hooks) == 0 {
		if err = puo.check(); err != nil {
			return nil, err
		}
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ProofMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = puo.check(); err != nil {
				return nil, err
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *ProofUpdateOne) SaveX(ctx context.Context) *Proof {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *ProofUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *ProofUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (puo *ProofUpdateOne) defaults() {
	if _, ok := puo.mutation.UpdateTime(); !ok {
		v := proof.UpdateDefaultUpdateTime()
		puo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (puo *ProofUpdateOne) check() error {
	if _, ok := puo.mutation.ChallengeID(); puo.mutation.ChallengeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Proof.challenge"`)
	}
	return nil
}

func (puo *ProofUpdateOne) sqlSave(ctx context.Context) (_node *Proof, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   proof.Table,
			Columns: proof.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: proof.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Proof.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, proof.FieldID)
		for _, f := range fields {
			if !proof.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != proof.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: proof.FieldUpdateTime,
		})
	}
	if value, ok := puo.mutation.Meta(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: proof.FieldMeta,
		})
	}
	if puo.mutation.MetaCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: proof.FieldMeta,
		})
	}
	if puo.mutation.ChallengeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proof.ChallengeTable,
			Columns: []string{proof.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.ChallengeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proof.ChallengeTable,
			Columns: []string{proof.ChallengeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: challenge.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Proof{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{proof.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}