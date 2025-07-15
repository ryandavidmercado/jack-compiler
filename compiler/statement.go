package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileStatement(indent int) (bool, error) {
	var statementcompilers = map[string]func(int) error{
		"let":    c.compileLetStatement,
		"if":     c.compileIfStatement,
		"while":  c.compileWhileStatement,
		"do":     c.compileDoStatement,
		"return": c.compileReturnStatement,
	}

	token, err := c.lexer.Expect(func(t *common.Token) bool {
		_, valid := statementcompilers[t.Value]
		return t.TokenType == common.TTKeyword && valid
	})
	c.lexer.TokenIsUnused = true // we relinquish this token to next compiler regardless of outcome
	if err != nil {
		// a statement is always optional; let compiler take relinquished token
		return true, nil
	}

	// we have a statement; proceed to statement compiler
	compiler := statementcompilers[token.Value]
	err = compiler(indent)
	if err != nil {
		return true, err
	}

	return false, nil
}
