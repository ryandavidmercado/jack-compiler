package compiler

import "github.com/ryandavidmercado/jack-compiler/common"

func (c *Compiler) compileStatements(indent int) error {
	c.writer.WriteOpeningTag("statements", indent)
	nextIndent := indent + common.BaseIndent

	// statement*
	for done, err := c.compileStatement(nextIndent); !done || err != nil; done, err = c.compileStatement(nextIndent) {
		if err != nil {
			return err
		}
	}

	c.writer.WriteClosingTag("statements", indent)
	return nil
}
