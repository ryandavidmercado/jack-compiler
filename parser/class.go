package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseClass(indent int) error {
	p.writer.WriteOpeningTag("class", indent)
	nextIndent := indent + common.BaseIndent

	// 'class'
	token, err := p.lexer.ExpectKeyword("class")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// className
	token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// '{'
	token, err = p.lexer.ExpectSymbol("{")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// classVarDec*
	for done, err := p.parseClassVarDec(nextIndent); !done || err != nil; done, err = p.parseClassVarDec(nextIndent) {
		if err != nil {
			return err
		}
	}

	// classSubroutineDec*
	for done, err := p.parseSubroutineDec(nextIndent); !done || err != nil; done, err = p.parseSubroutineDec(nextIndent) {
		if err != nil {
			return err
		}
	}

	// '}'
	token, err = p.lexer.ExpectSymbol("}")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	p.writer.WriteClosingTag("class", indent)
	return nil
}
