package token

type TokenType string

type TokenInstance struct {
	Type    TokenType
	Literal string
}

func NewToken(tokenType TokenType, value interface{}) TokenInstance {
	var literal string
	switch v := value.(type) {
	case byte:
		literal = string(v)
	case string:
		literal = v
	default:
		panic("Unsupported value type for NewToken")
	}

	return TokenInstance{
		Type:    tokenType,
		Literal: literal,
	}
}
