package transducer

import (
	"github.com/dipshit/compiler/ast"
)

type Transducer struct{}
type Screener struct{}
type Scanner struct{}
type Parser struct {
	//scanner screener tokenStack tableNumberStack treeStack left right tableNumber newTree
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) parse(text string) *ast.TreeNode {
	return &ast.TreeNode{}
}
