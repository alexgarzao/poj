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

	fmt.Println("Compiling ", inputFile)

	jasm, err := genCode(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = genFile(inputFile, jasm)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func genCode(inputFile string) (string, error) {
	inputStream, err := antlr.NewFileStream(inputFile + ".pas")
	if err != nil {
		return "", err
	}

	filename := path.Base(inputFile)

	lexer := parsing.NewPascalLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parsing.NewPascalParser(stream)
	p.AddErrorListener(antlr.NewDiagnosticErrorListener(true))

	listener := codegen.NewTreeShapeListener(filename)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	return listener.Code(), nil
}

func genFile(inputFile string, jasm string) error {
	filename := path.Base(inputFile)
	if err := os.WriteFile(filename+".jasm", []byte(jasm), 0666); err != nil {
		return err
	}

	return nil
}
