package parser

import (
	"bufio"
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

type parserWriter struct {
	*bufio.Writer
}

func (pw *parserWriter) WriteToken(token *common.Token, indent int) error {
	_, err := pw.WriteString(common.Indent(token.XML(), indent) + "\n")
	return err
}

func (pw *parserWriter) WriteWithIndent(s string, indent int) error {
	_, err := pw.WriteString(common.Indent(s, indent) + "\n")
	return err
}

func (pw *parserWriter) WriteOpeningTag(s string, indent int) error {
	tag := fmt.Sprintf("<%v>\n", s)
	_, err := pw.WriteString(common.Indent(tag, indent))
	return err
}

func (pw *parserWriter) WriteClosingTag(s string, indent int) error {
	tag := fmt.Sprintf("</%v>\n", s)
	_, err := pw.WriteString(common.Indent(tag, indent))
	return err
}
