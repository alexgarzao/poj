package main

import (
	"fmt"
	"os"
	"path"

	"github.com/alexgarzao/poj/codegen"
	"github.com/alexgarzao/poj/parsing"
	"github.com/antlr4-go/antlr/v4"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println("POJ: Pascal on the JVM")
		fmt.Println()
		fmt.Println("Usage:")
		fmt.Println("\tpoj <file>")
		os.Exit(1)
	}

	inputFile := os.Args[1]

	input, err := antlr.NewFileStream(inputFile + ".pas")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println("Compiling ", inputFile)

	lexer := parsing.NewPascalLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parsing.NewPascalParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	filename := path.Base(inputFile)
	listener := codegen.NewTreeShapeListener(filename)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	os.WriteFile(filename+".jasm", []byte(listener.Code()), 0666)
}
