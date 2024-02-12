package main

import (
	"fmt"
	"os"

	"github.com/alexgarzao/poj/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parsing.BasePascalListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewPascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewPascalParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Program()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
