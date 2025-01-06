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
		t.Error("Bad token type")
	}
	if nextToken.Value != expectedValue {
		t.Error("Bad token value")
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
		t.Error("Bad token type")
	}
	if nextToken.Value != expectedValue {
		t.Error("Bad token value")
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
		t.Error("Bad token type")
	}
	if nextToken.Value != expectedValue {
		t.Error("Bad token value")
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
		t.Error("Bad token type")
	}
	if nextToken.Value != expectedValue {
		t.Error("Bad token value")
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
		t.Error("Bad token type")
	}
	if nextToken.Value != expectedValue {
		t.Error("Bad token value")
	}
}
