package lexer_test

import (
	"testing"

	"github.com/atom-sgt/compiler-go/lexer"
)

func Test_GetNextToken_Number(t *testing.T) {
	// Arrange
	expectedTokenType := lexer.TokenType(lexer.LIT_NUMBER)
	expectedValue := "123"

	// Act
	var lexer = lexer.NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedTokenType, nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedValue, nextToken.Value)
	}
}

func Test_GetNextToken_Identifier(t *testing.T) {
	// Arrange
	var expectedTokenType = lexer.TokenType(lexer.IDENT)
	var expectedValue = "abc"

	// Act
	var lexer = lexer.NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedTokenType, nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedValue, nextToken.Value)
	}
}

func Test_GetNextToken_Punctuation(t *testing.T) {
	// Arrange
	var expectedTokenType = lexer.TokenType(lexer.Punctuation)
	var expectedValue = ","

	// Act
	var lexer = lexer.NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedTokenType, nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedValue, nextToken.Value)
	}
}

func Test_GetNextToken_Closure(t *testing.T) {
	// Arrange
	var expectedTokenType = lexer.TokenType(lexer.Closure)
	var expectedValue = "["

	// Act
	var lexer = lexer.NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedTokenType, nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedValue, nextToken.Value)
	}
}

func Test_GetNextToken_Operator(t *testing.T) {
	// Arrange
	var expectedTokenType = lexer.TokenType(lexer.Operator)
	var expectedValue = "="

	// Act
	var lexer = lexer.NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedTokenType, nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedValue, nextToken.Value)
	}
}

func Test_GetNextToken_DoubleSequences(t *testing.T) {
	// Arrange
	var input = "()"
	var expectedFirstValue = "("
	var expectedSecondValue = ")"

	// Act
	var lexer = lexer.NewTokenizer(input)
	var firstToken = lexer.GetNextToken()
	var secondToken = lexer.GetNextToken()

	// Assert
	if firstToken.Value != expectedFirstValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedFirstValue, firstToken.Value)
	}
	if secondToken.Value != expectedSecondValue {
		t.Errorf("Expected \"%v\", found \"%v\"", expectedSecondValue, secondToken.Value)
	}
}
