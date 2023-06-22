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
	"time"

	"go.infratographer.com/ipam-api/internal/ent/generated/ipaddress"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblock"
	"go.infratographer.com/ipam-api/internal/ent/generated/ipblocktype"
	"go.infratographer.com/ipam-api/internal/ent/schema"
	"go.infratographer.com/x/gidx"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	ipaddressMixin := schema.IPAddress{}.Mixin()
	ipaddressMixinFields0 := ipaddressMixin[0].Fields()
	_ = ipaddressMixinFields0
	ipaddressFields := schema.IPAddress{}.Fields()
	_ = ipaddressFields
	// ipaddressDescCreatedAt is the schema descriptor for created_at field.
	ipaddressDescCreatedAt := ipaddressMixinFields0[0].Descriptor()
	// ipaddress.DefaultCreatedAt holds the default value on creation for the created_at field.
	ipaddress.DefaultCreatedAt = ipaddressDescCreatedAt.Default.(func() time.Time)
	// ipaddressDescUpdatedAt is the schema descriptor for updated_at field.
	ipaddressDescUpdatedAt := ipaddressMixinFields0[1].Descriptor()
	// ipaddress.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ipaddress.DefaultUpdatedAt = ipaddressDescUpdatedAt.Default.(func() time.Time)
	// ipaddress.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ipaddress.UpdateDefaultUpdatedAt = ipaddressDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ipaddressDescIP is the schema descriptor for IP field.
	ipaddressDescIP := ipaddressFields[1].Descriptor()
	// ipaddress.IPValidator is a validator for the "IP" field. It is called by the builders before save.
	ipaddress.IPValidator = ipaddressDescIP.Validators[0].(func(string) error)
	// ipaddressDescReserved is the schema descriptor for reserved field.
	ipaddressDescReserved := ipaddressFields[5].Descriptor()
	// ipaddress.DefaultReserved holds the default value on creation for the reserved field.
	ipaddress.DefaultReserved = ipaddressDescReserved.Default.(bool)
	// ipaddressDescID is the schema descriptor for id field.
	ipaddressDescID := ipaddressFields[0].Descriptor()
	// ipaddress.DefaultID holds the default value on creation for the id field.
	ipaddress.DefaultID = ipaddressDescID.Default.(func() gidx.PrefixedID)
	ipblockMixin := schema.IPBlock{}.Mixin()
	ipblockMixinFields0 := ipblockMixin[0].Fields()
	_ = ipblockMixinFields0
	ipblockFields := schema.IPBlock{}.Fields()
	_ = ipblockFields
	// ipblockDescCreatedAt is the schema descriptor for created_at field.
	ipblockDescCreatedAt := ipblockMixinFields0[0].Descriptor()
	// ipblock.DefaultCreatedAt holds the default value on creation for the created_at field.
	ipblock.DefaultCreatedAt = ipblockDescCreatedAt.Default.(func() time.Time)
	// ipblockDescUpdatedAt is the schema descriptor for updated_at field.
	ipblockDescUpdatedAt := ipblockMixinFields0[1].Descriptor()
	// ipblock.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ipblock.DefaultUpdatedAt = ipblockDescUpdatedAt.Default.(func() time.Time)
	// ipblock.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ipblock.UpdateDefaultUpdatedAt = ipblockDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ipblockDescPrefix is the schema descriptor for prefix field.
	ipblockDescPrefix := ipblockFields[1].Descriptor()
	// ipblock.PrefixValidator is a validator for the "prefix" field. It is called by the builders before save.
	ipblock.PrefixValidator = ipblockDescPrefix.Validators[0].(func(string) error)
	// ipblockDescAllowAutoSubnet is the schema descriptor for allow_auto_subnet field.
	ipblockDescAllowAutoSubnet := ipblockFields[5].Descriptor()
	// ipblock.DefaultAllowAutoSubnet holds the default value on creation for the allow_auto_subnet field.
	ipblock.DefaultAllowAutoSubnet = ipblockDescAllowAutoSubnet.Default.(bool)
	// ipblockDescAllowAutoAllocate is the schema descriptor for allow_auto_allocate field.
	ipblockDescAllowAutoAllocate := ipblockFields[6].Descriptor()
	// ipblock.DefaultAllowAutoAllocate holds the default value on creation for the allow_auto_allocate field.
	ipblock.DefaultAllowAutoAllocate = ipblockDescAllowAutoAllocate.Default.(bool)
	// ipblockDescID is the schema descriptor for id field.
	ipblockDescID := ipblockFields[0].Descriptor()
	// ipblock.DefaultID holds the default value on creation for the id field.
	ipblock.DefaultID = ipblockDescID.Default.(func() gidx.PrefixedID)
	ipblocktypeMixin := schema.IPBlockType{}.Mixin()
	ipblocktypeMixinFields0 := ipblocktypeMixin[0].Fields()
	_ = ipblocktypeMixinFields0
	ipblocktypeFields := schema.IPBlockType{}.Fields()
	_ = ipblocktypeFields
	// ipblocktypeDescCreatedAt is the schema descriptor for created_at field.
	ipblocktypeDescCreatedAt := ipblocktypeMixinFields0[0].Descriptor()
	// ipblocktype.DefaultCreatedAt holds the default value on creation for the created_at field.
	ipblocktype.DefaultCreatedAt = ipblocktypeDescCreatedAt.Default.(func() time.Time)
	// ipblocktypeDescUpdatedAt is the schema descriptor for updated_at field.
	ipblocktypeDescUpdatedAt := ipblocktypeMixinFields0[1].Descriptor()
	// ipblocktype.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	ipblocktype.DefaultUpdatedAt = ipblocktypeDescUpdatedAt.Default.(func() time.Time)
	// ipblocktype.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	ipblocktype.UpdateDefaultUpdatedAt = ipblocktypeDescUpdatedAt.UpdateDefault.(func() time.Time)
	// ipblocktypeDescName is the schema descriptor for name field.
	ipblocktypeDescName := ipblocktypeFields[1].Descriptor()
	// ipblocktype.NameValidator is a validator for the "name" field. It is called by the builders before save.
	ipblocktype.NameValidator = ipblocktypeDescName.Validators[0].(func(string) error)
	// ipblocktypeDescID is the schema descriptor for id field.
	ipblocktypeDescID := ipblocktypeFields[0].Descriptor()
	// ipblocktype.DefaultID holds the default value on creation for the id field.
	ipblocktype.DefaultID = ipblocktypeDescID.Default.(func() gidx.PrefixedID)
}
