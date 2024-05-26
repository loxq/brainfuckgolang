package bfparser

import (
	"fmt"
	"os"
)

func defaultPrinter(s string) string {
	fmt.Print(s)
	return s
}

// RunBF takes a BF program and follows the execution code.
// Output is written to stdout. Input is read from Stdin.
// tapeSize is the size of the data memory for the machine. This tape should be
// sufficiently large to make complex programs possible.
//
// outputPrint is a function for write characters. If outputPrint is nil,
// consolePrinter is used.
func RunBF(program []Instruction, tapeSize int, outputPrint func(string) string) {
	var tape = make([]byte, tapeSize)
	var dataPointer int = tapeSize / 2

	if outputPrint == nil {
		outputPrint = defaultPrinter
	}

	var interpreter func(instructions []Instruction, tape []byte, dataPointer int)
	interpreter = func(instructions []Instruction, tape []byte, dataPointer int) {
		var readBuffer []byte = make([]byte, 1)
		for _, instruction := range instructions {
			switch instruction.Code {
			case InstructionIncPtr:
				dataPointer++
				if dataPointer >= tapeSize {
					panic("tape overflow error !")
				}
			case InstructionDecPtr:
				dataPointer--
				if dataPointer >= tapeSize {
					panic("tape overflow error !")
				}
			case InstructionInc:
				tape[dataPointer]++
			case InstructionDec:
				tape[dataPointer]--
			case InstructionWrite:
				outputPrint(string(tape[dataPointer]))
			case InstructionRead:
				os.Stdin.Read(readBuffer)
				tape[dataPointer] = readBuffer[0]
			case InstructionLoop:
				for tape[dataPointer] != 0 {
					interpreter(instruction.NestedInstructions, tape, dataPointer)
				}
			}
		}
	}
	interpreter(program, tape, dataPointer)
}
