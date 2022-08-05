// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/doc"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/schema/field"
)

// DocQuery is the builder for querying Doc entities.
type DocQuery struct {
	config
	limit        *int
	offset       *int
	unique       *bool
	order        []OrderFunc
	fields       []string
	predicates   []predicate.Doc
	withParent   *DocQuery
	withChildren *DocQuery
	withRelated  *DocQuery
	withFKs      bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the DocQuery builder.
func (dq *DocQuery) Where(ps ...predicate.Doc) *DocQuery {
	dq.predicates = append(dq.predicates, ps...)
	return dq
}

// Limit adds a limit step to the query.
func (dq *DocQuery) Limit(limit int) *DocQuery {
	dq.limit = &limit
	return dq
}

// Offset adds an offset step to the query.
func (dq *DocQuery) Offset(offset int) *DocQuery {
	dq.offset = &offset
	return dq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (dq *DocQuery) Unique(unique bool) *DocQuery {
	dq.unique = &unique
	return dq
}

// Order adds an order step to the query.
func (dq *DocQuery) Order(o ...OrderFunc) *DocQuery {
	dq.order = append(dq.order, o...)
	return dq
}

// QueryParent chains the current query on the "parent" edge.
func (dq *DocQuery) QueryParent() *DocQuery {
	query := &DocQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doc.Table, doc.FieldID, selector),
			sqlgraph.To(doc.Table, doc.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, doc.ParentTable, doc.ParentColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryChildren chains the current query on the "children" edge.
func (dq *DocQuery) QueryChildren() *DocQuery {
	query := &DocQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doc.Table, doc.FieldID, selector),
			sqlgraph.To(doc.Table, doc.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, doc.ChildrenTable, doc.ChildrenColumn),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryRelated chains the current query on the "related" edge.
func (dq *DocQuery) QueryRelated() *DocQuery {
	query := &DocQuery{config: dq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := dq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(doc.Table, doc.FieldID, selector),
			sqlgraph.To(doc.Table, doc.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, doc.RelatedTable, doc.RelatedPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(dq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Doc entity from the query.
// Returns a *NotFoundError when no Doc was found.
func (dq *DocQuery) First(ctx context.Context) (*Doc, error) {
	nodes, err := dq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{doc.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (dq *DocQuery) FirstX(ctx context.Context) *Doc {
	node, err := dq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Doc ID from the query.
// Returns a *NotFoundError when no Doc ID was found.
func (dq *DocQuery) FirstID(ctx context.Context) (id schema.DocID, err error) {
	var ids []schema.DocID
	if ids, err = dq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{doc.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (dq *DocQuery) FirstIDX(ctx context.Context) schema.DocID {
	id, err := dq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Doc entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Doc entity is found.
// Returns a *NotFoundError when no Doc entities are found.
func (dq *DocQuery) Only(ctx context.Context) (*Doc, error) {
	nodes, err := dq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{doc.Label}
	default:
		return nil, &NotSingularError{doc.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (dq *DocQuery) OnlyX(ctx context.Context) *Doc {
	node, err := dq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Doc ID in the query.
// Returns a *NotSingularError when more than one Doc ID is found.
// Returns a *NotFoundError when no entities are found.
func (dq *DocQuery) OnlyID(ctx context.Context) (id schema.DocID, err error) {
	var ids []schema.DocID
	if ids, err = dq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{doc.Label}
	default:
		err = &NotSingularError{doc.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (dq *DocQuery) OnlyIDX(ctx context.Context) schema.DocID {
	id, err := dq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Docs.
func (dq *DocQuery) All(ctx context.Context) ([]*Doc, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return dq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (dq *DocQuery) AllX(ctx context.Context) []*Doc {
	nodes, err := dq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Doc IDs.
func (dq *DocQuery) IDs(ctx context.Context) ([]schema.DocID, error) {
	var ids []schema.DocID
	if err := dq.Select(doc.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (dq *DocQuery) IDsX(ctx context.Context) []schema.DocID {
	ids, err := dq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (dq *DocQuery) Count(ctx context.Context) (int, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return dq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (dq *DocQuery) CountX(ctx context.Context) int {
	count, err := dq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (dq *DocQuery) Exist(ctx context.Context) (bool, error) {
	if err := dq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return dq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (dq *DocQuery) ExistX(ctx context.Context) bool {
	exist, err := dq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the DocQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (dq *DocQuery) Clone() *DocQuery {
	if dq == nil {
		return nil
	}
	return &DocQuery{
		config:       dq.config,
		limit:        dq.limit,
		offset:       dq.offset,
		order:        append([]OrderFunc{}, dq.order...),
		predicates:   append([]predicate.Doc{}, dq.predicates...),
		withParent:   dq.withParent.Clone(),
		withChildren: dq.withChildren.Clone(),
		withRelated:  dq.withRelated.Clone(),
		// clone intermediate query.
		sql:    dq.sql.Clone(),
		path:   dq.path,
		unique: dq.unique,
	}
}

// WithParent tells the query-builder to eager-load the nodes that are connected to
// the "parent" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DocQuery) WithParent(opts ...func(*DocQuery)) *DocQuery {
	query := &DocQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withParent = query
	return dq
}

// WithChildren tells the query-builder to eager-load the nodes that are connected to
// the "children" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DocQuery) WithChildren(opts ...func(*DocQuery)) *DocQuery {
	query := &DocQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withChildren = query
	return dq
}

// WithRelated tells the query-builder to eager-load the nodes that are connected to
// the "related" edge. The optional arguments are used to configure the query builder of the edge.
func (dq *DocQuery) WithRelated(opts ...func(*DocQuery)) *DocQuery {
	query := &DocQuery{config: dq.config}
	for _, opt := range opts {
		opt(query)
	}
	dq.withRelated = query
	return dq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Doc.Query().
//		GroupBy(doc.FieldText).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (dq *DocQuery) GroupBy(field string, fields ...string) *DocGroupBy {
	grbuild := &DocGroupBy{config: dq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := dq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return dq.sqlQuery(ctx), nil
	}
	grbuild.label = doc.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Text string `json:"text,omitempty"`
//	}
//
//	client.Doc.Query().
//		Select(doc.FieldText).
//		Scan(ctx, &v)
func (dq *DocQuery) Select(fields ...string) *DocSelect {
	dq.fields = append(dq.fields, fields...)
	selbuild := &DocSelect{DocQuery: dq}
	selbuild.label = doc.Label
	selbuild.flds, selbuild.scan = &dq.fields, selbuild.Scan
	return selbuild
}

func (dq *DocQuery) prepareQuery(ctx context.Context) error {
	for _, f := range dq.fields {
		if !doc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if dq.path != nil {
		prev, err := dq.path(ctx)
		if err != nil {
			return err
		}
		dq.sql = prev
	}
	return nil
}

func (dq *DocQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Doc, error) {
	var (
		nodes       = []*Doc{}
		withFKs     = dq.withFKs
		_spec       = dq.querySpec()
		loadedTypes = [3]bool{
			dq.withParent != nil,
			dq.withChildren != nil,
			dq.withRelated != nil,
		}
	)
	if dq.withParent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, doc.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*Doc).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &Doc{config: dq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, dq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := dq.withParent; query != nil {
		if err := dq.loadParent(ctx, query, nodes, nil,
			func(n *Doc, e *Doc) { n.Edges.Parent = e }); err != nil {
			return nil, err
		}
	}
	if query := dq.withChildren; query != nil {
		if err := dq.loadChildren(ctx, query, nodes,
			func(n *Doc) { n.Edges.Children = []*Doc{} },
			func(n *Doc, e *Doc) { n.Edges.Children = append(n.Edges.Children, e) }); err != nil {
			return nil, err
		}
	}
	if query := dq.withRelated; query != nil {
		if err := dq.loadRelated(ctx, query, nodes,
			func(n *Doc) { n.Edges.Related = []*Doc{} },
			func(n *Doc, e *Doc) { n.Edges.Related = append(n.Edges.Related, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (dq *DocQuery) loadParent(ctx context.Context, query *DocQuery, nodes []*Doc, init func(*Doc), assign func(*Doc, *Doc)) error {
	ids := make([]schema.DocID, 0, len(nodes))
	nodeids := make(map[schema.DocID][]*Doc)
	for i := range nodes {
		if nodes[i].doc_children == nil {
			continue
		}
		fk := *nodes[i].doc_children
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(doc.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "doc_children" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (dq *DocQuery) loadChildren(ctx context.Context, query *DocQuery, nodes []*Doc, init func(*Doc), assign func(*Doc, *Doc)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[schema.DocID]*Doc)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.Doc(func(s *sql.Selector) {
		s.Where(sql.InValues(doc.ChildrenColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.doc_children
		if fk == nil {
			return fmt.Errorf(`foreign-key "doc_children" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "doc_children" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (dq *DocQuery) loadRelated(ctx context.Context, query *DocQuery, nodes []*Doc, init func(*Doc), assign func(*Doc, *Doc)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[schema.DocID]*Doc)
	nids := make(map[schema.DocID]map[*Doc]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(doc.RelatedTable)
		s.Join(joinT).On(s.C(doc.FieldID), joinT.C(doc.RelatedPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(doc.RelatedPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(doc.RelatedPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	neighbors, err := query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
		assign := spec.Assign
		values := spec.ScanValues
		spec.ScanValues = func(columns []string) ([]interface{}, error) {
			values, err := values(columns[1:])
			if err != nil {
				return nil, err
			}
			return append([]interface{}{new(schema.DocID)}, values...), nil
		}
		spec.Assign = func(columns []string, values []interface{}) error {
			outValue := *values[0].(*schema.DocID)
			inValue := *values[1].(*schema.DocID)
			if nids[inValue] == nil {
				nids[inValue] = map[*Doc]struct{}{byID[outValue]: struct{}{}}
				return assign(columns[1:], values[1:])
			}
			nids[inValue][byID[outValue]] = struct{}{}
			return nil
		}
	})
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "related" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (dq *DocQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := dq.querySpec()
	_spec.Node.Columns = dq.fields
	if len(dq.fields) > 0 {
		_spec.Unique = dq.unique != nil && *dq.unique
	}
	return sqlgraph.CountNodes(ctx, dq.driver, _spec)
}

func (dq *DocQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := dq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("ent: check existence: %w", err)
	}
	return n > 0, nil
}

func (dq *DocQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   doc.Table,
			Columns: doc.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: doc.FieldID,
			},
		},
		From:   dq.sql,
		Unique: true,
	}
	if unique := dq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := dq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, doc.FieldID)
		for i := range fields {
			if fields[i] != doc.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := dq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := dq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := dq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := dq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (dq *DocQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(dq.driver.Dialect())
	t1 := builder.Table(doc.Table)
	columns := dq.fields
	if len(columns) == 0 {
		columns = doc.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if dq.sql != nil {
		selector = dq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if dq.unique != nil && *dq.unique {
		selector.Distinct()
	}
	for _, p := range dq.predicates {
		p(selector)
	}
	for _, p := range dq.order {
		p(selector)
	}
	if offset := dq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := dq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// DocGroupBy is the group-by builder for Doc entities.
type DocGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (dgb *DocGroupBy) Aggregate(fns ...AggregateFunc) *DocGroupBy {
	dgb.fns = append(dgb.fns, fns...)
	return dgb
}

// Scan applies the group-by query and scans the result into the given value.
func (dgb *DocGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := dgb.path(ctx)
	if err != nil {
		return err
	}
	dgb.sql = query
	return dgb.sqlScan(ctx, v)
}

func (dgb *DocGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range dgb.fields {
		if !doc.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := dgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := dgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (dgb *DocGroupBy) sqlQuery() *sql.Selector {
	selector := dgb.sql.Select()
	aggregation := make([]string, 0, len(dgb.fns))
	for _, fn := range dgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(dgb.fields)+len(dgb.fns))
		for _, f := range dgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(dgb.fields...)...)
}

// DocSelect is the builder for selecting fields of Doc entities.
type DocSelect struct {
	*DocQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ds *DocSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ds.prepareQuery(ctx); err != nil {
		return err
	}
	ds.sql = ds.DocQuery.sqlQuery(ctx)
	return ds.sqlScan(ctx, v)
}

func (ds *DocSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ds.sql.Query()
	if err := ds.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
