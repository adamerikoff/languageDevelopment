package gogo_parser

import (
	"github.com/adamerikoff/gogo/gogo_ast"
	"github.com/adamerikoff/gogo/gogo_lexer"
	"github.com/adamerikoff/gogo/gogo_token"
)

type Parser struct {
	l *gogo_lexer.Lexer

	currentToken gogo_token.Token
	peekToken    gogo_token.Token
}

func NewParser(l *gogo_lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currentTokenIs(t gogo_token.TokenType) bool {
	return p.currentToken.Type == t
}
func (p *Parser) peekTokenIs(t gogo_token.TokenType) bool {
	return p.peekToken.Type == t
}
func (p *Parser) expectPeek(t gogo_token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		return false
	}
}

func (p *Parser) parseLetStatement() *gogo_ast.LetStatement {
	statement := &gogo_ast.LetStatement{Token: p.currentToken}

	if !p.expectPeek(gogo_token.IDENTIFIER) {
		return nil
	}

	statement.Name = &gogo_ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(gogo_token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !p.currentTokenIs(gogo_token.SEMICOLON) {
		p.nextToken()
	}
	return statement

}

func (p *Parser) parseStatement() gogo_ast.Statement {
	switch p.currentToken.Type {
	case gogo_token.LET:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) ParseProgram() *gogo_ast.Program {
	program := &gogo_ast.Program{}
	program.Statements = []gogo_ast.Statement{}

	for p.currentToken.Type != gogo_token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}
