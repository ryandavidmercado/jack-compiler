package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseParameterList(indent int) error {
	p.writer.WriteOpeningTag("parameterList", indent)
	nextIndent := indent + common.BaseIndent

	// type
	err := p.parseType(false, nextIndent)
	if err != nil {
		p.lexer.TokenIsUnused = true
		p.writer.WriteClosingTag("parameterList", indent)
		return nil // terminate early; we don't have list contents
	}

	// varName
	token, err := p.lexer.ExpectTokenType(common.TTIdentifier)
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	for {
		// , (optional)
		token, err := p.lexer.ExpectSymbol(",")
		if err != nil {
			p.lexer.TokenIsUnused = true // we'll reuse this to check ;
			break                        // exit; no optional varName list here
		}
		p.writer.WriteToken(token, nextIndent)

		// type
		err = p.parseType(false, nextIndent)
		if err != nil {
			return err
		}

		// varName
		token, err = p.lexer.ExpectTokenType(common.TTIdentifier)
		if err != nil {
			return err
		}
		p.writer.WriteToken(token, nextIndent)
	}

	p.writer.WriteClosingTag("parameterList", indent)
	return nil
}
