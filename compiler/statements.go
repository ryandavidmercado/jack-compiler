package compiler

import (
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (c *Compiler) compileStatements(st *symbolTable) error {
	var statementcompilers = map[string]func(st *symbolTable) error{
		"let":    c.compileLetStatement,
		"if":     c.compileIfStatement,
		"while":  c.compileWhileStatement,
		"do":     c.compileDoStatement,
		"return": c.compileReturnStatement,
	}

	for {
		token, err := c.lexer.Expect(func(t *common.Token) bool {
			_, valid := statementcompilers[t.Value]
			return t.TokenType == common.TTKeyword && valid
		})

		c.lexer.TokenIsUnused = true // we relinquish this token to next compiler regardless of outcome
		if err != nil {
			// a statement is always optional; let compiler take relinquished token
			break
		}

		err = statementcompilers[token.Value](st)
		if err != nil {
			return err
		}
	}

	return nil
}
