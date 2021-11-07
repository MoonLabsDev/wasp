// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

package generator

import (
	"regexp"

	"github.com/iotaledger/wasp/tools/schema/generator/gotemplates"
)

var goTypes = StringMap{
	"Address":   "wasmlib.ScAddress",
	"AgentID":   "wasmlib.ScAgentID",
	"ChainID":   "wasmlib.ScChainID",
	"Color":     "wasmlib.ScColor",
	"Hash":      "wasmlib.ScHash",
	"Hname":     "wasmlib.ScHname",
	"Int16":     "int16",
	"Int32":     "int32",
	"Int64":     "int64",
	"RequestID": "wasmlib.ScRequestID",
	"String":    "string",
}

var goKeys = StringMap{
	"Address":   "key",
	"AgentID":   "key",
	"ChainID":   "key",
	"Color":     "key",
	"Hash":      "key",
	"Hname":     "key",
	"Int16":     "??TODO",
	"Int32":     "wasmlib.Key32(key)",
	"Int64":     "??TODO",
	"RequestID": "key",
	"String":    "wasmlib.Key(key)",
}

var goTypeIds = StringMap{
	"Address":   "wasmlib.TYPE_ADDRESS",
	"AgentID":   "wasmlib.TYPE_AGENT_ID",
	"ChainID":   "wasmlib.TYPE_CHAIN_ID",
	"Color":     "wasmlib.TYPE_COLOR",
	"Hash":      "wasmlib.TYPE_HASH",
	"Hname":     "wasmlib.TYPE_HNAME",
	"Int16":     "wasmlib.TYPE_INT16",
	"Int32":     "wasmlib.TYPE_INT32",
	"Int64":     "wasmlib.TYPE_INT64",
	"RequestID": "wasmlib.TYPE_REQUEST_ID",
	"String":    "wasmlib.TYPE_STRING",
}

type GoGenerator struct {
	GenBase
}

func NewGoGenerator() *GoGenerator {
	g := &GoGenerator{}
	g.extension = ".go"
	g.funcRegexp = regexp.MustCompile(`^func (\w+).+$`)
	g.language = "Go"
	g.rootFolder = "go"
	g.gen = g
	return g
}

func (g *GoGenerator) init(s *Schema) {
	g.GenBase.init(s, gotemplates.GoTemplates)
}

func (g *GoGenerator) funcName(f *Func) string {
	return f.FuncName
}

func (g *GoGenerator) generateLanguageSpecificFiles() error {
	if g.s.CoreContracts {
		return nil
	}
	return g.createSourceFile("../main", g.writeSpecialMain)
}

func (g *GoGenerator) writeInitialFuncs() {
	g.emit("funcs.go")
}

func (g *GoGenerator) writeSpecialMain() {
	g.emit("main.go")
}

func (g *GoGenerator) setFieldKeys() {
	g.GenBase.setFieldKeys()

	fldTypeID := goTypeIds[g.currentField.Type]
	if fldTypeID == "" {
		fldTypeID = "wasmlib.TYPE_BYTES"
	}
	g.keys["FldTypeID"] = fldTypeID
	g.keys["FldTypeKey"] = goKeys[g.currentField.Type]
	g.keys["FldLangType"] = goTypes[g.currentField.Type]
	g.keys["FldMapKeyLangType"] = goTypes[g.currentField.MapKey]
	g.keys["FldMapKeyKey"] = goKeys[g.currentField.MapKey]
}

func (g *GoGenerator) setFuncKeys() {
	g.GenBase.setFuncKeys()

	initFunc := ""
	initMap := ""
	if g.currentFunc.Type == InitFunc {
		initFunc = InitFunc
		initMap = ", keyMap[:], idxMap[:]"
	}
	g.keys["initFunc"] = initFunc
	g.keys["initMap"] = initMap
}
