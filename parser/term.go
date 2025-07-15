package parser

import "github.com/ryandavidmercado/jack-compiler/common"

var keywordConstants = map[string]struct{}{
	"true":  {},
	"false": {},
	"null":  {},
	"this":  {},
}

var unaryOps = map[string]struct{}{
	"-": {},
	"~": {},
}

func (p *Parser) parseTerm(indent int, toPrint string) error {
	nextIndent := indent + common.BaseIndent

	printToPrint := func() {
		if len(toPrint) == 0 {
			return
		}

		p.writer.WriteString(toPrint)
	}

	endParse := func(token *common.Token) error {
		if token != nil {
			p.writer.WriteToken(token, nextIndent)
		}
		p.writer.WriteClosingTag("term", indent)
		return nil
	}

	// integerConstant
	token, err := p.lexer.ExpectTokenType(common.TTIntegerConstant)
	if err == nil {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		return endParse(token)
	}
	p.lexer.TokenIsUnused = true

	// | stringConstant
	token, err = p.lexer.ExpectTokenType(common.TTStringConstant)
	if err == nil {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		return endParse(token)
	}
	p.lexer.TokenIsUnused = true

	// | keywordConstant
	token, err = p.lexer.Expect(func(t *common.Token) bool {
		_, valid := keywordConstants[t.Value]
		return t.TokenType == common.TTKeyword && valid
	})
	if err == nil {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		return endParse(token)
	}
	p.lexer.TokenIsUnused = true

	// | '(' expression ')'
	token, err = p.lexer.ExpectSymbol("(")
	if err == nil {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		p.writer.WriteToken(token, nextIndent)
		err = p.parseExpression(nextIndent)
		if err != nil {
			return err
		}
		token, err := p.lexer.ExpectSymbol(")")
		if err != nil {
			return err
		}
		return endParse(token)
	}
	p.lexer.TokenIsUnused = true

	// | (unaryOp term)
	token, err = p.lexer.Expect(func(t *common.Token) bool {
		_, valid := unaryOps[t.Value]
		return t.TokenType == common.TTSymbol && valid
	})
	if err == nil {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		p.writer.WriteToken(token, nextIndent)
		err := p.parseTerm(nextIndent, "")
		if err != nil {
			return err
		}
		return endParse(nil)
	}
	p.lexer.TokenIsUnused = true

	// varName | varName '[' expression ']' | subroutineCall
	token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	lookAheadToken, err := p.lexer.Advance()
	if err != nil {
		return err
	}

	// varName '[' expression ']'
	if lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "["}) {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		p.writer.WriteToken(token, nextIndent)
		p.writer.WriteToken(lookAheadToken, nextIndent)

		err := p.parseExpression(nextIndent)
		if err != nil {
			return err
		}

		token, err = p.lexer.ExpectSymbol("]")
		if err != nil {
			return err
		}
		return endParse(token)
	}

	// subRoutineCall
	if lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "("}) ||
		lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "."}) {
		printToPrint()
		p.writer.WriteOpeningTag("term", indent)
		err = p.parseSubroutineCall(nextIndent, token, lookAheadToken)
		if err != nil {
			return err
		}
		return endParse(nil)
	}

	// varName
	printToPrint()
	p.writer.WriteOpeningTag("term", indent)
	p.lexer.TokenIsUnused = true // relinquish lookAheadToken
	return endParse(token)
}
