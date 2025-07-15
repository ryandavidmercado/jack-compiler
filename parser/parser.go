package parser

import (
	"bufio"
	"github.com/ryandavidmercado/jack-compiler/lexer"
)

type Parser struct {
	lexer  *lexer.Lexer
	writer *parserWriter
}

func New(lexer *lexer.Lexer, writer *bufio.Writer) *Parser {
	return &Parser{lexer: lexer, writer: &parserWriter{writer}}
}

func (p *Parser) Parse() error {
	err := p.parseClass(0)
	if err != nil {
		return err
	}

	err = p.writer.Flush()
	return err
}
