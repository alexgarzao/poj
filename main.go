package main

import (
	"fmt"
	"os"

	"github.com/alexgarzao/poj/parsing"
	"github.com/antlr4-go/antlr/v4"
)

type TreeShapeListener struct {
	*parsing.BaseJSONListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
	fmt.Println(ctx.GetText())
}

func main() {
	input, _ := antlr.NewFileStream(os.Args[1])
	lexer := parsing.NewJSONLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parsing.NewJSONParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	tree := p.Json()
	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}
