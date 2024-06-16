package parser

import (
	"fmt"
	"strconv"

	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/lexer"
	"github.com/adamerikoff/ponGo/src/token"
)

type (
	prefixParseFn func() ast.Expression
	infixParseFn  func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer           *lexer.Lexer
	currentToken    token.Token
	subsequentToken token.Token

	errors []string

	prefixParseFns map[token.TokenType]prefixParseFn
	infixParseFns  map[token.TokenType]infixParseFn
}

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

var precedences = map[token.TokenType]int{
	token.EQUAL:            EQUALS,
	token.NOT_EQUAL:        EQUALS,
	token.INFERIOR:         LESSGREATER,
	token.SUPERIOR:         LESSGREATER,
	token.PLUS:             SUM,
	token.MINUS:            SUM,
	token.SLASH:            PRODUCT,
	token.ASTERISK:         PRODUCT,
	token.LEFT_PARENTHESIS: CALL,
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  lexer,
		errors: []string{},
	}

	parser.prefixParseFns = make(map[token.TokenType]prefixParseFn)
	parser.registerPrefix(token.IDENTIFIER, parser.parseIdentifier)
	parser.registerPrefix(token.INTEGER, parser.parseIntegerLiteral)
	parser.registerPrefix(token.EXCLAMATION, parser.parsePrefixExpression)
	parser.registerPrefix(token.MINUS, parser.parsePrefixExpression)
	parser.registerPrefix(token.TRUE, parser.parseBoolean)
	parser.registerPrefix(token.FALSE, parser.parseBoolean)
	parser.registerPrefix(token.LEFT_PARENTHESIS, parser.parseGroupedExpression)
	parser.registerPrefix(token.IF, parser.parseIfExpression)
	parser.registerPrefix(token.FUNCTION, parser.parseFunctionLiteral)

	parser.infixParseFns = make(map[token.TokenType]infixParseFn)
	parser.registerInfix(token.PLUS, parser.parseInfixExpression)
	parser.registerInfix(token.MINUS, parser.parseInfixExpression)
	parser.registerInfix(token.SLASH, parser.parseInfixExpression)
	parser.registerInfix(token.ASTERISK, parser.parseInfixExpression)
	parser.registerInfix(token.EQUAL, parser.parseInfixExpression)
	parser.registerInfix(token.NOT_EQUAL, parser.parseInfixExpression)
	parser.registerInfix(token.INFERIOR, parser.parseInfixExpression)
	parser.registerInfix(token.SUPERIOR, parser.parseInfixExpression)
	parser.registerInfix(token.LEFT_PARENTHESIS, parser.parseCallExpression)

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.subsequentToken
	parser.subsequentToken = parser.lexer.NextToken()
}

func (parser *Parser) currentTokenIs(t token.TokenType) bool {
	return parser.currentToken.Type == t
}

func (parser *Parser) subsequentTokenIs(t token.TokenType) bool {
	return parser.subsequentToken.Type == t
}

func (parser *Parser) expectSubsequentToken(t token.TokenType) bool {
	if parser.subsequentTokenIs(t) {
		parser.nextToken()
		return true
	} else {
		parser.expectSubsequentError(t)
		return false
	}
}

func (parser *Parser) Errors() []string {
	return parser.errors
}

func (parser *Parser) expectSubsequentError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, parser.subsequentToken.Type)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) noPrefixParseFnError(t token.TokenType) {
	msg := fmt.Sprintf("no prefix parse function for %s found", t)
	parser.errors = append(parser.errors, msg)
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for parser.currentToken.Type != token.END_OF_FILE {
		stmt := parser.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
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
	if !parser.expectSubsequentToken(token.IDENTIFIER) {
		return nil
	}
	stmt.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
	if !parser.expectSubsequentToken(token.ASSIGN) {
		return nil
	}
	parser.nextToken()
	stmt.Value = parser.parseExpression(LOWEST)
	if parser.subsequentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
}

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: parser.currentToken}
	parser.nextToken()
	stmt.ReturnValue = parser.parseExpression(LOWEST)
	if parser.subsequentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
}

func (parser *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: parser.currentToken}
	stmt.Expression = parser.parseExpression(LOWEST)
	if parser.subsequentTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
}

func (parser *Parser) parseExpression(precedence int) ast.Expression {
	prefix := parser.prefixParseFns[parser.currentToken.Type]
	if prefix == nil {
		parser.noPrefixParseFnError(parser.currentToken.Type)
		return nil
	}
	leftExp := prefix()
	for !parser.subsequentTokenIs(token.SEMICOLON) && precedence < parser.subsequentPrecedence() {
		infix := parser.infixParseFns[parser.subsequentToken.Type]
		if infix == nil {
			return leftExp
		}
		parser.nextToken()
		leftExp = infix(leftExp)
	}
	return leftExp
}

func (parser *Parser) subsequentPrecedence() int {
	if precedence, ok := precedences[parser.subsequentToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (parser *Parser) currentPrecedence() int {
	if precedence, ok := precedences[parser.currentToken.Type]; ok {
		return precedence
	}
	return LOWEST
}

func (parser *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{
		Token: parser.currentToken,
		Value: parser.currentToken.Literal,
	}
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

func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
		Left:     left,
	}
	precedence := parser.currentPrecedence()
	parser.nextToken()
	expression.Right = parser.parseExpression(precedence)
	return expression
}

func (parser *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{
		Token: parser.currentToken,
		Value: parser.currentTokenIs(token.TRUE),
	}
}

func (parser *Parser) parseGroupedExpression() ast.Expression {
	parser.nextToken()
	exp := parser.parseExpression(LOWEST)
	if !parser.expectSubsequentToken(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return exp
}

func (parser *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: parser.currentToken}
	if !parser.expectSubsequentToken(token.LEFT_PARENTHESIS) {
		return nil
	}
	parser.nextToken()
	expression.Condition = parser.parseExpression(LOWEST)
	if !parser.expectSubsequentToken(token.RIGHT_PARENTHESIS) {
		return nil
	}
	if !parser.expectSubsequentToken(token.LEFT_BRACE) {
		return nil
	}
	expression.Consequence = parser.parseBlockStatement()
	if parser.subsequentTokenIs(token.ELSE) {
		parser.nextToken()
		if !parser.expectSubsequentToken(token.LEFT_BRACE) {
			return nil
		}
		expression.Alternative = parser.parseBlockStatement()
	}
	return expression
}

func (parser *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: parser.currentToken}
	block.Statements = []ast.Statement{}
	parser.nextToken()
	for !parser.currentTokenIs(token.RIGHT_BRACE) && !parser.currentTokenIs(token.END_OF_FILE) {
		stmt := parser.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}
		parser.nextToken()
	}
	return block
}

func (parser *Parser) parseFunctionLiteral() ast.Expression {
	lit := &ast.FunctionLiteral{Token: parser.currentToken}
	if !parser.expectSubsequentToken(token.LEFT_PARENTHESIS) {
		return nil
	}
	lit.Parameters = parser.parseFunctionParameters()
	if !parser.expectSubsequentToken(token.LEFT_BRACE) {
		return nil
	}
	lit.Body = parser.parseBlockStatement()
	return lit
}

func (parser *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}
	if parser.subsequentTokenIs(token.RIGHT_PARENTHESIS) {
		parser.nextToken()
		return identifiers
	}
	parser.nextToken()
	ident := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
	identifiers = append(identifiers, ident)
	for parser.subsequentTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		ident := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
		identifiers = append(identifiers, ident)
	}
	if !parser.expectSubsequentToken(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return identifiers
}

func (parser *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: parser.currentToken, Function: function}
	exp.Arguments = parser.parseExpressionList(token.RIGHT_PARENTHESIS)
	return exp
}

func (parser *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if parser.subsequentTokenIs(end) {
		parser.nextToken()
		return list
	}
	parser.nextToken()
	list = append(list, parser.parseExpression(LOWEST))
	for parser.subsequentTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		list = append(list, parser.parseExpression(LOWEST))
	}
	if !parser.expectSubsequentToken(end) {
		return nil
	}
	return list
}

func (parser *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	parser.prefixParseFns[tokenType] = fn
}

func (parser *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	parser.infixParseFns[tokenType] = fn
}
