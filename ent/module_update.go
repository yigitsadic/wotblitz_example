// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/schema/field"
	"github.com/yigitsadic/wotblitz_example/ent/module"
	"github.com/yigitsadic/wotblitz_example/ent/predicate"
	"github.com/yigitsadic/wotblitz_example/ent/tank"
)

// ModuleUpdate is the builder for updating Module entities.
type ModuleUpdate struct {
	config
	hooks    []Hook
	mutation *ModuleMutation
}

// Where adds a new predicate for the builder.
func (mu *ModuleUpdate) Where(ps ...predicate.Module) *ModuleUpdate {
	mu.mutation.predicates = append(mu.mutation.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *ModuleUpdate) SetName(s string) *ModuleUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetModuleType sets the moduleType field.
func (mu *ModuleUpdate) SetModuleType(s string) *ModuleUpdate {
	mu.mutation.SetModuleType(s)
	return mu
}

// SetCreatedAt sets the createdAt field.
func (mu *ModuleUpdate) SetCreatedAt(t time.Time) *ModuleUpdate {
	mu.mutation.SetCreatedAt(t)
	return mu
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (mu *ModuleUpdate) SetNillableCreatedAt(t *time.Time) *ModuleUpdate {
	if t != nil {
		mu.SetCreatedAt(*t)
	}
	return mu
}

// AddTankIDs adds the tanks edge to Tank by ids.
func (mu *ModuleUpdate) AddTankIDs(ids ...int) *ModuleUpdate {
	mu.mutation.AddTankIDs(ids...)
	return mu
}

// AddTanks adds the tanks edges to Tank.
func (mu *ModuleUpdate) AddTanks(t ...*Tank) *ModuleUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.AddTankIDs(ids...)
}

// Mutation returns the ModuleMutation object of the builder.
func (mu *ModuleUpdate) Mutation() *ModuleMutation {
	return mu.mutation
}

// ClearTanks clears all "tanks" edges to type Tank.
func (mu *ModuleUpdate) ClearTanks() *ModuleUpdate {
	mu.mutation.ClearTanks()
	return mu
}

// RemoveTankIDs removes the tanks edge to Tank by ids.
func (mu *ModuleUpdate) RemoveTankIDs(ids ...int) *ModuleUpdate {
	mu.mutation.RemoveTankIDs(ids...)
	return mu
}

// RemoveTanks removes tanks edges to Tank.
func (mu *ModuleUpdate) RemoveTanks(t ...*Tank) *ModuleUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return mu.RemoveTankIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *ModuleUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		if err = mu.check(); err != nil {
			return 0, err
		}
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = mu.check(); err != nil {
				return 0, err
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *ModuleUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *ModuleUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *ModuleUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mu *ModuleUpdate) check() error {
	if v, ok := mu.mutation.Name(); ok {
		if err := module.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := mu.mutation.ModuleType(); ok {
		if err := module.ModuleTypeValidator(v); err != nil {
			return &ValidationError{Name: "moduleType", err: fmt.Errorf("ent: validator failed for field \"moduleType\": %w", err)}
		}
	}
	return nil
}

func (mu *ModuleUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   module.Table,
			Columns: module.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: module.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: module.FieldName,
		})
	}
	if value, ok := mu.mutation.ModuleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: module.FieldModuleType,
		})
	}
	if value, ok := mu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: module.FieldCreatedAt,
		})
	}
	if mu.mutation.TanksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   module.TanksTable,
			Columns: module.TanksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tank.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.RemovedTanksIDs(); len(nodes) > 0 && !mu.mutation.TanksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   module.TanksTable,
			Columns: module.TanksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.TanksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   module.TanksTable,
			Columns: module.TanksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{module.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// ModuleUpdateOne is the builder for updating a single Module entity.
type ModuleUpdateOne struct {
	config
	hooks    []Hook
	mutation *ModuleMutation
}

// SetName sets the name field.
func (muo *ModuleUpdateOne) SetName(s string) *ModuleUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetModuleType sets the moduleType field.
func (muo *ModuleUpdateOne) SetModuleType(s string) *ModuleUpdateOne {
	muo.mutation.SetModuleType(s)
	return muo
}

// SetCreatedAt sets the createdAt field.
func (muo *ModuleUpdateOne) SetCreatedAt(t time.Time) *ModuleUpdateOne {
	muo.mutation.SetCreatedAt(t)
	return muo
}

// SetNillableCreatedAt sets the createdAt field if the given value is not nil.
func (muo *ModuleUpdateOne) SetNillableCreatedAt(t *time.Time) *ModuleUpdateOne {
	if t != nil {
		muo.SetCreatedAt(*t)
	}
	return muo
}

// AddTankIDs adds the tanks edge to Tank by ids.
func (muo *ModuleUpdateOne) AddTankIDs(ids ...int) *ModuleUpdateOne {
	muo.mutation.AddTankIDs(ids...)
	return muo
}

// AddTanks adds the tanks edges to Tank.
func (muo *ModuleUpdateOne) AddTanks(t ...*Tank) *ModuleUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.AddTankIDs(ids...)
}

// Mutation returns the ModuleMutation object of the builder.
func (muo *ModuleUpdateOne) Mutation() *ModuleMutation {
	return muo.mutation
}

// ClearTanks clears all "tanks" edges to type Tank.
func (muo *ModuleUpdateOne) ClearTanks() *ModuleUpdateOne {
	muo.mutation.ClearTanks()
	return muo
}

// RemoveTankIDs removes the tanks edge to Tank by ids.
func (muo *ModuleUpdateOne) RemoveTankIDs(ids ...int) *ModuleUpdateOne {
	muo.mutation.RemoveTankIDs(ids...)
	return muo
}

// RemoveTanks removes tanks edges to Tank.
func (muo *ModuleUpdateOne) RemoveTanks(t ...*Tank) *ModuleUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return muo.RemoveTankIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *ModuleUpdateOne) Save(ctx context.Context) (*Module, error) {
	var (
		err  error
		node *Module
	)
	if len(muo.hooks) == 0 {
		if err = muo.check(); err != nil {
			return nil, err
		}
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ModuleMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = muo.check(); err != nil {
				return nil, err
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *ModuleUpdateOne) SaveX(ctx context.Context) *Module {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *ModuleUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *ModuleUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (muo *ModuleUpdateOne) check() error {
	if v, ok := muo.mutation.Name(); ok {
		if err := module.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := muo.mutation.ModuleType(); ok {
		if err := module.ModuleTypeValidator(v); err != nil {
			return &ValidationError{Name: "moduleType", err: fmt.Errorf("ent: validator failed for field \"moduleType\": %w", err)}
		}
	}
	return nil
}

func (muo *ModuleUpdateOne) sqlSave(ctx context.Context) (_node *Module, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   module.Table,
			Columns: module.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: module.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Module.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: module.FieldName,
		})
	}
	if value, ok := muo.mutation.ModuleType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: module.FieldModuleType,
		})
	}
	if value, ok := muo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: module.FieldCreatedAt,
		})
	}
	if muo.mutation.TanksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   module.TanksTable,
			Columns: module.TanksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tank.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.RemovedTanksIDs(); len(nodes) > 0 && !muo.mutation.TanksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   module.TanksTable,
			Columns: module.TanksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.TanksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   module.TanksTable,
			Columns: module.TanksPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tank.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Module{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{module.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
