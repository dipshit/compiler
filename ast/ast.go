package ast

// node of the ast
type TreeNode struct {
	children []*TreeNode // child nodes
	label    string      // the type of token this is
}
