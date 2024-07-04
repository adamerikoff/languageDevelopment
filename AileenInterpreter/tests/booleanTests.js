import assert from 'assert';

function runBooleanTests(aileen) {
    assert.strictEqual(aileen.evaluateExpression(["OR", true, true]), true);
    assert.strictEqual(aileen.evaluateExpression(["OR", true, false]), true);
    assert.strictEqual(aileen.evaluateExpression(["OR", false, true]), true);
    assert.strictEqual(aileen.evaluateExpression(["OR", false, false]), false);

    assert.strictEqual(aileen.evaluateExpression(["AND", true, true]), true);
    assert.strictEqual(aileen.evaluateExpression(["AND", true, false]), false);
    assert.strictEqual(aileen.evaluateExpression(["AND", false, true]), false);
    assert.strictEqual(aileen.evaluateExpression(["AND", false, false]), false);

    assert.strictEqual(aileen.evaluateExpression(["NOT", true]), false);
    assert.strictEqual(aileen.evaluateExpression(["NOT", false]), true);
}

export default runBooleanTests;