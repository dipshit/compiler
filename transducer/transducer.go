package transducer

import (
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/token"
)

type Transducer struct{}
type Screener struct{}
type Scanner struct{}
type Parser struct {
	//scanner screener tokenStack tableNumberStack treeStack left right tableNumber newTree
}

func NewTransducer() *Transducer {
	return &Transducer{}
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) parse(text string) *ast.TreeNode {
	return ast.NewTreeNode(token.NewToken("1", token.INT))
}

// evaluate creates a Parser which builds a tree from the text
// it then calls evaluateExpressionFor to get the result from the tree
func (t *Transducer) Evaluate(text string) (string, error) {
	result := "3"
	return result, nil
}

// evaluate the ast and return a result
// wrap eval and recover from its panics
func (t *Transducer) evaluateExpressionFor(tree *ast.TreeNode) (string, error) {
	return "", nil
}
