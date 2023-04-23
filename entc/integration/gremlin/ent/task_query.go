// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskQuery is the builder for querying Task entities.
type TaskQuery struct {
	config
	ctx        *QueryContext
	order      []enttask.OrderOption
	inters     []Interceptor
	predicates []predicate.Task
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the TaskQuery builder.
func (tq *TaskQuery) Where(ps ...predicate.Task) *TaskQuery {
	tq.predicates = append(tq.predicates, ps...)
	return tq
}

// Limit the number of records to be returned by this query.
func (tq *TaskQuery) Limit(limit int) *TaskQuery {
	tq.ctx.Limit = &limit
	return tq
}

// Offset to start from.
func (tq *TaskQuery) Offset(offset int) *TaskQuery {
	tq.ctx.Offset = &offset
	return tq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tq *TaskQuery) Unique(unique bool) *TaskQuery {
	tq.ctx.Unique = &unique
	return tq
}

// Order specifies how the records should be ordered.
func (tq *TaskQuery) Order(o ...enttask.OrderOption) *TaskQuery {
	tq.order = append(tq.order, o...)
	return tq
}

// First returns the first Task entity from the query.
// Returns a *NotFoundError when no Task was found.
func (tq *TaskQuery) First(ctx context.Context) (*Task, error) {
	nodes, err := tq.Limit(1).All(setContextOp(ctx, tq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enttask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tq *TaskQuery) FirstX(ctx context.Context) *Task {
	node, err := tq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Task ID from the query.
// Returns a *NotFoundError when no Task ID was found.
func (tq *TaskQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(1).IDs(setContextOp(ctx, tq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enttask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tq *TaskQuery) FirstIDX(ctx context.Context) string {
	id, err := tq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Task entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Task entity is found.
// Returns a *NotFoundError when no Task entities are found.
func (tq *TaskQuery) Only(ctx context.Context) (*Task, error) {
	nodes, err := tq.Limit(2).All(setContextOp(ctx, tq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enttask.Label}
	default:
		return nil, &NotSingularError{enttask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tq *TaskQuery) OnlyX(ctx context.Context) *Task {
	node, err := tq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Task ID in the query.
// Returns a *NotSingularError when more than one Task ID is found.
// Returns a *NotFoundError when no entities are found.
func (tq *TaskQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = tq.Limit(2).IDs(setContextOp(ctx, tq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enttask.Label}
	default:
		err = &NotSingularError{enttask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tq *TaskQuery) OnlyIDX(ctx context.Context) string {
	id, err := tq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tasks.
func (tq *TaskQuery) All(ctx context.Context) ([]*Task, error) {
	ctx = setContextOp(ctx, tq.ctx, "All")
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Task, *TaskQuery]()
	return withInterceptors[[]*Task](ctx, tq, qr, tq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tq *TaskQuery) AllX(ctx context.Context) []*Task {
	nodes, err := tq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Task IDs.
func (tq *TaskQuery) IDs(ctx context.Context) (ids []string, err error) {
	if tq.ctx.Unique == nil && tq.path != nil {
		tq.Unique(true)
	}
	ctx = setContextOp(ctx, tq.ctx, "IDs")
	if err = tq.Select(enttask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tq *TaskQuery) IDsX(ctx context.Context) []string {
	ids, err := tq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tq *TaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tq.ctx, "Count")
	if err := tq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tq, querierCount[*TaskQuery](), tq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tq *TaskQuery) CountX(ctx context.Context) int {
	count, err := tq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tq *TaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tq.ctx, "Exist")
	switch _, err := tq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tq *TaskQuery) ExistX(ctx context.Context) bool {
	exist, err := tq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tq *TaskQuery) Clone() *TaskQuery {
	if tq == nil {
		return nil
	}
	return &TaskQuery{
		config:     tq.config,
		ctx:        tq.ctx.Clone(),
		order:      append([]enttask.OrderOption{}, tq.order...),
		inters:     append([]Interceptor{}, tq.inters...),
		predicates: append([]predicate.Task{}, tq.predicates...),
		// clone intermediate query.
		gremlin: tq.gremlin.Clone(),
		path:    tq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Priority task.Priority `json:"priority,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Task.Query().
//		GroupBy(enttask.FieldPriority).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tq *TaskQuery) GroupBy(field string, fields ...string) *TaskGroupBy {
	tq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TaskGroupBy{build: tq}
	grbuild.flds = &tq.ctx.Fields
	grbuild.label = enttask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Priority task.Priority `json:"priority,omitempty"`
//	}
//
//	client.Task.Query().
//		Select(enttask.FieldPriority).
//		Scan(ctx, &v)
func (tq *TaskQuery) Select(fields ...string) *TaskSelect {
	tq.ctx.Fields = append(tq.ctx.Fields, fields...)
	sbuild := &TaskSelect{TaskQuery: tq}
	sbuild.label = enttask.Label
	sbuild.flds, sbuild.scan = &tq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TaskSelect configured with the given aggregations.
func (tq *TaskQuery) Aggregate(fns ...AggregateFunc) *TaskSelect {
	return tq.Select().Aggregate(fns...)
}

func (tq *TaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tq); err != nil {
				return err
			}
		}
	}
	if tq.path != nil {
		prev, err := tq.path(ctx)
		if err != nil {
			return err
		}
		tq.gremlin = prev
	}
	return nil
}

func (tq *TaskQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Task, error) {
	res := &gremlin.Response{}
	traversal := tq.gremlinQuery(ctx)
	if len(tq.ctx.Fields) > 0 {
		fields := make([]any, len(tq.ctx.Fields))
		for i, f := range tq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := tq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var ts Tasks
	if err := ts.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range ts {
		ts[i].config = tq.config
	}
	return ts, nil
}

func (tq *TaskQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := tq.gremlinQuery(ctx).Count().Query()
	if err := tq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (tq *TaskQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(enttask.Label)
	if tq.gremlin != nil {
		v = tq.gremlin.Clone()
	}
	for _, p := range tq.predicates {
		p(v)
	}
	if len(tq.order) > 0 {
		v.Order()
		for _, p := range tq.order {
			p(v)
		}
	}
	switch limit, offset := tq.ctx.Limit, tq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := tq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// TaskGroupBy is the group-by builder for Task entities.
type TaskGroupBy struct {
	selector
	build *TaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TaskGroupBy) Aggregate(fns ...AggregateFunc) *TaskGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, "GroupBy")
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskQuery, *TaskGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TaskGroupBy) gremlinScan(ctx context.Context, root *TaskQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range tgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *tgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*tgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := tgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*tgb.flds)+len(tgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// TaskSelect is the builder for selecting fields of Task entities.
type TaskSelect struct {
	*TaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TaskSelect) Aggregate(fns ...AggregateFunc) *TaskSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, "Select")
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskQuery, *TaskSelect](ctx, ts.TaskQuery, ts, ts.inters, v)
}

func (ts *TaskSelect) gremlinScan(ctx context.Context, root *TaskQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := ts.ctx.Fields; len(fields) == 1 {
		if fields[0] != enttask.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ts.ctx.Fields))
		for i, f := range ts.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ts.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
