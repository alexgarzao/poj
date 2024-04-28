package codegen

import (
	"fmt"
	"strings"
)

type SymbolType uint8

const (
	UndefinedSymbolType SymbolType = iota
	Variable
	Procedure
)

type Symbol struct {
	SymbolType SymbolType
	PascalType PascalType
	Index      int
	ParamTypes []string
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

	st.symbols[name] = Symbol{
		SymbolType: Variable,
		PascalType: ptype,
		Index:      st.count,
	}

	st.count++

	return nil
}

func (st *SymbolTable) AddProcedure(name string, paramTypes []string) error {
	name = strings.ToUpper(name)
	if _, ok := st.symbols[name]; ok {
		return fmt.Errorf("procedure %s already declared", name)
	}

	st.symbols[name] = Symbol{ // REFACTOR: Symbol is only to variables?
		SymbolType: Procedure,
		ParamTypes: paramTypes[:],
		// PascalType: ptype,
		// Index:      st.count,
	}

	return nil
}

func (st *SymbolTable) Get(name string) (Symbol, bool) {
	name = strings.ToUpper(name)
	symbol, ok := st.symbols[name]
	if !ok {
		return Symbol{
			SymbolType: UndefinedSymbolType,
			PascalType: Undefined,
			Index:      -1,
			ParamTypes: nil,
		}, false
	}

	return symbol, true
}
