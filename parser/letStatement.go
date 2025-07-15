package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseLetStatement(indent int) error {
	p.writer.WriteOpeningTag("letStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'let'
	token, err := p.lexer.ExpectKeyword("let")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// varName
	token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// ('[' expression ']')?
	token, err = p.lexer.ExpectSymbol("[")
	if err == nil {
		p.writer.WriteToken(token, nextIndent)

		// expression
		err = p.parseExpression(nextIndent)
		if err != nil {
			return err
		}

		// ']'
		token, err := p.lexer.ExpectSymbol("]")
		if err != nil {
			return err
		}
		p.writer.WriteToken(token, nextIndent)
	} else {
		// this is optional; relinquish token & skip
		p.lexer.TokenIsUnused = true
	}

	// '='
	token, err = p.lexer.ExpectSymbol("=")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

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

	p.writer.WriteClosingTag("letStatement", indent)
	return nil
}
