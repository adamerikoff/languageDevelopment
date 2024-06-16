package lexer

import (
	"testing"

	"github.com/adamerikoff/ponGo/src/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let temp = 22;
	let temp_two = 33;

	let add = function(x, y) {
		x + y;
	};

	let result = add(temp, temp_two);
	!-/*5;
	5 < 10 > 5;
	if (5 < 10){
		return true;
	} else {
		return false;
	}
	10 == 10;
	10 != 9;
	`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "temp"},
		{token.ASSIGN, "="},
		{token.INTEGER, "22"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "temp_two"},
		{token.ASSIGN, "="},
		{token.INTEGER, "33"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "function"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LEFT_PARENTHESIS, "("},
		{token.IDENTIFIER, "temp"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "temp_two"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.SEMICOLON, ";"},
		{token.EXCLAMATION, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "5"},
		{token.INFERIOR, "<"},
		{token.INTEGER, "10"},
		{token.SUPERIOR, ">"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LEFT_PARENTHESIS, "("},
		{token.INTEGER, "5"},
		{token.INFERIOR, "<"},
		{token.INTEGER, "10"},
		{token.RIGHT_PARENTHESIS, ")"},
		{token.LEFT_BRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.ELSE, "else"},
		{token.LEFT_BRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RIGHT_BRACE, "}"},
		{token.INTEGER, "10"},
		{token.EQUAL, "=="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "10"},
		{token.NOT_EQUAL, "!="},
		{token.INTEGER, "9"},
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
