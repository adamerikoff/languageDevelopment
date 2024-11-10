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

        // Recursively evaluate left and right operands to handle nested expressions
        const leftValue = this.eval(left);
        const rightValue = this.eval(right);

        // Ensure operands are numbers after evaluation
        if (!this.isNumber(leftValue) || !this.isNumber(rightValue)) {
            throw new Error(`Invalid operands for ${operator}: ${leftValue}, ${rightValue}`);
        }

        switch (operator) {
            case "+":
                return leftValue + rightValue;
            case "-":
                return leftValue - rightValue;
            case "*":
                return leftValue * rightValue;
            case "/":
                if (rightValue === 0) throw new Error("Division by zero error");
                return leftValue / rightValue;
            default:
                throw new Error(`Unsupported operator: ${operator}`);
        }
    }
}

module.exports = { Eva };
