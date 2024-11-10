const { Environment } = require("./Environment");

class Eva {
    constructor(global = new Environment()) {
        this.global = global;
    }

    eval(exp, env = this.global) {
        switch (true) {
            case this.isLiteral(exp):
                return this.parseLiteral(exp);
            case this.isBinaryExpression(exp):
                return this.parseBinaryExpression(exp, env);
            case this.isComparisonOperator(exp):
                return this.parseComparisonOperator(exp, env);
            case this.isVariableDeclaration(exp):
                return this.parseVariableDeclaration(exp, env);
            case this.isVariableRedeclaration(exp):
                return this.parseVariableRedeclaration(exp, env);
            case this.isVariable(exp):
                return env.lookup(exp);
            case this.isIf(exp):
                return this.parseIf(exp, env);
            case this.isWhile(exp):
                return this.parseWhile(exp, env);
            case this.isBlock(exp):
                return this.parseBlock(exp, env);
            default:
                throw new Error(`Unsupported expression type: ${JSON.stringify(exp)}`);
        }
    }

    // --- Type-checking Functions ---

    isWhile(exp) {
        return Array.isArray(exp) && exp[0] === "while";
    }

    isLiteral(exp) {
        return this.isNumber(exp) || this.isString(exp);
    }

    isNumber(exp) {
        return typeof exp === "number";
    }

    isString(exp) {
        return typeof exp === "string" && exp.startsWith(`"`) && exp.endsWith(`"`);
    }

    isBinaryExpression(exp) {
        return Array.isArray(exp) && ["+", "-", "*", "/"].includes(exp[0]);
    }

    isComparisonOperator(exp) {
        return Array.isArray(exp) && [">", "<", ">=", "<=", "==", "!="].includes(exp[0]);
    }

    isVariable(exp) {
        return typeof exp === "string" && /^[a-zA-Z][a-zA-Z0-9]*$/.test(exp);
    }

    isVariableDeclaration(exp) {
        return Array.isArray(exp) && exp[0] === "define";
    }

    isVariableRedeclaration(exp) {
        return Array.isArray(exp) && exp[0] === "assign";
    }

    isBlock(exp) {
        return Array.isArray(exp) && exp[0] === "begin";
    }

    isIf(exp) {
        return Array.isArray(exp) && exp[0] === "if";
    }

    // --- Parsing and Evaluation Functions ---

    parseLiteral(exp) {
        return this.isString(exp) ? exp.slice(1, -1) : exp;
    }

    parseBinaryExpression(exp, env) {
        const [operator, left, right] = exp;
        const leftValue = this.eval(left, env);
        const rightValue = this.eval(right, env);

        if (!this.isNumber(leftValue) || !this.isNumber(rightValue)) {
            throw new Error(`Invalid operands for ${operator}: ${leftValue}, ${rightValue}`);
        }

        switch (operator) {
            case "+": return leftValue + rightValue;
            case "-": return leftValue - rightValue;
            case "*": return leftValue * rightValue;
            case "/":
                if (rightValue === 0) throw new Error("Division by zero error");
                return leftValue / rightValue;
            default:
                throw new Error(`Unsupported operator: ${operator}`);
        }
    }

    parseComparisonOperator(exp, env) {
        const [operator, left, right] = exp;
        const leftValue = this.eval(left, env);
        const rightValue = this.eval(right, env);

        switch (operator) {
            case ">": return leftValue > rightValue;
            case "<": return leftValue < rightValue;
            case ">=": return leftValue >= rightValue;
            case "<=": return leftValue <= rightValue;
            case "!=": return leftValue != rightValue;
            case "==": return leftValue == rightValue;
            default:
                throw new Error(`Unsupported operator: ${operator}`);
        }
    }

    parseVariableDeclaration(exp, env) {
        const [_, name, value] = exp;
        const evaluatedValue = this.eval(value, env);
        return env.define(name, evaluatedValue);
    }

    parseVariableRedeclaration(exp, env) {
        const [_, name, value] = exp;
        const evaluatedValue = this.eval(value, env);
        return env.assign(name, evaluatedValue);
    }

    parseBlock(block, env) {
        const blockEnv = new Environment({}, env);
        let result;
        const [_tag, ...expressions] = block;
        expressions.forEach(exp => {
            result = this.eval(exp, blockEnv);
        });
        return result;
    }

    parseIf(exp, env) {
        const [_tag, condition, consequent, alternate] = exp;
        if (this.eval(condition)) {
            return this.eval(consequent, env);
        }
        return this.eval(alternate, env);
    }

    parseWhile(exp, env) {
        const [_tag, condition, body] = exp;
        let result;
        while (this.eval(condition, env)) {
            result = this.eval(body, env);
        }
        return result;
    }
}

module.exports = { Eva };
