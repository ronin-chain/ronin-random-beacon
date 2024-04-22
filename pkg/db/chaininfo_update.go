// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/chaininfo"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/predicate"
)

// ChainInfoUpdate is the builder for updating ChainInfo entities.
type ChainInfoUpdate struct {
	config
	hooks    []Hook
	mutation *ChainInfoMutation
}

// Where appends a list predicates to the ChainInfoUpdate builder.
func (ciu *ChainInfoUpdate) Where(ps ...predicate.ChainInfo) *ChainInfoUpdate {
	ciu.mutation.Where(ps...)
	return ciu
}

// SetProcessedBlock sets the "processed_block" field.
func (ciu *ChainInfoUpdate) SetProcessedBlock(u uint64) *ChainInfoUpdate {
	ciu.mutation.ResetProcessedBlock()
	ciu.mutation.SetProcessedBlock(u)
	return ciu
}

// SetNillableProcessedBlock sets the "processed_block" field if the given value is not nil.
func (ciu *ChainInfoUpdate) SetNillableProcessedBlock(u *uint64) *ChainInfoUpdate {
	if u != nil {
		ciu.SetProcessedBlock(*u)
	}
	return ciu
}

// AddProcessedBlock adds u to the "processed_block" field.
func (ciu *ChainInfoUpdate) AddProcessedBlock(u int64) *ChainInfoUpdate {
	ciu.mutation.AddProcessedBlock(u)
	return ciu
}

// SetFinalizedBlock sets the "finalized_block" field.
func (ciu *ChainInfoUpdate) SetFinalizedBlock(u uint64) *ChainInfoUpdate {
	ciu.mutation.ResetFinalizedBlock()
	ciu.mutation.SetFinalizedBlock(u)
	return ciu
}

// SetNillableFinalizedBlock sets the "finalized_block" field if the given value is not nil.
func (ciu *ChainInfoUpdate) SetNillableFinalizedBlock(u *uint64) *ChainInfoUpdate {
	if u != nil {
		ciu.SetFinalizedBlock(*u)
	}
	return ciu
}

// AddFinalizedBlock adds u to the "finalized_block" field.
func (ciu *ChainInfoUpdate) AddFinalizedBlock(u int64) *ChainInfoUpdate {
	ciu.mutation.AddFinalizedBlock(u)
	return ciu
}

// Mutation returns the ChainInfoMutation object of the builder.
func (ciu *ChainInfoUpdate) Mutation() *ChainInfoMutation {
	return ciu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ciu *ChainInfoUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ciu.sqlSave, ciu.mutation, ciu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciu *ChainInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := ciu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ciu *ChainInfoUpdate) Exec(ctx context.Context) error {
	_, err := ciu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciu *ChainInfoUpdate) ExecX(ctx context.Context) {
	if err := ciu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ciu *ChainInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(chaininfo.Table, chaininfo.Columns, sqlgraph.NewFieldSpec(chaininfo.FieldID, field.TypeString))
	if ps := ciu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciu.mutation.ProcessedBlock(); ok {
		_spec.SetField(chaininfo.FieldProcessedBlock, field.TypeUint64, value)
	}
	if value, ok := ciu.mutation.AddedProcessedBlock(); ok {
		_spec.AddField(chaininfo.FieldProcessedBlock, field.TypeUint64, value)
	}
	if value, ok := ciu.mutation.FinalizedBlock(); ok {
		_spec.SetField(chaininfo.FieldFinalizedBlock, field.TypeUint64, value)
	}
	if value, ok := ciu.mutation.AddedFinalizedBlock(); ok {
		_spec.AddField(chaininfo.FieldFinalizedBlock, field.TypeUint64, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ciu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chaininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ciu.mutation.done = true
	return n, nil
}

// ChainInfoUpdateOne is the builder for updating a single ChainInfo entity.
type ChainInfoUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChainInfoMutation
}

// SetProcessedBlock sets the "processed_block" field.
func (ciuo *ChainInfoUpdateOne) SetProcessedBlock(u uint64) *ChainInfoUpdateOne {
	ciuo.mutation.ResetProcessedBlock()
	ciuo.mutation.SetProcessedBlock(u)
	return ciuo
}

// SetNillableProcessedBlock sets the "processed_block" field if the given value is not nil.
func (ciuo *ChainInfoUpdateOne) SetNillableProcessedBlock(u *uint64) *ChainInfoUpdateOne {
	if u != nil {
		ciuo.SetProcessedBlock(*u)
	}
	return ciuo
}

// AddProcessedBlock adds u to the "processed_block" field.
func (ciuo *ChainInfoUpdateOne) AddProcessedBlock(u int64) *ChainInfoUpdateOne {
	ciuo.mutation.AddProcessedBlock(u)
	return ciuo
}

// SetFinalizedBlock sets the "finalized_block" field.
func (ciuo *ChainInfoUpdateOne) SetFinalizedBlock(u uint64) *ChainInfoUpdateOne {
	ciuo.mutation.ResetFinalizedBlock()
	ciuo.mutation.SetFinalizedBlock(u)
	return ciuo
}

// SetNillableFinalizedBlock sets the "finalized_block" field if the given value is not nil.
func (ciuo *ChainInfoUpdateOne) SetNillableFinalizedBlock(u *uint64) *ChainInfoUpdateOne {
	if u != nil {
		ciuo.SetFinalizedBlock(*u)
	}
	return ciuo
}

// AddFinalizedBlock adds u to the "finalized_block" field.
func (ciuo *ChainInfoUpdateOne) AddFinalizedBlock(u int64) *ChainInfoUpdateOne {
	ciuo.mutation.AddFinalizedBlock(u)
	return ciuo
}

// Mutation returns the ChainInfoMutation object of the builder.
func (ciuo *ChainInfoUpdateOne) Mutation() *ChainInfoMutation {
	return ciuo.mutation
}

// Where appends a list predicates to the ChainInfoUpdate builder.
func (ciuo *ChainInfoUpdateOne) Where(ps ...predicate.ChainInfo) *ChainInfoUpdateOne {
	ciuo.mutation.Where(ps...)
	return ciuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ciuo *ChainInfoUpdateOne) Select(field string, fields ...string) *ChainInfoUpdateOne {
	ciuo.fields = append([]string{field}, fields...)
	return ciuo
}

// Save executes the query and returns the updated ChainInfo entity.
func (ciuo *ChainInfoUpdateOne) Save(ctx context.Context) (*ChainInfo, error) {
	return withHooks(ctx, ciuo.sqlSave, ciuo.mutation, ciuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ciuo *ChainInfoUpdateOne) SaveX(ctx context.Context) *ChainInfo {
	node, err := ciuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ciuo *ChainInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := ciuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ciuo *ChainInfoUpdateOne) ExecX(ctx context.Context) {
	if err := ciuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ciuo *ChainInfoUpdateOne) sqlSave(ctx context.Context) (_node *ChainInfo, err error) {
	_spec := sqlgraph.NewUpdateSpec(chaininfo.Table, chaininfo.Columns, sqlgraph.NewFieldSpec(chaininfo.FieldID, field.TypeString))
	id, ok := ciuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`db: missing "ChainInfo.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ciuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chaininfo.FieldID)
		for _, f := range fields {
			if !chaininfo.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
			}
			if f != chaininfo.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ciuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ciuo.mutation.ProcessedBlock(); ok {
		_spec.SetField(chaininfo.FieldProcessedBlock, field.TypeUint64, value)
	}
	if value, ok := ciuo.mutation.AddedProcessedBlock(); ok {
		_spec.AddField(chaininfo.FieldProcessedBlock, field.TypeUint64, value)
	}
	if value, ok := ciuo.mutation.FinalizedBlock(); ok {
		_spec.SetField(chaininfo.FieldFinalizedBlock, field.TypeUint64, value)
	}
	if value, ok := ciuo.mutation.AddedFinalizedBlock(); ok {
		_spec.AddField(chaininfo.FieldFinalizedBlock, field.TypeUint64, value)
	}
	_node = &ChainInfo{config: ciuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ciuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chaininfo.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ciuo.mutation.done = true
	return _node, nil
}
