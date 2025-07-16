package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileVarDec(st *symbolTable) error {
	// var
	_, err := c.lexer.ExpectKeyword("var")
	if err != nil {
		return err
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

	st.Add(varName.Value, varType.Value, "local")

	// (, varName)* OR ';'
	for {
		// , OR ;
		sep, err := c.lexer.Expect(func(t *common.Token) bool {
			return t.TokenType == common.TTSymbol &&
				(t.Value == ";" || t.Value == ",")
		})

		if err != nil {
			return err
		}

		if sep.Value == ";" {
			break
		}

		// varName
		varName, err = c.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}

		st.Add(varName.Value, varType.Value, "local")
	}

	return nil
}
