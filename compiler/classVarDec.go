package compiler

import (
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (c *Compiler) compileClassVarDec(indent int) (bool, error) {
	nextIndent := indent + common.BaseIndent

	// static | field
	token, err := c.lexer.Expect(func(t *common.Token) bool {
		return t.TokenType == common.TTKeyword && (t.Value == "static" || t.Value == "field")
	})
	if err != nil {
		c.lexer.TokenIsUnused = true
		return true, nil
	}
	c.writer.WriteOpeningTag("classVarDec", indent)
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

	c.writer.WriteClosingTag("classVarDec", indent)
	return false, nil
}
