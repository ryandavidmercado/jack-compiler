package parser

import "github.com/ryandavidmercado/jack-compiler/common"

func (p *Parser) parseType(includeVoid bool, indent int) error {
	validKeywords := map[string]struct{}{
		"int":     {},
		"char":    {},
		"boolean": {},
	}

	if includeVoid {
		validKeywords["void"] = struct{}{}
	}

	token, err := p.lexer.Expect(func(t *common.Token) bool {
		switch t.TokenType {
		case common.TTKeyword:
			_, valid := validKeywords[t.Value]
			return valid
		case common.TTIdentifier:
			return true
		default:
			return false
		}
	})
	if err != nil {
		return err
	}

	p.writer.WriteToken(token, indent)

	return nil
}
