package token

const (
	ILLEGAL     = "ILLEGAL"
	END_OF_FILE = "END_OF_FILE"

	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	ASTERISK = "*"
	SLASH    = "/"

	INFERIOR = "<"
	SUPERIOR = ">"

	EXCLAMATION = "!"
	EQUAL       = "=="
	NOT_EQUAL   = "!="

	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var KEYWORDS = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
	"true":     TRUE,
	"false":    FALSE,
	"if":       IF,
	"else":     ELSE,
	"return":   RETURN,
}

func ClassifyToken(ident string) TokenType {
	if tok, ok := KEYWORDS[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
