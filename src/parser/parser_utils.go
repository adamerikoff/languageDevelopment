package parser

import "github.com/adamerikoff/ponGo/src/token"

func (parser *Parser) currentTokenIs(tokenType token.TokenType) bool {
	return parser.currentToken.Type == tokenType
}

func (parser *Parser) expectedTokenIs(tokenType token.TokenType) bool {
	return parser.expectedToken.Type == tokenType
}

func (parser *Parser) expectedNextToken(tokenType token.TokenType) bool {
	if parser.expectedTokenIs(tokenType) {
		parser.nextToken()
		return true
	} else {
		parser.expectedError(tokenType)
		return false
	}
}
