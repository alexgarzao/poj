package codegen

import "fmt"

type SymbolType uint8

const (
	UndefinedSymbolType SymbolType = iota
	Variable
	Procedure
	Function
)

type Symbol struct {
	Name       string
	SymbolType SymbolType
	PascalType PascalType
	Index      int
	ParamTypes []string
	Global     bool
}

func (s Symbol) GenLoadOpcode(className string) string {

	if s.SymbolType != Variable && s.Index != fixedFunctionVarIndex {
		return "invalid symbol type"
	}

	if s.Global {
		// global: getstatic GlobalAndLocalVars.globalVar I
		return fmt.Sprintf("getstatic %s.%s %s", className, s.Name, s.PascalType.JasmType())
	}

	// local:  iload 123
	switch s.PascalType {
	case Integer:
		return fmt.Sprintf("iload %d", s.Index)
	case Boolean:
		return fmt.Sprintf("iload %d", s.Index)
	case String:
		return fmt.Sprintf("aload %d", s.Index)
	default:
		return "undefined type"
	}
}

func (s Symbol) GenStoreOpcode(className string) string {

	if s.SymbolType != Variable && s.Index != fixedFunctionVarIndex {
		return "invalid symbol type"
	}

	if s.Global {
		// global: putstatic GlobalAndLocalVars.globalVar I
		return fmt.Sprintf("putstatic %s.%s %s", className, s.Name, s.PascalType.JasmType())
	}

	// local:  istore 123
	switch s.PascalType {
	case Integer:
		return fmt.Sprintf("istore %d", s.Index)
	case Boolean:
		return fmt.Sprintf("istore %d", s.Index)
	case String:
		return fmt.Sprintf("astore %d", s.Index)
	default:
		return "undefined type"
	}
}
