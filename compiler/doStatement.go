package compiler

func (c *Compiler) compileDoStatement(st *symbolTable) error {
	// 'do'
	_, err := c.lexer.ExpectKeyword("do")
	if err != nil {
		return err
	}

	// subroutineCall (handle as expression)
	err = c.compileExpression(st)
	if err != nil {
		return err
	}

	// ';'
	_, err = c.lexer.ExpectSymbol(";")
	if err != nil {
		return err
	}

	return c.writer.WritePop("temp", 0)
}
