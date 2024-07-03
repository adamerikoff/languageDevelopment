import assert from 'assert';

import Aileen from "../Aileen.js";

const aileen = new Aileen();

assert.strictEqual(aileen.evaluateExpression(["+", 1, 5]), 6);
assert.strictEqual(aileen.evaluateExpression(["+", ["+", 3, 2], 5]), 10);
assert.strictEqual(aileen.evaluateExpression(["-", 1, 5]), -4);
assert.strictEqual(aileen.evaluateExpression(["+", ["*", 3, 2], 5]), 11);
assert.strictEqual(aileen.evaluateExpression(["/", 20, 5]), 4);
assert.strictEqual(aileen.evaluateExpression(["modulus", 21, 5]), 1);

console.log("All assertions passed!");