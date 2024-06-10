package lexer

import (
	"testing"

	"github.com/adamerikoff/ponGo/src/token"
)

func TestNextToken(t *testing.T) {
	input := `@ five = 5.
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPARENTHESIS, "("},
		{token.RPPARENTHESIS, ")"},
		{token.LCBRACE, "{"},
		{token.RCBRACE, "}"},
		{token.COMMA, ","},
		{token.DOT, "."},
		{token.EOF, ""},
	}

	lexer := NewLexer(input)

	for index, test_element := range tests {
		token := lexer.NextToken()
		if token.Type != test_element.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, test_element.expectedType, token.Type)
		}
		if token.Literal != test_element.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, test_element.expectedLiteral, token.Literal)
		}
	}
}
