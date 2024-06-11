package ast

import "github.com/adamerikoff/ponGo/src/token"

type ExpressionStatement struct {
	Token      token.TokenInstance
	Expression Expression
}

func (expressionStatement *ExpressionStatement) statementNode() {}
func (expressionStatement *ExpressionStatement) TokenLiteral() string {
	return expressionStatement.Token.Literal
}
func (expressionStatement *ExpressionStatement) String() string {
	if expressionStatement.Expression != nil {
		return expressionStatement.Expression.String()
	}
	return ""
}
