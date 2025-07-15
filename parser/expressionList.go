package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseExpressionList(indent int) error {
	p.writer.WriteOpeningTag("expressionList", indent)
	nextIndent := indent + common.BaseIndent

	// expression
	err := p.parseExpression(nextIndent)
	if err != nil {
		// expression is optional here; relinquish token and return
		p.writer.WriteClosingTag("expressionList", indent)
		p.lexer.TokenIsUnused = true
		return nil
	}

	// (, expression)*
	for {
		token, err := p.lexer.ExpectSymbol(",")
		if err != nil {
			// no more expressions, break
			p.lexer.TokenIsUnused = true
			break
		}
		p.writer.WriteToken(token, nextIndent)

		err = p.parseExpression(nextIndent)
		if err != nil {
			// we should have gotten an expression here; fail
			return err
		}
	}

	p.writer.WriteClosingTag("expressionList", indent)
	return nil
}
