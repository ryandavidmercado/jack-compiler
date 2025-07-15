package compiler

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

func (c *Compiler) compileTerm(indent int, toPrint string) error {
	nextIndent := indent + common.BaseIndent

	printToPrint := func() {
		if len(toPrint) == 0 {
			return
		}

		c.writer.WriteString(toPrint)
	}

	endcompile := func(token *common.Token) error {
		if token != nil {
			c.writer.WriteToken(token, nextIndent)
		}
		c.writer.WriteClosingTag("term", indent)
		return nil
	}

	// integerConstant
	token, err := c.lexer.ExpectTokenType(common.TTIntegerConstant)
	if err == nil {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		return endcompile(token)
	}
	c.lexer.TokenIsUnused = true

	// | stringConstant
	token, err = c.lexer.ExpectTokenType(common.TTStringConstant)
	if err == nil {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		return endcompile(token)
	}
	c.lexer.TokenIsUnused = true

	// | keywordConstant
	token, err = c.lexer.Expect(func(t *common.Token) bool {
		_, valid := keywordConstants[t.Value]
		return t.TokenType == common.TTKeyword && valid
	})
	if err == nil {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		return endcompile(token)
	}
	c.lexer.TokenIsUnused = true

	// | '(' expression ')'
	token, err = c.lexer.ExpectSymbol("(")
	if err == nil {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		c.writer.WriteToken(token, nextIndent)
		err = c.compileExpression(nextIndent)
		if err != nil {
			return err
		}
		token, err := c.lexer.ExpectSymbol(")")
		if err != nil {
			return err
		}
		return endcompile(token)
	}
	c.lexer.TokenIsUnused = true

	// | (unaryOp term)
	token, err = c.lexer.Expect(func(t *common.Token) bool {
		_, valid := unaryOps[t.Value]
		return t.TokenType == common.TTSymbol && valid
	})
	if err == nil {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		c.writer.WriteToken(token, nextIndent)
		err := c.compileTerm(nextIndent, "")
		if err != nil {
			return err
		}
		return endcompile(nil)
	}
	c.lexer.TokenIsUnused = true

	// varName | varName '[' expression ']' | subroutineCall
	token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	lookAheadToken, err := c.lexer.Advance()
	if err != nil {
		return err
	}

	// varName '[' expression ']'
	if lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "["}) {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		c.writer.WriteToken(token, nextIndent)
		c.writer.WriteToken(lookAheadToken, nextIndent)

		err := c.compileExpression(nextIndent)
		if err != nil {
			return err
		}

		token, err = c.lexer.ExpectSymbol("]")
		if err != nil {
			return err
		}
		return endcompile(token)
	}

	// subRoutineCall
	if lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "("}) ||
		lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "."}) {
		printToPrint()
		c.writer.WriteOpeningTag("term", indent)
		err = c.compileSubroutineCall(nextIndent, token, lookAheadToken)
		if err != nil {
			return err
		}
		return endcompile(nil)
	}

	// varName
	printToPrint()
	c.writer.WriteOpeningTag("term", indent)
	c.lexer.TokenIsUnused = true // relinquish lookAheadToken
	return endcompile(token)
}
