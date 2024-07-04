import fs from "fs";

import Env from './Env.js';
import yyparse from './aileenParser.js';

const ExecutionStack = [];

const GlobalEnv = new Env({
    ZERO: null,

    TRUE: true,
    FALSE: false,

    VERSION: "0.2",

    "+"(op1, op2) {
        return op1 + op2;
    },
    "-"(op1, op2) {
        if (!op2) {
            return -op1;
        }
        return op1 - op2;
    },
    "*"(op1, op2) {
        return op1 * op2;
    },
    "/"(op1, op2) {
        return op1 / op2;
    },
    "<"(op1, op2) {
        return op1 < op2;
    },
    "<="(op1, op2) {
        return op1 <= op2;
    },
    ">"(op1, op2) {
        return op1 > op2;
    },
    ">="(op1, op2) {
        return op1 >= op2;
    },
    "="(op1, op2) {
        return op1 === op2;
    },
    modulus(op1, op2) {
        return op1 % op2;
    },
    display(...args) {
        console.log(...args);
    },
    print_stack_trace() {
        ExecutionStack.forEach((element, index, array) => {
            if (index !== array.length - 1) {
                console.log(element);
            }
        });
    },
});

class Aileen {
    constructor (globalEnv = GlobalEnv) {
        this.globalEnv = globalEnv;
    }

    evaluateGlobal(expression) {
        return this._evaluateBody(expression, this.globalEnv);
    }

    evaluateExpression(expression, env = this.globalEnv) {
        if (this._isBoolean(expression)) {
            return expression;
        }

        if (this._isNumber(expression)) {
            return expression;
        }

        if (this._isString(expression)) {
            return expression.slice(1, -1);
        }

        if (expression[0] === "++") {
            const [_, name] = expression;
            return env.reassignVariable(name, env.retrieveVariable(name) + 1);
        }

        if (expression[0] === "--") {
            const [_, name] = expression;
            return env.reassignVariable(name, env.retrieveVariable(name) - 1);
        }

        if (expression[0] === "OR") {
            return this.evaluateExpression(expression[1], env) || this.evaluateExpression(expression[2], env);
        }

        if (expression[0] === "AND") {
            return this.evaluateExpression(expression[1], env) && this.evaluateExpression(expression[2], env);
        }

        if (expression[0] === "NOT") {
            return !this.evaluateExpression(expression[1], env);
        }

        if (expression[0] === "SECTION") {
            const blockEnv = new Env({}, env);
            return this._evaluateBlock(expression, blockEnv);
        }

        if (expression[0] === "ASSIGN") {
            const [_, name, value] = expression;
            return env.assignVariable(name, this.evaluateExpression(value, env));
        }

        if (expression[0] === "REASSIGN") {
            const [_, ref, value] = expression;
            if (ref[0] === "PROPERTY") {
                const [_tag, instance, propName] = ref;
                const instanceEnv = this.evaluateExpression(instance, env);
                return instanceEnv.assignVariable(propName, this.evaluateExpression(value, env));
            }
            return env.reassignVariable(ref, this.evaluateExpression(value, env));
        }

        if (this._isVariableName(expression)) {
            return env.retrieveVariable(expression);
        }

        if (expression[0] === "CONDITION") {
            const [_tag, condition, consequent, alternate] = expression;
            if (this.evaluateExpression(condition, env)) {
                return this.evaluateExpression(consequent, env);
            } else {
                return this.evaluateExpression(alternate, env);
            }
        }

        if (expression[0] === "LOOP") {
            const [_tag, condition, innerExp] = expression;
            let result;
            while (this.evaluateExpression(condition, env)) {
                result = this.evaluateExpression(innerExp, env);
            }
            return result;
        }

        if (expression[0] === "FUNCTION") {
            const assignExp = this._transformFunctionToVarLambda(expression);
            return this.evaluateExpression(assignExp, env);
        }

        if (expression[0] === "SWITCH") {
            const conditionExp = this._transformSwitchToCondition(expression);
            return this.evaluateExpression(conditionExp, env);
        }

        if (expression[0] === "LOOPFOR") {
            const loopExp = this._transformForToWhile(expression);
            return this.evaluateExpression(loopExp, env);
        }

        if (expression[0] === "LAMBDA") {
            const [_tag, params, body] = expression;
            return {
                params,
                body,
                env,
            };
        }

        if (expression[0] === "CLASS") {
            const [_tag, name, parent, body] = expression;
            const parentEnv = this.evaluateExpression(parent, env) || env;
            const classEnv = new Env({}, parentEnv);
            this._evaluateBody(body, classEnv);
            return env.assignVariable(name, classEnv);
        }

        if (expression[0] === "PARENT") {
            const [_tag, className] = expression;
            return this.evaluateExpression(className, env).parentEnv;
        }

        if (expression[0] === "NEW") {
            const classEnv = this.evaluateExpression(expression[1], env);
            const instanceEnv = new Env({}, classEnv);
            const args = expression.slice(2).map((arg) => this.evaluateExpression(arg, env));
            this._callCustomFunction(classEnv.retrieveVariable("INIT"), [
                instanceEnv,
                ...args,
            ]);
            return instanceEnv;
        }

        if (expression[0] === "PROPERTY") {
            const [_tag, instance, name] = expression;
            const instanceEnv = this.evaluateExpression(instance, env);
            return instanceEnv.retrieveVariable(name);
        }

        if (expression[0] === "MODULE") {
            const [_tag, name, body] = expression;
            const moduleEnv = new Env({}, env);
            this._evaluateBody(body, moduleEnv);
            return env.assignVariable(name, moduleEnv);
        }

        if (expression[0] === "LOAD") {
            const [_tag, symbols, name] =
                expression.length === 2 ? [null, null, expression[1]] : expression;
            const moduleSrc = fs.readFileSync(`./LIBRARIES/${name}.aileen`, "utf-8");
            const body = yyparse.parse(`(SECTION ${moduleSrc})`);
            let moduleExp;
            if (expression.length === 2) {
                moduleExp = ["MODULE", name, body];
            } else {
                const filteredBody = body.filter(
                    (element, index) => index === 0 || symbols.includes(element[1])
                );
                moduleExp = ["MODULE", name, filteredBody];
            }
            return this.evaluateExpression(moduleExp, this.globalEnv);
        }

        if (Array.isArray(expression)) {
            let result;
            ExecutionStack.push(expression[0]);
            const fn = this.evaluateExpression(expression[0], env);
            const args = expression.slice(1).map((arg) => this.evaluateExpression(arg, env));
            if (typeof fn === "function") {
                result = fn(...args);
            } else {
                result = this._callCustomFunction(fn, args);
            }
            ExecutionStack.pop();
            return result;
        }

        throw new EvalError(`Unimplemented: ${JSON.stringify(expression)}`);
    }
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    _transformFunctionToVarLambda(functionExp) {
        const [_tag, name, params, body] = functionExp;
        return ["ASSIGN", name, ["LAMBDA", params, body]];
    }

    _transformSwitchToCondition(switchExp) {
        const [_tag, ...cases] = switchExp;
        const ifExp = ["CONDITION", null, null, null];
        let current = ifExp;
        for (let i = 0; i < cases.length - 1; i++) {
            const [currentCond, currentBlock] = cases[i];
            current[1] = currentCond;
            current[2] = currentBlock;
            const [nextCond, nextBlock] = cases[i + 1];
            current[3] = nextCond === "ELSE" ? nextBlock : ["CONDITION", null, null, null];
            current = current[3];
        }
        return ifExp;
    }

    _transformForToWhile(forExp) {
        const [_tag, init, cond, modifier, exp] = forExp;
        return ["SECTION", init, ["LOOP", cond, ["SECTION", exp, modifier]]];
    }
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    _callCustomFunction(fn, args) {
        const activationRecord = {};
        fn.params.forEach((param, index) => {
            activationRecord[param] = args[index];
        });
        const activationEnv = new Env(activationRecord, fn.env);
        return this._evaluateBody(fn.body, activationEnv);
    }
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    //-------------------------------------------------------------------
    _isBoolean(expression) {
        return typeof expression === 'boolean';
    }

    _isNumber(expression) {
        return typeof expression === 'number';
    }

    _isString(expression) {
        return typeof expression === "string" && expression[0] === '"' && expression.slice(-1) === '"';
    }

    _isVariableName(expression) {
        return typeof expression === "string" && /^[+\-*/<>=a-zA-Z0-9_]+$/.test(expression);
    }

    _evaluateBlock(block, env) {
        let result;
        const [_tag, ...expressions] = block;
        expressions.forEach((exp) => {
            result = this.evaluateExpression(exp, env);
        });
        return result;
    }

    _evaluateBody(body, env) {
        if (body[0] === "SECTION") {
            return this._evaluateBlock(body, env);
        }
        return this.evaluateExpression(body, env);
    }
}

export default Aileen;