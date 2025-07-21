package compiler

import (
	"fmt"
	"maps"
	"slices"
)

type symbolInfo struct {
	symbolType string
	kind       string
	number     uint16
}

type symbolTable struct {
	symbols map[string]*symbolInfo
	numbers map[string]uint16
}

func NewSymbolTable(kinds []string) *symbolTable {
	st := &symbolTable{
		symbols: map[string]*symbolInfo{},
		numbers: map[string]uint16{},
	}

	for _, kind := range kinds {
		st.numbers[kind] = 0
	}

	return st
}

func (st *symbolTable) Add(name string, symbolType string, kind string) error {
	number, ok := st.numbers[kind]
	if !ok {
		kinds := slices.Collect(maps.Keys(st.numbers))
		return fmt.Errorf("Tried to add variable %v with kind %v to symbol table with kinds %v", name, kind, kinds)
	}

	st.symbols[name] = &symbolInfo{symbolType, kind, number}
	st.numbers[kind] += 1

	return nil
}

func (st *symbolTable) Get(name string) (*symbolInfo, bool) {
	symbol, ok := st.symbols[name]
	return symbol, ok
}
