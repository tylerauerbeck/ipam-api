// Copyright 2023 The Infratographer Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// 	http://www.apache.org/licenses/LICENSE-2.0
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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/predicate"
)

// IPAddressUpdate is the builder for updating IPAddress entities.
type IPAddressUpdate struct {
	config
	hooks    []Hook
	mutation *IPAddressMutation
}

// Where appends a list predicates to the IPAddressUpdate builder.
func (iau *IPAddressUpdate) Where(ps ...predicate.IPAddress) *IPAddressUpdate {
	iau.mutation.Where(ps...)
	return iau
}

// SetIP sets the "IP" field.
func (iau *IPAddressUpdate) SetIP(s string) *IPAddressUpdate {
	iau.mutation.SetIP(s)
	return iau
}

// SetReserved sets the "reserved" field.
func (iau *IPAddressUpdate) SetReserved(b bool) *IPAddressUpdate {
	iau.mutation.SetReserved(b)
	return iau
}

// SetNillableReserved sets the "reserved" field if the given value is not nil.
func (iau *IPAddressUpdate) SetNillableReserved(b *bool) *IPAddressUpdate {
	if b != nil {
		iau.SetReserved(*b)
	}
	return iau
}

// Mutation returns the IPAddressMutation object of the builder.
func (iau *IPAddressUpdate) Mutation() *IPAddressMutation {
	return iau.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iau *IPAddressUpdate) Save(ctx context.Context) (int, error) {
	iau.defaults()
	return withHooks(ctx, iau.sqlSave, iau.mutation, iau.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iau *IPAddressUpdate) SaveX(ctx context.Context) int {
	affected, err := iau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iau *IPAddressUpdate) Exec(ctx context.Context) error {
	_, err := iau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iau *IPAddressUpdate) ExecX(ctx context.Context) {
	if err := iau.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iau *IPAddressUpdate) defaults() {
	if _, ok := iau.mutation.UpdatedAt(); !ok {
		v := ipaddress.UpdateDefaultUpdatedAt()
		iau.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iau *IPAddressUpdate) check() error {
	if v, ok := iau.mutation.IP(); ok {
		if err := ipaddress.IPValidator(v); err != nil {
			return &ValidationError{Name: "IP", err: fmt.Errorf(`generated: validator failed for field "IPAddress.IP": %w`, err)}
		}
	}
	if _, ok := iau.mutation.IPBlockID(); iau.mutation.IPBlockCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "IPAddress.ip_block"`)
	}
	return nil
}

func (iau *IPAddressUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := iau.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(ipaddress.Table, ipaddress.Columns, sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString))
	if ps := iau.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iau.mutation.UpdatedAt(); ok {
		_spec.SetField(ipaddress.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iau.mutation.IP(); ok {
		_spec.SetField(ipaddress.FieldIP, field.TypeString, value)
	}
	if value, ok := iau.mutation.Reserved(); ok {
		_spec.SetField(ipaddress.FieldReserved, field.TypeBool, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	iau.mutation.done = true
	return n, nil
}

// IPAddressUpdateOne is the builder for updating a single IPAddress entity.
type IPAddressUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IPAddressMutation
}

// SetIP sets the "IP" field.
func (iauo *IPAddressUpdateOne) SetIP(s string) *IPAddressUpdateOne {
	iauo.mutation.SetIP(s)
	return iauo
}

// SetReserved sets the "reserved" field.
func (iauo *IPAddressUpdateOne) SetReserved(b bool) *IPAddressUpdateOne {
	iauo.mutation.SetReserved(b)
	return iauo
}

// SetNillableReserved sets the "reserved" field if the given value is not nil.
func (iauo *IPAddressUpdateOne) SetNillableReserved(b *bool) *IPAddressUpdateOne {
	if b != nil {
		iauo.SetReserved(*b)
	}
	return iauo
}

// Mutation returns the IPAddressMutation object of the builder.
func (iauo *IPAddressUpdateOne) Mutation() *IPAddressMutation {
	return iauo.mutation
}

// Where appends a list predicates to the IPAddressUpdate builder.
func (iauo *IPAddressUpdateOne) Where(ps ...predicate.IPAddress) *IPAddressUpdateOne {
	iauo.mutation.Where(ps...)
	return iauo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iauo *IPAddressUpdateOne) Select(field string, fields ...string) *IPAddressUpdateOne {
	iauo.fields = append([]string{field}, fields...)
	return iauo
}

// Save executes the query and returns the updated IPAddress entity.
func (iauo *IPAddressUpdateOne) Save(ctx context.Context) (*IPAddress, error) {
	iauo.defaults()
	return withHooks(ctx, iauo.sqlSave, iauo.mutation, iauo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (iauo *IPAddressUpdateOne) SaveX(ctx context.Context) *IPAddress {
	node, err := iauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iauo *IPAddressUpdateOne) Exec(ctx context.Context) error {
	_, err := iauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iauo *IPAddressUpdateOne) ExecX(ctx context.Context) {
	if err := iauo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iauo *IPAddressUpdateOne) defaults() {
	if _, ok := iauo.mutation.UpdatedAt(); !ok {
		v := ipaddress.UpdateDefaultUpdatedAt()
		iauo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (iauo *IPAddressUpdateOne) check() error {
	if v, ok := iauo.mutation.IP(); ok {
		if err := ipaddress.IPValidator(v); err != nil {
			return &ValidationError{Name: "IP", err: fmt.Errorf(`generated: validator failed for field "IPAddress.IP": %w`, err)}
		}
	}
	if _, ok := iauo.mutation.IPBlockID(); iauo.mutation.IPBlockCleared() && !ok {
		return errors.New(`generated: clearing a required unique edge "IPAddress.ip_block"`)
	}
	return nil
}

func (iauo *IPAddressUpdateOne) sqlSave(ctx context.Context) (_node *IPAddress, err error) {
	if err := iauo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(ipaddress.Table, ipaddress.Columns, sqlgraph.NewFieldSpec(ipaddress.FieldID, field.TypeString))
	id, ok := iauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`generated: missing "IPAddress.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iauo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ipaddress.FieldID)
		for _, f := range fields {
			if !ipaddress.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("generated: invalid field %q for query", f)}
			}
			if f != ipaddress.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iauo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iauo.mutation.UpdatedAt(); ok {
		_spec.SetField(ipaddress.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := iauo.mutation.IP(); ok {
		_spec.SetField(ipaddress.FieldIP, field.TypeString, value)
	}
	if value, ok := iauo.mutation.Reserved(); ok {
		_spec.SetField(ipaddress.FieldReserved, field.TypeBool, value)
	}
	_node = &IPAddress{config: iauo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ipaddress.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	iauo.mutation.done = true
	return _node, nil
}
