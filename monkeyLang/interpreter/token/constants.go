package token

const (
	// SPECIAL TYPES
	ILLEGAL     = "ILLEGAL"
	END_OF_LINE = "END_OF_LINE"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LEFT_PARENTHESIS  = "("
	RIGHT_PARENTHESIS = ")"
	LEFT_BRACE        = "{"
	RIGHT_BRACE       = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	VARIABLE = "VARIABLE"
)
