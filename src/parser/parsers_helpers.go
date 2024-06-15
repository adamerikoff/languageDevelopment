package parser

import (
	"fmt"

	"github.com/adamerikoff/ponGo/src/token"
)

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) SubsequentTokenIs(t token.TokenType) bool {
	return parser.subsequentToken.Type == t
}

func (parser *Parser) expectSubsequentToken(t token.TokenType) bool {
	if parser.SubsequentTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.expectSubsequentError(t)
		return false
	}
}

func (parser *Parser) expectSubsequentError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead",
		t, parser.subsequentToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) subsequentPrecedence() int {
	if precedence, ok := precedences[parser.subsequentToken.Type]; ok {
		return precedence
	}
	return LOWEST
}
func (parser *Parser) currentPrecedence() int {
	if precedence, ok := precedences[parser.currentToken.Type]; ok {
		return precedence
	}
	return LOWEST
}
