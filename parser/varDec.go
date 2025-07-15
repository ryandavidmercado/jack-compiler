package parser

import (
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

func (p *Parser) parseVarDec(indent int) (bool, error) {
	nextIndent := indent + common.BaseIndent

	// var
	token, err := p.lexer.ExpectKeyword("var")
	if err != nil {
		p.lexer.TokenIsUnused = true
		return true, nil
	}
	p.writer.WriteOpeningTag("varDec", indent)
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

	p.writer.WriteClosingTag("varDec", indent)
	return false, nil
}
