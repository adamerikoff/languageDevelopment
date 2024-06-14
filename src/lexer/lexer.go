package lexer

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
