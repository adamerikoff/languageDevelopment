package parser

import (
	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/token"
)

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{
		Token: parser.currentToken,
	}

	if !parser.expectedNextToken(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}

	if !parser.expectedNextToken(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a DOT

	for !parser.currentTokenIs(token.DOT) {
		parser.nextToken()
	}
	return statement
}
