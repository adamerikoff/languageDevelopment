package lexer

func isLetter(character byte) bool {
	isLowercase := 'a' <= character && character <= 'z'
	isUppercase := 'A' <= character && character <= 'Z'
	isUnderscore := character == '_'
	isAt := character == '@'
	isCircumflex := character == '^'

	return isLowercase || isUppercase || isUnderscore || isAt || isCircumflex
}

func (lexer *Lexer) skipWhitespace() {
	for isWhitespace(lexer.currentChar) {
		lexer.readCharacter()
	}
}

func isWhitespace(character byte) bool {
	return character == ' ' || character == '\t' || character == '\n' || character == '\r'
}

func isDigit(character byte) bool {
	return '0' <= character && character <= '9'
}
