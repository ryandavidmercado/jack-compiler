package parser

import (
	"github.com/ryandavidmercado/jack-compiler/common"
)

var opTokens = map[string]struct{}{
	"+": {},
	"-": {},
	"*": {},
	"/": {},
	"&": {},
	"|": {},
	"<": {},
	">": {},
	"=": {},
}

func (p *Parser) parseExpression(indent int) error {
	nextIndent := indent + common.BaseIndent

	// term
	err := p.parseTerm(nextIndent, common.Indent("<expression>\n", indent))
	if err != nil {
		return err
	}

	// (op term)*
	for {
		// op
		token, err := p.lexer.Expect(func(t *common.Token) bool {
			_, valid := opTokens[t.Value]
			return t.TokenType == common.TTSymbol && valid
		})
		if err != nil {
			// since this is optional, relinquish token & break early
			p.lexer.TokenIsUnused = true
			break
		}
		p.writer.WriteToken(token, nextIndent)

		// term
		err = p.parseTerm(nextIndent, "")
		if err != nil {
			return err
		}
	}

	p.writer.WriteClosingTag("expression", indent)
	return nil
}
