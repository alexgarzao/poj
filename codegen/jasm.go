package codegen

import (
	"fmt"
	"strings"
)

type JASM struct {
	code *Code
	pst  *StackType
	st   *SymbolTable

	className                string
	procedureDeclarationName string
	labelID                  int
	endIfLabel               string
	elseLabel                string
	repeatLabel              string
	whileTestLabel           string
	nextStatementLabel       string
	forTestLabel             string
	forVariable              string
	forStep                  string

	// Stack contexts.
	ProcedureStatementContext Stack[string]
}

func NewJASM() *JASM {
	return &JASM{
		code:                     NewCode(),
		pst:                      NewStackType(),
		st:                       NewSymbolTable(),
		procedureDeclarationName: "main",
	}
}

func (j *JASM) StartMainClass(name string) {
	j.className = name
	j.addLine("// Code generated by POJ 0.1")
	j.addLine(fmt.Sprintf("public class %s {", name))
	j.incTab()
}

func (j *JASM) FinishMainClass() {
	j.finishMain()
	j.decTab()
	j.addLine("}")
}

func (j *JASM) StartProcedureDeclaration(name string, paramTypes []string) {
	j.procedureDeclarationName = name
	j.addLine(fmt.Sprintf("static %s(%s)V {", name, j.genSignature(paramTypes)))
	j.incTab()
}

func (j *JASM) StartFunctionDeclaration(name string, paramTypes []string, returnType string) {
	j.procedureDeclarationName = name
	j.addLine(fmt.Sprintf("static %s(%s)%s {", name, j.genSignature(paramTypes), j.genSignature([]string{returnType})))
	j.incTab()
}

func (j *JASM) genSignature(paramTypes []string) string {
	javaParams := make([]string, len(paramTypes))
	for i, p := range paramTypes {
		if p == "string" {
			javaParams[i] = "java/lang/String"
		} else if p == "integer" {
			javaParams[i] = "I"
		} else {
			javaParams[i] = "V"
		}
	}

	return strings.Join(javaParams, ", ")
}

func (j *JASM) FinishProcedureDeclaration() {
	j.addLine("return")
	j.decTab()
	j.addLine("}")
	j.procedureDeclarationName = "main"
}

func (j *JASM) FinishFunctionDeclaration() error {
	symbol, ok := j.st.Get(j.procedureDeclarationName)
	if !ok {
		return fmt.Errorf("bug: function %s not found", j.procedureDeclarationName)
	}

	switch symbol.PascalType {
	case String:
		j.addOpcode("aload", fmt.Sprintf("%d", symbol.Index))
		j.addOpcode("areturn")
	case Integer:
		j.addOpcode("iload", fmt.Sprintf("%d", symbol.Index))
		j.addOpcode("ireturn")
	default:
		return fmt.Errorf("invalid function type in return")
	}

	j.decTab()
	j.addLine("}")
	j.procedureDeclarationName = "main"

	return nil
}

func (j *JASM) StartProcedureStatement(name string) {
	j.ProcedureStatementContext.Push(name)
}

func (j *JASM) FinishProcedureStatement() error {
	procedureStatementName, exists := j.ProcedureStatementContext.Top()
	if !exists {
		return fmt.Errorf("during getting procedure statement context")
	}

	if procedureStatementName == "writeln" {
		j.addStaticPrintStream()
		j.addInvokeVirtual("java/io/PrintStream.println()V")
	} else if procedureStatementName != "write" {
		proc, ok := j.st.Get(procedureStatementName)
		if !ok {
			return fmt.Errorf("procedure %s not found", procedureStatementName)
		}

		j.addInvokeStatic(procedureStatementName, j.genSignature(proc.ParamTypes), j.genSignature([]string{proc.PascalType.String()}))
	}

	_, exists = j.ProcedureStatementContext.Pop()
	if !exists {
		return fmt.Errorf("during pop procedure statement context")
	}

	return nil
}

func (j *JASM) StartParameter() error {
	procedureStatementName, exists := j.ProcedureStatementContext.Top()
	if !exists {
		return fmt.Errorf("during getting procedure statement context")
	}

	if procedureStatementName == "write" || procedureStatementName == "writeln" {
		j.addStaticPrintStream()
	}

	return nil
}

func (j *JASM) FinishParameter() error {
	procedureStatementName, exists := j.ProcedureStatementContext.Top()
	if !exists {
		return fmt.Errorf("during getting procedure statement context")
	}

	if procedureStatementName == "write" || procedureStatementName == "writeln" {
		if err := j.addInvokeVirtualPrintWithType(); err != nil {
			return err
		}
	}

	return nil
}

func (j *JASM) StartMainBlock() {
	if j.procedureDeclarationName == "main" {
		// Main block.
		j.startMain()
		j.procedureDeclarationName = ""
	}
}

func (j *JASM) NewConstantString(constant string) {
	j.addLdcStringOpcode(constant)
}

func (j *JASM) NewConstantInteger(constant string) {
	j.addSiPushOpcode(constant)
}

func (j *JASM) StartIfStatement() {
	j.elseLabel = j.newLabel()
	j.endIfLabel = j.newLabel()
}

func (j *JASM) EnterThenStatement() {
	j.addIfEqOpcode(j.elseLabel)
}

func (j *JASM) FinishThenStatement() {
	j.addGotoOpcode(j.endIfLabel)
	j.addLabel(j.elseLabel)
}

func (j *JASM) FinishIfStatement() {
	j.addLabel(j.endIfLabel)
}

func (j *JASM) StartRepeatStatement() {
	j.repeatLabel = j.newLabel()
	j.addLabel(j.repeatLabel)
}

func (j *JASM) FinishRepeatStatement() {
	j.addIfEqOpcode(j.repeatLabel)
}

func (j *JASM) StartWhileStatement() {
	j.whileTestLabel = j.newLabel()
	j.nextStatementLabel = j.newLabel()
	j.addLabel(j.whileTestLabel)
}

func (j *JASM) FinishWhileStatement() {
	j.addGotoOpcode(j.whileTestLabel)
	j.addLabel(j.nextStatementLabel)
}

func (j *JASM) StartWhileBlock() {
	j.addIfEqOpcode(j.nextStatementLabel)
}

func (j *JASM) FinishForInit(varName string) error {
	if err := j.FinishAssignmentStatement(varName); err != nil {
		return err
	}

	j.forVariable = varName
	j.forTestLabel = j.newLabel()
	j.nextStatementLabel = j.newLabel()
	j.addLabel(j.forTestLabel)

	if err := j.LoadVarContent(j.forVariable); err != nil {
		return err
	}

	return nil
}

func (j *JASM) FinishForUntil(step string) {
	j.forStep = step

	if step == "to" {
		j.addOpcode("if_icmpgt", j.nextStatementLabel)
	} else {
		j.addOpcode("if_icmplt", j.nextStatementLabel)
	}
}

func (j *JASM) FinishForStatement() error {
	if err := j.LoadVarContent(j.forVariable); err != nil {
		return err
	}

	j.addSiPushOpcode("1")

	if j.forStep == "to" {
		j.addIAddOpcode()
	} else {
		j.addISubOpcode()
	}

	if err := j.FinishAssignmentStatement(j.forVariable); err != nil {
		return err
	}

	j.addGotoOpcode(j.forTestLabel)
	j.addLabel(j.nextStatementLabel)

	return nil
}

func (j *JASM) AddBooleanOperatorOpcode(op string) error {
	pt1 := j.pst.Pop()
	pt2 := j.pst.Pop()
	if pt1 != Boolean || pt2 != Boolean {
		return fmt.Errorf("invalid types in %s operator: %s and %s", op, pt1, pt2)
	}

	switch {
	case op == "and":
		j.addIAndOpcode()
	case op == "or":
		j.addIOrOpcode()
	default:
		return fmt.Errorf("invalid boolean operator: %s", op)
	}

	return nil
}

func (j *JASM) AddMulDivOperatorOpcode(op string) error {
	pt1 := j.pst.Pop()
	pt2 := j.pst.Pop()
	if pt1 != Integer || pt2 != Integer {
		return fmt.Errorf("invalid types in %s operator: %s and %s", op, pt1, pt2)
	}

	switch {
	case op == "*":
		j.addIMulOpcode()
	case op == "/":
		j.addIDivOpcode()
	}

	j.pst.Push(Integer)

	return nil
}

func (j *JASM) AddAddSubOperatorOpcode(op string) error {
	pt1 := j.pst.Pop()
	pt2 := j.pst.Pop()
	if pt1 != pt2 {
		return fmt.Errorf("invalid types in %s operator: %s and %s", op, pt1, pt2)
	}

	switch {
	case op == "+":
		switch pt1 {
		case String:
			j.startInvokeDynamic(`makeConcatWithConstants(java/lang/String, java/lang/String)java/lang/String`)
			j.addOpcode(`invokestatic java/lang/invoke/StringConcatFactory.makeConcatWithConstants(java/lang/invoke/MethodHandles$Lookup, java/lang/String, java/lang/invoke/MethodType, java/lang/String, [java/lang/Object)java/lang/invoke/CallSite`)
			j.addOpcode(`[""]`)
			j.finishInvokeDynamic()
			j.pst.Push(String)
		case Integer:
			j.addIAddOpcode()
			j.pst.Push(Integer)
		default:
			return fmt.Errorf("invalid type in %s operator: %s", op, pt1)
		}
	case op == "-":
		switch pt1 {
		case Integer:
			j.addISubOpcode()
			j.pst.Push(Integer)
		default:
			return fmt.Errorf("invalid type in %s operator: %s", op, pt1)
		}
	}

	return nil
}

func (j *JASM) AddRelationalOperatorOpcode(op string) error {
	pt1 := j.pst.Pop()
	pt2 := j.pst.Pop()
	if pt1 != pt2 {
		return fmt.Errorf("invalid types in %s operator: %s and %s", op, pt1, pt2)
	}

	if pt1 != Integer && pt1 != String {
		return fmt.Errorf("invalid type in %s operator: %s", op, pt1)
	}

	if pt1 == Integer {
		jmps := map[string]string{
			">":  "if_icmple",
			"<":  "if_icmpge",
			">=": "if_icmplt",
			"<=": "if_icmpgt",
			"=":  "if_icmpne",
			"<>": "if_icmpeq",
		}

		j.genBooleanOperatorTpl(jmps[op])
	} else if pt1 == String {
		jmps := map[string]string{
			">":  "iflt",
			"<":  "ifgt",
			">=": "iflt",
			"<=": "ifgt",
			"=":  "ifne",
			"<>": "ifeq",
		}

		j.addInvokeVirtual("java/lang/String.compareTo(java/lang/String)I")
		j.genBooleanOperatorTpl(jmps[op])
	}

	return nil
}

func (j *JASM) AddUnaryOperatorOpcode(op string) error {
	pt1 := j.pst.Pop()
	if pt1 != Boolean {
		return fmt.Errorf("invalid type in unary operator: %s", pt1)
	}

	switch op {
	case "not":
		lfalse := j.newLabel()
		lnext := j.newLabel()
		j.addIfNeOpcode(lfalse)
		j.addPushTrueOpcode()
		j.addGotoOpcode(lnext)
		j.addLabel(lfalse)
		j.addPushFalseOpcode()
		j.addLabel(lnext)
		j.pst.Push(Boolean)
	default:
		return fmt.Errorf("invalid unary operator: %s", op)
	}

	return nil
}

func (j *JASM) NewVariable(name, pst string) error {
	pt := ToPascalType(pst)
	if err := j.st.AddVariable(name, pt); err != nil {
		return err
	}

	return nil
}

func (j *JASM) FinishAssignmentStatement(varName string) error {
	symbol, ok := j.st.Get(varName)
	if !ok {
		return fmt.Errorf("variable %s not found", varName)
	}

	switch symbol.PascalType {
	case String:
		j.addOpcode("astore", fmt.Sprintf("%d", symbol.Index))
		j.pst.Pop()
	case Integer:
		j.addOpcode("istore", fmt.Sprintf("%d", symbol.Index))
		j.pst.Pop()
	default:
		return fmt.Errorf("invalid type in assignment")
	}

	return nil
}

func (j *JASM) LoadVarContent(varName string) error {
	symbol, ok := j.st.Get(varName)
	if !ok {
		return fmt.Errorf("variable %s not found", varName)
	}

	switch symbol.PascalType {
	case String:
		j.addOpcode("aload", fmt.Sprintf("%d", symbol.Index))
		j.pst.Push(String)
	case Integer:
		j.addOpcode("iload", fmt.Sprintf("%d", symbol.Index))
		j.pst.Push(Integer)
	default:
		return fmt.Errorf("invalid type %s in load var content", symbol.PascalType)
	}

	return nil
}

func (j *JASM) NewProcedure(name string, paramTypes []string) error {
	if err := j.st.AddProcedure(name, paramTypes); err != nil {
		return err
	}

	return nil
}

func (j *JASM) NewFunction(name string, paramTypes []string, returnType string) error {
	if err := j.st.AddFunction(name, paramTypes, returnType); err != nil {
		return err
	}

	return nil
}

func (j *JASM) CallFunction(name string) error {

	funcSymbol, ok := j.st.Get(name)
	if !ok {
		return fmt.Errorf("function %s not found", name)
	}

	j.addInvokeStatic(name, j.genSignature(funcSymbol.ParamTypes), j.genSignature([]string{funcSymbol.PascalType.String()}))
	j.pst.Push(funcSymbol.PascalType)

	return nil
}

func (j *JASM) Code() string {
	return j.code.Code()
}

func (j *JASM) startMain() {
	j.addLine("public static main([java/lang/String)V {")
	j.incTab()
}

func (j *JASM) addOpcode(opcode string, parameters ...string) {
	params := strings.Join(parameters, " ")

	j.addLine(fmt.Sprintf("%s %s", opcode, params))
}

func (j *JASM) addLdcStringOpcode(text string) {
	j.addOpcode("ldc", text)
	j.pst.Push(String)
}

func (j *JASM) addSiPushOpcode(number string) {
	j.addOpcode("sipush", number)
	j.pst.Push(Integer)
}

func (j *JASM) addPushTrueOpcode() {
	j.addOpcode("iconst 1")
}

func (j *JASM) addPushFalseOpcode() {
	j.addOpcode("iconst 0")
}

func (j *JASM) addIAddOpcode() {
	j.addOpcode("iadd")
}

func (j *JASM) addISubOpcode() {
	j.addOpcode("isub")
}

func (j *JASM) addIMulOpcode() {
	j.addOpcode("imul")
}

func (j *JASM) addIDivOpcode() {
	j.addOpcode("idiv")
}

func (j *JASM) addIAndOpcode() {
	j.addOpcode("iand")
	j.pst.Push(Boolean)
}

func (j *JASM) addIOrOpcode() {
	j.addOpcode("ior")
	j.pst.Push(Boolean)
}

func (j *JASM) addGotoOpcode(label string) {
	j.addOpcode("goto", label)
}

func (j *JASM) addIfEqOpcode(label string) {
	j.addOpcode("ifeq", label)
}

func (j *JASM) addIfNeOpcode(label string) {
	j.addOpcode("ifne", label)
}

func (j *JASM) addInvokeVirtual(method string) {
	j.addOpcode("invokevirtual", method)
}

func (j *JASM) addInvokeStatic(method, signature string, returnType string) {
	j.addOpcode(fmt.Sprintf("invokestatic %s.%s(%s)%s", j.className, method, signature, returnType))
}

func (j *JASM) addInvokeVirtualPrintWithType() error {
	pt := j.pst.Pop()

	if pt == String {
		j.addInvokeVirtual("java/io/PrintStream.print(java/lang/String)V")
	} else if pt == Integer {
		j.addInvokeVirtual("java/io/PrintStream.print(I)V")
	} else {
		return fmt.Errorf("undefined type %s in write/writeln", pt)
	}

	return nil
}

func (j *JASM) addStaticPrintStream() {
	j.addOpcode("getstatic", "java/lang/System.out", "java/io/PrintStream")
}

func (j *JASM) addLabel(label string) {
	j.addLine(fmt.Sprintf("%s:", label))
}

func (j *JASM) finishMain() {
	j.addLine("return")
	j.decTab()
	j.addLine("}")
}

func (j *JASM) startInvokeDynamic(param string) {
	j.addLine(fmt.Sprintf("invokedynamic %s {", param))
	j.incTab()
}

func (j *JASM) finishInvokeDynamic() {
	j.decTab()
	j.addLine("}")
}

func (j *JASM) newLabel() string {
	j.labelID++
	return fmt.Sprintf("L%d", j.labelID)
}

func (j *JASM) addLine(line string) {
	j.code.AddLine(line)
}

func (j *JASM) incTab() {
	j.code.IncTab()
}

func (j *JASM) decTab() {
	j.code.DecTab()
}

func (j *JASM) genBooleanOperatorTpl(ifOpcode string) {
	lfalse := j.newLabel()
	lnext := j.newLabel()
	j.addOpcode(ifOpcode, lfalse)
	j.addPushTrueOpcode()
	j.addGotoOpcode(lnext)
	j.addLabel(lfalse)
	j.addPushFalseOpcode()
	j.addLabel(lnext)
	j.pst.Push(Boolean)
}
