package compiler

import "fmt"

func getSymbolFromTables(name string, tables ...*symbolTable) (*symbolInfo, error) {
	for _, table := range tables {
		symbol, ok := table.Get(name)
		if ok {
			return symbol, nil
		}
	}

	return nil, fmt.Errorf("Tried to use undefined variable %v", name)
}
