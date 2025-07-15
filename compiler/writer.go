package compiler

import (
	"bufio"
	"fmt"
	"github.com/ryandavidmercado/jack-compiler/common"
)

type compilerWriter struct {
	*bufio.Writer
}

func (cw *compilerWriter) WriteToken(token *common.Token, indent int) error {
	_, err := cw.WriteString(common.Indent(token.XML(), indent) + "\n")
	return err
}

func (cw *compilerWriter) WriteWithIndent(s string, indent int) error {
	_, err := cw.WriteString(common.Indent(s, indent) + "\n")
	return err
}

func (cw *compilerWriter) WriteOpeningTag(s string, indent int) error {
	tag := fmt.Sprintf("<%v>\n", s)
	_, err := cw.WriteString(common.Indent(tag, indent))
	return err
}

func (cw *compilerWriter) WriteClosingTag(s string, indent int) error {
	tag := fmt.Sprintf("</%v>\n", s)
	_, err := cw.WriteString(common.Indent(tag, indent))
	return err
}
