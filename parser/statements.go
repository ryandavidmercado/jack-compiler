package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseStatements(indent int) error {
	p.writer.WriteOpeningTag("statements", indent)
	nextIndent := indent + common.BaseIndent

	// statement*
	for done, err := p.parseStatement(nextIndent); !done || err != nil; done, err = p.parseStatement(nextIndent) {
		if err != nil {
			return err
		}
	}

	p.writer.WriteClosingTag("statements", indent)
	return nil
}
