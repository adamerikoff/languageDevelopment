package parser

import (
	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/lexer"
	"github.com/adamerikoff/ponGo/src/token"
)

type (
	prefixParseFunction func() ast.Expression
	inflixParseFunction func(ast.Expression) ast.Expression
)

type Parser struct {
	lexer         *lexer.Lexer
	currentToken  token.TokenInstance
	expectedToken token.TokenInstance
	errors        []string

	prefixParseFunction map[token.TokenType]prefixParseFunction
	inflixParseFunction map[token.TokenType]inflixParseFunction
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

func (parser *Parser) registerPrefix(tokenType token.TokenType, function prefixParseFunction) {
	parser.prefixParseFunction[tokenType] = function
}
func (parser *Parser) registerInflix(tokenType token.TokenType, function inflixParseFunction) {
	parser.inflixParseFunction[tokenType] = function
}
