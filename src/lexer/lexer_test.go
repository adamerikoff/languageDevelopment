package lexer

import (
	"testing"

	"github.com/adamerikoff/ponGo/src/token"
)

func TestNextToken(t *testing.T) {
	input := `
	@ five = 5.
	@ ten = 10.

	@ add = fn(x, y) {
		x + y.
	}.
	@ result = add(five, ten).
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "@"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.DOT, "."},

		{token.LET, "@"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.DOT, "."},

		{token.LET, "@"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPPARENTHESIS, ")"},
		{token.LCBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.DOT, "."},
		{token.RCBRACE, "}"},
		{token.DOT, "."},

		{token.LET, "@"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPARENTHESIS, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPPARENTHESIS, ")"},
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
