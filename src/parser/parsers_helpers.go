package parser

import "github.com/adamerikoff/ponGo/src/token"

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
		return false
	}
}
