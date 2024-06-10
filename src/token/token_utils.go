package token

func LookupIdentifier(identifier string) TokenType {
	if tok, ok := keywords[identifier]; ok {
		return tok
	}
	return IDENTIFIER
}
