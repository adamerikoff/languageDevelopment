const { Parser } = require("../src/Parser");
const assert = require("assert");
const tests = [require("./literal_test")];

const parser = new Parser();

function test(program, expected_result) {
    const ast = parser.parse(program);
    assert.deepEqual(ast, expected_result);
}

tests.forEach(testRun => testRun(test));

console.log("All assertions are passed!");