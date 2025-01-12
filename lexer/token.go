package lexer

type TokenType int

const (
	Identifier TokenType = iota
	Number
	Punctuation
	Operator
	Closure
	Eof
	Unknown
)

type Token struct {
	Type  TokenType
	Value string
}

func NewToken(tokenType TokenType, value string) Token {
	return Token{
		Type:  tokenType,
		Value: value,
	}
}
