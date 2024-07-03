import Env from "./Env.js"

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
    constructor (global = GlobalEnv) {

    }
}

export default Aileen;