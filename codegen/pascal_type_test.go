package codegen_test

import (
	"testing"

	"github.com/alexgarzao/poj/codegen"
	"github.com/stretchr/testify/assert"
)

func TestPascalType_ToPascalType(t *testing.T) {
	pt := codegen.ToPascalType("string")
	assert.Equal(t, codegen.String, pt)

	pt = codegen.ToPascalType("integer")
	assert.Equal(t, codegen.Integer, pt)

	pt = codegen.ToPascalType("boolean")
	assert.Equal(t, codegen.Boolean, pt)

	pt = codegen.ToPascalType("integeer")
	assert.Equal(t, codegen.Undefined, pt)
}

func TestPascalType_String(t *testing.T) {
	pt := codegen.String
	assert.Equal(t, "string", pt.String())

	pt = codegen.Integer
	assert.Equal(t, "integer", pt.String())

	pt = codegen.Boolean
	assert.Equal(t, "boolean", pt.String())

	pt = codegen.Undefined
	assert.Equal(t, "undefined", pt.String())
}
