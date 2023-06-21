// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/schema/task"
	enttask "entgo.io/ent/entc/integration/ent/task"
	"entgo.io/ent/schema/field"
)

// TaskUpdate is the builder for updating Task entities.
type TaskUpdate struct {
	config
	hooks     []Hook
	mutation  *TaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TaskUpdate builder.
func (tu *TaskUpdate) Where(ps ...predicate.Task) *TaskUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetPriority sets the "priority" field.
func (tu *TaskUpdate) SetPriority(t task.Priority) *TaskUpdate {
	tu.mutation.ResetPriority()
	tu.mutation.SetPriority(t)
	return tu
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tu *TaskUpdate) SetNillablePriority(t *task.Priority) *TaskUpdate {
	if t != nil {
		tu.SetPriority(*t)
	}
	return tu
}

// AddPriority adds t to the "priority" field.
func (tu *TaskUpdate) AddPriority(t task.Priority) *TaskUpdate {
	tu.mutation.AddPriority(t)
	return tu
}

// SetPriorities sets the "priorities" field.
func (tu *TaskUpdate) SetPriorities(m map[string]task.Priority) *TaskUpdate {
	tu.mutation.SetPriorities(m)
	return tu
}

// ClearPriorities clears the value of the "priorities" field.
func (tu *TaskUpdate) ClearPriorities() *TaskUpdate {
	tu.mutation.ClearPriorities()
	return tu
}

// SetName sets the "name" field.
func (tu *TaskUpdate) SetName(s string) *TaskUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableName(s *string) *TaskUpdate {
	if s != nil {
		tu.SetName(*s)
	}
	return tu
}

// ClearName clears the value of the "name" field.
func (tu *TaskUpdate) ClearName() *TaskUpdate {
	tu.mutation.ClearName()
	return tu
}

// SetOwner sets the "owner" field.
func (tu *TaskUpdate) SetOwner(s string) *TaskUpdate {
	tu.mutation.SetOwner(s)
	return tu
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOwner(s *string) *TaskUpdate {
	if s != nil {
		tu.SetOwner(*s)
	}
	return tu
}

// ClearOwner clears the value of the "owner" field.
func (tu *TaskUpdate) ClearOwner() *TaskUpdate {
	tu.mutation.ClearOwner()
	return tu
}

// SetOrder sets the "order" field.
func (tu *TaskUpdate) SetOrder(i int) *TaskUpdate {
	tu.mutation.ResetOrder()
	tu.mutation.SetOrder(i)
	return tu
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOrder(i *int) *TaskUpdate {
	if i != nil {
		tu.SetOrder(*i)
	}
	return tu
}

// AddOrder adds i to the "order" field.
func (tu *TaskUpdate) AddOrder(i int) *TaskUpdate {
	tu.mutation.AddOrder(i)
	return tu
}

// ClearOrder clears the value of the "order" field.
func (tu *TaskUpdate) ClearOrder() *TaskUpdate {
	tu.mutation.ClearOrder()
	return tu
}

// SetOrderOption sets the "order_option" field.
func (tu *TaskUpdate) SetOrderOption(i int) *TaskUpdate {
	tu.mutation.ResetOrderOption()
	tu.mutation.SetOrderOption(i)
	return tu
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOrderOption(i *int) *TaskUpdate {
	if i != nil {
		tu.SetOrderOption(*i)
	}
	return tu
}

// AddOrderOption adds i to the "order_option" field.
func (tu *TaskUpdate) AddOrderOption(i int) *TaskUpdate {
	tu.mutation.AddOrderOption(i)
	return tu
}

// ClearOrderOption clears the value of the "order_option" field.
func (tu *TaskUpdate) ClearOrderOption() *TaskUpdate {
	tu.mutation.ClearOrderOption()
	return tu
}

// SetOp sets the "op" field.
func (tu *TaskUpdate) SetOp(s string) *TaskUpdate {
	tu.mutation.SetOpField(s)
	return tu
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (tu *TaskUpdate) SetNillableOp(s *string) *TaskUpdate {
	if s != nil {
		tu.SetOp(*s)
	}
	return tu
}

// Mutation returns the TaskMutation object of the builder.
func (tu *TaskUpdate) Mutation() *TaskMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TaskUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TaskUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TaskUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TaskUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tu *TaskUpdate) check() error {
	if v, ok := tu.mutation.Priority(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Task.priority": %w`, err)}
		}
	}
	if v, ok := tu.mutation.GetOp(); ok {
		if err := enttask.OpValidator(v); err != nil {
			return &ValidationError{Name: "op", err: fmt.Errorf(`ent: validator failed for field "Task.op": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TaskUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TaskUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(enttask.Table, enttask.Columns, sqlgraph.NewFieldSpec(enttask.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Priority(); ok {
		_spec.SetField(enttask.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedPriority(); ok {
		_spec.AddField(enttask.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tu.mutation.Priorities(); ok {
		_spec.SetField(enttask.FieldPriorities, field.TypeJSON, value)
	}
	if tu.mutation.PrioritiesCleared() {
		_spec.ClearField(enttask.FieldPriorities, field.TypeJSON)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(enttask.FieldName, field.TypeString, value)
	}
	if tu.mutation.NameCleared() {
		_spec.ClearField(enttask.FieldName, field.TypeString)
	}
	if value, ok := tu.mutation.Owner(); ok {
		_spec.SetField(enttask.FieldOwner, field.TypeString, value)
	}
	if tu.mutation.OwnerCleared() {
		_spec.ClearField(enttask.FieldOwner, field.TypeString)
	}
	if value, ok := tu.mutation.Order(); ok {
		_spec.SetField(enttask.FieldOrder, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedOrder(); ok {
		_spec.AddField(enttask.FieldOrder, field.TypeInt, value)
	}
	if tu.mutation.OrderCleared() {
		_spec.ClearField(enttask.FieldOrder, field.TypeInt)
	}
	if value, ok := tu.mutation.OrderOption(); ok {
		_spec.SetField(enttask.FieldOrderOption, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedOrderOption(); ok {
		_spec.AddField(enttask.FieldOrderOption, field.TypeInt, value)
	}
	if tu.mutation.OrderOptionCleared() {
		_spec.ClearField(enttask.FieldOrderOption, field.TypeInt)
	}
	if value, ok := tu.mutation.GetOp(); ok {
		_spec.SetField(enttask.FieldOp, field.TypeString, value)
	}
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enttask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TaskUpdateOne is the builder for updating a single Task entity.
type TaskUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TaskMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetPriority sets the "priority" field.
func (tuo *TaskUpdateOne) SetPriority(t task.Priority) *TaskUpdateOne {
	tuo.mutation.ResetPriority()
	tuo.mutation.SetPriority(t)
	return tuo
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillablePriority(t *task.Priority) *TaskUpdateOne {
	if t != nil {
		tuo.SetPriority(*t)
	}
	return tuo
}

// AddPriority adds t to the "priority" field.
func (tuo *TaskUpdateOne) AddPriority(t task.Priority) *TaskUpdateOne {
	tuo.mutation.AddPriority(t)
	return tuo
}

// SetPriorities sets the "priorities" field.
func (tuo *TaskUpdateOne) SetPriorities(m map[string]task.Priority) *TaskUpdateOne {
	tuo.mutation.SetPriorities(m)
	return tuo
}

// ClearPriorities clears the value of the "priorities" field.
func (tuo *TaskUpdateOne) ClearPriorities() *TaskUpdateOne {
	tuo.mutation.ClearPriorities()
	return tuo
}

// SetName sets the "name" field.
func (tuo *TaskUpdateOne) SetName(s string) *TaskUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableName(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetName(*s)
	}
	return tuo
}

// ClearName clears the value of the "name" field.
func (tuo *TaskUpdateOne) ClearName() *TaskUpdateOne {
	tuo.mutation.ClearName()
	return tuo
}

// SetOwner sets the "owner" field.
func (tuo *TaskUpdateOne) SetOwner(s string) *TaskUpdateOne {
	tuo.mutation.SetOwner(s)
	return tuo
}

// SetNillableOwner sets the "owner" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOwner(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetOwner(*s)
	}
	return tuo
}

// ClearOwner clears the value of the "owner" field.
func (tuo *TaskUpdateOne) ClearOwner() *TaskUpdateOne {
	tuo.mutation.ClearOwner()
	return tuo
}

// SetOrder sets the "order" field.
func (tuo *TaskUpdateOne) SetOrder(i int) *TaskUpdateOne {
	tuo.mutation.ResetOrder()
	tuo.mutation.SetOrder(i)
	return tuo
}

// SetNillableOrder sets the "order" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOrder(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetOrder(*i)
	}
	return tuo
}

// AddOrder adds i to the "order" field.
func (tuo *TaskUpdateOne) AddOrder(i int) *TaskUpdateOne {
	tuo.mutation.AddOrder(i)
	return tuo
}

// ClearOrder clears the value of the "order" field.
func (tuo *TaskUpdateOne) ClearOrder() *TaskUpdateOne {
	tuo.mutation.ClearOrder()
	return tuo
}

// SetOrderOption sets the "order_option" field.
func (tuo *TaskUpdateOne) SetOrderOption(i int) *TaskUpdateOne {
	tuo.mutation.ResetOrderOption()
	tuo.mutation.SetOrderOption(i)
	return tuo
}

// SetNillableOrderOption sets the "order_option" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOrderOption(i *int) *TaskUpdateOne {
	if i != nil {
		tuo.SetOrderOption(*i)
	}
	return tuo
}

// AddOrderOption adds i to the "order_option" field.
func (tuo *TaskUpdateOne) AddOrderOption(i int) *TaskUpdateOne {
	tuo.mutation.AddOrderOption(i)
	return tuo
}

// ClearOrderOption clears the value of the "order_option" field.
func (tuo *TaskUpdateOne) ClearOrderOption() *TaskUpdateOne {
	tuo.mutation.ClearOrderOption()
	return tuo
}

// SetOp sets the "op" field.
func (tuo *TaskUpdateOne) SetOp(s string) *TaskUpdateOne {
	tuo.mutation.SetOpField(s)
	return tuo
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (tuo *TaskUpdateOne) SetNillableOp(s *string) *TaskUpdateOne {
	if s != nil {
		tuo.SetOp(*s)
	}
	return tuo
}

// Mutation returns the TaskMutation object of the builder.
func (tuo *TaskUpdateOne) Mutation() *TaskMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TaskUpdate builder.
func (tuo *TaskUpdateOne) Where(ps ...predicate.Task) *TaskUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TaskUpdateOne) Select(field string, fields ...string) *TaskUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Task entity.
func (tuo *TaskUpdateOne) Save(ctx context.Context) (*Task, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TaskUpdateOne) SaveX(ctx context.Context) *Task {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TaskUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TaskUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TaskUpdateOne) check() error {
	if v, ok := tuo.mutation.Priority(); ok {
		if err := v.Validate(); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Task.priority": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.GetOp(); ok {
		if err := enttask.OpValidator(v); err != nil {
			return &ValidationError{Name: "op", err: fmt.Errorf(`ent: validator failed for field "Task.op": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TaskUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TaskUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TaskUpdateOne) sqlSave(ctx context.Context) (_node *Task, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(enttask.Table, enttask.Columns, sqlgraph.NewFieldSpec(enttask.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Task.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, enttask.FieldID)
		for _, f := range fields {
			if !enttask.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != enttask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Priority(); ok {
		_spec.SetField(enttask.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedPriority(); ok {
		_spec.AddField(enttask.FieldPriority, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.Priorities(); ok {
		_spec.SetField(enttask.FieldPriorities, field.TypeJSON, value)
	}
	if tuo.mutation.PrioritiesCleared() {
		_spec.ClearField(enttask.FieldPriorities, field.TypeJSON)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(enttask.FieldName, field.TypeString, value)
	}
	if tuo.mutation.NameCleared() {
		_spec.ClearField(enttask.FieldName, field.TypeString)
	}
	if value, ok := tuo.mutation.Owner(); ok {
		_spec.SetField(enttask.FieldOwner, field.TypeString, value)
	}
	if tuo.mutation.OwnerCleared() {
		_spec.ClearField(enttask.FieldOwner, field.TypeString)
	}
	if value, ok := tuo.mutation.Order(); ok {
		_spec.SetField(enttask.FieldOrder, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedOrder(); ok {
		_spec.AddField(enttask.FieldOrder, field.TypeInt, value)
	}
	if tuo.mutation.OrderCleared() {
		_spec.ClearField(enttask.FieldOrder, field.TypeInt)
	}
	if value, ok := tuo.mutation.OrderOption(); ok {
		_spec.SetField(enttask.FieldOrderOption, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedOrderOption(); ok {
		_spec.AddField(enttask.FieldOrderOption, field.TypeInt, value)
	}
	if tuo.mutation.OrderOptionCleared() {
		_spec.ClearField(enttask.FieldOrderOption, field.TypeInt)
	}
	if value, ok := tuo.mutation.GetOp(); ok {
		_spec.SetField(enttask.FieldOp, field.TypeString, value)
	}
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Task{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{enttask.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
