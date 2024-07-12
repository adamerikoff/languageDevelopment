package lexer

type Lexer struct {
	input               string
	currentCharPosition int  // current position in input (points to current char)
	readingCharPosition int  // current reading position in input (after current char)
	character           byte // current char under examination
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
