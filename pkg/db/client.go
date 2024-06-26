// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"
	"log"
	"reflect"

	"github.com/ronin-chain/ronin-random-beacon/pkg/db/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/chaininfo"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/randomrequest"
	"github.com/ronin-chain/ronin-random-beacon/pkg/db/task"

	stdsql "database/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// ChainInfo is the client for interacting with the ChainInfo builders.
	ChainInfo *ChainInfoClient
	// RandomRequest is the client for interacting with the RandomRequest builders.
	RandomRequest *RandomRequestClient
	// Task is the client for interacting with the Task builders.
	Task *TaskClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	client := &Client{config: newConfig(opts...)}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.ChainInfo = NewChainInfoClient(c.config)
	c.RandomRequest = NewRandomRequestClient(c.config)
	c.Task = NewTaskClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// newConfig creates a new config for the client.
func newConfig(opts ...Option) config {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	return cfg
}

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// ErrTxStarted is returned when trying to start a new transaction from a transactional client.
var ErrTxStarted = errors.New("db: cannot start a transaction within a transaction")

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, ErrTxStarted
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("db: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		ChainInfo:     NewChainInfoClient(cfg),
		RandomRequest: NewRandomRequestClient(cfg),
		Task:          NewTaskClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:           ctx,
		config:        cfg,
		ChainInfo:     NewChainInfoClient(cfg),
		RandomRequest: NewRandomRequestClient(cfg),
		Task:          NewTaskClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		ChainInfo.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.ChainInfo.Use(hooks...)
	c.RandomRequest.Use(hooks...)
	c.Task.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.ChainInfo.Intercept(interceptors...)
	c.RandomRequest.Intercept(interceptors...)
	c.Task.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *ChainInfoMutation:
		return c.ChainInfo.mutate(ctx, m)
	case *RandomRequestMutation:
		return c.RandomRequest.mutate(ctx, m)
	case *TaskMutation:
		return c.Task.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("db: unknown mutation type %T", m)
	}
}

// ChainInfoClient is a client for the ChainInfo schema.
type ChainInfoClient struct {
	config
}

// NewChainInfoClient returns a client for the ChainInfo from the given config.
func NewChainInfoClient(c config) *ChainInfoClient {
	return &ChainInfoClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `chaininfo.Hooks(f(g(h())))`.
func (c *ChainInfoClient) Use(hooks ...Hook) {
	c.hooks.ChainInfo = append(c.hooks.ChainInfo, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `chaininfo.Intercept(f(g(h())))`.
func (c *ChainInfoClient) Intercept(interceptors ...Interceptor) {
	c.inters.ChainInfo = append(c.inters.ChainInfo, interceptors...)
}

// Create returns a builder for creating a ChainInfo entity.
func (c *ChainInfoClient) Create() *ChainInfoCreate {
	mutation := newChainInfoMutation(c.config, OpCreate)
	return &ChainInfoCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of ChainInfo entities.
func (c *ChainInfoClient) CreateBulk(builders ...*ChainInfoCreate) *ChainInfoCreateBulk {
	return &ChainInfoCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *ChainInfoClient) MapCreateBulk(slice any, setFunc func(*ChainInfoCreate, int)) *ChainInfoCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &ChainInfoCreateBulk{err: fmt.Errorf("calling to ChainInfoClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*ChainInfoCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &ChainInfoCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for ChainInfo.
func (c *ChainInfoClient) Update() *ChainInfoUpdate {
	mutation := newChainInfoMutation(c.config, OpUpdate)
	return &ChainInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ChainInfoClient) UpdateOne(ci *ChainInfo) *ChainInfoUpdateOne {
	mutation := newChainInfoMutation(c.config, OpUpdateOne, withChainInfo(ci))
	return &ChainInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ChainInfoClient) UpdateOneID(id string) *ChainInfoUpdateOne {
	mutation := newChainInfoMutation(c.config, OpUpdateOne, withChainInfoID(id))
	return &ChainInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for ChainInfo.
func (c *ChainInfoClient) Delete() *ChainInfoDelete {
	mutation := newChainInfoMutation(c.config, OpDelete)
	return &ChainInfoDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *ChainInfoClient) DeleteOne(ci *ChainInfo) *ChainInfoDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *ChainInfoClient) DeleteOneID(id string) *ChainInfoDeleteOne {
	builder := c.Delete().Where(chaininfo.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ChainInfoDeleteOne{builder}
}

// Query returns a query builder for ChainInfo.
func (c *ChainInfoClient) Query() *ChainInfoQuery {
	return &ChainInfoQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeChainInfo},
		inters: c.Interceptors(),
	}
}

// Get returns a ChainInfo entity by its id.
func (c *ChainInfoClient) Get(ctx context.Context, id string) (*ChainInfo, error) {
	return c.Query().Where(chaininfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ChainInfoClient) GetX(ctx context.Context, id string) *ChainInfo {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ChainInfoClient) Hooks() []Hook {
	return c.hooks.ChainInfo
}

// Interceptors returns the client interceptors.
func (c *ChainInfoClient) Interceptors() []Interceptor {
	return c.inters.ChainInfo
}

func (c *ChainInfoClient) mutate(ctx context.Context, m *ChainInfoMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&ChainInfoCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&ChainInfoUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&ChainInfoUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&ChainInfoDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("db: unknown ChainInfo mutation op: %q", m.Op())
	}
}

// RandomRequestClient is a client for the RandomRequest schema.
type RandomRequestClient struct {
	config
}

// NewRandomRequestClient returns a client for the RandomRequest from the given config.
func NewRandomRequestClient(c config) *RandomRequestClient {
	return &RandomRequestClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `randomrequest.Hooks(f(g(h())))`.
func (c *RandomRequestClient) Use(hooks ...Hook) {
	c.hooks.RandomRequest = append(c.hooks.RandomRequest, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `randomrequest.Intercept(f(g(h())))`.
func (c *RandomRequestClient) Intercept(interceptors ...Interceptor) {
	c.inters.RandomRequest = append(c.inters.RandomRequest, interceptors...)
}

// Create returns a builder for creating a RandomRequest entity.
func (c *RandomRequestClient) Create() *RandomRequestCreate {
	mutation := newRandomRequestMutation(c.config, OpCreate)
	return &RandomRequestCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of RandomRequest entities.
func (c *RandomRequestClient) CreateBulk(builders ...*RandomRequestCreate) *RandomRequestCreateBulk {
	return &RandomRequestCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *RandomRequestClient) MapCreateBulk(slice any, setFunc func(*RandomRequestCreate, int)) *RandomRequestCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &RandomRequestCreateBulk{err: fmt.Errorf("calling to RandomRequestClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*RandomRequestCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &RandomRequestCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for RandomRequest.
func (c *RandomRequestClient) Update() *RandomRequestUpdate {
	mutation := newRandomRequestMutation(c.config, OpUpdate)
	return &RandomRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *RandomRequestClient) UpdateOne(rr *RandomRequest) *RandomRequestUpdateOne {
	mutation := newRandomRequestMutation(c.config, OpUpdateOne, withRandomRequest(rr))
	return &RandomRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *RandomRequestClient) UpdateOneID(id string) *RandomRequestUpdateOne {
	mutation := newRandomRequestMutation(c.config, OpUpdateOne, withRandomRequestID(id))
	return &RandomRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for RandomRequest.
func (c *RandomRequestClient) Delete() *RandomRequestDelete {
	mutation := newRandomRequestMutation(c.config, OpDelete)
	return &RandomRequestDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *RandomRequestClient) DeleteOne(rr *RandomRequest) *RandomRequestDeleteOne {
	return c.DeleteOneID(rr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *RandomRequestClient) DeleteOneID(id string) *RandomRequestDeleteOne {
	builder := c.Delete().Where(randomrequest.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &RandomRequestDeleteOne{builder}
}

// Query returns a query builder for RandomRequest.
func (c *RandomRequestClient) Query() *RandomRequestQuery {
	return &RandomRequestQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeRandomRequest},
		inters: c.Interceptors(),
	}
}

// Get returns a RandomRequest entity by its id.
func (c *RandomRequestClient) Get(ctx context.Context, id string) (*RandomRequest, error) {
	return c.Query().Where(randomrequest.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *RandomRequestClient) GetX(ctx context.Context, id string) *RandomRequest {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTask queries the task edge of a RandomRequest.
func (c *RandomRequestClient) QueryTask(rr *RandomRequest) *TaskQuery {
	query := (&TaskClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := rr.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(randomrequest.Table, randomrequest.FieldID, id),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, randomrequest.TaskTable, randomrequest.TaskColumn),
		)
		fromV = sqlgraph.Neighbors(rr.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *RandomRequestClient) Hooks() []Hook {
	return c.hooks.RandomRequest
}

// Interceptors returns the client interceptors.
func (c *RandomRequestClient) Interceptors() []Interceptor {
	return c.inters.RandomRequest
}

func (c *RandomRequestClient) mutate(ctx context.Context, m *RandomRequestMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&RandomRequestCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&RandomRequestUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&RandomRequestUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&RandomRequestDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("db: unknown RandomRequest mutation op: %q", m.Op())
	}
}

// TaskClient is a client for the Task schema.
type TaskClient struct {
	config
}

// NewTaskClient returns a client for the Task from the given config.
func NewTaskClient(c config) *TaskClient {
	return &TaskClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `task.Hooks(f(g(h())))`.
func (c *TaskClient) Use(hooks ...Hook) {
	c.hooks.Task = append(c.hooks.Task, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `task.Intercept(f(g(h())))`.
func (c *TaskClient) Intercept(interceptors ...Interceptor) {
	c.inters.Task = append(c.inters.Task, interceptors...)
}

// Create returns a builder for creating a Task entity.
func (c *TaskClient) Create() *TaskCreate {
	mutation := newTaskMutation(c.config, OpCreate)
	return &TaskCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Task entities.
func (c *TaskClient) CreateBulk(builders ...*TaskCreate) *TaskCreateBulk {
	return &TaskCreateBulk{config: c.config, builders: builders}
}

// MapCreateBulk creates a bulk creation builder from the given slice. For each item in the slice, the function creates
// a builder and applies setFunc on it.
func (c *TaskClient) MapCreateBulk(slice any, setFunc func(*TaskCreate, int)) *TaskCreateBulk {
	rv := reflect.ValueOf(slice)
	if rv.Kind() != reflect.Slice {
		return &TaskCreateBulk{err: fmt.Errorf("calling to TaskClient.MapCreateBulk with wrong type %T, need slice", slice)}
	}
	builders := make([]*TaskCreate, rv.Len())
	for i := 0; i < rv.Len(); i++ {
		builders[i] = c.Create()
		setFunc(builders[i], i)
	}
	return &TaskCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Task.
func (c *TaskClient) Update() *TaskUpdate {
	mutation := newTaskMutation(c.config, OpUpdate)
	return &TaskUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TaskClient) UpdateOne(t *Task) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTask(t))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TaskClient) UpdateOneID(id int) *TaskUpdateOne {
	mutation := newTaskMutation(c.config, OpUpdateOne, withTaskID(id))
	return &TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Task.
func (c *TaskClient) Delete() *TaskDelete {
	mutation := newTaskMutation(c.config, OpDelete)
	return &TaskDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TaskClient) DeleteOne(t *Task) *TaskDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TaskClient) DeleteOneID(id int) *TaskDeleteOne {
	builder := c.Delete().Where(task.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TaskDeleteOne{builder}
}

// Query returns a query builder for Task.
func (c *TaskClient) Query() *TaskQuery {
	return &TaskQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTask},
		inters: c.Interceptors(),
	}
}

// Get returns a Task entity by its id.
func (c *TaskClient) Get(ctx context.Context, id int) (*Task, error) {
	return c.Query().Where(task.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TaskClient) GetX(ctx context.Context, id int) *Task {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryRandomrequest queries the randomrequest edge of a Task.
func (c *TaskClient) QueryRandomrequest(t *Task) *RandomRequestQuery {
	query := (&RandomRequestClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(task.Table, task.FieldID, id),
			sqlgraph.To(randomrequest.Table, randomrequest.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, task.RandomrequestTable, task.RandomrequestColumn),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TaskClient) Hooks() []Hook {
	return c.hooks.Task
}

// Interceptors returns the client interceptors.
func (c *TaskClient) Interceptors() []Interceptor {
	return c.inters.Task
}

func (c *TaskClient) mutate(ctx context.Context, m *TaskMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TaskCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TaskUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TaskUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TaskDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("db: unknown Task mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		ChainInfo, RandomRequest, Task []ent.Hook
	}
	inters struct {
		ChainInfo, RandomRequest, Task []ent.Interceptor
	}
)

// ExecContext allows calling the underlying ExecContext method of the driver if it is supported by it.
// See, database/sql#DB.ExecContext for more information.
func (c *config) ExecContext(ctx context.Context, query string, args ...any) (stdsql.Result, error) {
	ex, ok := c.driver.(interface {
		ExecContext(context.Context, string, ...any) (stdsql.Result, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.ExecContext is not supported")
	}
	return ex.ExecContext(ctx, query, args...)
}

// QueryContext allows calling the underlying QueryContext method of the driver if it is supported by it.
// See, database/sql#DB.QueryContext for more information.
func (c *config) QueryContext(ctx context.Context, query string, args ...any) (*stdsql.Rows, error) {
	q, ok := c.driver.(interface {
		QueryContext(context.Context, string, ...any) (*stdsql.Rows, error)
	})
	if !ok {
		return nil, fmt.Errorf("Driver.QueryContext is not supported")
	}
	return q.QueryContext(ctx, query, args...)
}
