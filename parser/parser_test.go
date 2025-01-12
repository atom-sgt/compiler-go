package parser

import (
	"testing"

	"github.com/atom-sgt/compiler-go/lexer"
)

func Test_GetAbstractSyntaxTree_Expression(t *testing.T) {
	// Arrange
	input := "1 + 2"
	lexer := lexer.NewTokenizer(input)
	parser := NewParser(lexer)

	// Act
	parser.Parse()
}
