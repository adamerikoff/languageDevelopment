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

	!-/*5.
   	5 < 10 > 5.

	if (5 < 10) {
		<< T.
	} else {
		<< F.
 	}
	10 == 10.
	10 != 9.
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

		{token.EXCLAMATION, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.DOT, "."},

		{token.INTEGER, "5"},
		{token.LESSTHAN, "<"},
		{token.INTEGER, "10"},
		{token.GREATERTHAN, ">"},
		{token.INTEGER, "5"},
		{token.DOT, "."},

		{token.IF, "if"},
		{token.LPARENTHESIS, "("},
		{token.INTEGER, "5"},
		{token.LESSTHAN, "<"},
		{token.INTEGER, "10"},
		{token.RPPARENTHESIS, ")"},
		{token.LCBRACE, "{"},
		{token.RETURN, "<<"},
		{token.TRUE, "T"},
		{token.DOT, "."},
		{token.RCBRACE, "}"},
		{token.ELSE, "else"},
		{token.LCBRACE, "{"},
		{token.RETURN, "<<"},
		{token.FALSE, "F"},
		{token.DOT, "."},
		{token.RCBRACE, "}"},

		{token.INTEGER, "10"},
		{token.EQUAL, "=="},
		{token.INTEGER, "10"},
		{token.DOT, "."},

		{token.INTEGER, "10"},
		{token.NOTEQUAL, "!="},
		{token.INTEGER, "9"},
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
