package lexer

type TokenType int

const (
	// Special
	EOF TokenType = iota
	UNKNOWN
	IDENT

	// TODO: REMOVE
	Operator
	Closure
	Punctuation

	// Type literal
	LIT_NUMBER
	LIT_STRING

	// Binary Operator
	OP_ASSIGN
	OP_ADD
	OP_SUB
	OP_DIV
	OP_MULT
	OP_GT
	OP_LT
	OP_GTEQ
	OP_LTEQ
	OP_MOD
	// Unary Operator
	OP_AND
	OP_OR
	OP_XOR
	OP_SHIFTL
	OP_SHIFTR

	// Punctuation
	PUNC_PERIOD
	PUNC_COMMA
	PUNC_BANG
	PUNC_COLON
	PUNC_SEMICOLON
	PUNC_PARENL
	PUNC_PARENR
	PUNC_SQUAREL
	PUNC_SQUARER
	PUNC_CURLYL
	PUNC_CURLYR
	PUNC_ANGLEL
	PUNC_ANGLER

	// Keyword
	KEY_FUNC
	KEY_RETURN
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
