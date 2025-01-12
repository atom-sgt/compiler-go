package parser

import "fmt"

type AstNode struct {
	Type     NodeType
	Value    string
	Children []*AstNode
}

type NodeType string

const (
	Program  NodeType = "PROGRAM"
	Number   NodeType = "NUMBER"
	BinaryOp NodeType = "BINARYOP"
)

func (node *AstNode) PrintTree() {
	if node == nil {
		return
	}
	fmt.Println(node.Type, node.Value)
	for _, node := range node.Children {
		node.PrintTree()
	}
}
