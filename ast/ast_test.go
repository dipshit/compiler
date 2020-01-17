package ast

import (
	"github.com/dipshit/compiler/token"
	"testing"
)

func TestToString(t *testing.T) {
	ast := NewTreeNode(token.NewToken("3", token.INT))
	val := ast.ToString()
	if val != "3" {
		t.Fatalf("expected 3 but got %s", val)
	}
}

func TestToStringMultiple(t *testing.T) {
	ast := *NewTreeNode(token.NewToken("+", token.MULTIPLY))
	ast.AddChild(NewTreeNode(token.NewToken("1", token.INT)))
	ast.AddChild(NewTreeNode(token.NewToken("2", token.INT)))
	val := ast.ToString()
	t.Log(val)
	if val != "+ newline 1 2 " {
		t.Fatalf("expected + newline 1 2 but got %s", val)
	}
}
