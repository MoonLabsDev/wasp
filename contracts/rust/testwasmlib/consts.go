// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/vm/wasmlib"

const (
	ScName        = "testwasmlib"
	ScDescription = "Exercise all aspects of WasmLib"
	HScName       = wasmlib.ScHname(0x89703a45)
)

const (
	ParamAddress     = wasmlib.Key("address")
	ParamAgentID     = wasmlib.Key("agentID")
	ParamBlockIndex  = wasmlib.Key("blockIndex")
	ParamBytes       = wasmlib.Key("bytes")
	ParamChainID     = wasmlib.Key("chainID")
	ParamColor       = wasmlib.Key("color")
	ParamHash        = wasmlib.Key("hash")
	ParamHname       = wasmlib.Key("hname")
	ParamIndex       = wasmlib.Key("index")
	ParamInt16       = wasmlib.Key("int16")
	ParamInt32       = wasmlib.Key("int32")
	ParamInt64       = wasmlib.Key("int64")
	ParamName        = wasmlib.Key("name")
	ParamParam       = wasmlib.Key("@")
	ParamRecordIndex = wasmlib.Key("recordIndex")
	ParamRequestID   = wasmlib.Key("requestID")
	ParamString      = wasmlib.Key("string")
	ParamValue       = wasmlib.Key("value")
)

const (
	ResultCount  = wasmlib.Key("count")
	ResultIotas  = wasmlib.Key("iotas")
	ResultLength = wasmlib.Key("length")
	ResultRecord = wasmlib.Key("record")
	ResultValue  = wasmlib.Key("value")
)

const StateArrays = wasmlib.Key("arrays")

const (
	FuncArrayClear   = "arrayClear"
	FuncArrayCreate  = "arrayCreate"
	FuncArraySet     = "arraySet"
	FuncParamTypes   = "paramTypes"
	ViewArrayLength  = "arrayLength"
	ViewArrayValue   = "arrayValue"
	ViewBlockRecord  = "blockRecord"
	ViewBlockRecords = "blockRecords"
	ViewIotaBalance  = "iotaBalance"
)

const (
	HFuncArrayClear   = wasmlib.ScHname(0x88021821)
	HFuncArrayCreate  = wasmlib.ScHname(0x1ed5b23b)
	HFuncArraySet     = wasmlib.ScHname(0x2c4150b3)
	HFuncParamTypes   = wasmlib.ScHname(0x6921c4cd)
	HViewArrayLength  = wasmlib.ScHname(0x3a831021)
	HViewArrayValue   = wasmlib.ScHname(0x662dbd81)
	HViewBlockRecord  = wasmlib.ScHname(0xad13b2f8)
	HViewBlockRecords = wasmlib.ScHname(0x16e249ea)
	HViewIotaBalance  = wasmlib.ScHname(0x9d3920bd)
)
