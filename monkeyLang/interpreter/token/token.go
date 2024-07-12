package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, character byte) Token {
	return Token{
		Type:    tokenType,
		Literal: string(character),
	}
}
