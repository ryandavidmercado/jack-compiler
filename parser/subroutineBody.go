package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseSubroutineBody(indent int) error {
	p.writer.WriteOpeningTag("subroutineBody", indent)
	nextIndent := indent + common.BaseIndent

	// {
	token, err := p.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// varDec*
	for done, err := p.parseVarDec(nextIndent); !done || err != nil; done, err = p.parseVarDec(nextIndent) {
		if err != nil {
			return err
		}
	}

	// statements
	err = p.parseStatements(nextIndent)
	if err != nil {
		return err
	}

	// }
	token, err = p.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	p.writer.WriteClosingTag("subroutineBody", indent)
	return nil
}
