import Env from './Env.js';

class Eva {

    constructor(global = new Env()) {
        this.global = global;
    }

    eval(expression, env = this.global) {
        if (this.isNumber(expression)) {
            return expression;
        }
        if (this.isString(expression)) {
            return expression.slice(1, -1);
        }
        // ----------------------------------
        // ----------------------------------
        if (Array.isArray(expression)) {
            switch (expression[0]) {
                case '+':
                    return this.eval(expression[1], env) + this.eval(expression[2], env);
                case '-':
                    return this.eval(expression[1], env) - this.eval(expression[2], env);
                case '/':
                    return this.eval(expression[1], env) / this.eval(expression[2], env);
                case '*':
                    return this.eval(expression[1], env) * this.eval(expression[2], env);
                case '>':
                    return this.eval(expression[1], env) > this.eval(expression[2], env);
                case '>=':
                    return this.eval(expression[1], env) >= this.eval(expression[2], env);
                case '<':
                    return this.eval(expression[1], env) < this.eval(expression[2], env);
                case '<=':
                    return this.eval(expression[1], env) <= this.eval(expression[2], env);
                case '=':
                    return this.eval(expression[1], env) === this.eval(expression[2], env);
                case 'assign':
                    const [_assign, assignName, assignValue] = expression;
                    return env.eAssign(assignName, this.eval(assignValue, env));
                case 'reassign':
                    const [_reassign, reassignName, reassignValue] = expression;
                    return env.eReassign(reassignName, this.eval(reassignValue, env));
                case 'section':
                    const sectionEnv = new Env({}, env);
                    return this.evalSection(expression, sectionEnv);
                case 'assume':
                    const [_assume, assumeCondition, assumeConsequent, assumeAlternate] = expression;
                    if (this.eval(assumeCondition, env)) {
                        return this.eval(assumeConsequent, env);
                    }
                    return this.eval(assumeAlternate, env);
                case 'loop':
                    const [_loop, loopCondition, loopBody] = expression;
                    let loopResult;
                    while (this.eval(loopCondition, env)) {
                        loopResult = this.eval(loopBody, env);
                    }
                    return loopResult;
                default:
                    throw new Error(`Unimplemented ${JSON.stringify(expression)}`);
            }
        }
        // ----------------------------------
        // ----------------------------------
        if (this.isVariable(expression)) {
            return env.getVariableValue(expression);
        }
        // ----------------------------------
        // ----------------------------------
        throw new Error(`Invalid expression: ${JSON.stringify(expression)}`);
    }
    // ----------------------------------
    // ----------------------------------
    // ----------------------------------
    // ----------------------------------
    // ----------------------------------
    // ----------------------------------
    isNumber(expression) {
        return typeof expression === 'number';
    }

    isString(expression) {
        return typeof expression === 'string' && expression[0] === '"' && expression.slice(-1) === '"';
    }

    isVariable(expression) {
        return typeof expression === 'string' && /^[a-zA-Z_][a-zA-Z0-9_]*$/.test(expression);
    }

    evalSection(section, env) {
        let result;
        const [_tag, ...expressions] = section;
        expressions.forEach(expression => {
            result = this.eval(expression, env);
        });
        return result;
    }
}

export default Eva;
