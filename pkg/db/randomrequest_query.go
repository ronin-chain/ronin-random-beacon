// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/predicate"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/randomrequest"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/task"
)

// RandomRequestQuery is the builder for querying RandomRequest entities.
type RandomRequestQuery struct {
	config
	ctx        *QueryContext
	order      []randomrequest.OrderOption
	inters     []Interceptor
	predicates []predicate.RandomRequest
	withTask   *TaskQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the RandomRequestQuery builder.
func (rrq *RandomRequestQuery) Where(ps ...predicate.RandomRequest) *RandomRequestQuery {
	rrq.predicates = append(rrq.predicates, ps...)
	return rrq
}

// Limit the number of records to be returned by this query.
func (rrq *RandomRequestQuery) Limit(limit int) *RandomRequestQuery {
	rrq.ctx.Limit = &limit
	return rrq
}

// Offset to start from.
func (rrq *RandomRequestQuery) Offset(offset int) *RandomRequestQuery {
	rrq.ctx.Offset = &offset
	return rrq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (rrq *RandomRequestQuery) Unique(unique bool) *RandomRequestQuery {
	rrq.ctx.Unique = &unique
	return rrq
}

// Order specifies how the records should be ordered.
func (rrq *RandomRequestQuery) Order(o ...randomrequest.OrderOption) *RandomRequestQuery {
	rrq.order = append(rrq.order, o...)
	return rrq
}

// QueryTask chains the current query on the "task" edge.
func (rrq *RandomRequestQuery) QueryTask() *TaskQuery {
	query := (&TaskClient{config: rrq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := rrq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := rrq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(randomrequest.Table, randomrequest.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, randomrequest.TaskTable, randomrequest.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(rrq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first RandomRequest entity from the query.
// Returns a *NotFoundError when no RandomRequest was found.
func (rrq *RandomRequestQuery) First(ctx context.Context) (*RandomRequest, error) {
	nodes, err := rrq.Limit(1).All(setContextOp(ctx, rrq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{randomrequest.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (rrq *RandomRequestQuery) FirstX(ctx context.Context) *RandomRequest {
	node, err := rrq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first RandomRequest ID from the query.
// Returns a *NotFoundError when no RandomRequest ID was found.
func (rrq *RandomRequestQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rrq.Limit(1).IDs(setContextOp(ctx, rrq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{randomrequest.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (rrq *RandomRequestQuery) FirstIDX(ctx context.Context) string {
	id, err := rrq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single RandomRequest entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one RandomRequest entity is found.
// Returns a *NotFoundError when no RandomRequest entities are found.
func (rrq *RandomRequestQuery) Only(ctx context.Context) (*RandomRequest, error) {
	nodes, err := rrq.Limit(2).All(setContextOp(ctx, rrq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{randomrequest.Label}
	default:
		return nil, &NotSingularError{randomrequest.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (rrq *RandomRequestQuery) OnlyX(ctx context.Context) *RandomRequest {
	node, err := rrq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only RandomRequest ID in the query.
// Returns a *NotSingularError when more than one RandomRequest ID is found.
// Returns a *NotFoundError when no entities are found.
func (rrq *RandomRequestQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = rrq.Limit(2).IDs(setContextOp(ctx, rrq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{randomrequest.Label}
	default:
		err = &NotSingularError{randomrequest.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (rrq *RandomRequestQuery) OnlyIDX(ctx context.Context) string {
	id, err := rrq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of RandomRequests.
func (rrq *RandomRequestQuery) All(ctx context.Context) ([]*RandomRequest, error) {
	ctx = setContextOp(ctx, rrq.ctx, "All")
	if err := rrq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*RandomRequest, *RandomRequestQuery]()
	return withInterceptors[[]*RandomRequest](ctx, rrq, qr, rrq.inters)
}

// AllX is like All, but panics if an error occurs.
func (rrq *RandomRequestQuery) AllX(ctx context.Context) []*RandomRequest {
	nodes, err := rrq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of RandomRequest IDs.
func (rrq *RandomRequestQuery) IDs(ctx context.Context) (ids []string, err error) {
	if rrq.ctx.Unique == nil && rrq.path != nil {
		rrq.Unique(true)
	}
	ctx = setContextOp(ctx, rrq.ctx, "IDs")
	if err = rrq.Select(randomrequest.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (rrq *RandomRequestQuery) IDsX(ctx context.Context) []string {
	ids, err := rrq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (rrq *RandomRequestQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, rrq.ctx, "Count")
	if err := rrq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, rrq, querierCount[*RandomRequestQuery](), rrq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (rrq *RandomRequestQuery) CountX(ctx context.Context) int {
	count, err := rrq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (rrq *RandomRequestQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, rrq.ctx, "Exist")
	switch _, err := rrq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("db: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (rrq *RandomRequestQuery) ExistX(ctx context.Context) bool {
	exist, err := rrq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the RandomRequestQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (rrq *RandomRequestQuery) Clone() *RandomRequestQuery {
	if rrq == nil {
		return nil
	}
	return &RandomRequestQuery{
		config:     rrq.config,
		ctx:        rrq.ctx.Clone(),
		order:      append([]randomrequest.OrderOption{}, rrq.order...),
		inters:     append([]Interceptor{}, rrq.inters...),
		predicates: append([]predicate.RandomRequest{}, rrq.predicates...),
		withTask:   rrq.withTask.Clone(),
		// clone intermediate query.
		sql:  rrq.sql.Clone(),
		path: rrq.path,
	}
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (rrq *RandomRequestQuery) WithTask(opts ...func(*TaskQuery)) *RandomRequestQuery {
	query := (&TaskClient{config: rrq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	rrq.withTask = query
	return rrq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		BlockNumber uint64 `json:"blockNumber,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.RandomRequest.Query().
//		GroupBy(randomrequest.FieldBlockNumber).
//		Aggregate(db.Count()).
//		Scan(ctx, &v)
func (rrq *RandomRequestQuery) GroupBy(field string, fields ...string) *RandomRequestGroupBy {
	rrq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &RandomRequestGroupBy{build: rrq}
	grbuild.flds = &rrq.ctx.Fields
	grbuild.label = randomrequest.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		BlockNumber uint64 `json:"blockNumber,omitempty"`
//	}
//
//	client.RandomRequest.Query().
//		Select(randomrequest.FieldBlockNumber).
//		Scan(ctx, &v)
func (rrq *RandomRequestQuery) Select(fields ...string) *RandomRequestSelect {
	rrq.ctx.Fields = append(rrq.ctx.Fields, fields...)
	sbuild := &RandomRequestSelect{RandomRequestQuery: rrq}
	sbuild.label = randomrequest.Label
	sbuild.flds, sbuild.scan = &rrq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a RandomRequestSelect configured with the given aggregations.
func (rrq *RandomRequestQuery) Aggregate(fns ...AggregateFunc) *RandomRequestSelect {
	return rrq.Select().Aggregate(fns...)
}

func (rrq *RandomRequestQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range rrq.inters {
		if inter == nil {
			return fmt.Errorf("db: uninitialized interceptor (forgotten import db/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, rrq); err != nil {
				return err
			}
		}
	}
	for _, f := range rrq.ctx.Fields {
		if !randomrequest.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("db: invalid field %q for query", f)}
		}
	}
	if rrq.path != nil {
		prev, err := rrq.path(ctx)
		if err != nil {
			return err
		}
		rrq.sql = prev
	}
	return nil
}

func (rrq *RandomRequestQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*RandomRequest, error) {
	var (
		nodes       = []*RandomRequest{}
		withFKs     = rrq.withFKs
		_spec       = rrq.querySpec()
		loadedTypes = [1]bool{
			rrq.withTask != nil,
		}
	)
	if rrq.withTask != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, randomrequest.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*RandomRequest).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &RandomRequest{config: rrq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, rrq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := rrq.withTask; query != nil {
		if err := rrq.loadTask(ctx, query, nodes, nil,
			func(n *RandomRequest, e *Task) { n.Edges.Task = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (rrq *RandomRequestQuery) loadTask(ctx context.Context, query *TaskQuery, nodes []*RandomRequest, init func(*RandomRequest), assign func(*RandomRequest, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*RandomRequest)
	for i := range nodes {
		if nodes[i].task_randomrequest == nil {
			continue
		}
		fk := *nodes[i].task_randomrequest
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "task_randomrequest" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (rrq *RandomRequestQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := rrq.querySpec()
	_spec.Node.Columns = rrq.ctx.Fields
	if len(rrq.ctx.Fields) > 0 {
		_spec.Unique = rrq.ctx.Unique != nil && *rrq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, rrq.driver, _spec)
}

func (rrq *RandomRequestQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(randomrequest.Table, randomrequest.Columns, sqlgraph.NewFieldSpec(randomrequest.FieldID, field.TypeString))
	_spec.From = rrq.sql
	if unique := rrq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if rrq.path != nil {
		_spec.Unique = true
	}
	if fields := rrq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, randomrequest.FieldID)
		for i := range fields {
			if fields[i] != randomrequest.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := rrq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := rrq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := rrq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := rrq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (rrq *RandomRequestQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(rrq.driver.Dialect())
	t1 := builder.Table(randomrequest.Table)
	columns := rrq.ctx.Fields
	if len(columns) == 0 {
		columns = randomrequest.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if rrq.sql != nil {
		selector = rrq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if rrq.ctx.Unique != nil && *rrq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range rrq.predicates {
		p(selector)
	}
	for _, p := range rrq.order {
		p(selector)
	}
	if offset := rrq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := rrq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// RandomRequestGroupBy is the group-by builder for RandomRequest entities.
type RandomRequestGroupBy struct {
	selector
	build *RandomRequestQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (rrgb *RandomRequestGroupBy) Aggregate(fns ...AggregateFunc) *RandomRequestGroupBy {
	rrgb.fns = append(rrgb.fns, fns...)
	return rrgb
}

// Scan applies the selector query and scans the result into the given value.
func (rrgb *RandomRequestGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rrgb.build.ctx, "GroupBy")
	if err := rrgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RandomRequestQuery, *RandomRequestGroupBy](ctx, rrgb.build, rrgb, rrgb.build.inters, v)
}

func (rrgb *RandomRequestGroupBy) sqlScan(ctx context.Context, root *RandomRequestQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(rrgb.fns))
	for _, fn := range rrgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*rrgb.flds)+len(rrgb.fns))
		for _, f := range *rrgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*rrgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rrgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// RandomRequestSelect is the builder for selecting fields of RandomRequest entities.
type RandomRequestSelect struct {
	*RandomRequestQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (rrs *RandomRequestSelect) Aggregate(fns ...AggregateFunc) *RandomRequestSelect {
	rrs.fns = append(rrs.fns, fns...)
	return rrs
}

// Scan applies the selector query and scans the result into the given value.
func (rrs *RandomRequestSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, rrs.ctx, "Select")
	if err := rrs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*RandomRequestQuery, *RandomRequestSelect](ctx, rrs.RandomRequestQuery, rrs, rrs.inters, v)
}

func (rrs *RandomRequestSelect) sqlScan(ctx context.Context, root *RandomRequestQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(rrs.fns))
	for _, fn := range rrs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*rrs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := rrs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}