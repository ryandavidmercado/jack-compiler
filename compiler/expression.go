package compiler

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

func (c *Compiler) compileExpression(indent int) error {
	nextIndent := indent + common.BaseIndent

	// term
	err := c.compileTerm(nextIndent, common.Indent("<expression>\n", indent))
	if err != nil {
		return err
	}

	// (op term)*
	for {
		// op
		token, err := c.lexer.Expect(func(t *common.Token) bool {
			_, valid := opTokens[t.Value]
			return t.TokenType == common.TTSymbol && valid
		})
		if err != nil {
			// since this is optional, relinquish token & break early
			c.lexer.TokenIsUnused = true
			break
		}
		c.writer.WriteToken(token, nextIndent)

		// term
		err = c.compileTerm(nextIndent, "")
		if err != nil {
			return err
		}
	}

	c.writer.WriteClosingTag("expression", indent)
	return nil
}
