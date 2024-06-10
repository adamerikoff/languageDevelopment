package lexer

import (
	"github.com/adamerikoff/ponGo/src/token"
)

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	currentChar  byte // current char under examination
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.readCharacter()
	return lexer
}

func (lexer *Lexer) readCharacter() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.currentChar = 0
	} else {
		lexer.currentChar = lexer.input[lexer.readPosition]
	}
	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.TokenInstance {
	var tok token.TokenInstance

	switch lexer.currentChar {
	case '=':
		tok = newToken(token.ASSIGN, lexer.currentChar)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.currentChar)
	case '(':
		tok = newToken(token.LPARENTHESIS, lexer.currentChar)
	case ')':
		tok = newToken(token.RPPARENTHESIS, lexer.currentChar)
	case ',':
		tok = newToken(token.COMMA, lexer.currentChar)
	case '+':
		tok = newToken(token.PLUS, lexer.currentChar)
	case '{':
		tok = newToken(token.LCBRACE, lexer.currentChar)
	case '}':
		tok = newToken(token.RCBRACE, lexer.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	}
	lexer.readCharacter()
	return tok
}

func newToken(tokenType token.TokenType, character byte) token.TokenInstance {
	return token.TokenInstance{
		Type:    tokenType,
		Literal: string(character),
	}
}
