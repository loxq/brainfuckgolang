package bfparser

import (
	"fmt"
	"strings"
	"testing"
)

func TestLexHappyPath(t *testing.T) {
	program, err := Lex("><+- .\r\n\t,[]")
	if err != nil {
		t.Errorf("Lex() error")
	}
	expected := []Opcode{OpIncPtr, OpDecPtr, OpInc, OpDec, OpWrite, OpRead, OpLoopStart, OpLoopStop}

	for i, gotToken := range program {
		expectedToken := expected[i]

		if gotToken != expectedToken {
			t.Fatalf("Lex() : invalid instruction at index %d. Got '%s' expected '%s'.", i, fmt.Sprint(gotToken), fmt.Sprint(expectedToken))
		}
	}
}

func TestLexReturnsAnErrorOnInvalidCharacter(t *testing.T) {
	_, err := Lex("><+- 8.,[]#")
	if err == nil {
		t.Fatalf("Lex() didn't return an error as expected")
	}
	if !strings.Contains(err.Error(), "found '8' at index 5") {
		t.Fatalf("Lex() bad error message. Expected \"[...] found '8' at index 5 [...]\". Got \"%s\"", err.Error())
	}
}

func BenchmarkLex(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Lex("><+- 8.,[]#")
	}
}
