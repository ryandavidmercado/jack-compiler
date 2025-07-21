package compiler

func (c *Compiler) compileReturnStatement(st *symbolTable) error {
	// 'return'
	_, err := c.lexer.ExpectKeyword("return")
	if err != nil {
		return err
	}

	// ';'
	// look for this; if we don't find it we need an expression first
	_, err = c.lexer.ExpectSymbol(";")
	if err == nil {
		// found ;, return early
		err = c.writer.WritePushConstant(0)
		if err != nil {
			return err
		}
		return c.writer.WriteReturn()
	}
	c.lexer.TokenIsUnused = true

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

	return c.writer.WriteReturn()
}
