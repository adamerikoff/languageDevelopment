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
	lexer.readCharacter()
	return lexer
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		if lexer.inspectNextCharacter() == '=' {
			ch := lexer.character
			lexer.readCharacter()
			literal := string(ch) + string(lexer.character)
			tok = token.NewToken(token.EQUAL, literal)
		} else {
			tok = token.NewToken(token.ASSIGN, lexer.character)
		}
	case '+':
		tok = token.NewToken(token.PLUS, lexer.character)
	case '-':
		tok = token.NewToken(token.MINUS, lexer.character)
	case '*':
		tok = token.NewToken(token.ASTERISK, lexer.character)
	case '/':
		tok = token.NewToken(token.SLASH, lexer.character)
	case '!':
		if lexer.inspectNextCharacter() == '=' {
			ch := lexer.character
			lexer.readCharacter()
			literal := string(ch) + string(lexer.character)
			tok = token.NewToken(token.NOT_EQUAL, literal)
		} else {
			tok = token.NewToken(token.EXCLAMATION, lexer.character)
		}
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
	default:
		if isLetter(lexer.character) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.ClassifyToken(tok.Literal)
			return tok
		} else if isDigit(lexer.character) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.INTEGER
			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, lexer.character)
		}
	}
	lexer.readCharacter()
	return tok
}

// isLetter checks if the provided byte is an ASCII letter or an underscore.
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// isDigit checks if the provided byte is an ASCII digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readCharacter() {
	switch {
	case lexer.readPosition >= len(lexer.input):
		lexer.character = 0
	default:
		lexer.character = lexer.input[lexer.readPosition]
		lexer.currentPosition = lexer.readPosition
		lexer.readPosition += 1
	}
}

func (lexer *Lexer) inspectNextCharacter() byte {
	switch {
	case lexer.readPosition >= len(lexer.input):
		return 0
	default:
		return lexer.input[lexer.readPosition]
	}
}

func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.currentPosition
	for isLetter(lexer.character) {
		lexer.readCharacter()
	}
	endPosition := lexer.currentPosition
	return lexer.input[startPosition:endPosition]
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.currentPosition
	for isDigit(lexer.character) {
		lexer.readCharacter()
	}
	endPosition := lexer.currentPosition
	return lexer.input[startPosition:endPosition]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readCharacter()
	}
}
