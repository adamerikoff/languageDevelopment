package token

const (
	//BASE
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//IDENTIFIERS AND LITERALS
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	//OPERATORS
	ASSIGN = "="
	PLUS   = "+"

	//DELIMITERS
	COMMA         = ","
	DOT           = "."
	LPARENTHESIS  = "("
	RPPARENTHESIS = ")"
	LCBRACE       = "{"
	RCBRACE       = "}"

	//KEYWORDS
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn": FUNCTION,
	"@":  LET,
}
