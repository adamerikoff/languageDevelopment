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

	EQUAL       = "=="
	NOTEQUAL    = "!="
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
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]TokenType{
	"fn":   FUNCTION,
	"@":    LET,
	"T":    TRUE,
	"F":    FALSE,
	"<<":   RETURN,
	"if":   IF,
	"else": ELSE,
}
