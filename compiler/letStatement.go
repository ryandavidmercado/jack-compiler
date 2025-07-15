package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileLetStatement(indent int) error {
	c.writer.WriteOpeningTag("letStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'let'
	token, err := c.lexer.ExpectKeyword("let")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// varName
	token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// ('[' expression ']')?
	token, err = c.lexer.ExpectSymbol("[")
	if err == nil {
		c.writer.WriteToken(token, nextIndent)

		// expression
		err = c.compileExpression(nextIndent)
		if err != nil {
			return err
		}

		// ']'
		token, err := c.lexer.ExpectSymbol("]")
		if err != nil {
			return err
		}
		c.writer.WriteToken(token, nextIndent)
	} else {
		// this is optional; relinquish token & skip
		c.lexer.TokenIsUnused = true
	}

	// '='
	token, err = c.lexer.ExpectSymbol("=")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

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

	c.writer.WriteClosingTag("letStatement", indent)
	return nil
}
