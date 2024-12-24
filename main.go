package main

import (
	"github.com/atom-sgt/compiler-go/lexer"
)

func main() {
	input := "int main() { int a = 42; a++; }"
	lexer.GetTokens(input)
}
