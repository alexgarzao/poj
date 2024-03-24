package codegen

import (
	"github.com/alexgarzao/poj/parsing"
)

type TreeShapeListener struct {
	*parsing.BasePascalListener
	filename                string
	jasm                    *JASM
	procedureDefinitionName string
	pst                     *StackType
}

func NewTreeShapeListener(filename string) *TreeShapeListener {
	return &TreeShapeListener{
		filename: filename,
		jasm:     NewJASM(),
		pst:      NewStackType(),
	}
}

func (t *TreeShapeListener) EnterProgram(ctx *parsing.ProgramContext) {
	t.jasm.StartMainClass(t.filename)
}

func (t *TreeShapeListener) ExitProgram(ctx *parsing.ProgramContext) {
	t.jasm.FinishMainClass()
}

func (t *TreeShapeListener) EnterProcedureStatement(ctx *parsing.ProcedureStatementContext) {
	t.procedureDefinitionName = ctx.GetProcedureID().GetText()
}

func (t *TreeShapeListener) ExitProcedureStatement(ctx *parsing.ProcedureStatementContext) {
	t.procedureDefinitionName = ""
}

func (t *TreeShapeListener) ExitString(ctx *parsing.StringContext) {
	str := ctx.GetText()
	t.jasm.AddOpcode("ldc", "\""+str+"\"")
	t.pst.Push(String)
}

func (t *TreeShapeListener) EnterActualParameter(ctx *parsing.ActualParameterContext) {
	if t.procedureDefinitionName == "writeln" {
		t.jasm.AddOpcode("getstatic", "java/lang/System.out", "java/io/PrintStream")
	}
}

func (t *TreeShapeListener) ExitActualParameter(ctx *parsing.ActualParameterContext) {
	if t.procedureDefinitionName == "writeln" {
		pt := t.pst.Pop()
		if pt == String {
			t.jasm.AddOpcode("invokevirtual", "java/io/PrintStream.println(java/lang/String)V")
		} else if pt == Integer {
			t.jasm.AddOpcode("invokevirtual", "java/io/PrintStream.println(I)V")
		} else {
			t.jasm.AddOpcode("undefined type in writeln")
		}
	}
}

func (t *TreeShapeListener) ExitProcedureDeclaration(ctx *parsing.ProcedureDeclarationContext) {
	t.procedureDefinitionName = ctx.GetName().GetText()
}

func (t *TreeShapeListener) EnterBlock(ctx *parsing.BlockContext) {
	if t.procedureDefinitionName == "" {
		// Main block.
		t.jasm.StartMain()
	}
}

func (t *TreeShapeListener) ExitBlock(ctx *parsing.BlockContext) {
	if t.procedureDefinitionName == "" {
		// Main block.
		t.jasm.FinishMain()
	}

	t.procedureDefinitionName = ""
}

func (t *TreeShapeListener) ExitMulDivOp(ctx *parsing.MulDivOpContext) {
	pt1 := t.pst.Pop()
	pt2 := t.pst.Pop()
	if pt1 != pt2 {
		t.jasm.AddOpcode("invalid types")
		return
	}

	op := ctx.GetOp().GetText()
	switch {
	case op == "*":
		switch pt1 {
		case Integer:
			t.GenMulIntegers()
		default:
			t.jasm.AddOpcode("invalid type in mul")
		}
	case op == "/":
		switch pt1 {
		case Integer:
			t.GenDivIntegers()
		default:
			t.jasm.AddOpcode("invalid type in div")
		}
	}
}

func (t *TreeShapeListener) ExitAddOp(ctx *parsing.AddOpContext) {
	pt1 := t.pst.Pop()
	pt2 := t.pst.Pop()
	if pt1 != pt2 {
		t.jasm.AddOpcode("invalid types")
		return
	}

	op := ctx.GetOp().GetText()
	switch {
	case op == "+":
		switch pt1 {
		case String:
			t.GenAddStrings()
		case Integer:
			t.GenAddIntegers()
		default:
			t.jasm.AddOpcode("invalid type in add")
		}
	case op == "-":
		switch pt1 {
		case Integer:
			t.GenSubIntegers()
		default:
			t.jasm.AddOpcode("invalid type in sub")
		}
	}
}

func (t *TreeShapeListener) GenAddStrings() {
	t.jasm.StartInvokeDynamic(`makeConcatWithConstants(java/lang/String, java/lang/String)java/lang/String`)
	t.jasm.AddOpcode(`invokestatic java/lang/invoke/StringConcatFactory.makeConcatWithConstants(java/lang/invoke/MethodHandles$Lookup, java/lang/String, java/lang/invoke/MethodType, java/lang/String, [java/lang/Object)java/lang/invoke/CallSite`)
	t.jasm.AddOpcode(`[""]`)
	t.jasm.FinishInvokeDynamic()
	t.pst.Push(String)
}

func (t *TreeShapeListener) GenAddIntegers() {
	t.jasm.AddOpcode("iadd")
	t.pst.Push(Integer)
}

func (t *TreeShapeListener) GenSubIntegers() {
	t.jasm.AddOpcode("isub")
	t.pst.Push(Integer)
}

func (t *TreeShapeListener) GenMulIntegers() {
	t.jasm.AddOpcode("imul")
	t.pst.Push(Integer)
}

func (t *TreeShapeListener) GenDivIntegers() {
	t.jasm.AddOpcode("idiv")
	t.pst.Push(Integer)
}

func (t *TreeShapeListener) ExitUnsignedInteger(ctx *parsing.UnsignedIntegerContext) {
	t.jasm.AddOpcode("sipush", ctx.GetText())
	t.pst.Push(Integer)
}

func (t *TreeShapeListener) Code() string {
	return t.jasm.Code()
}
