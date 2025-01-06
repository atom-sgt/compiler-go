package lexer

import "unicode"

type TokenType string

const (
	Identifier  TokenType = "INDENTIFIER"
	Number                = "NUMBER"
	Punctuation           = "PUNCTUATION"
	Operator              = "OPERATOR"
	Closure               = "CLOSURE"
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
	position     int
	readPosition int
	char         rune
	input        string
}

func NewTokenizer(input string) Tokenizer {
	tokenizer := Tokenizer{
		position:     0,
		readPosition: 0,
		input:        input,
	}
	tokenizer.readChar()

	return tokenizer
}

func (tokenizer *Tokenizer) readChar() {
	if tokenizer.readPosition >= len(tokenizer.input) {
		tokenizer.char = '\x00'
	} else {
		tokenizer.char = rune(tokenizer.input[tokenizer.readPosition])
	}

	tokenizer.position = tokenizer.readPosition
	tokenizer.readPosition++
}

func (tokenizer *Tokenizer) GetNextToken() Token {
	// Skip whitespace
	for isWhitespace(tokenizer.char) {
		tokenizer.readChar()
	}

	if isNumber(tokenizer.char) {
		start := tokenizer.position
		for isNumber(tokenizer.char) {
			tokenizer.readChar()
		}

		return NewToken(Number, tokenizer.input[start:tokenizer.position])
	}

	if isAlpha(tokenizer.char) {
		start := tokenizer.position
		for isAlpha(tokenizer.char) {
			tokenizer.readChar()
		}

		return NewToken(Identifier, tokenizer.input[start:tokenizer.position])
	}

	if isPunctuation(tokenizer.char) {
		return NewToken(Punctuation, string(tokenizer.char))
	}

	if isOperator(tokenizer.char) {
		start := tokenizer.position
		for isOperator(tokenizer.char) {
			tokenizer.readChar()
		}

		return NewToken(Operator, tokenizer.input[start:tokenizer.position])
	}

	if isClosure(tokenizer.char) {
		return NewToken(Closure, string(tokenizer.char))
	}

	return NewToken(Unknown, string(tokenizer.char))
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
