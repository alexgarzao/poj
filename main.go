package main

import (
	"fmt"
	"os"

	"github.com/alexgarzao/poj/codegen"
	"github.com/alexgarzao/poj/parsing"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("POJ: Pascal on the JVM")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("\tpoj <file>.pas")
		os.Exit(1)
	}

	input, err := antlr.NewFileStream(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Compiling ", os.Args[1])

	lexer := parsing.NewPascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parsing.NewPascalParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))
	listener := codegen.NewTreeShapeListener()
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	fmt.Print(listener.Code())
}
