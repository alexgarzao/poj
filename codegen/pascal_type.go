package codegen

type PascalType uint8

const (
	Undefined PascalType = iota
	String
	Integer
	Boolean
)

func (pt PascalType) String() string {
	switch pt {
	case String:
		return "string"
	case Integer:
		return "integer"
	case Boolean:
		return "boolean"
	default:
		return "undefined"
	}
}

func ToPascalType(s string) PascalType {
	switch s {
	case "string":
		return String
	case "integer":
		return Integer
	case "boolean":
		return Boolean
	default:
		return Undefined
	}
}
