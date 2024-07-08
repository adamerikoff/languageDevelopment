class Parser {

    parse(string) {
        this._string = string;

        return this.Program();
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