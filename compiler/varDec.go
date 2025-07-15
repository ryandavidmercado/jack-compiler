package compiler

import (
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (c *Compiler) compileVarDec(indent int) (bool, error) {
	nextIndent := indent + common.BaseIndent

	// var
	token, err := c.lexer.ExpectKeyword("var")
	if err != nil {
		c.lexer.TokenIsUnused = true
		return true, nil
	}
	c.writer.WriteOpeningTag("varDec", indent)
	c.writer.WriteToken(token, nextIndent)

	// type
	err = c.compileType(false, nextIndent)
	if err != nil {
		return true, err
	}

	// varName
	token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return true, err
	}
	c.writer.WriteToken(token, nextIndent)

	// (, varName)* OR ';'
	for {
		// , OR ;
		token, err = c.lexer.ExpectTokenType(common.TTSymbol)
		if err != nil {
			return true, err
		}
		c.writer.WriteToken(token, nextIndent)

		if token.Value == ";" {
			break
		} else if token.Value != "," {
			return true, fmt.Errorf("Expected symbol ';' or ',', got %v", token)
		}

		// varName
		token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return true, err
		}
		c.writer.WriteToken(token, nextIndent)
	}

	c.writer.WriteClosingTag("varDec", indent)
	return false, nil
}
