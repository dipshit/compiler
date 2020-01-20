package stack

import (
	"github.com/dipshit/compiler/token"
	"testing"
)

func TestPeek(t *testing.T) {
	stack := NewStack()
	stack.Push(token.NewToken("+", token.PLUS))
	if stack.Peek().Label() != "+" {
		t.Fatalf("did not peek properly")
	}
	if stack.Len() != 1 {
		t.Fatalf("did not have expected len of 1")
	}
}

func TestPush(t *testing.T) {
	stack := NewStack()
	stack.Push(token.NewToken("+", token.PLUS))
	if stack.Peek().Label() != "+" {
		t.Fatalf("did not peek properly")
	}
	stack.Push(token.NewToken("*", token.MULTIPLY))
	if stack.Peek().Label() != "*" {
		t.Fatalf("did not peek properly")
	}
	if stack.Len() != 2 {
		t.Fatalf("did not have expected len of 2")
	}
}

func TestPop(t *testing.T) {
	stack := NewStack()
	stack.Push(token.NewToken("+", token.PLUS))
	if stack.Peek().Label() != "+" {
		t.Fatalf("did not peek properly")
	}
	stack.Push(token.NewToken("*", token.MULTIPLY))
	if stack.Peek().Label() != "*" {
		t.Fatalf("did not peek properly")
	}
	if stack.Pop().Label() != "*" {
		t.Fatalf("did not pop properly")
	}
	if stack.Pop().Label() != "+" {
		t.Fatalf("did not pop properly")
	}
	if stack.Len() != 0 {
		t.Fatalf("did not have expected len of 0")
	}
}
