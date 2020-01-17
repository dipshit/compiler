package ast

import (
	"github.com/dipshit/compiler/token"
)

// node of the ast
type TreeNode struct {
	children []*TreeNode // child nodes
	token    token.Token
}

func NewTreeNode(token token.Token) *TreeNode {
	return &TreeNode{token: token, children: make([]*TreeNode, 0)}
}

func (t *TreeNode) ToString() string {
	if !t.HasChildren() {
		return ""
	}
	str := t.token.Literal() + " newline "
	for _, node := range t.children {
		str += node.token.Literal() + " "
	}
	return str
}

func (t *TreeNode) AddChild(node *TreeNode) {
	t.children = append(t.children, node)
}

func (t *TreeNode) HasChildren() bool {
	return len(t.children) != 0
}
