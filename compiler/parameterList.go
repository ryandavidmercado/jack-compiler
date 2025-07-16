package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileParameterList(st *symbolTable) error {
	// type?
	varType, err := c.compileType(false)
	if err != nil {
		c.lexer.TokenIsUnused = true
		return nil
	}

	// varName
	varName, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	st.Add(varName.Value, varType.Value, "argument")

	for {
		// ,?
		_, err := c.lexer.ExpectSymbol(",")
		if err != nil {
			c.lexer.TokenIsUnused = true // we'll reuse this to check ;
			break                        // exit; no optional varName list here
		}

		// type
		varType, err := c.compileType(false)
		if err != nil {
			return err
		}

		// varName
		varName, err := c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}

		st.Add(varName.Value, varType.Value, "argument")
	}

	return nil
}
