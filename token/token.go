package token

import ()

type Token struct {
	label   TokenType
	literal string
}

type TokenType string

const (
	PLUS     = "+"
	MULTIPLY = "*"
	INT      = "INT"
	EOF      = "EOF"
)
