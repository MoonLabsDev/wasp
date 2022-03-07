// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

// (Re-)generated by schema tool
// >>>> DO NOT CHANGE THIS FILE! <<<<
// Change the json schema instead

package testwasmlib

import "github.com/iotaledger/wasp/packages/wasmvm/wasmlib/go/wasmlib/wasmtypes"

type ImmutableArrayOfArraysAddrAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfArraysAddrAppendParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfArraysAddrAppendParams) ValueAddr() ArrayOfImmutableAddress {
	return ArrayOfImmutableAddress{proxy: s.proxy.Root(ParamValueAddr)}
}

type MutableArrayOfArraysAddrAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfArraysAddrAppendParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfArraysAddrAppendParams) ValueAddr() ArrayOfMutableAddress {
	return ArrayOfMutableAddress{proxy: s.proxy.Root(ParamValueAddr)}
}

type ImmutableArrayOfArraysAddrSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfArraysAddrSetParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfArraysAddrSetParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

func (s ImmutableArrayOfArraysAddrSetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableArrayOfArraysAddrSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfArraysAddrSetParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfArraysAddrSetParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

func (s MutableArrayOfArraysAddrSetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableArrayOfArraysAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfArraysAppendParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfArraysAppendParams) Value() ArrayOfImmutableString {
	return ArrayOfImmutableString{proxy: s.proxy.Root(ParamValue)}
}

type MutableArrayOfArraysAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfArraysAppendParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfArraysAppendParams) Value() ArrayOfMutableString {
	return ArrayOfMutableString{proxy: s.proxy.Root(ParamValue)}
}

type ImmutableArrayOfArraysSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfArraysSetParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfArraysSetParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

func (s ImmutableArrayOfArraysSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableArrayOfArraysSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfArraysSetParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfArraysSetParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

func (s MutableArrayOfArraysSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableArrayOfMapsSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfMapsSetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfMapsSetParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableArrayOfMapsSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableArrayOfMapsSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfMapsSetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfMapsSetParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableArrayOfMapsSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableMapOfArraysAddrAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysAddrAppendParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s ImmutableMapOfArraysAddrAppendParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableMapOfArraysAddrAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysAddrAppendParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s MutableMapOfArraysAddrAppendParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableMapOfArraysAddrClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysAddrClearParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableMapOfArraysAddrClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysAddrClearParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableMapOfArraysAddrSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysAddrSetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableMapOfArraysAddrSetParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s ImmutableMapOfArraysAddrSetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableMapOfArraysAddrSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysAddrSetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableMapOfArraysAddrSetParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s MutableMapOfArraysAddrSetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableMapOfArraysAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysAppendParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableMapOfArraysAppendParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableMapOfArraysAppendParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysAppendParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableMapOfArraysAppendParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableMapOfArraysClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysClearParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapOfArraysClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysClearParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableMapOfArraysSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysSetParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableMapOfArraysSetParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableMapOfArraysSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableMapOfArraysSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysSetParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableMapOfArraysSetParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableMapOfArraysSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type ImmutableMapOfMapsAddrClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfMapsAddrClearParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableMapOfMapsAddrClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfMapsAddrClearParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableMapOfMapsAddrSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfMapsAddrSetParams) KeyAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s ImmutableMapOfMapsAddrSetParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s ImmutableMapOfMapsAddrSetParams) ValueAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamValueAddr))
}

type MutableMapOfMapsAddrSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfMapsAddrSetParams) KeyAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s MutableMapOfMapsAddrSetParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

func (s MutableMapOfMapsAddrSetParams) ValueAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamValueAddr))
}

type ImmutableMapOfMapsClearParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfMapsClearParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapOfMapsClearParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfMapsClearParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableMapOfMapsSetParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfMapsSetParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableMapOfMapsSetParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

func (s ImmutableMapOfMapsSetParams) Value() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamValue))
}

type MutableMapOfMapsSetParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfMapsSetParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableMapOfMapsSetParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

func (s MutableMapOfMapsSetParams) Value() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamValue))
}

type MapStringToImmutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToImmutableBytes) GetBytes(key string) wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type ImmutableParamTypesParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableParamTypesParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutableParamTypesParams) AgentID() wasmtypes.ScImmutableAgentID {
	return wasmtypes.NewScImmutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s ImmutableParamTypesParams) Bool() wasmtypes.ScImmutableBool {
	return wasmtypes.NewScImmutableBool(s.proxy.Root(ParamBool))
}

func (s ImmutableParamTypesParams) Bytes() wasmtypes.ScImmutableBytes {
	return wasmtypes.NewScImmutableBytes(s.proxy.Root(ParamBytes))
}

func (s ImmutableParamTypesParams) ChainID() wasmtypes.ScImmutableChainID {
	return wasmtypes.NewScImmutableChainID(s.proxy.Root(ParamChainID))
}

func (s ImmutableParamTypesParams) Color() wasmtypes.ScImmutableColor {
	return wasmtypes.NewScImmutableColor(s.proxy.Root(ParamColor))
}

func (s ImmutableParamTypesParams) Hash() wasmtypes.ScImmutableHash {
	return wasmtypes.NewScImmutableHash(s.proxy.Root(ParamHash))
}

func (s ImmutableParamTypesParams) Hname() wasmtypes.ScImmutableHname {
	return wasmtypes.NewScImmutableHname(s.proxy.Root(ParamHname))
}

func (s ImmutableParamTypesParams) Int16() wasmtypes.ScImmutableInt16 {
	return wasmtypes.NewScImmutableInt16(s.proxy.Root(ParamInt16))
}

func (s ImmutableParamTypesParams) Int32() wasmtypes.ScImmutableInt32 {
	return wasmtypes.NewScImmutableInt32(s.proxy.Root(ParamInt32))
}

func (s ImmutableParamTypesParams) Int64() wasmtypes.ScImmutableInt64 {
	return wasmtypes.NewScImmutableInt64(s.proxy.Root(ParamInt64))
}

func (s ImmutableParamTypesParams) Int8() wasmtypes.ScImmutableInt8 {
	return wasmtypes.NewScImmutableInt8(s.proxy.Root(ParamInt8))
}

func (s ImmutableParamTypesParams) Param() MapStringToImmutableBytes {
	//nolint:gosimple
	return MapStringToImmutableBytes{proxy: s.proxy}
}

func (s ImmutableParamTypesParams) RequestID() wasmtypes.ScImmutableRequestID {
	return wasmtypes.NewScImmutableRequestID(s.proxy.Root(ParamRequestID))
}

func (s ImmutableParamTypesParams) String() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamString))
}

func (s ImmutableParamTypesParams) Uint16() wasmtypes.ScImmutableUint16 {
	return wasmtypes.NewScImmutableUint16(s.proxy.Root(ParamUint16))
}

func (s ImmutableParamTypesParams) Uint32() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamUint32))
}

func (s ImmutableParamTypesParams) Uint64() wasmtypes.ScImmutableUint64 {
	return wasmtypes.NewScImmutableUint64(s.proxy.Root(ParamUint64))
}

func (s ImmutableParamTypesParams) Uint8() wasmtypes.ScImmutableUint8 {
	return wasmtypes.NewScImmutableUint8(s.proxy.Root(ParamUint8))
}

type MapStringToMutableBytes struct {
	proxy wasmtypes.Proxy
}

func (m MapStringToMutableBytes) Clear() {
	m.proxy.ClearMap()
}

func (m MapStringToMutableBytes) GetBytes(key string) wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(m.proxy.Key(wasmtypes.StringToBytes(key)))
}

type MutableParamTypesParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableParamTypesParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutableParamTypesParams) AgentID() wasmtypes.ScMutableAgentID {
	return wasmtypes.NewScMutableAgentID(s.proxy.Root(ParamAgentID))
}

func (s MutableParamTypesParams) Bool() wasmtypes.ScMutableBool {
	return wasmtypes.NewScMutableBool(s.proxy.Root(ParamBool))
}

func (s MutableParamTypesParams) Bytes() wasmtypes.ScMutableBytes {
	return wasmtypes.NewScMutableBytes(s.proxy.Root(ParamBytes))
}

func (s MutableParamTypesParams) ChainID() wasmtypes.ScMutableChainID {
	return wasmtypes.NewScMutableChainID(s.proxy.Root(ParamChainID))
}

func (s MutableParamTypesParams) Color() wasmtypes.ScMutableColor {
	return wasmtypes.NewScMutableColor(s.proxy.Root(ParamColor))
}

func (s MutableParamTypesParams) Hash() wasmtypes.ScMutableHash {
	return wasmtypes.NewScMutableHash(s.proxy.Root(ParamHash))
}

func (s MutableParamTypesParams) Hname() wasmtypes.ScMutableHname {
	return wasmtypes.NewScMutableHname(s.proxy.Root(ParamHname))
}

func (s MutableParamTypesParams) Int16() wasmtypes.ScMutableInt16 {
	return wasmtypes.NewScMutableInt16(s.proxy.Root(ParamInt16))
}

func (s MutableParamTypesParams) Int32() wasmtypes.ScMutableInt32 {
	return wasmtypes.NewScMutableInt32(s.proxy.Root(ParamInt32))
}

func (s MutableParamTypesParams) Int64() wasmtypes.ScMutableInt64 {
	return wasmtypes.NewScMutableInt64(s.proxy.Root(ParamInt64))
}

func (s MutableParamTypesParams) Int8() wasmtypes.ScMutableInt8 {
	return wasmtypes.NewScMutableInt8(s.proxy.Root(ParamInt8))
}

func (s MutableParamTypesParams) Param() MapStringToMutableBytes {
	//nolint:gosimple
	return MapStringToMutableBytes{proxy: s.proxy}
}

func (s MutableParamTypesParams) RequestID() wasmtypes.ScMutableRequestID {
	return wasmtypes.NewScMutableRequestID(s.proxy.Root(ParamRequestID))
}

func (s MutableParamTypesParams) String() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamString))
}

func (s MutableParamTypesParams) Uint16() wasmtypes.ScMutableUint16 {
	return wasmtypes.NewScMutableUint16(s.proxy.Root(ParamUint16))
}

func (s MutableParamTypesParams) Uint32() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamUint32))
}

func (s MutableParamTypesParams) Uint64() wasmtypes.ScMutableUint64 {
	return wasmtypes.NewScMutableUint64(s.proxy.Root(ParamUint64))
}

func (s MutableParamTypesParams) Uint8() wasmtypes.ScMutableUint8 {
	return wasmtypes.NewScMutableUint8(s.proxy.Root(ParamUint8))
}

type ImmutableTriggerEventParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableTriggerEventParams) Address() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamAddress))
}

func (s ImmutableTriggerEventParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableTriggerEventParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableTriggerEventParams) Address() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamAddress))
}

func (s MutableTriggerEventParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableArrayOfArraysAddrValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfArraysAddrValueParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfArraysAddrValueParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

type MutableArrayOfArraysAddrValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfArraysAddrValueParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfArraysAddrValueParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

type ImmutableArrayOfArraysValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfArraysValueParams) Index0() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex0))
}

func (s ImmutableArrayOfArraysValueParams) Index1() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex1))
}

type MutableArrayOfArraysValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfArraysValueParams) Index0() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex0))
}

func (s MutableArrayOfArraysValueParams) Index1() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex1))
}

type ImmutableArrayOfMapsValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableArrayOfMapsValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableArrayOfMapsValueParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

type MutableArrayOfMapsValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableArrayOfMapsValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableArrayOfMapsValueParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

type ImmutableBlockRecordParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBlockRecordParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

func (s ImmutableBlockRecordParams) RecordIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamRecordIndex))
}

type MutableBlockRecordParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBlockRecordParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

func (s MutableBlockRecordParams) RecordIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamRecordIndex))
}

type ImmutableBlockRecordsParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableBlockRecordsParams) BlockIndex() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamBlockIndex))
}

type MutableBlockRecordsParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableBlockRecordsParams) BlockIndex() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamBlockIndex))
}

type ImmutableMapOfArraysAddrLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysAddrLengthParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableMapOfArraysAddrLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysAddrLengthParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableMapOfArraysAddrValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysAddrValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableMapOfArraysAddrValueParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableMapOfArraysAddrValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysAddrValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableMapOfArraysAddrValueParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableMapOfArraysLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysLengthParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapOfArraysLengthParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysLengthParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableMapOfArraysValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfArraysValueParams) Index() wasmtypes.ScImmutableUint32 {
	return wasmtypes.NewScImmutableUint32(s.proxy.Root(ParamIndex))
}

func (s ImmutableMapOfArraysValueParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapOfArraysValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfArraysValueParams) Index() wasmtypes.ScMutableUint32 {
	return wasmtypes.NewScMutableUint32(s.proxy.Root(ParamIndex))
}

func (s MutableMapOfArraysValueParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}

type ImmutableMapOfMapsAddrValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfMapsAddrValueParams) KeyAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s ImmutableMapOfMapsAddrValueParams) NameAddr() wasmtypes.ScImmutableAddress {
	return wasmtypes.NewScImmutableAddress(s.proxy.Root(ParamNameAddr))
}

type MutableMapOfMapsAddrValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfMapsAddrValueParams) KeyAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamKeyAddr))
}

func (s MutableMapOfMapsAddrValueParams) NameAddr() wasmtypes.ScMutableAddress {
	return wasmtypes.NewScMutableAddress(s.proxy.Root(ParamNameAddr))
}

type ImmutableMapOfMapsValueParams struct {
	proxy wasmtypes.Proxy
}

func (s ImmutableMapOfMapsValueParams) Key() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamKey))
}

func (s ImmutableMapOfMapsValueParams) Name() wasmtypes.ScImmutableString {
	return wasmtypes.NewScImmutableString(s.proxy.Root(ParamName))
}

type MutableMapOfMapsValueParams struct {
	proxy wasmtypes.Proxy
}

func (s MutableMapOfMapsValueParams) Key() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamKey))
}

func (s MutableMapOfMapsValueParams) Name() wasmtypes.ScMutableString {
	return wasmtypes.NewScMutableString(s.proxy.Root(ParamName))
}
