package common

import (
	"fmt"
	"strings"
)

type TokenType int

const (
	TTKeyword TokenType = iota
	TTSymbol
	TTIdentifier
	TTIntegerConstant
	TTStringConstant
)

func (tt TokenType) String() string {
	return map[TokenType]string{
		TTKeyword:         "keyword",
		TTSymbol:          "symbol",
		TTIdentifier:      "identifier",
		TTIntegerConstant: "integerConstant",
		TTStringConstant:  "stringConstant",
	}[tt]
}

type Token struct {
	TokenType
	Value string
}

func (t *Token) XML() string {
	value := t.Value

	// replace banned XML tokens w/ escape sequences
	value = strings.ReplaceAll(value, "&", "&amp;")
	value = strings.ReplaceAll(value, "<", "&lt;")
	value = strings.ReplaceAll(value, ">", "&gt;")
	value = strings.ReplaceAll(value, "\"", "&quot;")

	return fmt.Sprintf("<%[1]v> %[2]v </%[1]v>", t.TokenType, value)
}

func (t *Token) String() string {
	return t.XML()
}

func (t *Token) Equals(other *Token) bool {
	return t.TokenType == other.TokenType && t.Value == other.Value
}
