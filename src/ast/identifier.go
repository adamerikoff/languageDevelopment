package ast

import "github.com/adamerikoff/ponGo/src/token"

type Identifier struct {
	Token token.TokenInstance // the token.IDENT token
	Value string
}

func (identifier *Identifier) expressionNode()      {}
func (identifier *Identifier) TokenLiteral() string { return identifier.Token.Literal }
func (identifier *Identifier) String() string       { return identifier.Value }
