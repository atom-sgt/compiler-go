package lexer

import "testing"

func Test_GetNextToken_Number(t *testing.T) {
	// Arrange
	var expectedTokenType = TokenType(Number)
	var expectedValue = "123"

	// Act
	var lexer = NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Bad token type: \"%s\"", nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Bad token value: \"%s\"", nextToken.Value)
	}
}

func Test_GetNextToken_Identifier(t *testing.T) {
	// Arrange
	var expectedTokenType = TokenType(Identifier)
	var expectedValue = "abc"

	// Act
	var lexer = NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Bad token type: \"%s\"", nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Bad token value: \"%s\"", nextToken.Value)
	}
}

func Test_GetNextToken_Punctuation(t *testing.T) {
	// Arrange
	var expectedTokenType = TokenType(Punctuation)
	var expectedValue = ","

	// Act
	var lexer = NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Bad token type: \"%s\"", nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Bad token value: \"%s\"", nextToken.Value)
	}
}

func Test_GetNextToken_Closure(t *testing.T) {
	// Arrange
	var expectedTokenType = TokenType(Closure)
	var expectedValue = "["

	// Act
	var lexer = NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Bad token type: \"%s\"", nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Bad token value: \"%s\"", nextToken.Value)
	}
}

func Test_GetNextToken_Operator(t *testing.T) {
	// Arrange
	var expectedTokenType = TokenType(Operator)
	var expectedValue = "="

	// Act
	var lexer = NewTokenizer(expectedValue)
	var nextToken = lexer.GetNextToken()

	// Assert
	if nextToken.Type != expectedTokenType {
		t.Errorf("Bad token type: \"%s\"", nextToken.Type)
	}
	if nextToken.Value != expectedValue {
		t.Errorf("Bad token value: \"%s\"", nextToken.Value)
	}
}

func Test_GetNextToken_DoubleSequences(t *testing.T) {
	// Arrange
	var input = "()"

	// Act
	var lexer = NewTokenizer(input)
	var firstToken = lexer.GetNextToken()
	var secondToken = lexer.GetNextToken()

	// Assert
	if firstToken.Value == "(" {
		t.Errorf("Bad token value: \"%s\"", firstToken.Value)
	}
	if secondToken.Value == ")" {
		t.Errorf("Bad token value: \"%s\"", secondToken.Value)
	}
}
