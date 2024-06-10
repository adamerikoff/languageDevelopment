package token

const (
	//BASE
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//IDENTIFIERS AND LITERALS
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	//OPERATORS
	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	ASTERISK    = "*"
	SLASH       = "/"
	EXCLAMATION = "!"

	LESSTHAN    = "<"
	GREATERTHAN = ">"

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
