class Eva {
    eval(expression) {
        if (this.isNumber(expression)) {
            return expression;
        }
        if (this.isString(expression)) {
            return expression.slice(1, -1);
        }
        if (expression[0] === '+') {
            return expression[1] + expression[2]
        }
        throw 'Unimplemented';
    }

    isNumber(expression) {
        return typeof expression === 'number';
    }

    isString(expression) {
        return typeof expression === 'string' && expression[0] === '"' && expression.slice(-1) === '"';
    }
}

export default Eva;
