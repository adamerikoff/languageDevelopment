package ast

import (
	"bytes"

	"github.com/adamerikoff/ponGo/src/token"
)

type ReturnStatement struct {
	Token       token.TokenInstance
	ReturnValue Expression
}

func (returnStatement *ReturnStatement) statementNode() {}
func (returnStatement *ReturnStatement) TokenLiteral() string {
	return returnStatement.Token.Literal
}
func (returnStatement *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(returnStatement.TokenLiteral() + " ")

	if returnStatement.ReturnValue != nil {
		out.WriteString(returnStatement.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}
