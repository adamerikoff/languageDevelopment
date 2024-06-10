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

func (lexer *Lexer) NextToken() token.TokenInstance {
	var tok token.TokenInstance

	lexer.skipWhitespace()

	switch lexer.currentChar {
	case '=':
		if lexer.peekCharacter() == '=' {
			lexer.readCharacter()
			tok = token.NewToken(token.EQUAL, "==")
		} else {
			tok = token.NewToken(token.ASSIGN, lexer.currentChar)
		}
	case '.':
		tok = token.NewToken(token.DOT, lexer.currentChar)
	case '(':
		tok = token.NewToken(token.LPARENTHESIS, lexer.currentChar)
	case ')':
		tok = token.NewToken(token.RPPARENTHESIS, lexer.currentChar)
	case ',':
		tok = token.NewToken(token.COMMA, lexer.currentChar)
	case '+':
		tok = token.NewToken(token.PLUS, lexer.currentChar)
	case '{':
		tok = token.NewToken(token.LCBRACE, lexer.currentChar)
	case '}':
		tok = token.NewToken(token.RCBRACE, lexer.currentChar)
	case '>':
		tok = token.NewToken(token.GREATERTHAN, lexer.currentChar)
	case '<':
		tok = token.NewToken(token.LESSTHAN, lexer.currentChar)
	case '-':
		tok = token.NewToken(token.MINUS, lexer.currentChar)
	case '!':
		if lexer.peekCharacter() == '=' {
			lexer.readCharacter()
			tok = token.NewToken(token.NOTEQUAL, "!=")
		} else {
			tok = token.NewToken(token.EXCLAMATION, lexer.currentChar)
		}
	case '*':
		tok = token.NewToken(token.ASTERISK, lexer.currentChar)
	case '/':
		tok = token.NewToken(token.SLASH, lexer.currentChar)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.currentChar) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(lexer.currentChar) {
			tok.Type = token.INTEGER
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, lexer.currentChar)
		}
	}
	lexer.readCharacter()
	return tok
}
