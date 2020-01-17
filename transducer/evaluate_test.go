package transducer

import (
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/token"
	"testing"
)

func TestEvaluateTreeWithInt(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("6", token.INT))
	val := evaluateTree(tree)
	if val != 6 {
		t.Fatalf("evaluateInt did not give 6, gave: %d", val)
	}
}

func TestEvaluateInt(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("6", token.INT))
	val := evaluateInt(tree)
	if val != 6 {
		t.Fatalf("evaluateInt did not give 6, gave: %d", val)
	}
}

func TestEvaluateMultiply(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("*", token.MULTIPLY))
	tree.AddChild(ast.NewTreeNode(token.NewToken("2", token.INT)))
	tree.AddChild(ast.NewTreeNode(token.NewToken("4", token.INT)))
	val := evaluateMultiply(tree)
	if val != 8 {
		t.Fatalf("evaluateMultiply did not give 8, gave: %d", val)
	}
}

func TestEvaluateTreeWithMultiply(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("*", token.MULTIPLY))
	tree.AddChild(ast.NewTreeNode(token.NewToken("3", token.INT)))
	tree.AddChild(ast.NewTreeNode(token.NewToken("4", token.INT)))
	val := evaluateTree(tree)
	if val != 12 {
		t.Fatalf("evaluateMultiply did not give 12, gave: %d", val)
	}
}

func TestEvaluatePlus(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("+", token.MULTIPLY))
	tree.AddChild(ast.NewTreeNode(token.NewToken("9", token.INT)))
	tree.AddChild(ast.NewTreeNode(token.NewToken("10", token.INT)))
	val := evaluatePlus(tree)
	if val != 19 {
		t.Fatalf("evaluateMultiply did not give 19, gave: %d", val)
	}
}

func TestEvaluateTreeWithPlus(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("+", token.PLUS))
	tree.AddChild(ast.NewTreeNode(token.NewToken("3", token.INT)))
	tree.AddChild(ast.NewTreeNode(token.NewToken("4", token.INT)))
	val := evaluateTree(tree)
	if val != 7 {
		t.Fatalf("evaluateTree did not give 7, gave: %d", val)
	}
}

func TestEvaluateDeepTree(t *testing.T) {
	tree := ast.NewTreeNode(token.NewToken("+", token.PLUS))
	tree.AddChild(ast.NewTreeNode(token.NewToken("3", token.INT)))
	tree2 := ast.NewTreeNode(token.NewToken("*", token.MULTIPLY))
	tree2.AddChild(ast.NewTreeNode(token.NewToken("5", token.INT)))
	tree2.AddChild(ast.NewTreeNode(token.NewToken("7", token.INT)))
	tree.AddChild(tree2)
	val := evaluateTree(tree)
	if val != 38 {
		t.Fatalf("evaluateTree did not give 38, gave: %d", val)
	}
}
