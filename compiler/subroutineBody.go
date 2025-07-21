package compiler

func (c *Compiler) compileSubroutineBody(st *symbolTable, subroutineName string, subroutineType string) error {
	// {
	_, err := c.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}

	// varDec*
	for {
		_, err := c.lexer.ExpectKeyword("var")
		c.lexer.TokenIsUnused = true

		if err == nil {
			err := c.compileVarDec(st)
			if err != nil {
				return err
			}
		} else {
			break
		}
	}

	// now that we've gathered local variables, we have the info we need to write the VM function declaration
	c.writer.WriteFunctionDeclaration(c.className, subroutineName, st.numbers["local"])

	switch subroutineType {
	case "constructor":
		// if we're compiling a constructor, make sure it handles allocating the new object when called
		instanceFields := c.symbolTable.numbers["field"]
		c.writer.WritePushConstant(int16(instanceFields))
		c.writer.WriteFunctionCall("Memory.alloc", 1)
		c.writer.WritePop("pointer", 0)
	case "method":
		// if we're compiling a method, make sure we align THIS w/ the callee object
		c.writer.WritePush("argument", 0)
		c.writer.WritePop("pointer", 0)
	}

	// statements
	err = c.compileStatements(st)
	if err != nil {
		return err
	}

	// }
	_, err = c.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}

	return nil
}
