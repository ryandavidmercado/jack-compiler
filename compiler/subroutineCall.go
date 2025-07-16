package compiler

import (
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (c *Compiler) compileSubroutineCall(firstValue string, lookAheadToken *common.Token, st *symbolTable) error {
	var className string
	var funcName string
	argCount := 0

	if lookAheadToken.Value == "." {
		// firstValue is either an object or a class
		next, err := c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}

		objSymbol, err := getSymbolFromTables(firstValue, st, c.symbolTable)

		if err == nil {
			// we are calling a method
			className = objSymbol.symbolType
			funcName = next.Value
			c.writer.WritePushSymbol(objSymbol) // push the receiver onto the stack as the callee's first argument
			argCount += 1
		} else {
			// we are calling a static function
			className = firstValue
			funcName = next.Value
		}

	} else {
		// we are calling a method on the current object
		// firstValue is the funcName; we should use this class as the className
		className = c.className
		funcName = firstValue
		c.writer.WritePush("pointer", 0)
		argCount += 1

		c.lexer.TokenIsUnused = true
	}

	_, err := c.lexer.ExpectSymbol("(")
	if err != nil {
		return err
	}

	expressionCount, err := c.compileExpressionList(st)
	if err != nil {
		return err
	}

	argCount += int(expressionCount)

	_, err = c.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}

	fullFunctionName := className + "." + funcName
	return c.writer.WriteFunctionCall(fullFunctionName, uint8(argCount))
}
