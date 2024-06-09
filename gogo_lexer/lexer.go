package gogo_lexer

import (
	"github.com/adamerikoff/gogo/gogo_token"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewToken(tokenType gogo_token.TokenType, ch interface{}) gogo_token.Token {
	switch v := ch.(type) {
	case byte:
		return gogo_token.Token{Type: tokenType, Literal: string(v)}
	case string:
		return gogo_token.Token{Type: tokenType, Literal: v}
	default:
		return gogo_token.Token{}
	}
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) NextToken() gogo_token.Token {
	var tok gogo_token.Token
	l.skipWhitespace()
	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = NewToken(gogo_token.EQ, literal)
		} else {
			tok = NewToken(gogo_token.ASSIGN, l.ch)
		}
	case '+':
		tok = NewToken(gogo_token.PLUS, l.ch)
	case '-':
		tok = NewToken(gogo_token.MINUS, l.ch)
	case '*':
		tok = NewToken(gogo_token.MULTIPLICATION, l.ch)
	case '/':
		tok = NewToken(gogo_token.DIVISION, l.ch)
	case ',':
		tok = NewToken(gogo_token.COMMA, l.ch)
	case ';':
		tok = NewToken(gogo_token.SEMICOLON, l.ch)
	case '(':
		tok = NewToken(gogo_token.LPARENTHESIS, l.ch)
	case ')':
		tok = NewToken(gogo_token.RPARENTHESIS, l.ch)
	case '{':
		tok = NewToken(gogo_token.LBRACE, l.ch)
	case '}':
		tok = NewToken(gogo_token.RBRACE, l.ch)
	case '<':
		tok = NewToken(gogo_token.LESSERTHAN, l.ch)
	case '>':
		tok = NewToken(gogo_token.GREATERTHAN, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = NewToken(gogo_token.NOT_EQ, literal)
		} else {
			tok = NewToken(gogo_token.EXCLAMATION, l.ch)
		}
	case 0:
		tok.Literal = ""
		tok.Type = gogo_token.EOF
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = gogo_token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Type = gogo_token.INTEGER
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = NewToken(gogo_token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return tok
}
