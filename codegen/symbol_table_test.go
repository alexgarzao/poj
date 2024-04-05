package codegen_test

import (
	"testing"

	"github.com/alexgarzao/poj/codegen"
	"github.com/stretchr/testify/assert"
)

func TestSymbolTable_AddVariable(t *testing.T) {
	st := codegen.NewSymbolTable()
	ok, symbol := st.Get("xpto")
	assert.Equal(t, false, ok)
	assert.Equal(t, codegen.UndefinedSymbolType, symbol.SymbolType)
	assert.Equal(t, codegen.Undefined, symbol.PascalType)
	assert.Equal(t, -1, symbol.Index)

	err := st.AddVariable("myvar", codegen.Integer)
	assert.Nil(t, err)

	ok, symbol = st.Get("myvar")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Variable, symbol.SymbolType)
	assert.Equal(t, codegen.Integer, symbol.PascalType)
	assert.Equal(t, 1, symbol.Index)

	ok, symbol = st.Get("MyVar")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Variable, symbol.SymbolType)
	assert.Equal(t, codegen.Integer, symbol.PascalType)
	assert.Equal(t, 1, symbol.Index)

	err = st.AddVariable("myvar", codegen.Integer)
	assert.NotNil(t, err)

	err = st.AddVariable("MyVar", codegen.Integer)
	assert.NotNil(t, err)
}
