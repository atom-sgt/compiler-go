package main

import (
	"fmt"

	"github.com/atom-sgt/compiler-go/lexer"
)

func main() {
	input := "int main() { int a = 42; a++; }"
	println(input)

	tokens := lexer.GetTokens(input)

	for _, token := range tokens {
		fmt.Println(token)
	}
}
