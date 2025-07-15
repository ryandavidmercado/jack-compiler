package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseDoStatement(indent int) error {
	p.writer.WriteOpeningTag("doStatement", indent)
	nextIndent := indent + common.BaseIndent

	// 'do'
	token, err := p.lexer.ExpectKeyword("do")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	// subroutineCall
	err = p.parseSubroutineCall(nextIndent, nil, nil)
	if err != nil {
		return err
	}

	// ';'
	token, err = p.lexer.ExpectSymbol(";")
	if err != nil {
		return err
	}
	p.writer.WriteToken(token, nextIndent)

	p.writer.WriteClosingTag("doStatement", indent)
	return nil
}
