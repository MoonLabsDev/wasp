// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

const (
	ScName        = "testwasmlib"
	ScDescription = "Exercise several aspects of WasmLib"
	HScName       = wasmtypes.ScHname(0x89703a45)
)

const (
	ParamAddress     = "address"
	ParamAgentID     = "agentID"
	ParamBlockIndex  = "blockIndex"
	ParamBool        = "bool"
	ParamBytes       = "bytes"
	ParamChainID     = "chainID"
	ParamColor       = "color"
	ParamHash        = "hash"
	ParamHname       = "hname"
	ParamIndex       = "index"
	ParamIndex0      = "index0"
	ParamIndex1      = "index1"
	ParamInt16       = "int16"
	ParamInt32       = "int32"
	ParamInt64       = "int64"
	ParamInt8        = "int8"
	ParamKey         = "key"
	ParamKeyAddr     = "keyAddr"
	ParamName        = "name"
	ParamNameAddr    = "nameAddr"
	ParamParam       = "this"
	ParamRecordIndex = "recordIndex"
	ParamRequestID   = "requestID"
	ParamString      = "string"
	ParamUint16      = "uint16"
	ParamUint32      = "uint32"
	ParamUint64      = "uint64"
	ParamUint8       = "uint8"
	ParamValue       = "value"
	ParamValueAddr   = "valueAddr"
)

const (
	ResultCount     = "count"
	ResultIotas     = "iotas"
	ResultLength    = "length"
	ResultRandom    = "random"
	ResultRecord    = "record"
	ResultValue     = "value"
	ResultValueAddr = "valueAddr"
)

const (
	StateAddrArrayOfArrays   = "addrArrayOfArrays"
	StateAddrArrayOfMaps     = "addrArrayOfMaps"
	StateAddrMapOfArrays     = "addrMapOfArrays"
	StateAddrMapOfMaps       = "addrMapOfMaps"
	StateLatLong             = "latLong"
	StateRandom              = "random"
	StateStringArrayOfArrays = "stringArrayOfArrays"
	StateStringArrayOfMaps   = "stringArrayOfMaps"
	StateStringMapOfArrays   = "stringMapOfArrays"
	StateStringMapOfMaps     = "stringMapOfMaps"
)

const (
	FuncArrayOfArraysAddrAppend = "arrayOfArraysAddrAppend"
	FuncArrayOfArraysAddrClear  = "arrayOfArraysAddrClear"
	FuncArrayOfArraysAddrSet    = "arrayOfArraysAddrSet"
	FuncArrayOfArraysAppend     = "arrayOfArraysAppend"
	FuncArrayOfArraysClear      = "arrayOfArraysClear"
	FuncArrayOfArraysSet        = "arrayOfArraysSet"
	FuncArrayOfMapsClear        = "arrayOfMapsClear"
	FuncArrayOfMapsSet          = "arrayOfMapsSet"
	FuncMapOfArraysAddrAppend   = "mapOfArraysAddrAppend"
	FuncMapOfArraysAddrClear    = "mapOfArraysAddrClear"
	FuncMapOfArraysAddrSet      = "mapOfArraysAddrSet"
	FuncMapOfArraysAppend       = "mapOfArraysAppend"
	FuncMapOfArraysClear        = "mapOfArraysClear"
	FuncMapOfArraysSet          = "mapOfArraysSet"
	FuncMapOfMapsAddrClear      = "mapOfMapsAddrClear"
	FuncMapOfMapsAddrSet        = "mapOfMapsAddrSet"
	FuncMapOfMapsClear          = "mapOfMapsClear"
	FuncMapOfMapsSet            = "mapOfMapsSet"
	FuncParamTypes              = "paramTypes"
	FuncRandom                  = "random"
	FuncTakeAllowance           = "takeAllowance"
	FuncTakeBalance             = "takeBalance"
	FuncTriggerEvent            = "triggerEvent"
	ViewArrayOfArraysAddrLength = "arrayOfArraysAddrLength"
	ViewArrayOfArraysAddrValue  = "arrayOfArraysAddrValue"
	ViewArrayOfArraysLength     = "arrayOfArraysLength"
	ViewArrayOfArraysValue      = "arrayOfArraysValue"
	ViewArrayOfMapsValue        = "arrayOfMapsValue"
	ViewBlockRecord             = "blockRecord"
	ViewBlockRecords            = "blockRecords"
	ViewGetRandom               = "getRandom"
	ViewIotaBalance             = "iotaBalance"
	ViewMapOfArraysAddrLength   = "mapOfArraysAddrLength"
	ViewMapOfArraysAddrValue    = "mapOfArraysAddrValue"
	ViewMapOfArraysLength       = "mapOfArraysLength"
	ViewMapOfArraysValue        = "mapOfArraysValue"
	ViewMapOfMapsAddrValue      = "mapOfMapsAddrValue"
	ViewMapOfMapsValue          = "mapOfMapsValue"
)

const (
	HFuncArrayOfArraysAddrAppend = wasmtypes.ScHname(0xfa499ba1)
	HFuncArrayOfArraysAddrClear  = wasmtypes.ScHname(0x951df3e0)
	HFuncArrayOfArraysAddrSet    = wasmtypes.ScHname(0x42e9eb75)
	HFuncArrayOfArraysAppend     = wasmtypes.ScHname(0x23f3a17e)
	HFuncArrayOfArraysClear      = wasmtypes.ScHname(0xc826f36f)
	HFuncArrayOfArraysSet        = wasmtypes.ScHname(0x74d77052)
	HFuncArrayOfMapsClear        = wasmtypes.ScHname(0x974adaae)
	HFuncArrayOfMapsSet          = wasmtypes.ScHname(0x9e7baa47)
	HFuncMapOfArraysAddrAppend   = wasmtypes.ScHname(0x4ba0fac9)
	HFuncMapOfArraysAddrClear    = wasmtypes.ScHname(0x0f9240f9)
	HFuncMapOfArraysAddrSet      = wasmtypes.ScHname(0x717f354d)
	HFuncMapOfArraysAppend       = wasmtypes.ScHname(0x072902d4)
	HFuncMapOfArraysClear        = wasmtypes.ScHname(0xdcdbc582)
	HFuncMapOfArraysSet          = wasmtypes.ScHname(0xcdbf9981)
	HFuncMapOfMapsAddrClear      = wasmtypes.ScHname(0x9cf5dec9)
	HFuncMapOfMapsAddrSet        = wasmtypes.ScHname(0x39fa7efb)
	HFuncMapOfMapsClear          = wasmtypes.ScHname(0xd02a5431)
	HFuncMapOfMapsSet            = wasmtypes.ScHname(0x353d577f)
	HFuncParamTypes              = wasmtypes.ScHname(0x6921c4cd)
	HFuncRandom                  = wasmtypes.ScHname(0xe86c97ca)
	HFuncTakeAllowance           = wasmtypes.ScHname(0x91e7bd00)
	HFuncTakeBalance             = wasmtypes.ScHname(0x8ad1cb27)
	HFuncTriggerEvent            = wasmtypes.ScHname(0xd5438ac6)
	HViewArrayOfArraysAddrLength = wasmtypes.ScHname(0x0e7d3af5)
	HViewArrayOfArraysAddrValue  = wasmtypes.ScHname(0xeae25944)
	HViewArrayOfArraysLength     = wasmtypes.ScHname(0x5e918d60)
	HViewArrayOfArraysValue      = wasmtypes.ScHname(0x41d5f686)
	HViewArrayOfMapsValue        = wasmtypes.ScHname(0x77e1ef85)
	HViewBlockRecord             = wasmtypes.ScHname(0xad13b2f8)
	HViewBlockRecords            = wasmtypes.ScHname(0x16e249ea)
	HViewGetRandom               = wasmtypes.ScHname(0x46263045)
	HViewIotaBalance             = wasmtypes.ScHname(0x9d3920bd)
	HViewMapOfArraysAddrLength   = wasmtypes.ScHname(0x3910ae79)
	HViewMapOfArraysAddrValue    = wasmtypes.ScHname(0x81b3c951)
	HViewMapOfArraysLength       = wasmtypes.ScHname(0x86379ff7)
	HViewMapOfArraysValue        = wasmtypes.ScHname(0x8dee3538)
	HViewMapOfMapsAddrValue      = wasmtypes.ScHname(0xf9538970)
	HViewMapOfMapsValue          = wasmtypes.ScHname(0x476c56e4)
)
