package codegen

import (
	"fmt"

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
	if err := t.jasm.StartParameter(); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitActualParameter(ctx *parsing.ActualParameterContext) {
	if err := t.jasm.FinishParameter(); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) EnterCompoundStatement(ctx *parsing.CompoundStatementContext) {
	if err := t.jasm.StartMainBlock(); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitNotOp(ctx *parsing.NotOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddUnaryOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitBoolOp(ctx *parsing.BoolOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddBooleanOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitMulDivOp(ctx *parsing.MulDivOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddMulDivOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitAddSubOp(ctx *parsing.AddSubOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddAddSubOperatorOpcode(op); err != nil {
		t.parserErrors.Add(err)
	}
}

func (t *TreeShapeListener) ExitRelOp(ctx *parsing.RelOpContext) {
	op := ctx.GetOp().GetText()
	if err := t.jasm.AddRelationalOperatorOpcode(op); err != nil {
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

func (t *TreeShapeListener) EnterProcedureDeclaration(ctx *parsing.ProcedureDeclarationContext) {
	procName := ctx.GetName().GetText()

	paramTypes := []string{}
	params := ctx.GetParamList()
	if params != nil {
		for _, param := range params.GetParams() {
			pascalType := param.GetParamType()
			for range param.GetParamNames().GetIds() {
				paramTypes = append(paramTypes, pascalType.GetText())
			}
		}
	}

	if err := t.jasm.NewProcedure(procName, paramTypes); err != nil {
		t.parserErrors.Add(err)
	}

	t.jasm.StartProcedureDeclaration(procName, paramTypes)
}

func (t *TreeShapeListener) ExitProcedureDeclaration(ctx *parsing.ProcedureDeclarationContext) {
	t.jasm.FinishProcedureDeclaration()
}

func (t *TreeShapeListener) EnterFunctionDeclaration(ctx *parsing.FunctionDeclarationContext) {
	funcName := ctx.GetName().GetText()

	paramTypes := []string{}
	params := ctx.GetParamList()
	if params != nil {
		for _, param := range params.GetParams() {
			// pascalType := param.GetParamType()
			for range param.GetParamNames().GetIds() {
				paramTypes = append(paramTypes, param.GetParamType().GetText())
			}
		}
	}

	returnType := ctx.GetReturnType().GetText()

	if err := t.jasm.NewFunction(funcName, paramTypes, returnType); err != nil {
		t.parserErrors.Add(err)
	}

	t.jasm.StartFunctionDeclaration(funcName, paramTypes, returnType)
}

func (t *TreeShapeListener) ExitFunctionDeclaration(ctx *parsing.FunctionDeclarationContext) {
	t.jasm.FinishFunctionDeclaration()
}

func (t *TreeShapeListener) EnterFormalParameterSection(ctx *parsing.FormalParameterSectionContext) {
	params := ctx.GetParamNames()
	pascalType := ctx.GetParamType().GetText()
	for _, id := range params.GetIds() {
		if err := t.jasm.NewVariable(id.GetText(), pascalType); err != nil {
			t.parserErrors.Add(err)
		}
	}
}

func (t *TreeShapeListener) EnterFunctionDesignator(ctx *parsing.FunctionDesignatorContext) {
	t.jasm.StartProcedureStatement(ctx.GetFunctionID().GetText())
}

func (t *TreeShapeListener) ExitFunctionDesignator(ctx *parsing.FunctionDesignatorContext) {
	funcName := ctx.GetFunctionID().GetText()
	if err := t.jasm.CallFunction(funcName); err != nil {
		t.parserErrors.Add(err)
	}

	_, exists := t.jasm.ProcedureStatementContext.Pop()
	if !exists {
		t.parserErrors.Add(fmt.Errorf("during pop procedure statement context"))
	}
}

func (t *TreeShapeListener) Code() string {
	return t.jasm.Code()
}
