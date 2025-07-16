package compiler

import (
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (c *Compiler) compileClassVarDec() error {
	// static | field
	kind, err := c.lexer.Expect(func(t *common.Token) bool {
		return t.TokenType == common.TTKeyword && (t.Value == "static" || t.Value == "field")
	})
	if err != nil {
		c.lexer.TokenIsUnused = true
		return nil
	}

	// type
	varType, err := c.compileType(false)
	if err != nil {
		return err
	}

	// varName
	varName, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	c.symbolTable.Add(varName.Value, varType.Value, kind.Value)

	// (, varName)* OR ';'
	for {
		// , OR ;
		sep, err := c.lexer.Expect(func(t *common.Token) bool {
			return t.TokenType == common.TTSymbol &&
				(t.Value == ";" || t.Value == ",")
		})

		if err != nil {
			return err
		}

		if sep.Value == ";" {
			break
		}

		// varName
		varName, err := c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}

		c.symbolTable.Add(varName.Value, varType.Value, kind.Value)
	}

	return nil
}
