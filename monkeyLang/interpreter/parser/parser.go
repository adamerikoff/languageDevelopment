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
	INDEX       // array[index]
)

var precedences = map[token.TokenType]int{
	token.EQUAL:            EQUALS,
	token.NOT_EQUAL:        EQUALS,
	token.LESS_THAN:        LESSGREATER,
	token.GREATER_THAN:     LESSGREATER,
	token.PLUS:             SUM,
	token.MINUS:            SUM,
	token.SLASH:            PRODUCT,
	token.STAR:             PRODUCT,
	token.LEFT_PARENTHESIS: CALL,
	token.LBRACKET:         INDEX,
}

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
	parser.registerPrefix(token.TRUE, parser.parseBoolean)
	parser.registerPrefix(token.FALSE, parser.parseBoolean)
	parser.registerPrefix(token.LEFT_PARENTHESIS, parser.parseGroupedExpression)
	parser.registerPrefix(token.IF, parser.parseIfExpression)
	parser.registerPrefix(token.FUNCTION, parser.parseFunctionLiteral)
	parser.registerPrefix(token.STRING, parser.parseStringLiteral)
	parser.registerPrefix(token.LBRACKET, parser.parseArrayLiteral)
	parser.registerPrefix(token.LEFT_BRACE, parser.parseHashLiteral)
	parser.registerPrefix(token.MACRO, parser.parseMacroLiteral)

	parser.infixParseFns = make(map[token.TokenType]infixParseFn)
	parser.registerInfix(token.PLUS, parser.parseInfixExpression)
	parser.registerInfix(token.MINUS, parser.parseInfixExpression)
	parser.registerInfix(token.SLASH, parser.parseInfixExpression)
	parser.registerInfix(token.STAR, parser.parseInfixExpression)
	parser.registerInfix(token.EQUAL, parser.parseInfixExpression)
	parser.registerInfix(token.NOT_EQUAL, parser.parseInfixExpression)
	parser.registerInfix(token.LESS_THAN, parser.parseInfixExpression)
	parser.registerInfix(token.GREATER_THAN, parser.parseInfixExpression)
	parser.registerInfix(token.LEFT_PARENTHESIS, parser.parseCallExpression)
	parser.registerInfix(token.LBRACKET, parser.parseIndexExpression)

	parser.nextToken()
	parser.nextToken()

	return parser
}

func (parser *Parser) parseMacroLiteral() ast.Expression {
	lit := &ast.MacroLiteral{Token: parser.currentToken}
	if !parser.expectPeek(token.LEFT_PARENTHESIS) {
		return nil
	}
	lit.Parameters = parser.parseFunctionParameters()
	if !parser.expectPeek(token.LEFT_BRACE) {
		return nil
	}
	lit.Body = parser.parseBlockStatement()
	return lit
}

func (parser *Parser) parseHashLiteral() ast.Expression {
	hash := &ast.HashLiteral{Token: parser.currentToken}
	hash.Pairs = make(map[ast.Expression]ast.Expression)
	for !parser.peekTokenIs(token.RIGHT_BRACE) {
		parser.nextToken()
		key := parser.parseExpression(LOWEST)
		if !parser.expectPeek(token.COLON) {
			return nil
		}
		parser.nextToken()
		value := parser.parseExpression(LOWEST)
		hash.Pairs[key] = value
		if !parser.peekTokenIs(token.RIGHT_BRACE) && !parser.expectPeek(token.COMMA) {
			return nil
		}
	}
	if !parser.expectPeek(token.RIGHT_BRACE) {
		return nil
	}
	return hash
}

func (parser *Parser) parseIndexExpression(left ast.Expression) ast.Expression {
	exp := &ast.IndexExpression{Token: parser.currentToken, Left: left}
	parser.nextToken()
	exp.Index = parser.parseExpression(LOWEST)
	if !parser.expectPeek(token.RBRACKET) {
		return nil
	}
	return exp
}

func (parser *Parser) parseArrayLiteral() ast.Expression {
	array := &ast.ArrayLiteral{Token: parser.currentToken}
	array.Elements = parser.parseExpressionList(token.RBRACKET)
	return array
}

func (parser *Parser) parseExpressionList(end token.TokenType) []ast.Expression {
	list := []ast.Expression{}
	if parser.peekTokenIs(end) {
		parser.nextToken()
		return list
	}
	parser.nextToken()
	list = append(list, parser.parseExpression(LOWEST))
	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		list = append(list, parser.parseExpression(LOWEST))
	}
	if !parser.expectPeek(end) {
		return nil
	}
	return list
}

func (parser *Parser) parseStringLiteral() ast.Expression {
	return &ast.StringLiteral{Token: parser.currentToken, Value: parser.currentToken.Literal}
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

func (parser *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: parser.currentToken}
	parser.nextToken()
	stmt.ReturnValue = parser.parseExpression(LOWEST)
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
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
	parser.nextToken()
	stmt.Value = parser.parseExpression(LOWEST)
	if parser.peekTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}
	return stmt
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
	for !parser.peekTokenIs(token.SEMICOLON) && precedence < parser.peekPrecedence() {
		infix := parser.infixParseFns[parser.peekToken.Type]
		if infix == nil {
			return leftExp
		}
		parser.nextToken()
		leftExp = infix(leftExp)
	}
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

func (parser *Parser) parseInfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.InfixExpression{
		Token:    parser.currentToken,
		Operator: parser.currentToken.Literal,
		Left:     left,
	}
	precedence := parser.curPrecedence()
	parser.nextToken()
	expression.Right = parser.parseExpression(precedence)
	return expression
}

func (parser *Parser) parseBoolean() ast.Expression {
	return &ast.Boolean{Token: parser.currentToken, Value: parser.currentTokenIs(token.TRUE)}
}

func (parser *Parser) parseGroupedExpression() ast.Expression {
	parser.nextToken()
	exp := parser.parseExpression(LOWEST)
	if !parser.expectPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return exp
}

func (parser *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: parser.currentToken}
	if !parser.expectPeek(token.LEFT_PARENTHESIS) {
		return nil
	}
	parser.nextToken()
	expression.Condition = parser.parseExpression(LOWEST)
	if !parser.expectPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}
	if !parser.expectPeek(token.LEFT_BRACE) {
		return nil
	}
	expression.Consequence = parser.parseBlockStatement()
	if parser.peekTokenIs(token.ELSE) {
		parser.nextToken()
		if !parser.expectPeek(token.LEFT_BRACE) {
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
	for !parser.currentTokenIs(token.RIGHT_BRACE) && !parser.currentTokenIs(token.END_OF_LINE) {
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
	if !parser.expectPeek(token.LEFT_PARENTHESIS) {
		return nil
	}
	lit.Parameters = parser.parseFunctionParameters()
	if !parser.expectPeek(token.LEFT_BRACE) {
		return nil
	}
	lit.Body = parser.parseBlockStatement()
	return lit
}

func (parser *Parser) parseFunctionParameters() []*ast.Identifier {
	identifiers := []*ast.Identifier{}

	if parser.peekTokenIs(token.RIGHT_PARENTHESIS) {
		parser.nextToken()
		return identifiers
	}
	parser.nextToken()

	ident := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
	identifiers = append(identifiers, ident)
	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		ident := &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}
		identifiers = append(identifiers, ident)
	}
	if !parser.expectPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return identifiers
}

func (parser *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	exp := &ast.CallExpression{Token: parser.currentToken, Function: function}
	exp.Arguments = parser.parseExpressionList(token.RIGHT_PARENTHESIS)
	return exp
}

func (parser *Parser) parseCallArguments() []ast.Expression {
	args := []ast.Expression{}
	if parser.peekTokenIs(token.RIGHT_PARENTHESIS) {
		parser.nextToken()
		return args
	}
	parser.nextToken()
	args = append(args, parser.parseExpression(LOWEST))
	for parser.peekTokenIs(token.COMMA) {
		parser.nextToken()
		parser.nextToken()
		args = append(args, parser.parseExpression(LOWEST))
	}
	if !parser.expectPeek(token.RIGHT_PARENTHESIS) {
		return nil
	}
	return args
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

func (parser *Parser) peekPrecedence() int {
	if parser, ok := precedences[parser.peekToken.Type]; ok {
		return parser
	}
	return LOWEST
}

func (parser *Parser) curPrecedence() int {
	if parser, ok := precedences[parser.currentToken.Type]; ok {
		return parser
	}
	return LOWEST
}
