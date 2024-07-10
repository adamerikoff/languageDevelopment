import assert from "assert";

import Parser from "../src/Parser.js";


const parser = new Parser();

function test(program, expected) {
    const ast = parser.parse(program);
    assert.deepEqual(ast, expected);
}

const tests = [
    ["42;", {
        type: "Program",
        body: [
            {
                type: "ExpressionStatement",
                expression: {
                    type: "NumericLiteral",
                    value: 42,
                }
            }
        ],
    }],
    ["33;", {
        type: "Program",
        body: [
            {
                type: "ExpressionStatement",
                expression: {
                    type: "NumericLiteral",
                    value: 33,
                }
            }
        ],
    }],
    ["'hi there';", {
        type: "Program",
        body: [
            {
                type: "ExpressionStatement",
                expression: {
                    type: "StringLiteral",
                    value: "hi there",
                }
            }
        ],
    }],
    [`{
        42;
        'hello';
    }`, {
        type: "Program",
        body: [
            {
                type: "BlockStatement",
                body: [
                    {
                        type: "ExpressionStatement",
                        expression: {
                            type: "NumericLiteral",
                            value: 42,
                        }
                    },
                    {
                        type: "ExpressionStatement",
                        expression: {
                            type: "StringLiteral",
                            value: "hello",
                        }
                    }
                ]
            }
        ],
    }],
    [`5 + 5;`, {
        type: "Program",
        body: [
            {
                type: "ExpressionStatement",
                expression: [
                    {
                        type: "BinaryExpression",
                        operator: "+",
                        left: {
                            type: "NumericLiteral",
                            value: 5,
                        },
                        right: {
                            type: "NumericLiteral",
                            value: 5,
                        },
                    }
                ]
            }
        ],
    }],
    [`5 + 4 - 3;`, {
        type: "Program",
        body: [
            {
                type: "ExpressionStatement",
                expression: [
                    {
                        type: "BinaryExpression",
                        operator: "-",
                        left: {
                            type: "BinaryExpression",
                            operator: "+",
                            left: {
                                type: "NumericLiteral",
                                value: 5,
                            },
                            right: {
                                type: "NumericLiteral",
                                value: 4,
                            },
                        },
                        right: {
                            type: "NumericLiteral",
                            value: 3,
                        },
                    }
                ]
            }
        ],
    }],
];

tests.forEach((t) => {
    test(t[0], t[1]);
})

console.log("ALL TESTS ARE PASSED!")