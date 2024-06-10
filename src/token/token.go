package token

type TokenType string

type TokenInstance struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, character byte) TokenInstance {
	return TokenInstance{
		Type:    tokenType,
		Literal: string(character),
	}
}
