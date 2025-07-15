package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseSubroutineDec(indent int) (bool, error) {
	nextIndent := indent + common.BaseIndent

	// constructor | function | method
	token, err := p.lexer.Expect(func(t *common.Token) bool {
		return t.TokenType == common.TTKeyword && (t.Value == "constructor" || t.Value == "function" || t.Value == "method")
	})
	if err != nil {
		p.lexer.TokenIsUnused = true
		return true, nil
	}

	p.writer.WriteOpeningTag("subroutineDec", indent)
	p.writer.WriteToken(token, nextIndent)

	// type || 'void'
	err = p.parseType(true, nextIndent)
	if err != nil {
		return true, err
	}

	// subroutineName
	token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return true, err
	}
	p.writer.WriteToken(token, nextIndent)

	// '(';
	token, err = p.lexer.ExpectSymbol("(")
	if err != nil {
		return true, err
	}
	p.writer.WriteToken(token, nextIndent)

	// parameterList
	err = p.parseParameterList(nextIndent)
	if err != nil {
		return true, err
	}

	// ')';
	token, err = p.lexer.ExpectSymbol(")")
	if err != nil {
		return true, err
	}
	p.writer.WriteToken(token, nextIndent)

	// subroutineBody
	err = p.parseSubroutineBody(nextIndent)
	if err != nil {
		return true, err
	}

	p.writer.WriteClosingTag("subroutineDec", indent)
	return false, nil
}
