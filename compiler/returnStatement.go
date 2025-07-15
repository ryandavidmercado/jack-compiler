package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileReturnStatement(indent int) error {
	c.writer.WriteOpeningTag("returnStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'return'
	token, err := c.lexer.ExpectKeyword("return")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// ';'
	// look for this; if we don't find it we need an expression first
	token, err = c.lexer.ExpectSymbol(";")
	if err != nil {
		// didn't find it; relinqiush token & proceed to find expression
		c.lexer.TokenIsUnused = true
	} else {
		// found it; terminate early
		c.writer.WriteToken(token, nextIndent)
		c.writer.WriteClosingTag("returnStatement", indent)
		return nil
	}

	// expression
	err = c.compileExpression(nextIndent)
	if err != nil {
		return err
	}

	// ';'
	token, err = c.lexer.ExpectSymbol(";")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	c.writer.WriteClosingTag("returnStatement", indent)
	return nil
}
