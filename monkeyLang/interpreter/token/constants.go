package token

const (
	// SPECIAL TYPES
	ILLEGAL     = "ILLEGAL"
	END_OF_LINE = "END_OF_LINE"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	// Operators
	ASSIGN       = "="
	PLUS         = "+"
	MINUS        = "-"
	EXCLAMATION  = "!"
	STAR         = "*"
	SLASH        = "/"
	LESS_THAN    = "<"
	GREATER_THAN = ">"

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
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(identifier string) TokenType {
	if tokenType, ok := keywords[identifier]; ok {
		return tokenType
	}
	return IDENTIFIER
}
