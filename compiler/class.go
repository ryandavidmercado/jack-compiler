package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileClass() error {
	// 'class'
	_, err := c.lexer.ExpectKeyword("class")
	if err != nil {
		return err
	}

	// className
	className, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}
	c.className = className.Value

	// '{'
	_, err = c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}

	// classVarDec*
	for {
		_, err = c.lexer.Expect(func(t *common.Token) bool {
			return t.TokenType == common.TTKeyword && (t.Value == "static" || t.Value == "field")
		})
		c.lexer.TokenIsUnused = true

		if err != nil {
			break
		} else {
			c.compileClassVarDec()
		}
	}

	// classSubroutineDec*
	for {
		_, err := c.lexer.Expect(func(t *common.Token) bool {
			return t.TokenType == common.TTKeyword && (t.Value == "constructor" || t.Value == "function" || t.Value == "method")
		})
		c.lexer.TokenIsUnused = true

		if err != nil {
			break
		} else {
			c.compileSubroutineDec()
		}
	}

	// '}'
	_, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}

	return nil
}
