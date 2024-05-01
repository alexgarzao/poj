package codegen_test

import (
	"testing"

	"github.com/alexgarzao/poj/codegen"
	"github.com/stretchr/testify/assert"
)

func TestStackCompleteBasicTest(t *testing.T) {
	assert := assert.New(t)

	s := codegen.NewStack[codegen.PascalType]()
	_, ok := s.Top()
	assert.Equal(false, ok)

	s.Push(codegen.String)
	pt, ok := s.Top()
	assert.Equal(true, ok)
	assert.Equal(codegen.String, pt)

	s.Push(codegen.Integer)
	pt, ok = s.Top()
	assert.Equal(true, ok)
	assert.Equal(codegen.Integer, pt)

	pt, ok = s.Pop()
	assert.Equal(true, ok)
	assert.Equal(codegen.Integer, pt)

	pt, ok = s.Top()
	assert.Equal(true, ok)
	assert.Equal(codegen.String, pt)

	pt, ok = s.Pop()
	assert.Equal(true, ok)
	assert.Equal(codegen.String, pt)

	_, ok = s.Top()
	assert.Equal(false, ok)

	_, ok = s.Pop()
	assert.Equal(false, ok)
}
