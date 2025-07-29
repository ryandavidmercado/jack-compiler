package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

type RunMode int

const (
	RunModeAnalyze RunMode = iota
	RunModeCompile
)

var runModeMap = map[string]RunMode{
	"analyze": RunModeAnalyze,
	"compile": RunModeCompile,
}

type CliFlags struct {
	filename string
	dirname  string
	mode     RunMode
	inline   bool
}

func parseCliFlags() (*CliFlags, error) {
	cwd, _ := os.Getwd()

	filePtr := flag.String("file", "", "input file for single-file translation (optional; takes priority over dir)")
	dirPtr := flag.String("dir", cwd, "input directory for multi-file translation (default cwd)")
	modePtr := flag.String("mode", "compile", "action to perform (\"analyze\" or \"compile\"). Defaults to \"compile\".")
	inlinePtr := flag.Bool("inline", false, "set to true to target output directory directly instead of /dist.")

	flag.Parse()

	if len(*filePtr) == 0 && len(*dirPtr) == 0 {
		flag.Usage()
		return nil, errors.New("bad args")
	}

	runMode, ok := runModeMap[*modePtr]
	if !ok {
		return nil, fmt.Errorf("Expected run mode \"analyze\" or \"compile\"; got \"%v\"", *modePtr)
	}

	return &CliFlags{filename: *filePtr, dirname: *dirPtr, mode: runMode, inline: *inlinePtr}, nil
}
