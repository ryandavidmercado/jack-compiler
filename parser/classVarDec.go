package parser

import (
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (p *Parser) parseClassVarDec(indent int) (bool, error) {
	nextIndent := indent + common.BaseIndent

	// static | field
	token, err := p.lexer.Expect(func(t *common.Token) bool {
		return t.TokenType == common.TTKeyword && (t.Value == "static" || t.Value == "field")
	})
	if err != nil {
		p.lexer.TokenIsUnused = true
		return true, nil
	}
	p.writer.WriteOpeningTag("classVarDec", indent)
	p.writer.WriteToken(token, nextIndent)

	// type
	err = p.parseType(false, nextIndent)
	if err != nil {
		return true, err
	}

	// varName
	token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return true, err
	}
	p.writer.WriteToken(token, nextIndent)

	// (, varName)* OR ';'
	for {
		// , OR ;
		token, err = p.lexer.ExpectTokenType(common.TTSymbol)
		if err != nil {
			return true, err
		}
		p.writer.WriteToken(token, nextIndent)

		if token.Value == ";" {
			break
		} else if token.Value != "," {
			return true, fmt.Errorf("Expected symbol ';' or ',', got %v", token)
		}

		// varName
		token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return true, err
		}
		p.writer.WriteToken(token, nextIndent)
	}

	p.writer.WriteClosingTag("classVarDec", indent)
	return false, nil
}
