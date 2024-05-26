package bfparser

import "fmt"

// ToOpcode converts a raw BF symbol to its corresponding Opcode ignoring
// whitespace-like character. An error is returned if any other invalid symbol
// is detected.
func (s Symbol) ToOpcode() (Opcode, error) {
	switch s {
	case '>':
		return OpIncPtr, nil
	case '<':
		return OpDecPtr, nil
	case '+':
		return OpInc, nil
	case '-':
		return OpDec, nil
	case '.':
		return OpWrite, nil
	case ',':
		return OpRead, nil
	case '[':
		return OpLoopStart, nil
	case ']':
		return OpLoopStop, nil
	case '\t', ' ', '\r', '\n':
		return OpNoOp, nil
	default:
		return OpInvalid, fmt.Errorf("invalid symbol in source code: %s", fmt.Sprint(s))
	}
}

// Lex converts a sequence of raw BF tokens to an abstract sequence of instructions.
// This function should become a generator to produce on-the-fly opcodes instead of
// inefficiently building the entire program in memory.
func Lex(symbols string) ([]Opcode, error) {
	operations := []Opcode{}

	for i, symbol := range symbols {
		instruction, err := Symbol(symbol).ToOpcode()
		if err != nil {
			return nil, fmt.Errorf("invalid symbol in source code: found '%s' at index %d", string(symbol), i)
		}
		if instruction != OpNoOp {
			operations = append(operations, instruction)
		}
	}
	return operations, nil
}
