package parser

import (
	"fmt"

	"github.com/atom-sgt/compiler-go/lexer"
)

type Parser struct {
	lexer lexer.Tokenizer
}

func NewParser(lexer lexer.Tokenizer) Parser {
	return Parser{
		lexer: lexer,
	}
}

type Node struct {
	Type     string
	Value    string
	Children []*Node
}

type Constant struct {
	Value string
}

type Operator struct {
	Value string
}

type Operation struct {
	LeftValue  Constant
	RightValue Constant
	Operator   Operator
}

type Return struct {
	Value Operation
}

func (parser Parser) GetAbstractSyntaxTree() {
	for {
		t := parser.lexer.GetNextToken()
		if t.Type == lexer.Unknown {
			return
		}
		fmt.Println(t)
	}
}
