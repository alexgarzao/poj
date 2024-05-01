package codegen

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	size := len(s.elements)
	if size == 0 {
		var undefined T
		return undefined, false
	}

	elem := s.elements[size-1]
	s.elements = s.elements[:size-1]

	return elem, true
}

func (s *Stack[T]) Top() (T, bool) {
	size := len(s.elements)
	if size == 0 {
		var undefined T
		return undefined, false
	}

	elem := s.elements[size-1]

	return elem, true
}
