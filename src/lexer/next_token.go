package lexer

import "github.com/adamerikoff/ponGo/src/token"

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
			tok = token.NewToken(token.ILLEGAL, "")
		}
	}
	lexer.readCharacter()
	return tok
}
