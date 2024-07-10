import assert from "assert";

import Parser from "../src/Parser.js";


const parser = new Parser();

function test(program, expected) {
    const ast = parser.parse(program);
    assert.deepEqual(ast, expected);
}

const tests = [
    ["42", {
        type: "Program",
        body: {
            type: "NumericLiteral",
            value: 42,
        }
    }],
    ["52", {
        type: "Program",
        body: {
            type: "NumericLiteral",
            value: 52,
        }
    }],
    ["'string string'", {
        type: "Program",
        body: {
            type: "StringLiteral",
            value: "string string",
        }
    }],
];

tests.forEach((t) => {
    test(t[0], t[1]);
})