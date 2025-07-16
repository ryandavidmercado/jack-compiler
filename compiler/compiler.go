package compiler

import (
	"bufio"
	"fmt"

	"github.com/ryandavidmercado/jack-compiler/lexer"
)

type Compiler struct {
	lexer       *lexer.Lexer
	writer      *compilerWriter
	symbolTable *symbolTable
	className   string
	labelNum    int
}

func New(lexer *lexer.Lexer, writer *bufio.Writer) *Compiler {
	return &Compiler{
		lexer:       lexer,
		writer:      &compilerWriter{writer},
		symbolTable: NewSymbolTable([]string{"field", "static"}),
		labelNum:    -1,
	}
}

func (c *Compiler) Compile() error {
	err := c.compileClass()
	if err != nil {
		return err
	}

	err = c.writer.Flush()
	return err
}

func (c *Compiler) label() string {
	return fmt.Sprintf("%v_%d", c.className, c.labelNum)
}

func (c *Compiler) incLabel() {
	c.labelNum += 1
}
