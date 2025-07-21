package compiler

import (
	"fmt"
	"strconv"

	"github.com/ryandavidmercado/jack-compiler/common"
)

var keywordConstants = map[string]struct{}{
	"true":  {},
	"false": {},
	"null":  {},
	"this":  {},
}

var unaryOps = map[string]struct{}{
	"-": {},
	"~": {},
}

func (c *Compiler) compileTerm(st *symbolTable) error {
	// integerConstant
	intToken, err := c.lexer.ExpectTokenType(common.TTIntegerConstant)
	if err == nil {
		asInt, err := strconv.Atoi(intToken.Value)
		if err != nil {
			return err
		}
		if asInt < -32_767 || asInt > 32_767 {
			return fmt.Errorf("Invalid integer constant %d", asInt)
		}

		return c.writer.WritePushConstant(int16(asInt))
	}
	c.lexer.TokenIsUnused = true

	// | stringConstant
	strToken, err := c.lexer.ExpectTokenType(common.TTStringConstant)
	if err == nil {
		return c.writer.WriteStringConstant(strToken.Value)
	}
	c.lexer.TokenIsUnused = true

	// | keywordConstant
	keywordToken, err := c.lexer.Expect(func(t *common.Token) bool {
		_, valid := keywordConstants[t.Value]
		return t.TokenType == common.TTKeyword && valid
	})
	if err == nil {
		return c.writer.WriteKeywordConstant(keywordToken.Value)
	}
	c.lexer.TokenIsUnused = true

	// | '(' expression ')'
	_, err = c.lexer.ExpectSymbol("(")
	if err == nil {
		err = c.compileExpression(st)
		if err != nil {
			return err
		}
		_, err = c.lexer.ExpectSymbol(")")
		return err
	}
	c.lexer.TokenIsUnused = true

	// | (unaryOp term)
	unaryOpToken, err := c.lexer.Expect(func(t *common.Token) bool {
		_, valid := unaryOps[t.Value]
		return t.TokenType == common.TTSymbol && valid
	})
	if err == nil {
		err = c.compileTerm(st)
		if err != nil {
			return err
		}
		return c.writer.WriteUnaryOp(unaryOpToken.Value)
	}
	c.lexer.TokenIsUnused = true

	// varName | varName '[' expression ']' | subroutineCall
	varNameToken, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	lookAheadToken, err := c.lexer.Advance()
	if err != nil {
		return err
	}

	// varName '[' expression ']'
	if lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "["}) {
		symbol, err := getSymbolFromTables(varNameToken.Value, st, c.symbolTable)
		if err != nil {
			return err
		}

		// ** push arr **
		c.writer.WritePushSymbol(symbol)

		// ** compute/push index **
		err = c.compileExpression(st)
		if err != nil {
			return err
		}

		// add base arr address to access index
		c.writer.WriteBody("add")

		// set THAT to the memory address of the element
		c.writer.WritePop("pointer", 1)

		// finally, push the value at that memory address to top of stack
		c.writer.WritePush("that", 0)

		_, err = c.lexer.ExpectSymbol("]")
		return nil
	}

	// subRoutineCall
	if lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "("}) ||
		lookAheadToken.Equals(&common.Token{TokenType: common.TTSymbol, Value: "."}) {
		return c.compileSubroutineCall(varNameToken.Value, lookAheadToken, st)
	}

	// varName
	symbol, err := getSymbolFromTables(varNameToken.Value, st, c.symbolTable)
	if err != nil {
		return err
	}

	err = c.writer.WritePushSymbol(symbol)
	c.lexer.TokenIsUnused = true // relinquish lookAheadToken

	return nil
}
