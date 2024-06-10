package token

const (
	//BASE
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//IDENTIFIERS AND LITERALS
	IDENT = "IDENT"
	INT   = "INT"

	//OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	//DELIMITERS
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
