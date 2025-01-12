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

func (p *Parser) Parse() {
	p.ParseProgram().PrintTree()
}

func (p *Parser) ParseProgram() *AstNode {
	x := p.ParseExpression()

	return &AstNode{
		Type:     Program,
		Children: []*AstNode{x},
	}
}

func (p *Parser) ParseExpression() *AstNode {
	println("begin exp")

	left := p.ParseNumber()
	println(left.Value)
	op := p.lexer.GetNextToken()
	println(op.Value)
	right := p.ParseNumber()
	println(right.Value)

	expectTokenType(op, lexer.Operator)
	return &AstNode{
		Type:     BinaryOp,
		Value:    op.Value,
		Children: []*AstNode{left, right},
	}
}

func (p *Parser) ParseNumber() *AstNode {
	println("parse num")
	t := p.lexer.GetNextToken()

	expectTokenType(t, lexer.LIT_NUMBER)
	return &AstNode{
		Type:  Number,
		Value: t.Value,
	}
}

func matchTokenTypes(t lexer.Token, types ...lexer.TokenType) bool {
	for _, type_ := range types {
		if t.Type == type_ {
			return true
		}
	}
	return false
}

func expectTokenType(t lexer.Token, type_ lexer.TokenType) {
	if t.Type != type_ {
		panic(fmt.Sprintf("Expected \"%s\", found \"%s\"", string(type_), string(t.Type)))
	}
}
