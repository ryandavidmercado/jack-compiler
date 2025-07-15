package lexer

import (
	"bufio"
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
	"slices"
	"unicode"
)

type Lexer struct {
	reader        *bufio.Reader
	next          *lexerAdvance
	Token         *common.Token
	TokenIsUnused bool
}

func New(reader *bufio.Reader) *Lexer {
	return &Lexer{reader: reader, next: nil, Token: nil}
}

func (l *Lexer) Advance() (*common.Token, error) {
	l.TokenIsUnused = false

	next, err := l.getNext()
	if err != nil {
		return nil, err
	}

	token, err := next("")
	if err != nil {
		return nil, err
	}

	l.Token = token
	return token, nil
}

func (l *Lexer) getToken() (*common.Token, error) {
	var token *common.Token
	var err error

	if l.TokenIsUnused {
		token = l.Token
		l.TokenIsUnused = false
	} else {
		token, err = l.Advance()
		l.Token = token

		if err != nil {
			return token, err
		}
	}

	return token, err
}

func (l *Lexer) Expect(check func(*common.Token) bool) (*common.Token, error) {
	token, err := l.getToken()
	if err != nil {
		return token, err
	}

	if !check(token) {
		return token, fmt.Errorf("Unexpected token %v", token)
	}

	return token, nil
}

func (l *Lexer) ExpectToken(expectedToken *common.Token) (*common.Token, error) {
	token, err := l.getToken()
	if err != nil {
		return token, err
	}

	if !token.Equals(expectedToken) {
		return token, fmt.Errorf("Expected %v, got %v", expectedToken, token)
	}

	return token, nil
}

func (l *Lexer) ExpectSymbol(expectedValue string) (*common.Token, error) {
	return l.ExpectToken(&common.Token{TokenType: common.TTSymbol, Value: expectedValue})
}

func (l *Lexer) ExpectKeyword(expectedValue string) (*common.Token, error) {
	return l.ExpectToken(&common.Token{TokenType: common.TTKeyword, Value: expectedValue})
}

func (l *Lexer) ExpectTokenType(expectedTokenType common.TokenType) (*common.Token, error) {
	token, err := l.getToken()
	if err != nil {
		return token, err
	}

	if token.TokenType != expectedTokenType {
		return token, fmt.Errorf("Expected token type %v, got %v", expectedTokenType, token)
	}

	return token, nil
}

func (l *Lexer) ExpectTokenTypes(expectedTokenTypes []common.TokenType) (*common.Token, error) {
	token, err := l.getToken()
	if err != nil {
		return token, err
	}

	if !slices.Contains(expectedTokenTypes, token.TokenType) {
		return token, fmt.Errorf("Expected one of token types %v, got %v", expectedTokenTypes, token)
	}

	return token, nil
}

type lexerAdvance func(buffer string) (*common.Token, error)

func (l *Lexer) getNext() (lexerAdvance, error) {
	char, _, err := l.reader.ReadRune()

	if err != nil {
		return nil, err
	}

	if unicode.IsSpace(char) {
		return l.getNext()
	}

	if char == '/' {
		// check next char; if it's another /, we're in a comment
		nextChar, _, _ := l.reader.ReadRune()
		switch nextChar {
		case '/':
			return l.advanceComment, nil
		case '*':
			return l.advanceMultilineComment, nil
		default:
			l.reader.UnreadRune()
			return l.advanceReadBackslash, nil
		}
	}

	if char == '"' {
		return l.advanceStringConstant, nil
	} else if unicode.IsDigit(char) {
		l.reader.UnreadRune()
		return l.advanceIntegerConstant, nil
	} else if slices.Contains(symbols, char) {
		l.reader.UnreadRune()
		return l.advanceSymbol, nil
	} else {
		l.reader.UnreadRune()
		return l.advanceWord, nil
	}
}

func (l *Lexer) advanceStringConstant(buffer string) (*common.Token, error) {
	char, _, err := l.reader.ReadRune()

	if err != nil {
		return nil, err
	}

	if char == '"' {
		return &common.Token{TokenType: common.TTStringConstant, Value: buffer}, nil
	}

	return l.advanceStringConstant(buffer + string(char))
}

func (l *Lexer) advanceIntegerConstant(buffer string) (*common.Token, error) {
	char, _, err := l.reader.ReadRune()

	if err != nil {
		return nil, err
	}

	if !unicode.IsDigit(char) {
		l.reader.UnreadRune()
		return &common.Token{TokenType: common.TTIntegerConstant, Value: buffer}, nil
	}

	return l.advanceIntegerConstant(buffer + string(char))
}

func (l *Lexer) advanceWord(buffer string) (*common.Token, error) {
	char, _, err := l.reader.ReadRune()

	if err != nil {
		return nil, err
	}

	if !(unicode.IsDigit(char) || unicode.IsLetter(char) || char == '_') || unicode.IsSpace(char) {
		l.reader.UnreadRune()

		if _, isKeyword := common.StringToKeyword[buffer]; isKeyword {
			return &common.Token{TokenType: common.TTKeyword, Value: buffer}, nil
		} else {
			return &common.Token{TokenType: common.TTIdentifier, Value: buffer}, nil
		}
	}

	return l.advanceWord(buffer + string(char))
}

func (l *Lexer) advanceSymbol(buffer string) (*common.Token, error) {
	char, _, err := l.reader.ReadRune()
	if err != nil {
		return nil, err
	}

	return &common.Token{TokenType: common.TTSymbol, Value: string(char)}, nil
}

func (l *Lexer) advanceReadBackslash(_ string) (*common.Token, error) {
	return &common.Token{TokenType: common.TTSymbol, Value: "/"}, nil
}

func (l *Lexer) advanceComment(_ string) (*common.Token, error) {
	char, _, err := l.reader.ReadRune()
	if err != nil {
		return nil, err
	}

	if char == '\n' {
		next, err := l.getNext()
		if err != nil {
			return nil, err
		}

		return next("")
	}

	return l.advanceComment("")
}

func (l *Lexer) advanceMultilineComment(_ string) (*common.Token, error) {
	char, _, err := l.reader.ReadRune()
	if err != nil {
		return nil, err
	}

	if char == '*' {
		nextChar, _, err := l.reader.ReadRune()
		if err != nil {
			return nil, err
		}

		if nextChar == '/' {
			next, err := l.getNext()
			if err != nil {
				return nil, err
			}

			return next("")
		}
	}

	return l.advanceMultilineComment("")
}
