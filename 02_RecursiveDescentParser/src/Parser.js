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


    Statement(){
        switch (this._lookahead.type) {
            case ";": 
                return this.EmptyStatement();
            case "{": 
                return this.BlockStatement();
            default:
                return this.ExpressionStatement();
        }
    }

    EmptyStatement() {
        this._eat(";");
        return factory.EmptyStatement();
    }

    BlockStatement() {
        this._eat("{");
        const body = this._lookahead.type !== "}" ? this.StatementList("}"): [];
        this._eat("}");
        return factory.BlockStatement(body);
    }

    ExpressionStatement() {
        const expression = this.Expression();
        this._eat(";");
        return factory.ExpressionStatement(expression);
    }

    Expression() {
        return this.AdditiveExpression();
    }

    AdditiveExpression() {
        return this._BinaryExpression(
            "MultiplicativeExpression",
            "ADDITIVE_OPERATOR"
        )
    }

    MultiplicativeExpression() {
        return this._BinaryExpression(
            "PrimaryExpression",
            "MULTIPLICATIVE_OPERATOR"
        )
    }

    _BinaryExpression(builderName, operatorToken) {
        let left = this[builderName]();

        while (this._lookahead != null && this._lookahead.type === operatorToken) {
            const operator = this._eat(operatorToken).value;
            const right = this[builderName]();
            left = {
                type: "BinaryExpression",
                operator,
                left,
                right
            };
        }
        return left;
    }

    PrimaryExpression() {
        switch (this._lookahead.type) {
            case "(":
                return this.ParenthesizedExpression();
            default:
                return this.Literal();
        }
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
        throw new SyntaxError(`LITERAL: Unexpected literal production "${this._lookahead.type}" !`);
    }

    StringLiteral() {
        const token = this._eat("STRING");
        return factory.StringLiteral(token.value.slice(1,-1));
    }

    NumericLiteral() {
        const token = this._eat("NUMBER");
        return factory.NumericLiteral(Number(token.value));
    }

    _eat(tokenType) {
        const token = this._lookahead;

        if (token == null) {
            throw new SyntaxError(`Unexpected end of input, expected: "${tokenType}" !`);
        }

        if (token.type != tokenType) {
            throw new SyntaxError(`Unexpected token: "${token.value}" , expected: "${tokenType}" !`);
        }

        this._lookahead = this._tokenizer.getNextToken();

        return token;
    }
}

module.exports = {
    Parser,
}