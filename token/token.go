package token

import ()

type Token struct {
	// one of plus, multiply, int etc.
	label TokenType
	// actual value, like "1", "*" etc.
	literal string
}

type TokenAction string
type TokenType string

const (
	// TokenTypes
	PLUS        = "+"     // 2 children
	MULTIPLY    = "*"     // 2 children
	INT         = "INT"   // no children
	LEFT_PAREN  = "("     // 1 child
	RIGHT_PAREN = ")"     // no children
	BAR         = "-|"    // ?
	COMMA       = ","     // ?
	SEMI        = ";"     // no children
	ASSIGN      = "="     // 2 children
	IDENT       = "IDENT" // ?

	// TokenActions
	RK = "RK"
	L  = "L"
	R  = "R"
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
