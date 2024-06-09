package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INTEGER    = "INTEGER"

	ASSIGN         = "="
	PLUS           = "+"
	MINUS          = "-"
	MULTIPLICATION = "*"
	DIVISION       = "/"

	COMMA        = ","
	SEMICOLON    = ";"
	LPARENTHESIS = "("
	RPARENTHESIS = ")"
	LBRACE       = "{"
	RBRACE       = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)
