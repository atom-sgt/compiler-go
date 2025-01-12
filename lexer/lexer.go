package lexer

import (
	"unicode"
)

func (tokenizer *Tokenizer) GetNextToken() Token {
	tokenizer.Step()
	var nextToken Token

	for isWhitespace(tokenizer.Current()) {
		tokenizer.Step()
	}

	if isNumber(tokenizer.Current()) {
		runes := []rune{tokenizer.Current()}
		for isNumber(tokenizer.Peek()) {
			tokenizer.Step()
			runes = append(runes, tokenizer.Current())
		}

		nextToken = NewToken(LIT_NUMBER, string(runes))
	} else if isAlpha(tokenizer.Current()) {
		runes := []rune{tokenizer.Current()}
		for isAlpha(tokenizer.Peek()) {
			tokenizer.Step()
			runes = append(runes, tokenizer.Current())
		}

		nextToken = NewToken(IDENT, string(runes))
	} else if isOperator(tokenizer.Current()) {
		runes := []rune{tokenizer.Current()}
		for isOperator(tokenizer.Peek()) {
			tokenizer.Step()
			runes = append(runes, tokenizer.Current())
		}

		nextToken = NewToken(Operator, string(runes))
	} else if isPunctuation(tokenizer.Current()) {
		nextToken = NewToken(Punctuation, string(tokenizer.Current()))
	} else if isClosure(tokenizer.Current()) {
		nextToken = NewToken(Closure, string(tokenizer.Current()))
	} else if isEof(tokenizer.Current()) {
		nextToken = NewToken(EOF, string(tokenizer.Current()))
	} else {
		nextToken = NewToken(UNKNOWN, string(tokenizer.Current()))
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
