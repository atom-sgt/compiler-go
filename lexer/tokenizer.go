package lexer

const NULL_RUNE = '\x00'

type Tokenizer struct {
	walkPosition int
	lookPosition int
	currentRune  rune
	input        string
}

func NewTokenizer(input string) Tokenizer {
	tokenizer := Tokenizer{
		walkPosition: 0,
		lookPosition: 0,
		input:        input,
	}

	return tokenizer
}

func (t *Tokenizer) Step() {
	// Check for end
	if t.lookPosition >= len(t.input) {
		t.currentRune = NULL_RUNE
		return
	}

	// Consume next char
	t.currentRune = rune(t.input[t.lookPosition])
	t.walkPosition = t.lookPosition
	t.lookPosition++
}

func (t *Tokenizer) Peek() rune {
	// Check for end
	if t.lookPosition >= len(t.input) {
		return NULL_RUNE
	}

	// Look but don't touch
	return rune(t.input[t.lookPosition])
}

func (t *Tokenizer) Current() rune {
	return t.currentRune
}
