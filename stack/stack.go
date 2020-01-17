package stack

import (
	"github.com/dipshit/compiler/token"
)

type (
	Stack struct {
		top    *StackNode
		length int
	}
	StackNode struct {
		next  *StackNode
		token token.Token
	}
)

func NewStack() *Stack {
	return Stack{nil, 0}
}

func NewStackNode(t token.Token) *StackNode {
	return StackNode{t}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() token.Token {
	if s.Len() == 0 {
		return nil
	}
	return s.top.token
}

func (s *Stack) Push(t token.Token) {
	node := NewStackNode(t)
	node.next = s.top
	s.length++
	s.top = node
}

func (s *Stack) Pop() token.Token {
	if s.Len() == 0 {
		return nil
	}

	node := s.top
	s.top = node.prev
	this.length--
	return node.token
}
