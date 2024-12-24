package lexer

import (
	"fmt"
	"unicode"
)

type TokenType string

const (
	TK_INDENTIFIER TokenType = "INDENTIFIER"
	TK_NUMBER                = "NUMBER"
	TK_OPERATOR              = "OPERATOR"
	TK_PUNCTUATION           = "PUNCTUATION"
	TK_UNKNOWN               = "UNKNOWN"
)

type Token struct {
	Type  TokenType
	Value string
}

func newToken(t TokenType, val string) Token {
	return Token{
		Type:  t,
		Value: val,
	}
}

func printToken(token Token) {
	fmt.Printf("Token: Type=\"%v\", Value=\"%v\"\n", token.Type, token.Value)
}

func GetTokens(input string) {
	println(input)
	for i := 0; i < len(input); {
		c := rune(input[i])

		// Skip whitespace
		if unicode.IsSpace(c) {
			i++
			continue
		}

		// Indentify numbers
		if isNumber(c) {
			start := i
			for isNumber(rune(input[i])) || i >= len(input) {
				i++
			}

			value := input[start:i]
			token := newToken(TK_NUMBER, value)
			printToken(token)
			continue
		}

		// Identify identifiers
		if isAlpha(c) {
			start := i
			for isAlphanumeric(rune(input[i])) || i >= len(input) {
				i++
			}

			value := input[start:i]
			token := newToken(TK_INDENTIFIER, value)
			printToken(token)
			continue
		}

		// Identify operators
		if isOperator(c) {
			start := i
			i++
			if i < len(input) && isCompoundOperator(rune(input[start]), rune(input[i])) {
				i++
			}

			value := string(input[start:i])
			token := newToken(TK_OPERATOR, value)
			printToken(token)
			continue
		}

		// Identify punctuation
		if isPunctuation(c) {
			value := string(c)
			token := newToken(TK_PUNCTUATION, value)
			printToken(token)
			i++
			continue
		}

		// Unknown
		token := newToken(TK_UNKNOWN, "")
		printToken(token)
	}
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
		c == ',' ||
		c == '(' ||
		c == ')' ||
		c == '{' ||
		c == '}'
}
