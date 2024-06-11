package ast

import (
	"testing"

	"github.com/adamerikoff/ponGo/src/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.TokenInstance{
					Type:    token.LET,
					Literal: "@",
				},
				Name: &Identifier{
					Token: token.TokenInstance{
						Type:    token.IDENTIFIER,
						Literal: "testVar",
					},
					Value: "testVar",
				},
				Value: &Identifier{
					Token: token.TokenInstance{
						Type:    token.IDENTIFIER,
						Literal: "anotherVar",
					},
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != "@ testVar = anotherVar." {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
