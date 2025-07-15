package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileSubroutineDec(indent int) (bool, error) {
	nextIndent := indent + common.BaseIndent

	// constructor | function | method
	token, err := c.lexer.Expect(func(t *common.Token) bool {
		return t.TokenType == common.TTKeyword && (t.Value == "constructor" || t.Value == "function" || t.Value == "method")
	})
	if err != nil {
		c.lexer.TokenIsUnused = true
		return true, nil
	}

	c.writer.WriteOpeningTag("subroutineDec", indent)
	c.writer.WriteToken(token, nextIndent)

	// type || 'void'
	err = c.compileType(true, nextIndent)
	if err != nil {
		return true, err
	}

	// subroutineName
	token, err = c.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return true, err
	}
	c.writer.WriteToken(token, nextIndent)

	// '(';
	token, err = c.lexer.ExpectSymbol("(")
	if err != nil {
		return true, err
	}
	c.writer.WriteToken(token, nextIndent)

	// parameterList
	err = c.compileParameterList(nextIndent)
	if err != nil {
		return true, err
	}

	// ')';
	token, err = c.lexer.ExpectSymbol(")")
	if err != nil {
		return true, err
	}
	c.writer.WriteToken(token, nextIndent)

	// subroutineBody
	err = c.compileSubroutineBody(nextIndent)
	if err != nil {
		return true, err
	}

	c.writer.WriteClosingTag("subroutineDec", indent)
	return false, nil
}
