package codegen

type PascalType uint8

const (
	Undefined PascalType = iota
	String
	Integer
	Boolean
)

type StackType struct {
	types []PascalType
}

func NewStackType() *StackType {
	return &StackType{}
}

func (s *StackType) Push(ptype PascalType) {
	s.types = append(s.types, ptype)
}

func (s *StackType) Pop() PascalType {
	size := len(s.types)
	if size == 0 {
		return Undefined
	}

	ptype := s.types[size-1]
	s.types = s.types[:size-1]
	return ptype
}

func (s *StackType) Top() PascalType {
	size := len(s.types)
	if size == 0 {
		return Undefined
	}

	ptype := s.types[size-1]
	return ptype
}
