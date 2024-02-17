package codegen_test

import (
	"testing"

	"github.com/alexgarzao/poj/codegen"
	"github.com/stretchr/testify/assert"
)

func TestJASM_CompleteBasicTest(t *testing.T) {
	assert := assert.New(t)

	j := codegen.NewJASM()
	assert.Empty(j.Code())

	j.StartMainClass("XPTO")
	expected := "// Code generated by POJ 0.1\npublic class XPTO {\n"
	assert.Equal(expected, j.Code())

	j.StartMain()
	expected += "\tpublic static main([java/lang/String)V {\n"
	assert.Equal(expected, j.Code())

	j.AddOpcode("op1", "param1", "param2", "param3")
	expected += "\t\top1 param1 param2 param3\n"
	assert.Equal(expected, j.Code())

	j.FinishMain()
	expected += "\t\treturn\n"
	expected += "\t}\n"
	assert.Equal(expected, j.Code())

	j.FinishMainClass()
	expected += "}\n"
	assert.Equal(expected, j.Code())
}

func TestJASM_HelloWorld(t *testing.T) {
	assert := assert.New(t)

	j := codegen.NewJASM()
	j.StartMainClass("HelloWorld")
	j.StartMain()
	j.AddOpcode("getstatic", "java/lang/System.out", "java/io/PrintStream")
	j.AddOpcode("ldc", "\"Hello, World\"")
	j.AddOpcode("invokevirtual", "java/io/PrintStream.println(java/lang/String)V")
	j.FinishMain()
	j.FinishMainClass()
	expected := `// Code generated by POJ 0.1
public class HelloWorld {
	public static main([java/lang/String)V {
		getstatic java/lang/System.out java/io/PrintStream
		ldc "Hello, World"
		invokevirtual java/io/PrintStream.println(java/lang/String)V
		return
	}
}
`
	assert.Equal(expected, j.Code())
}