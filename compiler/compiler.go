package compiler

import (
	"bufio"
	"github.com/ryandavidmercado/jack-compiler/lexer"
)

type Compiler struct {
	lexer  *lexer.Lexer
	writer *compilerWriter
}

func New(lexer *lexer.Lexer, writer *bufio.Writer) *Compiler {
	return &Compiler{lexer: lexer, writer: &compilerWriter{writer}}
}

func (c *Compiler) Compile() error {
	err := c.compileClass(0)
	if err != nil {
		return err
	}

	err = c.writer.Flush()
	return err
}
