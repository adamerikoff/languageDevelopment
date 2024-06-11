package parser

import (
	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/lexer"
	"github.com/adamerikoff/ponGo/src/token"
)

type Parser struct {
	lexer         *lexer.Lexer
	currentToken  token.TokenInstance
	expectedToken token.TokenInstance
	errors        []string
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer:  lexer,
		errors: []string{},
	}
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.expectedToken
	parser.expectedToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		statement := parser.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		parser.nextToken()
	}
	return program
}
