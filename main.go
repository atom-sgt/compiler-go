package main

import (
	"github.com/atom-sgt/compiler-go/lexer"
	"github.com/atom-sgt/compiler-go/parser"
)

func main() {
	// input := "int main() { int a = 42; a++; }"
	input := "1 + 2"
	println(input)

	lexer := lexer.NewTokenizer(input)
	parser := parser.NewParser(lexer)
	parser.Parse()
}
