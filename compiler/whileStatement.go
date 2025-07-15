package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileWhileStatement(indent int) error {
	c.writer.WriteOpeningTag("whileStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'while'
	token, err := c.lexer.ExpectKeyword("while")
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
	if err := c.compileExpression(nextIndent); err != nil {
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
	if err := c.compileStatements(nextIndent); err != nil {
		return err
	}

	// '}'
	token, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	c.writer.WriteClosingTag("whileStatement", indent)
	return nil
}
