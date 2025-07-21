package compiler

func (c *Compiler) compileWhileStatement(st *symbolTable) error {
	c.incLabel()
	l1 := c.label()
	c.incLabel()
	l2 := c.label()

	// ** label L1 **
	err := c.writer.WriteLabel(l1)
	if err != nil {
		return err
	}

	// 'while'
	_, err = c.lexer.ExpectKeyword("while")
	if err != nil {
		return err
	}

	// '('
	_, err = c.lexer.ExpectSymbol("(")
	if err != nil {
		return err
	}

	// ** compiled (expression) **
	err = c.compileExpression(st)
	if err != nil {
		return err
	}

	// ')'
	_, err = c.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}

	// ** not        **
	// ** if-goto L2 **
	c.writer.WriteBody("not")
	c.writer.WriteBody("if-goto %v", l2)

	// '{'
	_, err = c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}

	// ** compiled (statements) **
	err = c.compileStatements(st)
	if err != nil {
		return err
	}

	// '}'
	_, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}

	// ** goto L1 **
	// ** label L2 **
	c.writer.WriteBody("goto %v", l1)
	c.writer.WriteLabel(l2)

	return nil
}
