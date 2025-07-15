package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseReturnStatement(indent int) error {
	p.writer.WriteOpeningTag("returnStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'return'
	token, err := p.lexer.ExpectKeyword("return")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// ';'
	// look for this; if we don't find it we need an expression first
	token, err = p.lexer.ExpectSymbol(";")
	if err != nil {
		// didn't find it; relinqiush token & proceed to find expression
		p.lexer.TokenIsUnused = true
	} else {
		// found it; terminate early
		p.writer.WriteToken(token, nextIndent)
		p.writer.WriteClosingTag("returnStatement", indent)
		return nil
	}

	// expression
	err = p.parseExpression(nextIndent)
	if err != nil {
		return err
	}

	// ';'
	token, err = p.lexer.ExpectSymbol(";")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	p.writer.WriteClosingTag("returnStatement", indent)
	return nil
}
