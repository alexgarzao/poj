package codegen

import (
	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type CustomSyntaxError struct {
	line, column int
	msg          string
}

func (c *CustomSyntaxError) Error() string {
	return c.msg
}

type CustomErrorListener struct {
	*antlr.DefaultErrorListener // Embed default which ensures we fit the interface
	Errors                      []error
}

func (c *CustomErrorListener) Add(line int, err error) {
	c.Errors = append(c.Errors, fmt.Errorf("line %d: %s", line, err))
}

func (c *CustomErrorListener) String() string {
	var s string

	for _, err := range c.Errors {
		s += err.Error() + "\n"
	}

	return s
}

func (c *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	c.Errors = append(c.Errors, &CustomSyntaxError{
		line:   line,
		column: column,
		msg:    msg,
	})
}
