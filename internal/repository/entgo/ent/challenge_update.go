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
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/prediction"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/proof"
	"github.com/DanielTitkov/predictor/internal/repository/entgo/ent/user"
	"github.com/google/uuid"
)

// ChallengeUpdate is the builder for updating Challenge entities.
type ChallengeUpdate struct {
	config
	hooks    []Hook
	mutation *ChallengeMutation
}

// Where appends a list predicates to the ChallengeUpdate builder.
func (cu *ChallengeUpdate) Where(ps ...predicate.Challenge) *ChallengeUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *ChallengeUpdate) SetUpdateTime(t time.Time) *ChallengeUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetContent sets the "content" field.
func (cu *ChallengeUpdate) SetContent(s string) *ChallengeUpdate {
	cu.mutation.SetContent(s)
	return cu
}

// SetDescription sets the "description" field.
func (cu *ChallengeUpdate) SetDescription(s string) *ChallengeUpdate {
	cu.mutation.SetDescription(s)
	return cu
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableDescription(s *string) *ChallengeUpdate {
	if s != nil {
		cu.SetDescription(*s)
	}
	return cu
}

// ClearDescription clears the value of the "description" field.
func (cu *ChallengeUpdate) ClearDescription() *ChallengeUpdate {
	cu.mutation.ClearDescription()
	return cu
}

// SetOutcome sets the "outcome" field.
func (cu *ChallengeUpdate) SetOutcome(b bool) *ChallengeUpdate {
	cu.mutation.SetOutcome(b)
	return cu
}

// SetNillableOutcome sets the "outcome" field if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableOutcome(b *bool) *ChallengeUpdate {
	if b != nil {
		cu.SetOutcome(*b)
	}
	return cu
}

// ClearOutcome clears the value of the "outcome" field.
func (cu *ChallengeUpdate) ClearOutcome() *ChallengeUpdate {
	cu.mutation.ClearOutcome()
	return cu
}

// SetPublished sets the "published" field.
func (cu *ChallengeUpdate) SetPublished(b bool) *ChallengeUpdate {
	cu.mutation.SetPublished(b)
	return cu
}

// SetNillablePublished sets the "published" field if the given value is not nil.
func (cu *ChallengeUpdate) SetNillablePublished(b *bool) *ChallengeUpdate {
	if b != nil {
		cu.SetPublished(*b)
	}
	return cu
}

// SetStartTime sets the "start_time" field.
func (cu *ChallengeUpdate) SetStartTime(t time.Time) *ChallengeUpdate {
	cu.mutation.SetStartTime(t)
	return cu
}

// SetEndTime sets the "end_time" field.
func (cu *ChallengeUpdate) SetEndTime(t time.Time) *ChallengeUpdate {
	cu.mutation.SetEndTime(t)
	return cu
}

// AddPredictionIDs adds the "predictions" edge to the Prediction entity by IDs.
func (cu *ChallengeUpdate) AddPredictionIDs(ids ...uuid.UUID) *ChallengeUpdate {
	cu.mutation.AddPredictionIDs(ids...)
	return cu
}

// AddPredictions adds the "predictions" edges to the Prediction entity.
func (cu *ChallengeUpdate) AddPredictions(p ...*Prediction) *ChallengeUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddPredictionIDs(ids...)
}

// AddProofIDs adds the "proofs" edge to the Proof entity by IDs.
func (cu *ChallengeUpdate) AddProofIDs(ids ...uuid.UUID) *ChallengeUpdate {
	cu.mutation.AddProofIDs(ids...)
	return cu
}

// AddProofs adds the "proofs" edges to the Proof entity.
func (cu *ChallengeUpdate) AddProofs(p ...*Proof) *ChallengeUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProofIDs(ids...)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (cu *ChallengeUpdate) SetAuthorID(id uuid.UUID) *ChallengeUpdate {
	cu.mutation.SetAuthorID(id)
	return cu
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (cu *ChallengeUpdate) SetNillableAuthorID(id *uuid.UUID) *ChallengeUpdate {
	if id != nil {
		cu = cu.SetAuthorID(*id)
	}
	return cu
}

// SetAuthor sets the "author" edge to the User entity.
func (cu *ChallengeUpdate) SetAuthor(u *User) *ChallengeUpdate {
	return cu.SetAuthorID(u.ID)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cu *ChallengeUpdate) Mutation() *ChallengeMutation {
	return cu.mutation
}

// ClearPredictions clears all "predictions" edges to the Prediction entity.
func (cu *ChallengeUpdate) ClearPredictions() *ChallengeUpdate {
	cu.mutation.ClearPredictions()
	return cu
}

// RemovePredictionIDs removes the "predictions" edge to Prediction entities by IDs.
func (cu *ChallengeUpdate) RemovePredictionIDs(ids ...uuid.UUID) *ChallengeUpdate {
	cu.mutation.RemovePredictionIDs(ids...)
	return cu
}

// RemovePredictions removes "predictions" edges to Prediction entities.
func (cu *ChallengeUpdate) RemovePredictions(p ...*Prediction) *ChallengeUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemovePredictionIDs(ids...)
}

// ClearProofs clears all "proofs" edges to the Proof entity.
func (cu *ChallengeUpdate) ClearProofs() *ChallengeUpdate {
	cu.mutation.ClearProofs()
	return cu
}

// RemoveProofIDs removes the "proofs" edge to Proof entities by IDs.
func (cu *ChallengeUpdate) RemoveProofIDs(ids ...uuid.UUID) *ChallengeUpdate {
	cu.mutation.RemoveProofIDs(ids...)
	return cu
}

// RemoveProofs removes "proofs" edges to Proof entities.
func (cu *ChallengeUpdate) RemoveProofs(p ...*Proof) *ChallengeUpdate {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProofIDs(ids...)
}

// ClearAuthor clears the "author" edge to the User entity.
func (cu *ChallengeUpdate) ClearAuthor() *ChallengeUpdate {
	cu.mutation.ClearAuthor()
	return cu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChallengeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChallengeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChallengeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChallengeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChallengeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChallengeUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := challenge.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChallengeUpdate) check() error {
	if v, ok := cu.mutation.Content(); ok {
		if err := challenge.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Challenge.content": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Description(); ok {
		if err := challenge.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Challenge.description": %w`, err)}
		}
	}
	return nil
}

func (cu *ChallengeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: challenge.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: challenge.FieldContent,
		})
	}
	if value, ok := cu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: challenge.FieldDescription,
		})
	}
	if cu.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: challenge.FieldDescription,
		})
	}
	if value, ok := cu.mutation.Outcome(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: challenge.FieldOutcome,
		})
	}
	if cu.mutation.OutcomeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: challenge.FieldOutcome,
		})
	}
	if value, ok := cu.mutation.Published(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: challenge.FieldPublished,
		})
	}
	if value, ok := cu.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldStartTime,
		})
	}
	if value, ok := cu.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldEndTime,
		})
	}
	if cu.mutation.PredictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedPredictionsIDs(); len(nodes) > 0 && !cu.mutation.PredictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.PredictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.ProofsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProofsIDs(); len(nodes) > 0 && !cu.mutation.ProofsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProofsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   challenge.AuthorTable,
			Columns: []string{challenge.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   challenge.AuthorTable,
			Columns: []string{challenge.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challenge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChallengeUpdateOne is the builder for updating a single Challenge entity.
type ChallengeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChallengeMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *ChallengeUpdateOne) SetUpdateTime(t time.Time) *ChallengeUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetContent sets the "content" field.
func (cuo *ChallengeUpdateOne) SetContent(s string) *ChallengeUpdateOne {
	cuo.mutation.SetContent(s)
	return cuo
}

// SetDescription sets the "description" field.
func (cuo *ChallengeUpdateOne) SetDescription(s string) *ChallengeUpdateOne {
	cuo.mutation.SetDescription(s)
	return cuo
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableDescription(s *string) *ChallengeUpdateOne {
	if s != nil {
		cuo.SetDescription(*s)
	}
	return cuo
}

// ClearDescription clears the value of the "description" field.
func (cuo *ChallengeUpdateOne) ClearDescription() *ChallengeUpdateOne {
	cuo.mutation.ClearDescription()
	return cuo
}

// SetOutcome sets the "outcome" field.
func (cuo *ChallengeUpdateOne) SetOutcome(b bool) *ChallengeUpdateOne {
	cuo.mutation.SetOutcome(b)
	return cuo
}

// SetNillableOutcome sets the "outcome" field if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableOutcome(b *bool) *ChallengeUpdateOne {
	if b != nil {
		cuo.SetOutcome(*b)
	}
	return cuo
}

// ClearOutcome clears the value of the "outcome" field.
func (cuo *ChallengeUpdateOne) ClearOutcome() *ChallengeUpdateOne {
	cuo.mutation.ClearOutcome()
	return cuo
}

// SetPublished sets the "published" field.
func (cuo *ChallengeUpdateOne) SetPublished(b bool) *ChallengeUpdateOne {
	cuo.mutation.SetPublished(b)
	return cuo
}

// SetNillablePublished sets the "published" field if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillablePublished(b *bool) *ChallengeUpdateOne {
	if b != nil {
		cuo.SetPublished(*b)
	}
	return cuo
}

// SetStartTime sets the "start_time" field.
func (cuo *ChallengeUpdateOne) SetStartTime(t time.Time) *ChallengeUpdateOne {
	cuo.mutation.SetStartTime(t)
	return cuo
}

// SetEndTime sets the "end_time" field.
func (cuo *ChallengeUpdateOne) SetEndTime(t time.Time) *ChallengeUpdateOne {
	cuo.mutation.SetEndTime(t)
	return cuo
}

// AddPredictionIDs adds the "predictions" edge to the Prediction entity by IDs.
func (cuo *ChallengeUpdateOne) AddPredictionIDs(ids ...uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.AddPredictionIDs(ids...)
	return cuo
}

// AddPredictions adds the "predictions" edges to the Prediction entity.
func (cuo *ChallengeUpdateOne) AddPredictions(p ...*Prediction) *ChallengeUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddPredictionIDs(ids...)
}

// AddProofIDs adds the "proofs" edge to the Proof entity by IDs.
func (cuo *ChallengeUpdateOne) AddProofIDs(ids ...uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.AddProofIDs(ids...)
	return cuo
}

// AddProofs adds the "proofs" edges to the Proof entity.
func (cuo *ChallengeUpdateOne) AddProofs(p ...*Proof) *ChallengeUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProofIDs(ids...)
}

// SetAuthorID sets the "author" edge to the User entity by ID.
func (cuo *ChallengeUpdateOne) SetAuthorID(id uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.SetAuthorID(id)
	return cuo
}

// SetNillableAuthorID sets the "author" edge to the User entity by ID if the given value is not nil.
func (cuo *ChallengeUpdateOne) SetNillableAuthorID(id *uuid.UUID) *ChallengeUpdateOne {
	if id != nil {
		cuo = cuo.SetAuthorID(*id)
	}
	return cuo
}

// SetAuthor sets the "author" edge to the User entity.
func (cuo *ChallengeUpdateOne) SetAuthor(u *User) *ChallengeUpdateOne {
	return cuo.SetAuthorID(u.ID)
}

// Mutation returns the ChallengeMutation object of the builder.
func (cuo *ChallengeUpdateOne) Mutation() *ChallengeMutation {
	return cuo.mutation
}

// ClearPredictions clears all "predictions" edges to the Prediction entity.
func (cuo *ChallengeUpdateOne) ClearPredictions() *ChallengeUpdateOne {
	cuo.mutation.ClearPredictions()
	return cuo
}

// RemovePredictionIDs removes the "predictions" edge to Prediction entities by IDs.
func (cuo *ChallengeUpdateOne) RemovePredictionIDs(ids ...uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.RemovePredictionIDs(ids...)
	return cuo
}

// RemovePredictions removes "predictions" edges to Prediction entities.
func (cuo *ChallengeUpdateOne) RemovePredictions(p ...*Prediction) *ChallengeUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemovePredictionIDs(ids...)
}

// ClearProofs clears all "proofs" edges to the Proof entity.
func (cuo *ChallengeUpdateOne) ClearProofs() *ChallengeUpdateOne {
	cuo.mutation.ClearProofs()
	return cuo
}

// RemoveProofIDs removes the "proofs" edge to Proof entities by IDs.
func (cuo *ChallengeUpdateOne) RemoveProofIDs(ids ...uuid.UUID) *ChallengeUpdateOne {
	cuo.mutation.RemoveProofIDs(ids...)
	return cuo
}

// RemoveProofs removes "proofs" edges to Proof entities.
func (cuo *ChallengeUpdateOne) RemoveProofs(p ...*Proof) *ChallengeUpdateOne {
	ids := make([]uuid.UUID, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProofIDs(ids...)
}

// ClearAuthor clears the "author" edge to the User entity.
func (cuo *ChallengeUpdateOne) ClearAuthor() *ChallengeUpdateOne {
	cuo.mutation.ClearAuthor()
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChallengeUpdateOne) Select(field string, fields ...string) *ChallengeUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Challenge entity.
func (cuo *ChallengeUpdateOne) Save(ctx context.Context) (*Challenge, error) {
	var (
		err  error
		node *Challenge
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChallengeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChallengeUpdateOne) SaveX(ctx context.Context) *Challenge {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChallengeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChallengeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChallengeUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := challenge.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChallengeUpdateOne) check() error {
	if v, ok := cuo.mutation.Content(); ok {
		if err := challenge.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Challenge.content": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Description(); ok {
		if err := challenge.DescriptionValidator(v); err != nil {
			return &ValidationError{Name: "description", err: fmt.Errorf(`ent: validator failed for field "Challenge.description": %w`, err)}
		}
	}
	return nil
}

func (cuo *ChallengeUpdateOne) sqlSave(ctx context.Context) (_node *Challenge, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   challenge.Table,
			Columns: challenge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: challenge.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Challenge.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, challenge.FieldID)
		for _, f := range fields {
			if !challenge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != challenge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.Content(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: challenge.FieldContent,
		})
	}
	if value, ok := cuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: challenge.FieldDescription,
		})
	}
	if cuo.mutation.DescriptionCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: challenge.FieldDescription,
		})
	}
	if value, ok := cuo.mutation.Outcome(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: challenge.FieldOutcome,
		})
	}
	if cuo.mutation.OutcomeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: challenge.FieldOutcome,
		})
	}
	if value, ok := cuo.mutation.Published(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: challenge.FieldPublished,
		})
	}
	if value, ok := cuo.mutation.StartTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldStartTime,
		})
	}
	if value, ok := cuo.mutation.EndTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: challenge.FieldEndTime,
		})
	}
	if cuo.mutation.PredictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedPredictionsIDs(); len(nodes) > 0 && !cuo.mutation.PredictionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.PredictionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.PredictionsTable,
			Columns: []string{challenge.PredictionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: prediction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.ProofsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProofsIDs(); len(nodes) > 0 && !cuo.mutation.ProofsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProofsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   challenge.ProofsTable,
			Columns: []string{challenge.ProofsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: proof.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.AuthorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   challenge.AuthorTable,
			Columns: []string{challenge.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.AuthorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   challenge.AuthorTable,
			Columns: []string{challenge.AuthorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Challenge{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{challenge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
