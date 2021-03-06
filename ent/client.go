// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/yigitsadic/wotblitz_example/ent/migrate"

	"github.com/yigitsadic/wotblitz_example/ent/module"
	"github.com/yigitsadic/wotblitz_example/ent/tank"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Module is the client for interacting with the Module builders.
	Module *ModuleClient
	// Tank is the client for interacting with the Tank builders.
	Tank *TankClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Module = NewModuleClient(c.config)
	c.Tank = NewTankClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Module: NewModuleClient(cfg),
		Tank:   NewTankClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Module: NewModuleClient(cfg),
		Tank:   NewTankClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Module.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Module.Use(hooks...)
	c.Tank.Use(hooks...)
}

// ModuleClient is a client for the Module schema.
type ModuleClient struct {
	config
}

// NewModuleClient returns a client for the Module from the given config.
func NewModuleClient(c config) *ModuleClient {
	return &ModuleClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `module.Hooks(f(g(h())))`.
func (c *ModuleClient) Use(hooks ...Hook) {
	c.hooks.Module = append(c.hooks.Module, hooks...)
}

// Create returns a create builder for Module.
func (c *ModuleClient) Create() *ModuleCreate {
	mutation := newModuleMutation(c.config, OpCreate)
	return &ModuleCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Module entities.
func (c *ModuleClient) CreateBulk(builders ...*ModuleCreate) *ModuleCreateBulk {
	return &ModuleCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Module.
func (c *ModuleClient) Update() *ModuleUpdate {
	mutation := newModuleMutation(c.config, OpUpdate)
	return &ModuleUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ModuleClient) UpdateOne(m *Module) *ModuleUpdateOne {
	mutation := newModuleMutation(c.config, OpUpdateOne, withModule(m))
	return &ModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ModuleClient) UpdateOneID(id int) *ModuleUpdateOne {
	mutation := newModuleMutation(c.config, OpUpdateOne, withModuleID(id))
	return &ModuleUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Module.
func (c *ModuleClient) Delete() *ModuleDelete {
	mutation := newModuleMutation(c.config, OpDelete)
	return &ModuleDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ModuleClient) DeleteOne(m *Module) *ModuleDeleteOne {
	return c.DeleteOneID(m.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ModuleClient) DeleteOneID(id int) *ModuleDeleteOne {
	builder := c.Delete().Where(module.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ModuleDeleteOne{builder}
}

// Query returns a query builder for Module.
func (c *ModuleClient) Query() *ModuleQuery {
	return &ModuleQuery{config: c.config}
}

// Get returns a Module entity by its id.
func (c *ModuleClient) Get(ctx context.Context, id int) (*Module, error) {
	return c.Query().Where(module.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ModuleClient) GetX(ctx context.Context, id int) *Module {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryTanks queries the tanks edge of a Module.
func (c *ModuleClient) QueryTanks(m *Module) *TankQuery {
	query := &TankQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := m.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(module.Table, module.FieldID, id),
			sqlgraph.To(tank.Table, tank.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, module.TanksTable, module.TanksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(m.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *ModuleClient) Hooks() []Hook {
	return c.hooks.Module
}

// TankClient is a client for the Tank schema.
type TankClient struct {
	config
}

// NewTankClient returns a client for the Tank from the given config.
func NewTankClient(c config) *TankClient {
	return &TankClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `tank.Hooks(f(g(h())))`.
func (c *TankClient) Use(hooks ...Hook) {
	c.hooks.Tank = append(c.hooks.Tank, hooks...)
}

// Create returns a create builder for Tank.
func (c *TankClient) Create() *TankCreate {
	mutation := newTankMutation(c.config, OpCreate)
	return &TankCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Tank entities.
func (c *TankClient) CreateBulk(builders ...*TankCreate) *TankCreateBulk {
	return &TankCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Tank.
func (c *TankClient) Update() *TankUpdate {
	mutation := newTankMutation(c.config, OpUpdate)
	return &TankUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TankClient) UpdateOne(t *Tank) *TankUpdateOne {
	mutation := newTankMutation(c.config, OpUpdateOne, withTank(t))
	return &TankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TankClient) UpdateOneID(id int) *TankUpdateOne {
	mutation := newTankMutation(c.config, OpUpdateOne, withTankID(id))
	return &TankUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Tank.
func (c *TankClient) Delete() *TankDelete {
	mutation := newTankMutation(c.config, OpDelete)
	return &TankDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *TankClient) DeleteOne(t *Tank) *TankDeleteOne {
	return c.DeleteOneID(t.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *TankClient) DeleteOneID(id int) *TankDeleteOne {
	builder := c.Delete().Where(tank.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TankDeleteOne{builder}
}

// Query returns a query builder for Tank.
func (c *TankClient) Query() *TankQuery {
	return &TankQuery{config: c.config}
}

// Get returns a Tank entity by its id.
func (c *TankClient) Get(ctx context.Context, id int) (*Tank, error) {
	return c.Query().Where(tank.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TankClient) GetX(ctx context.Context, id int) *Tank {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryFromTankId queries the fromTankId edge of a Tank.
func (c *TankClient) QueryFromTankId(t *Tank) *TankQuery {
	query := &TankQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tank.Table, tank.FieldID, id),
			sqlgraph.To(tank.Table, tank.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tank.FromTankIdTable, tank.FromTankIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNextTanks queries the nextTanks edge of a Tank.
func (c *TankClient) QueryNextTanks(t *Tank) *TankQuery {
	query := &TankQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tank.Table, tank.FieldID, id),
			sqlgraph.To(tank.Table, tank.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, tank.NextTanksTable, tank.NextTanksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNextTankId queries the nextTankId edge of a Tank.
func (c *TankClient) QueryNextTankId(t *Tank) *TankQuery {
	query := &TankQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tank.Table, tank.FieldID, id),
			sqlgraph.To(tank.Table, tank.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, tank.NextTankIdTable, tank.NextTankIdPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryPreviousTanks queries the previousTanks edge of a Tank.
func (c *TankClient) QueryPreviousTanks(t *Tank) *TankQuery {
	query := &TankQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tank.Table, tank.FieldID, id),
			sqlgraph.To(tank.Table, tank.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, tank.PreviousTanksTable, tank.PreviousTanksPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryModules queries the modules edge of a Tank.
func (c *TankClient) QueryModules(t *Tank) *ModuleQuery {
	query := &ModuleQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := t.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(tank.Table, tank.FieldID, id),
			sqlgraph.To(module.Table, module.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, tank.ModulesTable, tank.ModulesPrimaryKey...),
		)
		fromV = sqlgraph.Neighbors(t.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *TankClient) Hooks() []Hook {
	return c.hooks.Tank
}
