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
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
	"go.infratographer.com/x/gidx"
)

// IPBlockCreate is the builder for creating a IPBlock entity.
type IPBlockCreate struct {
	config
	mutation *IPBlockMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (ibc *IPBlockCreate) SetCreatedAt(t time.Time) *IPBlockCreate {
	ibc.mutation.SetCreatedAt(t)
	return ibc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ibc *IPBlockCreate) SetNillableCreatedAt(t *time.Time) *IPBlockCreate {
	if t != nil {
		ibc.SetCreatedAt(*t)
	}
	return ibc
}

// SetUpdatedAt sets the "updated_at" field.
func (ibc *IPBlockCreate) SetUpdatedAt(t time.Time) *IPBlockCreate {
	ibc.mutation.SetUpdatedAt(t)
	return ibc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (ibc *IPBlockCreate) SetNillableUpdatedAt(t *time.Time) *IPBlockCreate {
	if t != nil {
		ibc.SetUpdatedAt(*t)
	}
	return ibc
}

// SetPrefix sets the "prefix" field.
func (ibc *IPBlockCreate) SetPrefix(s string) *IPBlockCreate {
	ibc.mutation.SetPrefix(s)
	return ibc
}

// SetBlockTypeID sets the "block_type_id" field.
func (ibc *IPBlockCreate) SetBlockTypeID(gi gidx.PrefixedID) *IPBlockCreate {
	ibc.mutation.SetBlockTypeID(gi)
	return ibc
}

// SetLocationID sets the "location_id" field.
func (ibc *IPBlockCreate) SetLocationID(gi gidx.PrefixedID) *IPBlockCreate {
	ibc.mutation.SetLocationID(gi)
	return ibc
}

// SetParentBlockID sets the "parent_block_id" field.
func (ibc *IPBlockCreate) SetParentBlockID(gi gidx.PrefixedID) *IPBlockCreate {
	ibc.mutation.SetParentBlockID(gi)
	return ibc
}

// SetAllowAutoSubnet sets the "allow_auto_subnet" field.
func (ibc *IPBlockCreate) SetAllowAutoSubnet(b bool) *IPBlockCreate {
	ibc.mutation.SetAllowAutoSubnet(b)
	return ibc
}

// SetNillableAllowAutoSubnet sets the "allow_auto_subnet" field if the given value is not nil.
func (ibc *IPBlockCreate) SetNillableAllowAutoSubnet(b *bool) *IPBlockCreate {
	if b != nil {
		ibc.SetAllowAutoSubnet(*b)
	}
	return ibc
}

// SetAllowAutoAllocate sets the "allow_auto_allocate" field.
func (ibc *IPBlockCreate) SetAllowAutoAllocate(b bool) *IPBlockCreate {
	ibc.mutation.SetAllowAutoAllocate(b)
	return ibc
}

// SetNillableAllowAutoAllocate sets the "allow_auto_allocate" field if the given value is not nil.
func (ibc *IPBlockCreate) SetNillableAllowAutoAllocate(b *bool) *IPBlockCreate {
	if b != nil {
		ibc.SetAllowAutoAllocate(*b)
	}
	return ibc
}

// SetID sets the "id" field.
func (ibc *IPBlockCreate) SetID(gi gidx.PrefixedID) *IPBlockCreate {
	ibc.mutation.SetID(gi)
	return ibc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ibc *IPBlockCreate) SetNillableID(gi *gidx.PrefixedID) *IPBlockCreate {
	if gi != nil {
		ibc.SetID(*gi)
	}
	return ibc
}

// SetIPBlockTypeID sets the "ip_block_type" edge to the IPBlockType entity by ID.
func (ibc *IPBlockCreate) SetIPBlockTypeID(id gidx.PrefixedID) *IPBlockCreate {
	ibc.mutation.SetIPBlockTypeID(id)
	return ibc
}

// SetIPBlockType sets the "ip_block_type" edge to the IPBlockType entity.
func (ibc *IPBlockCreate) SetIPBlockType(i *IPBlockType) *IPBlockCreate {
	return ibc.SetIPBlockTypeID(i.ID)
}

// AddIPAddresIDs adds the "ip_address" edge to the IPAddress entity by IDs.
func (ibc *IPBlockCreate) AddIPAddresIDs(ids ...gidx.PrefixedID) *IPBlockCreate {
	ibc.mutation.AddIPAddresIDs(ids...)
	return ibc
}

// AddIPAddress adds the "ip_address" edges to the IPAddress entity.
func (ibc *IPBlockCreate) AddIPAddress(i ...*IPAddress) *IPBlockCreate {
	ids := make([]gidx.PrefixedID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return ibc.AddIPAddresIDs(ids...)
}

// Mutation returns the IPBlockMutation object of the builder.
func (ibc *IPBlockCreate) Mutation() *IPBlockMutation {
	return ibc.mutation
}

// Save creates the IPBlock in the database.
func (ibc *IPBlockCreate) Save(ctx context.Context) (*IPBlock, error) {
	ibc.defaults()
	return withHooks(ctx, ibc.sqlSave, ibc.mutation, ibc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (ibc *IPBlockCreate) SaveX(ctx context.Context) *IPBlock {
	v, err := ibc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibc *IPBlockCreate) Exec(ctx context.Context) error {
	_, err := ibc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibc *IPBlockCreate) ExecX(ctx context.Context) {
	if err := ibc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ibc *IPBlockCreate) defaults() {
	if _, ok := ibc.mutation.CreatedAt(); !ok {
		v := ipblock.DefaultCreatedAt()
		ibc.mutation.SetCreatedAt(v)
	}
	if _, ok := ibc.mutation.UpdatedAt(); !ok {
		v := ipblock.DefaultUpdatedAt()
		ibc.mutation.SetUpdatedAt(v)
	}
	if _, ok := ibc.mutation.AllowAutoSubnet(); !ok {
		v := ipblock.DefaultAllowAutoSubnet
		ibc.mutation.SetAllowAutoSubnet(v)
	}
	if _, ok := ibc.mutation.AllowAutoAllocate(); !ok {
		v := ipblock.DefaultAllowAutoAllocate
		ibc.mutation.SetAllowAutoAllocate(v)
	}
	if _, ok := ibc.mutation.ID(); !ok {
		v := ipblock.DefaultID()
		ibc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ibc *IPBlockCreate) check() error {
	if _, ok := ibc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`generated: missing required field "IPBlock.created_at"`)}
	}
	if _, ok := ibc.mutation.UpdatedAt(); !ok {
		return &ValidationError{Name: "updated_at", err: errors.New(`generated: missing required field "IPBlock.updated_at"`)}
	}
	if _, ok := ibc.mutation.Prefix(); !ok {
		return &ValidationError{Name: "prefix", err: errors.New(`generated: missing required field "IPBlock.prefix"`)}
	}
	if v, ok := ibc.mutation.Prefix(); ok {
		if err := ipblock.PrefixValidator(v); err != nil {
			return &ValidationError{Name: "prefix", err: fmt.Errorf(`generated: validator failed for field "IPBlock.prefix": %w`, err)}
		}
	}
	if _, ok := ibc.mutation.BlockTypeID(); !ok {
		return &ValidationError{Name: "block_type_id", err: errors.New(`generated: missing required field "IPBlock.block_type_id"`)}
	}
	if _, ok := ibc.mutation.LocationID(); !ok {
		return &ValidationError{Name: "location_id", err: errors.New(`generated: missing required field "IPBlock.location_id"`)}
	}
	if _, ok := ibc.mutation.ParentBlockID(); !ok {
		return &ValidationError{Name: "parent_block_id", err: errors.New(`generated: missing required field "IPBlock.parent_block_id"`)}
	}
	if _, ok := ibc.mutation.AllowAutoSubnet(); !ok {
		return &ValidationError{Name: "allow_auto_subnet", err: errors.New(`generated: missing required field "IPBlock.allow_auto_subnet"`)}
	}
	if _, ok := ibc.mutation.AllowAutoAllocate(); !ok {
		return &ValidationError{Name: "allow_auto_allocate", err: errors.New(`generated: missing required field "IPBlock.allow_auto_allocate"`)}
	}
	if _, ok := ibc.mutation.IPBlockTypeID(); !ok {
		return &ValidationError{Name: "ip_block_type", err: errors.New(`generated: missing required edge "IPBlock.ip_block_type"`)}
	}
	return nil
}

func (ibc *IPBlockCreate) sqlSave(ctx context.Context) (*IPBlock, error) {
	if err := ibc.check(); err != nil {
		return nil, err
	}
	_node, _spec := ibc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ibc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*gidx.PrefixedID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	ibc.mutation.id = &_node.ID
	ibc.mutation.done = true
	return _node, nil
}

func (ibc *IPBlockCreate) createSpec() (*IPBlock, *sqlgraph.CreateSpec) {
	var (
		_node = &IPBlock{config: ibc.config}
		_spec = sqlgraph.NewCreateSpec(ipblock.Table, sqlgraph.NewFieldSpec(ipblock.FieldID, field.TypeString))
	)
	if id, ok := ibc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ibc.mutation.CreatedAt(); ok {
		_spec.SetField(ipblock.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := ibc.mutation.UpdatedAt(); ok {
		_spec.SetField(ipblock.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	if value, ok := ibc.mutation.Prefix(); ok {
		_spec.SetField(ipblock.FieldPrefix, field.TypeString, value)
		_node.Prefix = value
	}
	if value, ok := ibc.mutation.LocationID(); ok {
		_spec.SetField(ipblock.FieldLocationID, field.TypeString, value)
		_node.LocationID = value
	}
	if value, ok := ibc.mutation.ParentBlockID(); ok {
		_spec.SetField(ipblock.FieldParentBlockID, field.TypeString, value)
		_node.ParentBlockID = value
	}
	if value, ok := ibc.mutation.AllowAutoSubnet(); ok {
		_spec.SetField(ipblock.FieldAllowAutoSubnet, field.TypeBool, value)
		_node.AllowAutoSubnet = value
	}
	if value, ok := ibc.mutation.AllowAutoAllocate(); ok {
		_spec.SetField(ipblock.FieldAllowAutoAllocate, field.TypeBool, value)
		_node.AllowAutoAllocate = value
	}
	if nodes := ibc.mutation.IPBlockTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   ipblock.IPBlockTypeTable,
			Columns: []string{ipblock.IPBlockTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipblocktype.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.BlockTypeID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ibc.mutation.IPAddressIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   ipblock.IPAddressTable,
			Columns: []string{ipblock.IPAddressColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// IPBlockCreateBulk is the builder for creating many IPBlock entities in bulk.
type IPBlockCreateBulk struct {
	config
	builders []*IPBlockCreate
}

// Save creates the IPBlock entities in the database.
func (ibcb *IPBlockCreateBulk) Save(ctx context.Context) ([]*IPBlock, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ibcb.builders))
	nodes := make([]*IPBlock, len(ibcb.builders))
	mutators := make([]Mutator, len(ibcb.builders))
	for i := range ibcb.builders {
		func(i int, root context.Context) {
			builder := ibcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*IPBlockMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ibcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ibcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ibcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ibcb *IPBlockCreateBulk) SaveX(ctx context.Context) []*IPBlock {
	v, err := ibcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ibcb *IPBlockCreateBulk) Exec(ctx context.Context) error {
	_, err := ibcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ibcb *IPBlockCreateBulk) ExecX(ctx context.Context) {
	if err := ibcb.Exec(ctx); err != nil {
		panic(err)
	}
}
