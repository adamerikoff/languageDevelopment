package lexer

import (
	"testing"

	"github.com/adamerikoff/ponGo/src/token"
)

func TestNextToken(t *testing.T) {
	input := "=+(){},;"

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LEFT_PARENTHESIS, "("},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RIGHT_BRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
		{token.END_OF_FILE, ""},
	}
	lexer := NewLexer(input)

	for index, test_element := range tests {
		tok := lexer.NextToken()
		if tok.Type != test_element.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", index, test_element.expectedType, tok.Type)
		}

		if tok.Literal != test_element.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", index, test_element.expectedLiteral, tok.Literal)
		}
	}
}
