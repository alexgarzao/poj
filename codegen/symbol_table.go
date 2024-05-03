package codegen

import (
	"fmt"
	"strings"
)

const fixedFunctionVarIndex = 100

type SymbolTable struct {
	gsymbols map[string]Symbol
	lsymbols map[string]Symbol
	lcount   int
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		gsymbols: make(map[string]Symbol),
		lsymbols: make(map[string]Symbol),
		lcount:   0,
	}
}

func (st *SymbolTable) AddGlobalVariable(name string, ptype PascalType) error {
	return st.addVariable(true, name, ptype)
}

func (st *SymbolTable) AddLocalVariable(name string, ptype PascalType) error {
	return st.addVariable(false, name, ptype)
}

func (st *SymbolTable) AddProcedure(name string, paramTypes []string) error {
	name = strings.ToUpper(name)
	if _, ok := st.gsymbols[name]; ok {
		return fmt.Errorf("procedure %s already declared", name)
	}

	st.gsymbols[name] = Symbol{ // REFACTOR: Symbol is only to variables?
		SymbolType: Procedure,
		ParamTypes: paramTypes[:],
		// PascalType: ptype,
		// Index:      st.count,
	}

	return nil
}

func (st *SymbolTable) AddFunction(name string, paramTypes []string, returnType string) error {
	name = strings.ToUpper(name)
	if _, ok := st.gsymbols[name]; ok {
		return fmt.Errorf("function %s already declared", name)
	}

	st.gsymbols[name] = Symbol{
		SymbolType: Function,
		ParamTypes: paramTypes[:],
		PascalType: ToPascalType(returnType),
		Index:      fixedFunctionVarIndex,
	}

	return nil
}

func (st *SymbolTable) Get(name string) (Symbol, bool) {
	name = strings.ToUpper(name)

	// Local symbol?
	symbol, ok := st.lsymbols[name]
	if ok {
		return symbol, true
	}

	// Global symbol?
	symbol, ok = st.gsymbols[name]
	if ok {
		return symbol, true
	}

	// Undefined symbol
	return Symbol{
		SymbolType: UndefinedSymbolType,
		PascalType: Undefined,
		Index:      -1,
		ParamTypes: nil,
	}, false
}

func (st *SymbolTable) CleanLocalSymbols() {
	st.lsymbols = make(map[string]Symbol)
	st.lcount = 0
}

func (st *SymbolTable) addVariable(globalScope bool, name string, ptype PascalType) error {
	var symbols *map[string]Symbol

	if globalScope {
		symbols = &st.gsymbols
	} else {
		symbols = &st.lsymbols
	}

	nameToSearch := strings.ToUpper(name)
	if _, ok := (*symbols)[nameToSearch]; ok {
		return fmt.Errorf("variable %s already declared", name)
	}

	(*symbols)[nameToSearch] = Symbol{
		Name:       name,
		SymbolType: Variable,
		PascalType: ptype,
		Index:      st.lcount,
		Global:     globalScope,
	}

	if !globalScope {
		st.lcount++
	}

	return nil
}
