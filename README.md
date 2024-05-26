# Brainfuck Go Interpreter
A Brainfuck interpreter in Go

## Usage
BF code example :
```Brainfuck
++++++++++
[
    >+++++++
    >++++++++++
    >+++
    >+<<<<-
]   >++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.
--------.>+.>.
```

```bash
go build ./cmd/bfgo/

./bfgo helloworld.bf
```

```
Hello World!
```

## Brainfuck

Brainfuck is a simple language for educational purposes. With only 8 instructions it can perform any computation that a Turing machine can (given an infinite memory space).

The language operates on an array of memory cells, all initially set to zero, with a data pointer that can move left and right. Commands increment or decrement the data pointer or the byte at the data pointer, output or input a byte, and create loops for control flow.