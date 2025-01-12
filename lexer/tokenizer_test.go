package lexer_test

import (
	"testing"

	"github.com/atom-sgt/compiler-go/lexer"
)

func Test_NextChar_SingleStep(t *testing.T) {
	// Arrange
	tokenizer := lexer.NewTokenizer("X")
	expectedResult := 'X'

	// Act
	tokenizer.Step()
	result := tokenizer.Current()

	// Assert
	if result != expectedResult {
		t.Errorf("Expected \"%c\", found \"%c\"", expectedResult, result)
	}
}

func Test_NextChar_MultipleSteps(t *testing.T) {
	// Arrange
	tokenizer := lexer.NewTokenizer("ABC")
	expectedResult := 'C'

	// Act
	tokenizer.Step()
	tokenizer.Step()
	tokenizer.Step()
	result := tokenizer.Current()

	// Assert
	if result != expectedResult {
		t.Errorf("Expected \"%c\", found \"%c\"", expectedResult, result)
	}
}

func Test_NextChar_EndOfFile(t *testing.T) {
	// Arrange
	tokenizer := lexer.NewTokenizer("A")
	expectedResult := '\x00'

	// Act
	tokenizer.Step()
	tokenizer.Step()
	result := tokenizer.Current()

	// Assert
	if result != expectedResult {
		t.Errorf("Expected \"%c\", found \"%c\"", expectedResult, result)
	}
}
