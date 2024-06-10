package lexer

func (lexer *Lexer) peekCharacter() byte {
	if lexer.readPosition >= len(lexer.input) {
		return 0
	} else {
		return lexer.input[lexer.readPosition]
	}
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

func (lexer *Lexer) readIdentifier() string {
	position := lexer.position
	for isLetter(lexer.currentChar) {
		lexer.readCharacter()
	}
	return lexer.input[position:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	position := lexer.position
	for isDigit(lexer.currentChar) {
		lexer.readCharacter()
	}
	return lexer.input[position:lexer.position]
}
