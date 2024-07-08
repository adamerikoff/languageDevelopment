class Parser {

    parse(string) {
        this._string = string;
    }

    Program() {
        return this.NumericLiteral();
    }

    NumericLiteral() {
        return {
            type: 'NumericLiteral',
            value: Number(this._string)
        };
    }
}

export default Parser;