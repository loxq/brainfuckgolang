package bfparser

// A Symbol represents a raw BF token.
type Symbol rune

// An Opcode is an abstract representation of a BF instruction. BF is a simple
// language so there is a direct equivalence between Symbols and Opcodes.
type Opcode string

// BF only has 8 distinct opcodes
// 2 extra Opcodes (noop and invalid) are added for parsing purposes
const (
	OpInvalid   Opcode = "invalid"
	OpNoOp      Opcode = "noop"
	OpIncPtr    Opcode = "incptr"
	OpDecPtr    Opcode = "decptr"
	OpInc       Opcode = "inc"
	OpDec       Opcode = "dec"
	OpWrite     Opcode = "write"
	OpRead      Opcode = "read"
	OpLoopStart Opcode = "loopstart"
	OpLoopStop  Opcode = "loopstop"
)

// An InstructionCode represents one of the 7 actions of a BF program.
type InstructionCode string

// An instruction is an interpretable token.
// Loop instructions contain other instructions, recursively
type Instruction struct {
	Code               InstructionCode
	NestedInstructions []Instruction
}

const (
	InstructionInvalid InstructionCode = "invalid"
	InstructionIncPtr  InstructionCode = "incptr"
	InstructionDecPtr  InstructionCode = "decptr"
	InstructionInc     InstructionCode = "inc"
	InstructionDec     InstructionCode = "dec"
	InstructionWrite   InstructionCode = "write"
	InstructionRead    InstructionCode = "read"
	InstructionLoop    InstructionCode = "loop"
)
