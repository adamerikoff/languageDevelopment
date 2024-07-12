package lexer

import "github.com/adamerikoff/lesLanguesDevs/monkeyLang/interpreter/token"

type Lexer struct {
	input               string
	currentCharPosition int  // current position in input (points to current char)
	readingCharPosition int  // current reading position in input (after current char)
	character           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readCharacter()
	return lexer
}

func (lexer *Lexer) readCharacter() {
	if lexer.readingCharPosition >= len(lexer.input) {
		lexer.character = 0
	} else {
		lexer.character = lexer.input[lexer.readingCharPosition]
	}
	lexer.currentCharPosition = lexer.readingCharPosition
	lexer.readingCharPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		tok = token.NewToken(token.ASSIGN, lexer.character)
	case ';':
		tok = token.NewToken(token.SEMICOLON, lexer.character)
	case ',':
		tok = token.NewToken(token.COMMA, lexer.character)
	case '+':
		tok = token.NewToken(token.PLUS, lexer.character)
	case '(':
		tok = token.NewToken(token.LEFT_PARENTHESIS, lexer.character)
	case ')':
		tok = token.NewToken(token.RIGHT_PARENTHESIS, lexer.character)
	case '{':
		tok = token.NewToken(token.LEFT_BRACE, lexer.character)
	case '}':
		tok = token.NewToken(token.RIGHT_BRACE, lexer.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.END_OF_LINE
	default:
		if isLetter(lexer.character) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isDigit(lexer.character) {
			tok.Type = token.INTEGER
			tok.Literal = lexer.readNumber()
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readCharacter()
	return tok
}

func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.currentCharPosition
	for isLetter(lexer.character) {
		lexer.readCharacter()
	}
	return lexer.input[startPosition:lexer.currentCharPosition]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readCharacter()
	}
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.currentCharPosition
	for isDigit(lexer.character) {
		lexer.readCharacter()
	}
	return lexer.input[startPosition:lexer.currentCharPosition]
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}

func isLetter(character byte) bool {
	return 'a' <= character && character <= 'z' || 'A' <= character && character <= 'Z' || character == '_'
}
