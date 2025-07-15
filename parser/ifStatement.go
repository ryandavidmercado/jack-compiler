package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseIfStatement(indent int) error {
	p.writer.WriteOpeningTag("ifStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'if'
	token, err := p.lexer.ExpectKeyword("if")
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
	err = p.parseExpression(nextIndent)
	if err != nil {
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
	err = p.parseStatements(nextIndent)
	if err != nil {
		return err
	}

	// '}'
	token, err = p.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// else?
	token, err = p.lexer.ExpectKeyword("else")
	if err != nil {
		// if we don't have else, relinquish token & return early
		p.lexer.TokenIsUnused = true
		p.writer.WriteClosingTag("ifStatement", indent)
		return nil
	}
	p.writer.WriteToken(token, nextIndent)

	// '{'
	token, err = p.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// statements
	err = p.parseStatements(nextIndent)
	if err != nil {
		return err
	}

	// '}'
	token, err = p.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	p.writer.WriteClosingTag("ifStatement", indent)
	return nil
}
