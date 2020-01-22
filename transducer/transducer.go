package transducer

import (
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/token"
)

type Transducer struct{}
type Screener struct{}
type Scanner struct {
	input          string
	token          token.Token
	keptCharacters string
}
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

// peek in stream
func (t *Transducer) PeekInput() string {
	// TODO
	return ""
}

func (t *Transducer) BuildToken(label token.TokenType) {
	// TODO
	//t.token = NewToken(t.keptCharacters, label)
	//t.keptCharacters = ""
}

func (t *Transducer) BuildTree() {
	// TODO
}

func (t *Transducer) Transitions(char string) (token.TokenAction, int) {
	// TODO
	// return a token action and the jump from the transition
	return token.L, 1
}

func (t *Transducer) KeepChar(char string) {
	// TODO
}

func (t *Transducer) DiscardInput() {
	// TODO
}
