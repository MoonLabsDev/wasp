// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package inccounter

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

type ImmutableInitParams struct {
	id int32
}

func (s ImmutableInitParams) Counter() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamCounter])
}

type MutableInitParams struct {
	id int32
}

func (s MutableInitParams) Counter() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamCounter])
}

type ImmutableRepeatManyParams struct {
	id int32
}

func (s ImmutableRepeatManyParams) NumRepeats() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamNumRepeats])
}

type MutableRepeatManyParams struct {
	id int32
}

func (s MutableRepeatManyParams) NumRepeats() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamNumRepeats])
}

type ImmutableWhenMustIncrementParams struct {
	id int32
}

func (s ImmutableWhenMustIncrementParams) Dummy() wasmlib.ScImmutableInt64 {
	return wasmlib.NewScImmutableInt64(s.id, idxMap[IdxParamDummy])
}

type MutableWhenMustIncrementParams struct {
	id int32
}

func (s MutableWhenMustIncrementParams) Dummy() wasmlib.ScMutableInt64 {
	return wasmlib.NewScMutableInt64(s.id, idxMap[IdxParamDummy])
}
