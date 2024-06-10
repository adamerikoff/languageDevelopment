package lexer

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
