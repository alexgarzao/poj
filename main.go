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

	jasm, lexerErrors, parserErrors, err := genCode(inputFile)
	if err != nil {
		if lexerErrors != nil && len(lexerErrors.Errors) > 0 {
			fmt.Printf("Lexer: %d errors found\n", len(lexerErrors.Errors))
			for _, e := range lexerErrors.Errors {
				fmt.Println("\t", e.Error())
			}
		}

		if parserErrors != nil && len(parserErrors.Errors) > 0 {
			fmt.Printf("Parser: %d errors found\n", len(parserErrors.Errors))
			for _, e := range parserErrors.Errors {
				fmt.Println("\t", e.Error())
			}
		}

		fmt.Printf("error %s\n", err)
		os.Exit(1)
	}

	err = genFile(inputFile, jasm)
	if err != nil {
		fmt.Printf("error %s\n", err)
		os.Exit(1)
	}
}

func genCode(inputFile string) (string, *codegen.CustomErrorListener, *codegen.CustomErrorListener, error) {
	inputStream, err := antlr.NewFileStream(inputFile + ".pas")
	if err != nil {
		return "", nil, nil, err
	}

	filename := path.Base(inputFile)

	lexerErrors := &codegen.CustomErrorListener{}
	lexer := parsing.NewPascalLexer(inputStream)
	lexer.RemoveErrorListeners()
	lexer.AddErrorListener(lexerErrors)

	parserErrors := &codegen.CustomErrorListener{}
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	p := parsing.NewPascalParser(stream)
	p.RemoveErrorListeners()
	p.AddErrorListener(parserErrors)

	listener := codegen.NewTreeShapeListener(filename, parserErrors)
	antlr.ParseTreeWalkerDefault.Walk(listener, p.Program())

	parserError := len(lexerErrors.Errors) > 0 || len(parserErrors.Errors) > 0
	if parserError {
		return "", lexerErrors, parserErrors, fmt.Errorf("during parsing")
	}

	return listener.Code(), nil, nil, nil
}

func genFile(inputFile string, jasm string) error {
	filename := path.Base(inputFile)
	if err := os.WriteFile(filename+".jasm", []byte(jasm), 0666); err != nil {
		return fmt.Errorf("during file generation: %w", err)
	}

	return nil
}
