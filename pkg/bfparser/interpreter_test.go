package bfparser

import (
	"testing"
)

type outputBuffer struct {
	buffer string
}

func (b *outputBuffer) captureOutput(s string) string {
	b.buffer += s
	return b.buffer
}

func TestRunBFHappyPath(t *testing.T) {
	bfcode := `++++++++++
[
    >+++++++
    >++++++++++
    >+++
    >+<<<<-
]   >++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.
--------.>+.>.`
	opcodes, _ := Lex(string(bfcode))
	if opcodes != nil {
		program, _ := ParseBF(opcodes)
		if program != nil {
			var output = outputBuffer{}
			RunBF(program, 4096, output.captureOutput)
			if output.buffer != "Hello World!\n" {
				t.Fatalf("Got '%s'. Expected 'Hello World!\\n'", output.buffer)
			}
		}
	}

}
