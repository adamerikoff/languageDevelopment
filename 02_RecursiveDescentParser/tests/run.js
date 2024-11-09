const { Parser } = require("../src/Parser");
const assert = require("assert");
const tests = [
    require("./literal_test"),
    require("./statement_list_test"),
    require("./block_test"),
    require("./empty_test"),
    require("./math_test"),
    require("./assignment_test"),
    require("./variable_test"),
    require("./control_test"),
    require("./relational_test"),
];

const parser = new Parser();

function test(program, expected_result) {
    const ast = parser.parse(program);
    assert.deepEqual(ast, expected_result);
}

tests.forEach(testRun => testRun(test));

console.log("All assertions are passed!");

function manual_test() {
    const program = `
        x = 42;
    `;

    const ast = parser.parse(program);
    console.log(JSON.stringify(ast, null, 2));
}
manual_test();