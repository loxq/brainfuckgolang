package main

import (
	"errors"
	"fmt"
	"os"

	"bfgo/pkg/bfparser"
)

const TAPESIZE int = 4096

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s usage: \n\t%s filename\n", os.Args[0], os.Args[0])
		os.Exit(22)
	}
	filename := os.Args[1]
	// Check file exists
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Printf("no such file '%s'\n", filename)
		os.Exit(2)
	}

	// Read BF source code from file.
	bfcode, err := os.ReadFile(filename)
	if err != nil {
		panic(fmt.Errorf("cannot read file '%s': %w", filename, err))
	}

	// Convert BF symbols to opcodes and clean whitespaces.
	// Returns an error on invalid symbols.
	opcodes, err := bfparser.Lex(string(bfcode))
	if err != nil {
		panic(fmt.Errorf("cannot analyse file './helloworld.bf': %w", err))
	}

	// Convert opcodes to a tree of instructions, maybe recursive
	program, err := bfparser.ParseBF(opcodes)
	if err != nil {
		panic(fmt.Errorf("cannot parse file './helloworld.bf': %w", err))
	}

	// Run the BF code.
	// OutputPrint is nil to keep the default ConsolePrinter.
	bfparser.RunBF(program, TAPESIZE, nil)

}
