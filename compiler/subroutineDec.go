package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileSubroutineDec() error {
	// constructor | function | method
	subroutineType, err := c.lexer.Expect(func(t *common.Token) bool {
		return t.TokenType == common.TTKeyword && (t.Value == "constructor" || t.Value == "function" || t.Value == "method")
	})
	if err != nil {
		return err
	}

	// ------------- Initialize Subroutine Symbol Table --------------
	symbolTable := NewSymbolTable([]string{"argument", "local"})
	if subroutineType.Value == "method" {
		symbolTable.Add("this", c.className, "argument")
	}
	// ---------------------------------------------------------------

	// type || 'void'
	_, err = c.compileType(true)
	if err != nil {
		return err
	}

	// subroutineName
	subroutineName, err := c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}

	// '(';
	_, err = c.lexer.ExpectSymbol("(")
	if err != nil {
		return err
	}

	// parameterList
	err = c.compileParameterList(symbolTable)
	if err != nil {
		return err
	}

	// ')';
	_, err = c.lexer.ExpectSymbol(")")
	if err != nil {
		return err
	}

	// subroutineBody
	err = c.compileSubroutineBody(symbolTable, subroutineName.Value, subroutineType.Value)
	if err != nil {
		return err
	}

	return nil
}
