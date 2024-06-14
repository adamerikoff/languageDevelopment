package token

const (
	ILLEGAL     = "ILLEGAL"
	END_OF_FILE = "END_OF_FILE"

	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	ASSIGN      = "="
	PLUS        = "+"
	MINUS       = "-"
	ASTERISK    = "*"
	SLASH       = "/"
	INFERIOR    = "<"
	SUPERIOR    = ">"
	EXCLAMATION = "!"
	//AMPERSAND = "&"

	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var KEYWORDS = map[string]TokenType{
	"function": FUNCTION,
	"let":      LET,
}

func ClassifyToken(ident string) TokenType {
	if tok, ok := KEYWORDS[ident]; ok {
		return tok
	}
	return IDENTIFIER
}
