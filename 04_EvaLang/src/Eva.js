class Eva {
    eval(exp) {
        switch (true) {
            case this.isNumber(exp):
                return exp;
            case this.isString(exp):
                return this.parseString(exp);
            case this.isExpression(exp):
                return this.parseExpression(exp);
            default:
                throw new Error(`Unsupported expression type: ${typeof exp}`);
        }
    }

    isNumber(exp) {
        return typeof exp === "number";
    }

    isString(exp) {
        return typeof exp === "string" && exp.startsWith(`"`) && exp.endsWith(`"`);
    }

    isExpression(exp) {
        return Array.isArray(exp) && ["+", "-", "*", "/"].includes(exp[0]);
    }

    parseString(exp) {
        return exp.slice(1, -1); // Remove surrounding quotes
    }

    parseExpression(exp) {
        const [operator, left, right] = exp;

        if (!this.isNumber(left) || !this.isNumber(right)) {
            throw new Error(`Invalid operands for ${operator}: ${left}, ${right}`);
        }

        switch (operator) {
            case "+":
                return left + right;
            case "-":
                return left - right;
            case "*":
                return left * right;
            case "/":
                if (right === 0) throw new Error("Division by zero error");
                return left / right;
            default:
                throw new Error(`Unsupported operator: ${operator}`);
        }
    }
}

module.exports = { Eva };
