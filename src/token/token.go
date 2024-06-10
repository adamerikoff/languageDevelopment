package token

type TokenType string

type TokenInstance struct {
	Type    TokenType
	Literal string
}
