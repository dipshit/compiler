package transducer

import (
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/token"
	"testing"
)

// test evaluation for tree input
func TestEvaluateFor(t *testing.T) {
	transducer := NewTransducer()
	tree := ast.NewTreeNode(token.NewToken("1", token.INT))
	val, err := transducer.evaluateExpressionFor(tree)
	if err != nil {
		t.Fatal(err)
	}
	if val != "1" {
		t.Fatalf("expected 1 got %s", val)
	}
}

func TestEvaluateExpressionFor(t *testing.T) {
	transducer := NewTransducer()
	tree := ast.NewTreeNode(token.NewToken("+", token.INT))
	tree.AddChild(ast.NewTreeNode(token.NewToken("1", token.INT)))
	tree.AddChild(ast.NewTreeNode(token.NewToken("2", token.INT)))
	t.Log(tree.ToString())
	val, err := transducer.evaluateExpressionFor(tree)
	if err != nil {
		t.Fatal(err)
	}
	if val != "3" {
		t.Fatalf("expected 1 got %s", val)
	}
}
