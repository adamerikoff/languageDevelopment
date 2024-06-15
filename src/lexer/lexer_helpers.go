package lexer

// isLetter checks if the provided byte is an ASCII letter or an underscore.
func isLetter(ch byte) bool {
	return ('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z') || ch == '_'
}

// isDigit checks if the provided byte is an ASCII digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
