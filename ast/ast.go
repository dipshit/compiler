package ast

import (
	"github.com/dipshit/compiler/token"
)

// node of the ast
type TreeNode struct {
	children []*TreeNode // child nodes
	label    string      // the type of token this is
	token    token.Token
}
