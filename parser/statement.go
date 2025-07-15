package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseStatement(indent int) (bool, error) {
	var statementParsers = map[string]func(int) error{
		"let":    p.parseLetStatement,
		"if":     p.parseIfStatement,
		"while":  p.parseWhileStatement,
		"do":     p.parseDoStatement,
		"return": p.parseReturnStatement,
	}

	token, err := p.lexer.Expect(func(t *common.Token) bool {
		_, valid := statementParsers[t.Value]
		return t.TokenType == common.TTKeyword && valid
	})
	p.lexer.TokenIsUnused = true // we relinquish this token to next parser regardless of outcome
	if err != nil {
		// a statement is always optional; let parser take relinquished token
		return true, nil
	}

	// we have a statement; proceed to statement parser
	parser := statementParsers[token.Value]
	err = parser(indent)
	if err != nil {
		return true, err
	}

	return false, nil
}
