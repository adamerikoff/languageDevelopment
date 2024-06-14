package lexer

import "github.com/adamerikoff/ponGo/src/token"

type Lexer struct {
	input           string
	currentPosition int
	readPosition    int
	character       byte
}

func NewLexer(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.ReadCharacter()
	return lexer
}

func (lexer *Lexer) ReadCharacter() {
	switch {
	case lexer.readPosition >= len(lexer.input):
		lexer.character = 0
	default:
		lexer.character = lexer.input[lexer.readPosition]
		lexer.currentPosition = lexer.readPosition
		lexer.readPosition += 1
	}
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	switch lexer.character {
	case '=':
		tok = token.NewToken(token.ASSIGN, lexer.character)
	case '+':
		tok = token.NewToken(token.PLUS, lexer.character)
	case '-':
		tok = token.NewToken(token.MINUS, lexer.character)
	case '*':
		tok = token.NewToken(token.ASTERISK, lexer.character)
	case '/':
		tok = token.NewToken(token.SLASH, lexer.character)
	case '<':
		tok = token.NewToken(token.INFERIOR, lexer.character)
	case '>':
		tok = token.NewToken(token.SUPERIOR, lexer.character)
	case ';':
		tok = token.NewToken(token.SEMICOLON, lexer.character)
	case '(':
		tok = token.NewToken(token.LEFT_PARENTHESIS, lexer.character)
	case ')':
		tok = token.NewToken(token.RIGHT_PARENTHESIS, lexer.character)
	case '{':
		tok = token.NewToken(token.LEFT_BRACE, lexer.character)
	case '}':
		tok = token.NewToken(token.RIGHT_BRACE, lexer.character)
	case ',':
		tok = token.NewToken(token.COMMA, lexer.character)
	case 0:
		tok = token.NewToken(token.END_OF_FILE, "")
	}
	lexer.ReadCharacter()
	return tok
}
