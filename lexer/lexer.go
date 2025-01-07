package lexer

import (
	"unicode"
)

type TokenType string

const (
	Identifier  TokenType = "INDENTIFIER"
	Number                = "NUMBER"
	Punctuation           = "PUNCTUATION"
	Operator              = "OPERATOR"
	Closure               = "CLOSURE"
	Eof                   = "EOF"
	Unknown               = "UNKNOWN"
)

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(t TokenType, v string) Token {
	return Token{
		Type:  t,
		Value: v,
	}
}

type Tokenizer struct {
	currentPosition int
	readPosition    int
	char            rune
	input           string
}

func NewTokenizer(input string) Tokenizer {
	tokenizer := Tokenizer{
		currentPosition: 0,
		readPosition:    0,
		input:           input,
	}

	return tokenizer
}

func (tokenizer *Tokenizer) nextChar() {
	if tokenizer.readPosition >= len(tokenizer.input) {
		tokenizer.char = '\x00'
	} else {
		tokenizer.char = rune(tokenizer.input[tokenizer.readPosition])
	}

	tokenizer.currentPosition = tokenizer.readPosition
	tokenizer.readPosition++
}

func (tokenizer *Tokenizer) prevChar() {
	if tokenizer.readPosition < 0 {
		tokenizer.char = '\x00'
	} else {
		tokenizer.char = rune(tokenizer.input[tokenizer.readPosition])
	}

	tokenizer.currentPosition = tokenizer.readPosition
	tokenizer.readPosition--
}

func (tokenizer *Tokenizer) peekNextChar() rune {
	if tokenizer.readPosition >= len(tokenizer.input) {
		return '\x00'
	} else {
		return rune(tokenizer.input[tokenizer.readPosition])
	}
}

func (tokenizer *Tokenizer) GetNextToken() Token {
	tokenizer.nextChar()
	var nextToken Token

	// Skip whitespace
	for isWhitespace(tokenizer.char) {
		tokenizer.nextChar()
	}

	if isNumber(tokenizer.char) {
		start := tokenizer.currentPosition
		for isNumber(tokenizer.peekNextChar()) {
			tokenizer.nextChar()
		}

		nextToken = NewToken(Number, tokenizer.input[start:tokenizer.readPosition])
	} else if isAlpha(tokenizer.char) {
		start := tokenizer.currentPosition
		for isAlpha(tokenizer.peekNextChar()) {
			tokenizer.nextChar()
		}

		nextToken = NewToken(Identifier, tokenizer.input[start:tokenizer.readPosition])
	} else if isOperator(tokenizer.char) {
		start := tokenizer.currentPosition
		for isOperator(tokenizer.peekNextChar()) {
			tokenizer.nextChar()
		}

		nextToken = NewToken(Operator, tokenizer.input[start:tokenizer.readPosition])
	} else if isPunctuation(tokenizer.char) {
		nextToken = NewToken(Punctuation, string(tokenizer.char))
	} else if isClosure(tokenizer.char) {
		nextToken = NewToken(Closure, string(tokenizer.char))
	} else if isEof(tokenizer.char) {
		nextToken = NewToken(Eof, string(tokenizer.char))
	} else {
		nextToken = NewToken(Unknown, string(tokenizer.char))
	}

	return nextToken
}

func isWhitespace(c rune) bool {
	return unicode.IsSpace(c)
}

func isNumber(c rune) bool {
	return unicode.IsDigit(c)
}

func isAlpha(c rune) bool {
	return unicode.IsLetter(c)
}

func isAlphanumeric(c rune) bool {
	return unicode.IsLetter(c) || unicode.IsDigit(c)
}

func isOperator(c rune) bool {
	return c == '+' ||
		c == '-' ||
		c == '*' ||
		c == '/' ||
		c == '=' ||
		c == '!'
}

func isCompoundOperator(c1 rune, c2 rune) bool {
	return (c1 == '+' && c2 == '+') ||
		(c1 == '-' && c2 == '-') ||
		(c1 == '=' && c2 == '=') ||
		(c1 == '!' && c2 == '=')
}

func isPunctuation(c rune) bool {
	return c == ';' ||
		c == '.' ||
		c == ',' ||
		c == '?'
}

func isClosure(c rune) bool {
	return c == '(' ||
		c == ')' ||
		c == '[' ||
		c == ']' ||
		c == '{' ||
		c == '}'
}

func isEof(c rune) bool {
	return c == '\x00'
}
