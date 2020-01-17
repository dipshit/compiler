package token

import ()

type Token struct {
	// one of plus, multiply, int etc.
	label TokenType
	// actual value, like "1", "*" etc.
	literal string
}

type TokenType string

const (
	PLUS     = "+"   // 2 children
	MULTIPLY = "*"   // 2 children
	INT      = "INT" // no children
)

func NewToken(literal string, label TokenType) Token {
	return Token{label, literal}
}

func (t Token) Label() TokenType {
	return t.label
}

func (t Token) Literal() string {
	return t.literal
}
