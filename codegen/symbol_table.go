package codegen

import (
	"fmt"
	"strings"
)

type SymbolType uint8

const (
	UndefinedSymbolType SymbolType = iota
	Variable
)

type Symbol struct {
	SymbolType SymbolType
	PascalType PascalType
	Index      int
}

type SymbolTable struct {
	symbols map[string]Symbol
	count   int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		symbols: make(map[string]Symbol),
		count:   0,
	}
}

func (st *SymbolTable) AddVariable(name string, ptype PascalType) error {
	name = strings.ToUpper(name)
	if _, ok := st.symbols[name]; ok {
		return fmt.Errorf("variable %s already declared", name)
	}

	st.count++
	st.symbols[name] = Symbol{
		SymbolType: Variable,
		PascalType: ptype,
		Index:      st.count,
	}

	return nil
}

func (st *SymbolTable) Get(name string) (bool, Symbol) {
	name = strings.ToUpper(name)
	symbol, ok := st.symbols[name]
	if !ok {
		return false, Symbol{
			SymbolType: UndefinedSymbolType,
			PascalType: Undefined,
			Index:      -1,
		}
	}

	return true, symbol
}
