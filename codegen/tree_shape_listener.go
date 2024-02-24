package codegen

import (
	"github.com/alexgarzao/poj/parsing"
)

type TreeShapeListener struct {
	*parsing.BasePascalListener
	filename                string
	jasm                    *JASM
	procedureDefinitionName string
}

func NewTreeShapeListener(filename string) *TreeShapeListener {
	return &TreeShapeListener{
		filename: filename,
		jasm:     NewJASM(),
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
	t.jasm.AddOpcode("ldc", "\""+ctx.GetText()+"\"")
}

func (t *TreeShapeListener) ExitProcedureStatement(ctx *parsing.ProcedureStatementContext) {
	if ctx.GetProcedureID().GetText() == "writeln" {
		t.jasm.AddOpcode("invokevirtual", "java/io/PrintStream.println(java/lang/String)V")
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
			// Only works for string types.
			t.jasm.StartInvokeDynamic(`makeConcatWithConstants(java/lang/String, java/lang/String)java/lang/String`)
			t.jasm.AddOpcode(`invokestatic java/lang/invoke/StringConcatFactory.makeConcatWithConstants(java/lang/invoke/MethodHandles$Lookup, java/lang/String, java/lang/invoke/MethodType, java/lang/String, [java/lang/Object)java/lang/invoke/CallSite`)
			t.jasm.AddOpcode(`[""]`)
			t.jasm.FinishInvokeDynamic()
		}
	}
}

func (t *TreeShapeListener) Code() string {
	return t.jasm.Code()
}
