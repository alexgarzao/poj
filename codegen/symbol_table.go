package codegen

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

func (st *SymbolTable) AddVariable(name string, ptype PascalType) {
	st.count++
	st.symbols[name] = Symbol{
		SymbolType: Variable,
		PascalType: ptype,
		Index:      st.count,
	}
}

func (st *SymbolTable) Get(name string) (bool, Symbol) {
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
