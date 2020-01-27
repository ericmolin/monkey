package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const {
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"
	COMMENT = "#"

	IDENT = "IDENT"
	INT = "INT"

	ASSIGN = "="
	PLUS = "+"
	MINUS = "-"
	MULTI = "*"
	DIV = "/"

	COMMA = ","
	SEMICOLON = ";"

	LPAR = "("
	RPAR = ")"
	LBRACE = "{"
	RBRACE = "{"
	LBRACK = "["
	RBRACK = "]"

	FUNCTION = "FUNCTION"
	LET = "LET"
}