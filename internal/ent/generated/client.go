// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by entc, DO NOT EDIT.

package generated

import (
	"context"
	"errors"
	"fmt"
	"log"

	"go.infratographer.com/ipam-api/internal/ent/generated/migrate"
	"go.infratographer.com/x/gidx"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// IPAddress is the client for interacting with the IPAddress builders.
	IPAddress *IPAddressClient
	// IPBlock is the client for interacting with the IPBlock builders.
	IPBlock *IPBlockClient
	// IPBlockType is the client for interacting with the IPBlockType builders.
	IPBlockType *IPBlockTypeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.IPAddress = NewIPAddressClient(c.config)
	c.IPBlock = NewIPBlockClient(c.config)
	c.IPBlockType = NewIPBlockTypeClient(c.config)
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

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("generated: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("generated: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:         ctx,
		config:      cfg,
		IPAddress:   NewIPAddressClient(cfg),
		IPBlock:     NewIPBlockClient(cfg),
		IPBlockType: NewIPBlockTypeClient(cfg),
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
		ctx:         ctx,
		config:      cfg,
		IPAddress:   NewIPAddressClient(cfg),
		IPBlock:     NewIPBlockClient(cfg),
		IPBlockType: NewIPBlockTypeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		IPAddress.
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
	c.IPAddress.Use(hooks...)
	c.IPBlock.Use(hooks...)
	c.IPBlockType.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.IPAddress.Intercept(interceptors...)
	c.IPBlock.Intercept(interceptors...)
	c.IPBlockType.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *IPAddressMutation:
		return c.IPAddress.mutate(ctx, m)
	case *IPBlockMutation:
		return c.IPBlock.mutate(ctx, m)
	case *IPBlockTypeMutation:
		return c.IPBlockType.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("generated: unknown mutation type %T", m)
	}
}

// IPAddressClient is a client for the IPAddress schema.
type IPAddressClient struct {
	config
}

// NewIPAddressClient returns a client for the IPAddress from the given config.
func NewIPAddressClient(c config) *IPAddressClient {
	return &IPAddressClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ipaddress.Hooks(f(g(h())))`.
func (c *IPAddressClient) Use(hooks ...Hook) {
	c.hooks.IPAddress = append(c.hooks.IPAddress, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ipaddress.Intercept(f(g(h())))`.
func (c *IPAddressClient) Intercept(interceptors ...Interceptor) {
	c.inters.IPAddress = append(c.inters.IPAddress, interceptors...)
}

// Create returns a builder for creating a IPAddress entity.
func (c *IPAddressClient) Create() *IPAddressCreate {
	mutation := newIPAddressMutation(c.config, OpCreate)
	return &IPAddressCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of IPAddress entities.
func (c *IPAddressClient) CreateBulk(builders ...*IPAddressCreate) *IPAddressCreateBulk {
	return &IPAddressCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for IPAddress.
func (c *IPAddressClient) Update() *IPAddressUpdate {
	mutation := newIPAddressMutation(c.config, OpUpdate)
	return &IPAddressUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IPAddressClient) UpdateOne(ia *IPAddress) *IPAddressUpdateOne {
	mutation := newIPAddressMutation(c.config, OpUpdateOne, withIPAddress(ia))
	return &IPAddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IPAddressClient) UpdateOneID(id gidx.PrefixedID) *IPAddressUpdateOne {
	mutation := newIPAddressMutation(c.config, OpUpdateOne, withIPAddressID(id))
	return &IPAddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for IPAddress.
func (c *IPAddressClient) Delete() *IPAddressDelete {
	mutation := newIPAddressMutation(c.config, OpDelete)
	return &IPAddressDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IPAddressClient) DeleteOne(ia *IPAddress) *IPAddressDeleteOne {
	return c.DeleteOneID(ia.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IPAddressClient) DeleteOneID(id gidx.PrefixedID) *IPAddressDeleteOne {
	builder := c.Delete().Where(ipaddress.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IPAddressDeleteOne{builder}
}

// Query returns a query builder for IPAddress.
func (c *IPAddressClient) Query() *IPAddressQuery {
	return &IPAddressQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIPAddress},
		inters: c.Interceptors(),
	}
}

// Get returns a IPAddress entity by its id.
func (c *IPAddressClient) Get(ctx context.Context, id gidx.PrefixedID) (*IPAddress, error) {
	return c.Query().Where(ipaddress.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IPAddressClient) GetX(ctx context.Context, id gidx.PrefixedID) *IPAddress {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIPBlock queries the ip_block edge of a IPAddress.
func (c *IPAddressClient) QueryIPBlock(ia *IPAddress) *IPBlockQuery {
	query := (&IPBlockClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ia.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ipaddress.Table, ipaddress.FieldID, id),
			sqlgraph.To(ipblock.Table, ipblock.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ipaddress.IPBlockTable, ipaddress.IPBlockColumn),
		)
		fromV = sqlgraph.Neighbors(ia.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IPAddressClient) Hooks() []Hook {
	return c.hooks.IPAddress
}

// Interceptors returns the client interceptors.
func (c *IPAddressClient) Interceptors() []Interceptor {
	return c.inters.IPAddress
}

func (c *IPAddressClient) mutate(ctx context.Context, m *IPAddressMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IPAddressCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IPAddressUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IPAddressUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IPAddressDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown IPAddress mutation op: %q", m.Op())
	}
}

// IPBlockClient is a client for the IPBlock schema.
type IPBlockClient struct {
	config
}

// NewIPBlockClient returns a client for the IPBlock from the given config.
func NewIPBlockClient(c config) *IPBlockClient {
	return &IPBlockClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ipblock.Hooks(f(g(h())))`.
func (c *IPBlockClient) Use(hooks ...Hook) {
	c.hooks.IPBlock = append(c.hooks.IPBlock, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ipblock.Intercept(f(g(h())))`.
func (c *IPBlockClient) Intercept(interceptors ...Interceptor) {
	c.inters.IPBlock = append(c.inters.IPBlock, interceptors...)
}

// Create returns a builder for creating a IPBlock entity.
func (c *IPBlockClient) Create() *IPBlockCreate {
	mutation := newIPBlockMutation(c.config, OpCreate)
	return &IPBlockCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of IPBlock entities.
func (c *IPBlockClient) CreateBulk(builders ...*IPBlockCreate) *IPBlockCreateBulk {
	return &IPBlockCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for IPBlock.
func (c *IPBlockClient) Update() *IPBlockUpdate {
	mutation := newIPBlockMutation(c.config, OpUpdate)
	return &IPBlockUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IPBlockClient) UpdateOne(ib *IPBlock) *IPBlockUpdateOne {
	mutation := newIPBlockMutation(c.config, OpUpdateOne, withIPBlock(ib))
	return &IPBlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IPBlockClient) UpdateOneID(id gidx.PrefixedID) *IPBlockUpdateOne {
	mutation := newIPBlockMutation(c.config, OpUpdateOne, withIPBlockID(id))
	return &IPBlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for IPBlock.
func (c *IPBlockClient) Delete() *IPBlockDelete {
	mutation := newIPBlockMutation(c.config, OpDelete)
	return &IPBlockDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IPBlockClient) DeleteOne(ib *IPBlock) *IPBlockDeleteOne {
	return c.DeleteOneID(ib.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IPBlockClient) DeleteOneID(id gidx.PrefixedID) *IPBlockDeleteOne {
	builder := c.Delete().Where(ipblock.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IPBlockDeleteOne{builder}
}

// Query returns a query builder for IPBlock.
func (c *IPBlockClient) Query() *IPBlockQuery {
	return &IPBlockQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIPBlock},
		inters: c.Interceptors(),
	}
}

// Get returns a IPBlock entity by its id.
func (c *IPBlockClient) Get(ctx context.Context, id gidx.PrefixedID) (*IPBlock, error) {
	return c.Query().Where(ipblock.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IPBlockClient) GetX(ctx context.Context, id gidx.PrefixedID) *IPBlock {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIPBlockType queries the ip_block_type edge of a IPBlock.
func (c *IPBlockClient) QueryIPBlockType(ib *IPBlock) *IPBlockTypeQuery {
	query := (&IPBlockTypeClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ib.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ipblock.Table, ipblock.FieldID, id),
			sqlgraph.To(ipblocktype.Table, ipblocktype.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ipblock.IPBlockTypeTable, ipblock.IPBlockTypeColumn),
		)
		fromV = sqlgraph.Neighbors(ib.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryIPAddress queries the ip_address edge of a IPBlock.
func (c *IPBlockClient) QueryIPAddress(ib *IPBlock) *IPAddressQuery {
	query := (&IPAddressClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ib.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ipblock.Table, ipblock.FieldID, id),
			sqlgraph.To(ipaddress.Table, ipaddress.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ipblock.IPAddressTable, ipblock.IPAddressColumn),
		)
		fromV = sqlgraph.Neighbors(ib.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IPBlockClient) Hooks() []Hook {
	return c.hooks.IPBlock
}

// Interceptors returns the client interceptors.
func (c *IPBlockClient) Interceptors() []Interceptor {
	return c.inters.IPBlock
}

func (c *IPBlockClient) mutate(ctx context.Context, m *IPBlockMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IPBlockCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IPBlockUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IPBlockUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IPBlockDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown IPBlock mutation op: %q", m.Op())
	}
}

// IPBlockTypeClient is a client for the IPBlockType schema.
type IPBlockTypeClient struct {
	config
}

// NewIPBlockTypeClient returns a client for the IPBlockType from the given config.
func NewIPBlockTypeClient(c config) *IPBlockTypeClient {
	return &IPBlockTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `ipblocktype.Hooks(f(g(h())))`.
func (c *IPBlockTypeClient) Use(hooks ...Hook) {
	c.hooks.IPBlockType = append(c.hooks.IPBlockType, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `ipblocktype.Intercept(f(g(h())))`.
func (c *IPBlockTypeClient) Intercept(interceptors ...Interceptor) {
	c.inters.IPBlockType = append(c.inters.IPBlockType, interceptors...)
}

// Create returns a builder for creating a IPBlockType entity.
func (c *IPBlockTypeClient) Create() *IPBlockTypeCreate {
	mutation := newIPBlockTypeMutation(c.config, OpCreate)
	return &IPBlockTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of IPBlockType entities.
func (c *IPBlockTypeClient) CreateBulk(builders ...*IPBlockTypeCreate) *IPBlockTypeCreateBulk {
	return &IPBlockTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for IPBlockType.
func (c *IPBlockTypeClient) Update() *IPBlockTypeUpdate {
	mutation := newIPBlockTypeMutation(c.config, OpUpdate)
	return &IPBlockTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *IPBlockTypeClient) UpdateOne(ibt *IPBlockType) *IPBlockTypeUpdateOne {
	mutation := newIPBlockTypeMutation(c.config, OpUpdateOne, withIPBlockType(ibt))
	return &IPBlockTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *IPBlockTypeClient) UpdateOneID(id gidx.PrefixedID) *IPBlockTypeUpdateOne {
	mutation := newIPBlockTypeMutation(c.config, OpUpdateOne, withIPBlockTypeID(id))
	return &IPBlockTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for IPBlockType.
func (c *IPBlockTypeClient) Delete() *IPBlockTypeDelete {
	mutation := newIPBlockTypeMutation(c.config, OpDelete)
	return &IPBlockTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *IPBlockTypeClient) DeleteOne(ibt *IPBlockType) *IPBlockTypeDeleteOne {
	return c.DeleteOneID(ibt.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *IPBlockTypeClient) DeleteOneID(id gidx.PrefixedID) *IPBlockTypeDeleteOne {
	builder := c.Delete().Where(ipblocktype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &IPBlockTypeDeleteOne{builder}
}

// Query returns a query builder for IPBlockType.
func (c *IPBlockTypeClient) Query() *IPBlockTypeQuery {
	return &IPBlockTypeQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeIPBlockType},
		inters: c.Interceptors(),
	}
}

// Get returns a IPBlockType entity by its id.
func (c *IPBlockTypeClient) Get(ctx context.Context, id gidx.PrefixedID) (*IPBlockType, error) {
	return c.Query().Where(ipblocktype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *IPBlockTypeClient) GetX(ctx context.Context, id gidx.PrefixedID) *IPBlockType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryIPBlock queries the ip_block edge of a IPBlockType.
func (c *IPBlockTypeClient) QueryIPBlock(ibt *IPBlockType) *IPBlockQuery {
	query := (&IPBlockClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := ibt.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(ipblocktype.Table, ipblocktype.FieldID, id),
			sqlgraph.To(ipblock.Table, ipblock.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ipblocktype.IPBlockTable, ipblocktype.IPBlockColumn),
		)
		fromV = sqlgraph.Neighbors(ibt.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *IPBlockTypeClient) Hooks() []Hook {
	return c.hooks.IPBlockType
}

// Interceptors returns the client interceptors.
func (c *IPBlockTypeClient) Interceptors() []Interceptor {
	return c.inters.IPBlockType
}

func (c *IPBlockTypeClient) mutate(ctx context.Context, m *IPBlockTypeMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&IPBlockTypeCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&IPBlockTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&IPBlockTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&IPBlockTypeDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("generated: unknown IPBlockType mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		IPAddress, IPBlock, IPBlockType []ent.Hook
	}
	inters struct {
		IPAddress, IPBlock, IPBlockType []ent.Interceptor
	}
)
