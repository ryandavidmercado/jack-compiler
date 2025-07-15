package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseWhileStatement(indent int) error {
	p.writer.WriteOpeningTag("whileStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'while'
	token, err := p.lexer.ExpectKeyword("while")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// '('
	token, err = p.lexer.ExpectSymbol("(")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// expression
	if err := p.parseExpression(nextIndent); err != nil {
		return err
	}

	// ')'
	token, err = p.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// '{'
	token, err = p.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// statements
	if err := p.parseStatements(nextIndent); err != nil {
		return err
	}

	// '}'
	token, err = p.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	p.writer.WriteClosingTag("whileStatement", indent)
	return nil
}
