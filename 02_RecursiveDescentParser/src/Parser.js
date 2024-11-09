const { Tokenizer } = require("./Tokenizer");
const { DefaultFactory, SExpressionFactory } = require("./Factories");

const AST_MODE = "default";
const factory = AST_MODE === "default" ? DefaultFactory : SExpressionFactory;

class Parser {
    constructor() {
        this._string = "";
        this._tokenizer = new Tokenizer();
    }

    parse(string) {
        this._string = string;
        this._tokenizer.init(string);
        this._lookahead = this._tokenizer.getNextToken();
        return this.Program();
    }

    Program() {
        return factory.Program(this.StatementList());
    }

    StatementList(stopLookahead = null) {
        const statementList = [this.Statement()];
        while (this._lookahead != null && this._lookahead.type !== stopLookahead) {
            statementList.push(this.Statement());
        }
        return statementList;
    }

    Statement() {
        switch (this._lookahead.type) {
            case ";": 
                return this.EmptyStatement();
            case "if": 
                return this.IfStatement();
            case "{": 
                return this.BlockStatement();
            case "let": 
                return this.VariableStatement();
            default:
                return this.ExpressionStatement();
        }
    }

    IfStatement() {
        this._eat("if");
    
        // Parse the test condition inside parentheses
        this._eat("(");
        const test = this.Expression();
        this._eat(")");
    
        // Parse the consequent (the "then" block)
        const consequent = this.Statement();
    
        // Parse the optional else clause
        let alternate = null;
        if (this._lookahead != null && this._lookahead.type === "else") {
            this._eat("else"); // Consume "else"
            alternate = this.Statement();
        }
    
        return factory.IfStatement(test, consequent, alternate);
    }

    VariableStatement() {
        this._eat("let");
        const declarations = this.VariableDeclarationList();
        this._eat(";");
        return factory.VariableStatement(declarations);
    }

    VariableDeclarationList() {
        const declarations = [];
        do {
            declarations.push(this.VariableDeclaration());
        } while (this._lookahead.type === "," && this._eat(","));
        return declarations;
    }

    VariableDeclaration() {
        const id = this.Identifier();
        const init = this._lookahead && this._lookahead.type === "SIMPLE_ASSIGN" 
            ? this.VariableInitializer() 
            : null;
        return factory.VariableDeclaration(id, init);
    }

    VariableInitializer() {
        this._eat("SIMPLE_ASSIGN");
        return this.AssignmentExpression();
    }

    EmptyStatement() {
        this._eat(";");
        return factory.EmptyStatement();
    }

    BlockStatement() {
        this._eat("{");
        const body = this._lookahead.type !== "}" ? this.StatementList("}") : [];
        this._eat("}");
        return factory.BlockStatement(body);
    }

    ExpressionStatement() {
        const expression = this.Expression();
        this._eat(";");
        return factory.ExpressionStatement(expression);
    }

    Expression() {
        return this.AssignmentExpression();
    }

    AssignmentExpression() {
        const left = this.RelationalExpression();
        if (!this._isAssignmentOperator(this._lookahead.type)) {
            return left;
        }
        const operator = this.AssignmentOperator().value;
        const right = this.AssignmentExpression();
        return factory.AssignmentExpression(operator, this._checkValidAssignmentTarget(left), right);
    }

    LeftHandSideExpression() {
        return this.Identifier();
    }

    Identifier() {
        const name = this._eat("IDENTIFIER").value;
        return factory.Identifier(name);
    }

    _checkValidAssignmentTarget(node) {
        if (node.type === "Identifier") {
            return node;
        }
        throw new SyntaxError(`Invalid left-hand side in assignment expression!`);
    }

    _isAssignmentOperator(tokenType) {
        return tokenType === "SIMPLE_ASSIGN" || tokenType === "COMPLEX_ASSIGN";
    }

    AssignmentOperator() {
        if (this._lookahead.type === "SIMPLE_ASSIGN") {
            return this._eat("SIMPLE_ASSIGN");
        }
        return this._eat("COMPLEX_ASSIGN");
    }

    AdditiveExpression() {
        return this._BinaryExpression("MultiplicativeExpression", "ADDITIVE_OPERATOR");
    }

    MultiplicativeExpression() {
        return this._BinaryExpression("PrimaryExpression", "MULTIPLICATIVE_OPERATOR");
    }

    _BinaryExpression(builderName, operatorToken) {
        let left = this[builderName]();
        while (this._lookahead != null && this._lookahead.type === operatorToken) {
            const operator = this._eat(operatorToken).value;
            const right = this[builderName]();
            left = factory.BinaryExpression(operator, left, right);
        }
        return left;
    }

    RelationalExpression() {
        return this._BinaryExpression("AdditiveExpression", "RELATIONAL_OPERATOR");
    }

    PrimaryExpression() {
        if (this._isLiteral(this._lookahead.type)) {
            return this.Literal();
        }
        switch (this._lookahead.type) {
            case "(":
                return this.ParenthesizedExpression();
            default:
                return this.LeftHandSideExpression();
        }
    }

    _isLiteral(tokenType) {
        return tokenType === "NUMBER" || tokenType === "STRING";
    }

    ParenthesizedExpression() {
        this._eat("(");
        const expression = this.Expression();
        this._eat(")");
        return expression;
    }

    Literal() {
        switch (this._lookahead.type) {
            case "NUMBER":
                return this.NumericLiteral();
            case "STRING":
                return this.StringLiteral();
        }
        throw new SyntaxError(`Unexpected literal "${this._lookahead.type}"!`);
    }

    StringLiteral() {
        const token = this._eat("STRING");
        return factory.StringLiteral(token.value.slice(1, -1));
    }

    NumericLiteral() {
        const token = this._eat("NUMBER");
        return factory.NumericLiteral(Number(token.value));
    }

    _eat(tokenType) {
        const token = this._lookahead;
        if (token == null) {
            throw new SyntaxError(`Unexpected end of input, expected: "${tokenType}"!`);
        }
        if (token.type !== tokenType) {
            throw new SyntaxError(`Unexpected token: "${token.value}", expected: "${tokenType}"!`);
        }
        this._lookahead = this._tokenizer.getNextToken();
        return token;
    }
}

module.exports = {
    Parser,
};
