package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileParameterList(indent int) error {
	c.writer.WriteOpeningTag("parameterList", indent)
	nextIndent := indent + common.BaseIndent

	// type
	err := c.compileType(false, nextIndent)
	if err != nil {
		c.lexer.TokenIsUnused = true
		c.writer.WriteClosingTag("parameterList", indent)
		return nil // terminate early; we don't have list contents
	}

	// varName
	token, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	for {
		// , (optional)
		token, err := c.lexer.ExpectSymbol(",")
		if err != nil {
			c.lexer.TokenIsUnused = true // we'll reuse this to check ;
			break                        // exit; no optional varName list here
		}
		c.writer.WriteToken(token, nextIndent)

		// type
		err = c.compileType(false, nextIndent)
		if err != nil {
			return err
		}

		// varName
		token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}
		c.writer.WriteToken(token, nextIndent)
	}

	c.writer.WriteClosingTag("parameterList", indent)
	return nil
}
