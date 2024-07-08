import Tokenizer from "./Tokenizer.js";

class Parser {

    constructor() {
        this._tokenizer = new Tokenizer();
        this._string = "";
    }
    parse(string) {
        this._string = string;
        this._tokenizer.init(string);

        this._lookahead = this._tokenizer.getNextToken();
        return this.Program();
    }

    Program() {
        return {
            type: "Program",
            body: this.NumericLiteral(),
        };
    }

    NumericLiteral() {
        const token = this._eat("NUMBER");
        return {
            type: "NumericLiteral",
            value: Number(token.value)
        };
    }

    _eat(tokenType) {
        const token = this._lookahead;
        if (token == null) {
            throw new SyntaxError(
                `Unexpected end of input, expected: "${tokenType}"`,
            );
        }

        if (token.type !== tokenType) {
            throw new SyntaxError(
                `Unexpected token: "${token.value}", expected "${tokenType}"`,
            );
        }

        this._lookahead = this._tokenizer.getNextToken();

        return token;
    }
}

export default Parser;