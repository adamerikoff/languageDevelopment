package parser

import (
	"github.com/adamerikoff/ponGo/src/ast"
	"github.com/adamerikoff/ponGo/src/token"

	"github.com/adamerikoff/ponGo/src/lexer"
)

type Parser struct {
	lexer           *lexer.Lexer
	currentToken    token.Token
	subsequentToken token.Token
}

func NewParser(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}
	// Read two tokens, so curToken and peekToken are both set
	parser.nextToken()
	parser.nextToken()
	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.subsequentToken
	parser.subsequentToken = parser.lexer.NextToken()
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
	default:
		return nil
	}
}
