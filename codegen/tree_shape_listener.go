package codegen

import (
	"github.com/alexgarzao/poj/parsing"
)

type TreeShapeListener struct {
	*parsing.BasePascalListener
	filename string
	jasm     *JASM
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
	t.jasm.FinishParameter()
}

func (t *TreeShapeListener) EnterBlock(ctx *parsing.BlockContext) {
	t.jasm.StartBlock()
}

func (t *TreeShapeListener) ExitBlock(ctx *parsing.BlockContext) {
	t.jasm.FinishBlock()
}

func (t *TreeShapeListener) ExitNotOp(ctx *parsing.NotOpContext) {
	op := ctx.GetOp().GetText()
	t.jasm.AddUnaryOperatorOpcode(op)
}

func (t *TreeShapeListener) ExitBoolOp(ctx *parsing.BoolOpContext) {
	op := ctx.GetOp().GetText()
	t.jasm.AddOperatorOpcode(op)
}

func (t *TreeShapeListener) ExitMulDivOp(ctx *parsing.MulDivOpContext) {
	op := ctx.GetOp().GetText()
	t.jasm.AddOperatorOpcode(op)
}

func (t *TreeShapeListener) ExitAddOp(ctx *parsing.AddOpContext) {
	op := ctx.GetOp().GetText()
	t.jasm.AddOperatorOpcode(op)
}

func (t *TreeShapeListener) ExitRelOp(ctx *parsing.RelOpContext) {
	op := ctx.GetOp().GetText()
	t.jasm.AddOperatorOpcode(op)
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

func (t *TreeShapeListener) Code() string {
	return t.jasm.Code()
}
