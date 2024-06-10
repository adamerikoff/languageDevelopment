package gogo_ast

import (
	"testing"

	"github.com/adamerikoff/gogo/gogo_token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: gogo_token.Token{Type: gogo_token.LET, Literal: "let"},
				Name: &Identifier{
					Token: gogo_token.Token{Type: gogo_token.IDENTIFIER, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: gogo_token.Token{Type: gogo_token.IDENTIFIER, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}
	if program.String() != "let myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
