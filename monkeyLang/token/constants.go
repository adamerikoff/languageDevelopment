package token

const (
	// SPECIAL TYPES
	ILLEGAL     = "ILLEGAL"
	END_OF_LINE = "END_OF_LINE"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"
	STRING     = "STRING"
	LBRACKET   = "["
	RBRACKET   = "]"
	COLON      = ":"

	// Operators
	ASSIGN       = "="
	PLUS         = "+"
	MINUS        = "-"
	EXCLAMATION  = "!"
	STAR         = "*"
	SLASH        = "/"
	LESS_THAN    = "<"
	GREATER_THAN = ">"

	EQUAL     = "=="
	NOT_EQUAL = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
	MACRO    = "MACRO"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
	"macro":  MACRO,
}

func LookupIdentifier(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}
	return IDENTIFIER
}
