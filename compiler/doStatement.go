package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileDoStatement(indent int) error {
	c.writer.WriteOpeningTag("doStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'do'
	token, err := c.lexer.ExpectKeyword("do")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// subroutineCall
	err = c.compileSubroutineCall(nextIndent, nil, nil)
	if err != nil {
		return err
	}

	// ';'
	token, err = c.lexer.ExpectSymbol(";")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	c.writer.WriteClosingTag("doStatement", indent)
	return nil
}
