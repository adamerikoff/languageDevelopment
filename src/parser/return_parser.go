package parser

import (
	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/token"
)

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{
		Token: parser.currentToken,
	}

	parser.nextToken()

	// TODO: We're skipping the expressions until we
	// encounter a DOT

	for !parser.currentTokenIs(token.DOT) {
		parser.nextToken()
	}
	return statement
}
