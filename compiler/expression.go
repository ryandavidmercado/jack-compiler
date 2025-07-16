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

func (c *Compiler) compileExpression(st *symbolTable) error {
	// term
	err := c.compileTerm(st)
	if err != nil {
		return err
	}

	// (op term)*
	for {
		// op
		opToken, err := c.lexer.Expect(func(t *common.Token) bool {
			_, valid := opTokens[t.Value]
			return t.TokenType == common.TTSymbol && valid
		})
		if err != nil {
			// since this is optional, relinquish token & break early
			c.lexer.TokenIsUnused = true
			break
		}

		// term
		err = c.compileTerm(st)
		if err != nil {
			return err
		}

		err = c.writer.WriteBinaryOp(opToken.Value)
		if err != nil {
			return err
		}
	}

	return nil
}
