package main

import (
	"bufio"
	"github.com/ryandavidmercado/jack-compiler/lexer"
	"github.com/ryandavidmercado/jack-compiler/parser"
	"log"
	"os"
	"path"
	"strings"
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
		log.SetFlags(log.Flags())

		input, err := os.Open(file)
		if err != nil {
			log.Fatal(err)
		}

		log.SetFlags(0)

		outputPath := strings.TrimSuffix(file, path.Ext(file)) + ".xml"
		output, err := os.Create(outputPath)

		writer := bufio.NewWriter(output)

		lexer := lexer.New(bufio.NewReader(input))
		parser := parser.New(lexer, writer)

		err = parser.Parse()

		inputName := path.Base(file)
		outputName := path.Base(outputPath)

		writer.Flush()
		output.Close()
		if err != nil {
			log.Printf("✖ | %v → %v\n\t%v", inputName, outputName, err)
		}

		log.Printf("✔ | %v → %v", inputName, outputName)
	}
}
