package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileExpressionList(indent int) error {
	c.writer.WriteOpeningTag("expressionList", indent)
	nextIndent := indent + common.BaseIndent

	// expression
	err := c.compileExpression(nextIndent)
	if err != nil {
		// expression is optional here; relinquish token and return
		c.writer.WriteClosingTag("expressionList", indent)
		c.lexer.TokenIsUnused = true
		return nil
	}

	// (, expression)*
	for {
		token, err := c.lexer.ExpectSymbol(",")
		if err != nil {
			// no more expressions, break
			c.lexer.TokenIsUnused = true
			break
		}
		c.writer.WriteToken(token, nextIndent)

		err = c.compileExpression(nextIndent)
		if err != nil {
			// we should have gotten an expression here; fail
			return err
		}
	}

	c.writer.WriteClosingTag("expressionList", indent)
	return nil
}
