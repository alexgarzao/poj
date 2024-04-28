package codegen_test

import (
	"testing"

	"github.com/alexgarzao/poj/codegen"
	"github.com/stretchr/testify/assert"
)

func TestSymbolTable_AddVariable(t *testing.T) {
	st := codegen.NewSymbolTable()
	symbol, ok := st.Get("xpto")
	assert.Equal(t, false, ok)
	assert.Equal(t, codegen.UndefinedSymbolType, symbol.SymbolType)
	assert.Equal(t, codegen.Undefined, symbol.PascalType)
	assert.Equal(t, -1, symbol.Index)

	err := st.AddVariable("myvar", codegen.Integer)
	assert.Nil(t, err)

	symbol, ok = st.Get("myvar")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Variable, symbol.SymbolType)
	assert.Equal(t, codegen.Integer, symbol.PascalType)
	assert.Equal(t, 0, symbol.Index)

	symbol, ok = st.Get("MyVar")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Variable, symbol.SymbolType)
	assert.Equal(t, codegen.Integer, symbol.PascalType)
	assert.Equal(t, 0, symbol.Index)

	err = st.AddVariable("myvar", codegen.Integer)
	assert.NotNil(t, err)

	err = st.AddVariable("MyVar", codegen.Integer)
	assert.NotNil(t, err)
}

func TestSymbolTable_AddProcedure(t *testing.T) {
	st := codegen.NewSymbolTable()
	symbol, ok := st.Get("xpto")
	assert.Equal(t, false, ok)
	assert.Equal(t, codegen.UndefinedSymbolType, symbol.SymbolType)
	assert.Equal(t, codegen.Undefined, symbol.PascalType)
	assert.Equal(t, -1, symbol.Index)

	err := st.AddProcedure("myproc", []string{})
	assert.Nil(t, err)

	symbol, ok = st.Get("myproc")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Procedure, symbol.SymbolType)
	assert.Equal(t, codegen.Undefined, symbol.PascalType)
	assert.Equal(t, 0, symbol.Index)

	symbol, ok = st.Get("MyProc")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Procedure, symbol.SymbolType)
	assert.Equal(t, codegen.Undefined, symbol.PascalType)
	assert.Equal(t, 0, symbol.Index)

	err = st.AddProcedure("myproc", []string{})
	assert.NotNil(t, err)

	err = st.AddProcedure("MyProc", []string{})
	assert.NotNil(t, err)

	err = st.AddVariable("MyProc", codegen.Integer)
	assert.NotNil(t, err)

	err = st.AddProcedure("myproc2", []string{"integer", "string"})
	assert.Nil(t, err)

	symbol, ok = st.Get("myproc2")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Procedure, symbol.SymbolType)
	assert.Equal(t, []string{"integer", "string"}, symbol.ParamTypes)
	assert.Equal(t, codegen.Undefined, symbol.PascalType)
	assert.Equal(t, 0, symbol.Index)
}
