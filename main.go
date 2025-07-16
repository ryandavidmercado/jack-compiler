package main

import (
	"bufio"
	"log"
	"os"
	"path"
	"strings"

	"github.com/ryandavidmercado/jack-compiler/compiler"
	"github.com/ryandavidmercado/jack-compiler/lexer"
	"github.com/ryandavidmercado/jack-compiler/parser"
)

func main() {
	flags, err := parseCliFlags()
	if err != nil {
		log.Fatal(err)
	}

	files := []string{}

	if len(flags.filename) != 0 {
		files = append(files, flags.filename)
	} else {
		dirEntries, err := os.ReadDir(flags.dirname)
		if err != nil {
			log.Fatal(err)
		}

		for _, dirEntry := range dirEntries {
			if dirEntry.IsDir() {
				continue
			}

			if path.Ext(dirEntry.Name()) != ".jack" {
				continue
			}

			path := path.Join(flags.dirname, dirEntry.Name())
			files = append(files, path)
		}
	}

	for _, file := range files {
		inputName := path.Base(file)
		var outputName string

		log.SetFlags(log.Flags())

		input, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		log.SetFlags(0)

		lexer := lexer.New(bufio.NewReader(input))

		switch flags.mode {
		case RunModeAnalyze:
			outputPath := strings.TrimSuffix(file, path.Ext(file)) + ".xml"
			output, fileerr := os.Create(outputPath)
			if fileerr != nil {
				log.Fatal(fileerr)
			}
			outputName = path.Base(outputPath)

			writer := bufio.NewWriter(output)
			parser := parser.New(lexer, writer)

			err = parser.Parse()

			writer.Flush()
			output.Close()
		case RunModeCompile:
			outputPath := strings.TrimSuffix(file, path.Ext(file)) + ".vm"
			output, fileerr := os.Create(outputPath)
			if fileerr != nil {
				log.Fatal(fileerr)
			}
			outputName = path.Base(outputPath)

			writer := bufio.NewWriter(output)
			compiler := compiler.New(lexer, writer)

			err = compiler.Compile()

			writer.Flush()
			output.Close()
		}

		if err != nil {
			log.Printf("✖ | %v → %v\n\t%v", inputName, outputName, err)
		} else {
			log.Printf("✔ | %v → %v", inputName, outputName)
		}
	}
}
