package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileSubroutineBody(indent int) error {
	c.writer.WriteOpeningTag("subroutineBody", indent)
	nextIndent := indent + common.BaseIndent

	// {
	token, err := c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// varDec*
	for done, err := c.compileVarDec(nextIndent); !done || err != nil; done, err = c.compileVarDec(nextIndent) {
		if err != nil {
			return err
		}
	}

	// statements
	err = c.compileStatements(nextIndent)
	if err != nil {
		return err
	}

	// }
	token, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	c.writer.WriteClosingTag("subroutineBody", indent)
	return nil
}
