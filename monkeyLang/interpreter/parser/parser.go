package parser

import (
	"fmt"
	"github.com/adamerikoff/lesLanguesDevs/monkeyLang/interpreter/ast"
	"github.com/adamerikoff/lesLanguesDevs/monkeyLang/interpreter/lexer"
	"github.com/adamerikoff/lesLanguesDevs/monkeyLang/interpreter/token"
	"strconv"
)

const (
	_ int = iota
	LOWEST
	EQUALS      // ==
	LESSGREATER // > or <
	SUM         // +
	PRODUCT     // *
	PREFIX      // -X or !X
	CALL        // myFunction(X)
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}
func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}

func NewParser(l *lexer.Lexer) *Parser {
	parser := &Parser{lexer: l}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.IDENTIFIER, parser.parseIdentifier)
	parser.registerPrefix(token.INTEGER, parser.parseIntegerLiteral)
	parser.registerPrefix(token.EXCLAMATION, parser.parsePrefixExpression)
	parser.registerPrefix(token.MINUS, parser.parsePrefixExpression)

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.END_OF_LINE {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}
	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	case token.RETURN:
		return parser.parseReturnStatement()
	default:
		return parser.parseExpressionStatement()
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: parser.currentToken}
	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	statement := &ast.ReturnStatement{Token: parser.currentToken}
	parser.nextToken()
	for !parser.currentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return statement
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	statement := &ast.ExpressionStatement{Token: parser.currentToken}
	statement.Expression = parser.parseExpression(LOWEST)
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return statement
}
func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.currentToken.Type]
	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}
	leftExp := prefix()
	return leftExp
}

func (parser *Parser) parseIntegerLiteral() ast.Expression {
	lit := &ast.IntegerLiteral{Token: parser.currentToken}
	value, err := strconv.ParseInt(parser.currentToken.Literal, 0, 64)
	if err != nil {
		msg := fmt.Sprintf("could not parse %q as integer", parser.currentToken.Literal)
		parser.errors = append(parser.errors, msg)
		return nil
	}
	lit.Value = value
	return lit
}

func (parser *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
	}
	parser.nextToken()
	expression.Right = parser.parseExpression(PREFIX)
	return expression
}

func (parser *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) peekTokenIs(t token.TokenType) bool {
	return parser.peekToken.Type == t
}

func (parser *Parser) expectPeek(t token.TokenType) bool {
	if parser.peekTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.peekError(t)
		return false
	}
}

func (parser *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, parser.peekToken.Type)
	parser.errors = append(parser.errors, msg)
}
