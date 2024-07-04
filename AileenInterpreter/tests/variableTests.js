import assert from "assert";

function runVariableTests(aileen) {
    assert.strictEqual(aileen.evaluateExpression(1), 1);
    assert.strictEqual(aileen.evaluateExpression(1), 1);
    assert.strictEqual(aileen.evaluateExpression('"hello"'), "hello");

    assert.strictEqual(aileen.evaluateExpression(["ASSIGN", "x", 10]), 10);
    assert.strictEqual(aileen.evaluateExpression("x"), 10);

    assert.strictEqual(aileen.evaluateExpression(["ASSIGN", "y", 100]), 100);
    assert.strictEqual(aileen.evaluateExpression("y"), 100);

    assert.strictEqual(aileen.evaluateExpression("VERSION"), "0.2");
    assert.strictEqual(aileen.evaluateExpression(["ASSIGN", "isUser", "TRUE"]), true);

}

export default runVariableTests;