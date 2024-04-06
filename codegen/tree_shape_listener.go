package codegen

import (
	"github.com/alexgarzao/poj/parsing"
)

type TreeShapeListener struct {
	*parsing.BasePascalListener
	filename     string
	jasm         *JASM
	parserErrors *CustomErrorListener
}

func NewTreeShapeListener(filename string, parserErrors *CustomErrorListener) *TreeShapeListener {
	return &TreeShapeListener{
		filename:     filename,
		jasm:         NewJASM(),
		parserErrors: parserErrors,
	}
}

func (t *TreeShapeListener) EnterProgram(ctx *parsing.ProgramContext) {
	t.jasm.StartMainClass(t.filename)
}

func (t *TreeShapeListener) ExitProgram(ctx *parsing.ProgramContext) {
	t.jasm.FinishMainClass()
}

func (t *TreeShapeListener) EnterProcedureStatement(ctx *parsing.ProcedureStatementContext) {
	t.jasm.StartProcedureStatement(ctx.GetProcedureID().GetText())
}

func (t *TreeShapeListener) ExitProcedureStatement(ctx *parsing.ProcedureStatementContext) {
	t.jasm.FinishProcedureStatement()
}

func (t *TreeShapeListener) ExitString(ctx *parsing.StringContext) {
	str := ctx.GetText()
	t.jasm.NewConstantString("\"" + str + "\"")
}

func (t *TreeShapeListener) EnterActualParameter(ctx *parsing.ActualParameterContext) {
	t.jasm.StartParameter()
}

func (t *TreeShapeListener) ExitActualParameter(ctx *parsing.ActualParameterContext) {
	if err := t.jasm.FinishParameter(); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) EnterBlock(ctx *parsing.BlockContext) {
	t.jasm.StartBlock()
}

func (t *TreeShapeListener) ExitBlock(ctx *parsing.BlockContext) {
	t.jasm.FinishBlock()
}

func (t *TreeShapeListener) ExitNotOp(ctx *parsing.NotOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddUnaryOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitBoolOp(ctx *parsing.BoolOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitMulDivOp(ctx *parsing.MulDivOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitAddOp(ctx *parsing.AddOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitRelOp(ctx *parsing.RelOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) EnterIfStatement(ctx *parsing.IfStatementContext) {
	t.jasm.StartIfStatement()
}

func (t *TreeShapeListener) ExitIfStatement(ctx *parsing.IfStatementContext) {
	t.jasm.FinishIfStatement()
}

func (t *TreeShapeListener) EnterThenStatement(ctx *parsing.ThenStatementContext) {
	t.jasm.EnterThenStatement()
}

func (t *TreeShapeListener) ExitThenStatement(ctx *parsing.ThenStatementContext) {
	t.jasm.FinishThenStatement()
}

func (t *TreeShapeListener) ExitUnsignedInteger(ctx *parsing.UnsignedIntegerContext) {
	t.jasm.NewConstantInteger(ctx.GetText())
}

func (t *TreeShapeListener) ExitVariableDeclaration(ctx *parsing.VariableDeclarationContext) {
	varNames := ctx.GetVarNames()
	pascalType := ctx.GetPascalType().GetText()
	for _, id := range varNames.GetIds() {
		if err := t.jasm.NewVariable(id.GetText(), pascalType); err != nil {
			t.parserErrors.Add(err)
		}
	}
}

func (t *TreeShapeListener) ExitAssignmentStatement(ctx *parsing.AssignmentStatementContext) {
	varName := ctx.GetVarName().GetText()
	if err := t.jasm.FinishAssignmentStatement(varName); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitFactorVariable(ctx *parsing.FactorVariableContext) {
	varName := ctx.GetId().GetText()
	if err := t.jasm.LoadVarContent(varName); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) EnterRepeatStatement(ctx *parsing.RepeatStatementContext) {
	t.jasm.StartRepeatStatement()
}

func (t *TreeShapeListener) ExitRepeatStatement(ctx *parsing.RepeatStatementContext) {
	t.jasm.FinishRepeatStatement()
}

func (t *TreeShapeListener) EnterWhileStatement(ctx *parsing.WhileStatementContext) {
	t.jasm.StartWhileStatement()
}

func (t *TreeShapeListener) ExitWhileStatement(ctx *parsing.WhileStatementContext) {
	t.jasm.FinishWhileStatement()
}

func (t *TreeShapeListener) EnterWhileBlock(ctx *parsing.WhileBlockContext) {
	t.jasm.StartWhileBlock()
}

func (t *TreeShapeListener) ExitForInit(ctx *parsing.ForInitContext) {
	t.jasm.FinishForInit(ctx.GetVarName().GetText())
}

func (t *TreeShapeListener) ExitForUntil(ctx *parsing.ForUntilContext) {
	t.jasm.FinishForUntil(ctx.GetStep().GetText())
}

func (t *TreeShapeListener) ExitForStatement(ctx *parsing.ForStatementContext) {
	t.jasm.FinishForStatement()
}

func (t *TreeShapeListener) Code() string {
	return t.jasm.Code()
}
