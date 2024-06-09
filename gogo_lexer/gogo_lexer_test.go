package gogo_lexer

import (
	"testing"

	"github.com/adamerikoff/gogo/gogo_token"
)

func TestNextToken(t *testing.T) {
	input := `
	let ten = 10;
	let three = 3;
	let add = fn(x,y){
		x+y;
	};
	let result = add(ten, three);
	!-/*5;
   	5 < 10 > 5;
	if (5 < 10) {
		return true;
	} else {
		return false;
 	}
	`

	tests := []struct {
		expectedType    gogo_token.TokenType
		expectedLiteral string
	}{
		{gogo_token.LET, "let"},
		{gogo_token.IDENTIFIER, "ten"},
		{gogo_token.ASSIGN, "="},
		{gogo_token.INTEGER, "10"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.LET, "let"},
		{gogo_token.IDENTIFIER, "three"},
		{gogo_token.ASSIGN, "="},
		{gogo_token.INTEGER, "3"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.LET, "let"},
		{gogo_token.IDENTIFIER, "add"},
		{gogo_token.ASSIGN, "="},
		{gogo_token.FUNCTION, "fn"},
		{gogo_token.LPARENTHESIS, "("},
		{gogo_token.IDENTIFIER, "x"},
		{gogo_token.COMMA, ","},
		{gogo_token.IDENTIFIER, "y"},
		{gogo_token.RPARENTHESIS, ")"},
		{gogo_token.LBRACE, "{"},
		{gogo_token.IDENTIFIER, "x"},
		{gogo_token.PLUS, "+"},
		{gogo_token.IDENTIFIER, "y"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.RBRACE, "}"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.LET, "let"},
		{gogo_token.IDENTIFIER, "result"},
		{gogo_token.ASSIGN, "="},
		{gogo_token.IDENTIFIER, "add"},
		{gogo_token.LPARENTHESIS, "("},
		{gogo_token.IDENTIFIER, "ten"},
		{gogo_token.COMMA, ","},
		{gogo_token.IDENTIFIER, "three"},
		{gogo_token.RPARENTHESIS, ")"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.EXCLAMATION, "!"},
		{gogo_token.MINUS, "-"},
		{gogo_token.DIVISION, "/"},
		{gogo_token.MULTIPLICATION, "*"},
		{gogo_token.INTEGER, "5"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.INTEGER, "5"},
		{gogo_token.LESSERTHAN, "<"},
		{gogo_token.INTEGER, "10"},
		{gogo_token.GREATERTHAN, ">"},
		{gogo_token.INTEGER, "5"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.IF, "if"},
		{gogo_token.LPARENTHESIS, "("},
		{gogo_token.INTEGER, "5"},
		{gogo_token.LESSERTHAN, "<"},
		{gogo_token.INTEGER, "10"},
		{gogo_token.RPARENTHESIS, ")"},
		{gogo_token.LBRACE, "{"},
		{gogo_token.RETURN, "return"},
		{gogo_token.TRUE, "true"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.RBRACE, "}"},
		{gogo_token.ELSE, "else"},
		{gogo_token.LBRACE, "{"},
		{gogo_token.RETURN, "return"},
		{gogo_token.FALSE, "false"},
		{gogo_token.SEMICOLON, ";"},
		{gogo_token.RBRACE, "}"},
		{gogo_token.EOF, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}

		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
