package codegen

import (
	"fmt"
	"strings"
)

type Code struct {
	code []string
	tabs string
}

func NewCode() *Code {
	return &Code{}
}

func (c *Code) IncTab() {
	c.tabs += "\t"
}

func (c *Code) DecTab() {
	if len(c.tabs) == 0 {
		return
	}

	c.tabs = c.tabs[:len(c.tabs)-1]
}

func (c *Code) AddLine(line string) {
	c.code = append(c.code, fmt.Sprintf("%s%s\n", c.tabs, line))
}

func (c *Code) Code() string {
	return strings.Join(c.code, "")
}
