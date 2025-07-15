package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileClass(indent int) error {
	c.writer.WriteOpeningTag("class", indent)
	nextIndent := indent + common.BaseIndent

	// 'class'
	token, err := c.lexer.ExpectKeyword("class")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	// className
	token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
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

	// classVarDec*
	for done, err := c.compileClassVarDec(nextIndent); !done || err != nil; done, err = c.compileClassVarDec(nextIndent) {
		if err != nil {
			return err
		}
	}

	// classSubroutineDec*
	for done, err := c.compileSubroutineDec(nextIndent); !done || err != nil; done, err = c.compileSubroutineDec(nextIndent) {
		if err != nil {
			return err
		}
	}

	// '}'
	token, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, nextIndent)

	c.writer.WriteClosingTag("class", indent)
	return nil
}
