package parser

import (
	"fmt"

	"github.com/adamerikoff/ponGo/src/token"
)

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) expectedError(tokenType token.TokenType) {
	message := fmt.Sprintf("expected next token to be %s, got %s instead", tokenType, parser.expectedToken.Type)
	parser.errors = append(parser.errors, message)
}
