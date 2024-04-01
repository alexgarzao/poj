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

	st.AddVariable("myvar", codegen.Integer)
	ok, symbol = st.Get("myvar")
	assert.Equal(t, true, ok)
	assert.Equal(t, codegen.Variable, symbol.SymbolType)
	assert.Equal(t, codegen.Integer, symbol.PascalType)
	assert.Equal(t, 1, symbol.Index)
}
