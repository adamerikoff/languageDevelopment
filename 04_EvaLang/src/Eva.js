const { Environment } = require("./Environment");

class Eva {
    constructor(global = new Environment()) {
        this.global = global;
    }

    eval(exp, env = this.global) {
        switch (true) {
            case this.isNumber(exp):
                return exp;
            case this.isString(exp):
                return this.parseString(exp);
            case this.isExpression(exp):
                return this.parseExpression(exp, env);
            case this.isVariableDeclaration(exp):
                return this.parseVariableDeclaration(exp, env);
            case this.isVariableName(exp):
                return this.parseVariable(exp, env);
            default:
                throw new Error(`Unsupported expression type: ${JSON.stringify(exp)}`);
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
    
    isVariableName(exp) {
        return typeof exp === "string" && /^[a-zA-Z][a-zA-Z0-9]*$/.test(exp);
    }

    isVariableDeclaration(exp) {
        return exp[0] === "declare";
    }

    parseString(exp) {
        return exp.slice(1, -1); // Remove surrounding quotes
    }

    parseVariableDeclaration(exp, env) {
        const [_, name, value] = exp;
        const evaluatedValue = this.eval(value, env);
        return env.define(name, evaluatedValue);
    }
    
    parseVariable(exp, env) {
        return env.lookup(exp);
    }

    parseExpression(exp, env) {
        const [operator, left, right] = exp;

        // Recursively evaluate left and right operands to handle nested expressions
        const leftValue = this.eval(left, env);
        const rightValue = this.eval(right, env);

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
