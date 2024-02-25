package codegen_test

import (
	"testing"

	"github.com/alexgarzao/poj/codegen"
	"github.com/stretchr/testify/assert"
)

func TestStackType_CompleteBasicTest(t *testing.T) {
	assert := assert.New(t)

	s := codegen.NewStackType()
	assert.Equal(codegen.Undefined, s.Top())

	s.Push(codegen.String)
	assert.Equal(codegen.String, s.Top())

	s.Push(codegen.Integer)
	assert.Equal(codegen.Integer, s.Top())

	pt := s.Pop()
	assert.Equal(codegen.Integer, pt)

	assert.Equal(codegen.String, s.Top())
	pt = s.Pop()
	assert.Equal(codegen.String, pt)

	assert.Equal(codegen.Undefined, s.Top())
	pt = s.Pop()
	assert.Equal(codegen.Undefined, pt)
}
