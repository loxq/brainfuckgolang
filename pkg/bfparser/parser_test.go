package bfparser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func TestParseBFHappyPath(t *testing.T) {
	input := []Opcode{
		OpIncPtr,
		OpDecPtr,
		OpInc,
		OpDec,
		OpRead,
		OpWrite,

		OpInc, OpInc,
		OpLoopStart,
		OpDec,
		OpDec,
		OpLoopStop,
	}
	expected := []Instruction{
		{Code: InstructionIncPtr},
		{Code: InstructionDecPtr},
		{Code: InstructionInc},
		{Code: InstructionDec},
		{Code: InstructionRead},
		{Code: InstructionWrite},
		{Code: InstructionInc},
		{Code: InstructionInc},
		{Code: InstructionLoop,
			NestedInstructions: []Instruction{
				{Code: InstructionDec},
				{Code: InstructionDec},
			},
		},
	}
	output, err := ParseBF(input)
	if err != nil {
		t.Fail()
	}

	if !reflect.DeepEqual(output, expected) {
		t.Fatalf("instructions slice doesn't match expected slice. Got '%s' expected '%s'", fmt.Sprint(output), fmt.Sprint(expected))
	}

}

func TestParseBFReturnsAnErrorOnUnbalancedLoopOpcodes(t *testing.T) {
	input := []Opcode{
		OpLoopStart,
		OpLoopStart,
		OpDec,
		OpDec,
		OpLoopStop, // missing second OpLoopStop
	}
	_, err := ParseBF(input)
	if err == nil {
		t.Fatalf("ParseBF() didn't return an error as expected")
	}
	if !strings.Contains(err.Error(), "unbalanced loop end at index") {
		t.Fatalf("ParseBF() didn't return the expected unbalanced loop error message")
	}

	input2 := []Opcode{
		OpDec,
		OpDec,
		OpLoopStop, // missing matching OpLoopStart
	}
	_, err2 := ParseBF(input2)
	if err2 == nil {
		t.Fatalf("ParseBF() didn't return an error as expected")
	}
	if !strings.Contains(err2.Error(), "unbalanced loop end at index") {
		t.Fatalf("ParseBF() didn't return the expected unbalanced loop error message")
	}

}

func BenchmarkParseBF(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ParseBF([]Opcode{OpInc, OpDec, OpIncPtr, OpDecPtr})
	}
}
