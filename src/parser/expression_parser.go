package parser

import (
	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/token"
)

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{
		Token: parser.currentToken,
	}
	statement.Expression = parser.parseExpression(LOWEST)
	if parser.expectedTokenIs(token.DOT) {
		parser.nextToken()
	}
	return statement
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFunction[parser.currentToken.Type]
	if prefix == nil {
		return nil
	}
	leftExp := prefix()
	return leftExp
}
