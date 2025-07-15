package compiler

import (
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (c *Compiler) compileSubroutineCall(indent int, token1 *common.Token, token2 *common.Token) error {
	if token1 == nil {
		res, err := c.lexer.Advance()
		if err != nil {
			return err
		}
		token1 = res
	}

	if token2 == nil {
		res, err := c.lexer.Advance()
		if err != nil {
			return err
		}
		token2 = res
	}

	// subroutineName | className | varName
	if token1.TokenType != common.TTIdentifier {
		return fmt.Errorf("Expected identifier token, got %v", token1)
	}
	c.writer.WriteToken(token1, indent)

	// '.' | '('
	valid := token2.Value == "." || token2.Value == "("
	if token2.TokenType != common.TTSymbol || !valid {
		return fmt.Errorf("Expected symbol token '.' or '(', got %v", token2)
	}
	c.writer.WriteToken(token2, indent)

	if token2.Value == "." {
		// subroutineName
		token, err := c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}
		c.writer.WriteToken(token, indent)

		// (
		token, err = c.lexer.ExpectSymbol("(")
		if err != nil {
			return err
		}
		c.writer.WriteToken(token, indent)
	}

	err := c.compileExpressionList(indent)
	if err != nil {
		return err
	}

	token, err := c.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}
	c.writer.WriteToken(token, indent)

	return nil
}
