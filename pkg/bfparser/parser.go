package bfparser

import "fmt"

// Convert opcode to instruction
func (o Opcode) ToInstruction() (Instruction, error) {
	switch o {
	case OpIncPtr:
		return Instruction{Code: InstructionIncPtr}, nil
	case OpDecPtr:
		return Instruction{Code: InstructionDecPtr}, nil
	case OpInc:
		return Instruction{Code: InstructionInc}, nil
	case OpDec:
		return Instruction{Code: InstructionDec}, nil
	case OpWrite:
		return Instruction{Code: InstructionWrite}, nil
	case OpRead:
		return Instruction{Code: InstructionRead}, nil
	default:
		return Instruction{Code: InstructionInvalid}, fmt.Errorf("invalid opcode")
	}
}

// ParseBF converts Opcodes to interpretable instructions.
// Loop instructions are parsed recursively.
// Loops could be nested up to Go runtime's available stack size.
// ParseBF returns an error if loop opcodes are not well balanced.
func ParseBF(operations []Opcode) ([]Instruction, error) {
	program := []Instruction{}
	var loopStart int
	var loopStack int

	for i, operation := range operations {
		if loopStack == 0 {
			// We are at the top level of the program.

			if operation == OpLoopStart {
				// Track loop starting point.
				loopStart = i
				// Count nested loops.
				loopStack += 1
			} else if operation == OpLoopStop {
				return nil, fmt.Errorf("unbalanced loop end at index %d", i)
			} else {
				instruction, err := operation.ToInstruction()
				if err != nil {
					return nil, fmt.Errorf("error parsing opcode at index %d: %w", i, err)
				}
				program = append(program, instruction)
			}
		} else {
			// We are inside a loop.
			if operation == OpLoopStart {
				// Dive inside a nested loop
				loopStack++
			} else if operation == OpLoopStop {
				// Return from a loop
				loopStack--

				if loopStack == 0 {
					// Parse nested loop once we've returned to the caller loop
					nestedOperations := operations[loopStart+1 : i]
					nestedInstructions, err := ParseBF(nestedOperations)
					if err != nil {
						return nil, fmt.Errorf("error parsing loop between index %d and index %d: %w", loopStart, i, err)
					}
					instruction := Instruction{Code: InstructionLoop, NestedInstructions: nestedInstructions}
					program = append(program, instruction)
				}
			}
		}
	}
	if loopStack != 0 {
		return nil, fmt.Errorf("unbalanced loop end at index %d", loopStart)
	}

	return program, nil
}
