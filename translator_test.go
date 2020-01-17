package main

import (
	"github.com/dipshit/compiler/ast"
	"github.com/dipshit/compiler/token"
	"testing"
)

// test evaluation for string input
func TestEvaluate(t *testing.T) {
	sample := NewSampleTranslator()
	val, err := sample.evaluate("1+2")
	if err != nil {
		t.Fatal(err)
	}
	if val != "3" {
		t.Fatalf("expected 3 got %s", val)
	}
}

// test evaluation for tree input
func TestEvaluateFor(t *testing.T) {
	sample := NewSampleTranslator()
	tree := ast.NewTreeNode(token.NewToken("1", token.INT))
	val, err := sample.evaluateExpressionFor(tree)
	if err != nil {
		t.Fatal(err)
	}
	if val != "1" {
		t.Fatalf("expected 1 got %s", val)
	}
}
