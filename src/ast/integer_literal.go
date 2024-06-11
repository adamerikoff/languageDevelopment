package ast

import "github.com/adamerikoff/ponGo/src/token"

type IntegerLiteral struct {
	Token token.TokenInstance
	Value int64
}

func (integerLiteral *IntegerLiteral) expressionNode()      {}
func (integerLiteral *IntegerLiteral) TokenLiteral() string { return integerLiteral.Token.Literal }
func (integerLiteral *IntegerLiteral) String() string       { return integerLiteral.Token.Literal }
