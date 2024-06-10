package gogo_parser

import (
	"fmt"
	"strconv"

	"github.com/adamerikoff/gogo/gogo_ast"
	"github.com/adamerikoff/gogo/gogo_lexer"
	"github.com/adamerikoff/gogo/gogo_token"
)

const (
	_ int = iota
	LOWEST
	EQUALS  // == LESSGREATER // > or <
	SUM     // +
	PRODUCT // *
	PREFIX  // -X or !X
	CALL    // myFunction(X)
)

type Parser struct {
	l *gogo_lexer.Lexer

	currentToken gogo_token.Token
	peekToken    gogo_token.Token

	errors []string

	prefixParseFns map[gogo_token.TokenType]prefixParseFn
	infixParseFns  map[gogo_token.TokenType]infixParseFn
}

func (p *Parser) registerPrefix(tokenType gogo_token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}
func (p *Parser) registerInfix(tokenType gogo_token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}

func NewParser(l *gogo_lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []string{},
	}
	p.prefixParseFns = make(map[gogo_token.TokenType]prefixParseFn)
	p.registerPrefix(gogo_token.IDENTIFIER, p.parseIdentifier)
	p.registerPrefix(gogo_token.INTEGER, p.parseIntegerLiteral)
	p.registerPrefix(gogo_token.EXCLAMATION, p.parsePrefixExpression)
	p.registerPrefix(gogo_token.MINUS, p.parsePrefixExpression)
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) parsePrefixExpression() gogo_ast.Expression {
	expression := &gogo_ast.PrefixExpression{
		Token:    p.currentToken,
		Operator: p.currentToken.Literal,
	}
	p.nextToken()
	expression.Right = p.parseExpression(PREFIX)
	return expression
}

func (p *Parser) parseIdentifier() gogo_ast.Expression {
	return &gogo_ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t gogo_token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
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
		p.peekError(t)
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

func (p *Parser) parseReturnStatement() *gogo_ast.ReturnStatement {
	statement := &gogo_ast.ReturnStatement{Token: p.currentToken}
	p.nextToken()

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
	case gogo_token.RETURN:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
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

type (
	prefixParseFn func() gogo_ast.Expression
	infixParseFn  func(gogo_ast.Expression) gogo_ast.Expression
)

func (p *Parser) noPrefixParseFnError(t gogo_token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	p.errors = append(p.errors, msg)
}

func (p *Parser) parseExpression(precedence int) gogo_ast.Expression {
	prefix := p.prefixParseFns[p.currentToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.currentToken.Type)
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (p *Parser) parseExpressionStatement() *gogo_ast.ExpressionStatement {
	statement := &gogo_ast.ExpressionStatement{Token: p.currentToken}
	statement.Expression = p.parseExpression(LOWEST)
	if p.peekTokenIs(gogo_token.SEMICOLON) {
		p.nextToken()
	}
	return statement
}

func (p *Parser) parseIntegerLiteral() gogo_ast.Expression {
	lit := &gogo_ast.IntegerLiteral{Token: p.currentToken}
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", p.currentToken.Literal)
		p.errors = append(p.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}
