package compiler

func (c *Compiler) compileExpressionList(st *symbolTable) (uint8, error) {
	var count uint8 = 0

	// expression
	err := c.compileExpression(st)
	if err != nil {
		// expression is optional here; relinquish token and return
		c.lexer.TokenIsUnused = true
		return 0, nil
	}
	count += 1

	// (, expression)*
	for {
		_, err := c.lexer.ExpectSymbol(",")
		if err != nil {
			// no more expressions, break
			c.lexer.TokenIsUnused = true
			break
		}

		err = c.compileExpression(st)
		if err != nil {
			// we should have gotten an expression here; fail
			return 0, err
		}
		count += 1
	}

	return count, nil
}
