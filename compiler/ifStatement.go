package compiler

func (c *Compiler) compileIfStatement(st *symbolTable) error {
	c.incLabel()
	l1 := c.label()

	// 'if'
	_, err := c.lexer.ExpectKeyword("if")
	if err != nil {
		return err
	}

	// '('
	_, err = c.lexer.ExpectSymbol("(")
	if err != nil {
		return err
	}

	// expression
	err = c.compileExpression(st)
	if err != nil {
		return err
	}

	err = c.writer.WriteBody("not")
	if err != nil {
		return err
	}
	err = c.writer.WriteBody("if-goto %v", l1)
	if err != nil {
		return err
	}

	// ')'
	_, err = c.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}

	// '{'
	_, err = c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}

	// statements
	err = c.compileStatements(st)
	if err != nil {
		return err
	}

	// '}'
	_, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}

	// else?
	_, err = c.lexer.ExpectKeyword("else")
	if err != nil {
		// if we don't have else, relinquish token & return early
		c.lexer.TokenIsUnused = true
		return c.writer.WriteLabel(l1)
	}

	c.incLabel()
	l2 := c.label()

	err = c.writer.WriteBody("goto %v", l2)
	if err != nil {
		return err
	}
	err = c.writer.WriteLabel(l1)
	if err != nil {
		return err
	}

	// '{'
	_, err = c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}

	// statements
	err = c.compileStatements(st)
	if err != nil {
		return err
	}

	// '}'
	_, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}

	return c.writer.WriteLabel(l2)
}
