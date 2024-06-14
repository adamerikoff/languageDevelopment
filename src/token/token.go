package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, value interface{}) Token {
	switch v := value.(type) {
	case string:
		return Token{Type: tokenType, Literal: v}
	case byte:
		return Token{Type: tokenType, Literal: string(v)}
	default:
		return Token{Type: tokenType, Literal: ""}
	}
}
