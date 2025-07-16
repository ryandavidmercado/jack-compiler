package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileLetStatement(st *symbolTable) error {
	// 'let'
	_, err := c.lexer.ExpectKeyword("let")
	if err != nil {
		return err
	}

	// varName
	varName, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	symbol, err := getSymbolFromTables(varName.Value, st, c.symbolTable)
	if err != nil {
		return err
	}

	// ('[' expression ']')?
	_, err = c.lexer.ExpectSymbol("[")
	if err == nil {
		// ** push arr **
		c.writer.WritePushSymbol(symbol)

		// ** push access value **
		err = c.compileExpression(st)
		if err != nil {
			return err
		}

		// ** add **
		// now the relevant addr is top of stack
		c.writer.WriteBody("add")

		// ']'
		_, err := c.lexer.ExpectSymbol("]")
		if err != nil {
			return err
		}

		// '='
		_, err = c.lexer.ExpectSymbol("=")
		if err != nil {
			return err
		}

		// ** push the value to assign to arr[x] **
		err = c.compileExpression(st)
		if err != nil {
			return err
		}

		// put value in temp[0]
		// save addr goes back to top of stack
		c.writer.WritePop("temp", 0)

		// point THAT to save addr
		c.writer.WritePop("pointer", 1)
		// get the value to save
		c.writer.WritePush("temp", 0)
		// save the value to THAT
		c.writer.WritePop("that", 0)

		// ';'
		_, err = c.lexer.ExpectSymbol(";")
		if err != nil {
			return err
		}

		return nil
	}
	c.lexer.TokenIsUnused = true

	// '='
	_, err = c.lexer.ExpectSymbol("=")
	if err != nil {
		return err
	}

	// expression
	err = c.compileExpression(st)
	if err != nil {
		return err
	}

	// ';'
	_, err = c.lexer.ExpectSymbol(";")
	if err != nil {
		return err
	}

	return c.writer.WritePopSymbol(symbol)
}
