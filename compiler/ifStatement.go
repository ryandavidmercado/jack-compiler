package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileIfStatement(indent int) error {
	c.writer.WriteOpeningTag("ifStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'if'
	token, err := c.lexer.ExpectKeyword("if")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// '('
	token, err = c.lexer.ExpectSymbol("(")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// expression
	err = c.compileExpression(nextIndent)
	if err != nil {
		return err
	}

	// ')'
	token, err = c.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// '{'
	token, err = c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// statements
	err = c.compileStatements(nextIndent)
	if err != nil {
		return err
	}

	// '}'
	token, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// else?
	token, err = c.lexer.ExpectKeyword("else")
	if err != nil {
		// if we don't have else, relinquish token & return early
		c.lexer.TokenIsUnused = true
		c.writer.WriteClosingTag("ifStatement", indent)
		return nil
	}
	c.writer.WriteToken(token, nextIndent)

	// '{'
	token, err = c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// statements
	err = c.compileStatements(nextIndent)
	if err != nil {
		return err
	}

	// '}'
	token, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	c.writer.WriteClosingTag("ifStatement", indent)
	return nil
}
