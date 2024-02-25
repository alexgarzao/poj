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
	if ctx.GetProcedureID().GetText() == "writeln" {
		t.jasm.AddOpcode("getstatic", "java/lang/System.out", "java/io/PrintStream")
	}
}

func (t *TreeShapeListener) ExitString(ctx *parsing.StringContext) {
	str := ctx.GetText()
	t.jasm.AddOpcode("ldc", "\""+str+"\"")
	t.pst.Push(String)
}

func (t *TreeShapeListener) ExitProcedureStatement(ctx *parsing.ProcedureStatementContext) {
	if ctx.GetProcedureID().GetText() == "writeln" {
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

func (t *TreeShapeListener) ExitSimpleExpression(ctx *parsing.SimpleExpressionContext) {
	if ctx.GetT2() != nil {
		if ctx.GetOp() == ctx.Additiveoperator() {
			pt1 := t.pst.Pop()
			pt2 := t.pst.Pop()
			if pt1 != pt2 {
				t.jasm.AddOpcode("invalid types in add")
				return
			}

			if pt1 == String {
				// String types.
				t.jasm.StartInvokeDynamic(`makeConcatWithConstants(java/lang/String, java/lang/String)java/lang/String`)
				t.jasm.AddOpcode(`invokestatic java/lang/invoke/StringConcatFactory.makeConcatWithConstants(java/lang/invoke/MethodHandles$Lookup, java/lang/String, java/lang/invoke/MethodType, java/lang/String, [java/lang/Object)java/lang/invoke/CallSite`)
				t.jasm.AddOpcode(`[""]`)
				t.jasm.FinishInvokeDynamic()
				t.pst.Push(String)
			} else if pt1 == Integer {
				// Integer types.
				t.jasm.AddOpcode(`iadd`)
				t.pst.Push(Integer)
			} else {
				// Undefined types.
				t.jasm.AddOpcode("undefined type in add")
			}
		}
	}
}

func (t *TreeShapeListener) ExitUnsignedInteger(ctx *parsing.UnsignedIntegerContext) {
	t.jasm.AddOpcode("sipush", ctx.GetText())
	t.pst.Push(Integer)
}

func (t *TreeShapeListener) Code() string {
	return t.jasm.Code()
}
