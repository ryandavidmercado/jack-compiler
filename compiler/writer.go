package compiler

import (
	"bufio"
	"fmt"
)

type compilerWriter struct {
	*bufio.Writer
}

func (cw *compilerWriter) WriteBody(str string, args ...any) error {
	_, err := cw.WriteString(
		"    " +
			fmt.Sprintf(str, args...) +
			"\n",
	)

	return err
}

func (cw *compilerWriter) WriteLabel(label string) error {
	_, err := fmt.Fprintf(cw, "label %v\n", label)

	return err
}

func (cw *compilerWriter) WriteFunctionDeclaration(className string, subroutineName string, localVarCount uint16) error {
	_, err := fmt.Fprintf(cw, "function %v.%v %d\n", className, subroutineName, localVarCount)

	return err
}

func (cw *compilerWriter) WriteFunctionCall(functionName string, argCount uint8) error {
	return cw.WriteBody("call %v %d", functionName, argCount)
}

func (cw *compilerWriter) WritePushSymbol(symbol *symbolInfo) error {
	kind := symbol.kind
	if kind == "field" {
		kind = "this"
	}

	return cw.WriteBody("push %v %d", kind, symbol.number)
}

func (cw *compilerWriter) WritePopSymbol(symbol *symbolInfo) error {
	kind := symbol.kind
	if kind == "field" {
		kind = "this"
	}

	return cw.WriteBody("pop %v %d", kind, symbol.number)
}

func (cw *compilerWriter) WritePushConstant(num int16) error {
	asPositive := num
	isNegative := false
	if num < 0 {
		asPositive *= -1
		isNegative = true
	}

	err := cw.WriteBody("push constant %d", asPositive)
	if err != nil {
		return err
	}

	if isNegative {
		err := cw.WriteBody("neg")
		if err != nil {
			return err
		}
	}

	return nil
}

func (cw *compilerWriter) WriteStringConstant(str string) error {
	err := cw.WriteBody("call String.new 0")
	if err != nil {
		return err
	}

	// do the rest

	return nil
}

var unaryOpOutputs = map[string]string{
	"-": "neg",
	"~": "not",
}

var binaryOpOutputs = map[string]string{
	"+": "add",
	"-": "sub",
	"*": "call Math.multiply 2",
	"/": "call Math.divide 2",
	"&": "and",
	"|": "or",
	"<": "lt",
	">": "gt",
	"=": "eq",
}

func (cw *compilerWriter) WriteUnaryOp(op string) error {
	output, ok := unaryOpOutputs[op]
	if !ok {
		return fmt.Errorf("Unexpected unary op %v", op)
	}

	return cw.WriteBody("%s", output)
}

func (cw *compilerWriter) WriteBinaryOp(op string) error {
	output, ok := binaryOpOutputs[op]
	if !ok {
		return fmt.Errorf("Unexpected binary op %v", op)
	}

	return cw.WriteBody("%s", output)
}

var keywordConstantOutputs = map[string]string{
	"true":  "push constant 1\n    neg",
	"false": "push constant 0",
	"null":  "push constant 0",
	"this":  "push pointer 0",
}

func (cw *compilerWriter) WriteKeywordConstant(keyword string) error {
	output, ok := keywordConstantOutputs[keyword]
	if !ok {
		return fmt.Errorf("Unexpected keyword constant %v", keyword)
	}

	return cw.WriteBody("%s", output)
}

func (cw *compilerWriter) WriteReturn() error {
	return cw.WriteBody("return")
}

func (cw *compilerWriter) WritePush(segment string, number int) error {
	return cw.WriteBody("push %v %d", segment, number)
}

func (cw *compilerWriter) WritePop(segment string, number int) error {
	return cw.WriteBody("pop %v %d", segment, number)
}
