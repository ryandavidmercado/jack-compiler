package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileType(includeVoid bool, indent int) error {
	validKeywords := map[string]struct{}{
		"int":     {},
		"char":    {},
		"boolean": {},
	}

	if includeVoid {
		validKeywords["void"] = struct{}{}
	}

	token, err := c.lexer.Expect(func(t *common.Token) bool {
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

	c.writer.WriteToken(token, indent)

	return nil
}
