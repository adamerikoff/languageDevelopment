const assert = require("assert");
const { Eva } = require("../src/Eva");

const eva = new Eva();

assert.strictEqual(eva.eval(1), 1);
assert.strictEqual(eva.eval(`"hello"`), "hello");
assert.strictEqual(eva.eval(["+", 5, 6]), 11);
assert.strictEqual(eva.eval(["-", 5, 6]), -1);
assert.strictEqual(eva.eval(["*", 5, 6]), 30);
assert.strictEqual(eva.eval(["/", 5, 2]), 2.5);
assert.strictEqual(eva.eval(["+", ["-", 5, 4], 6]), 7);

console.log("ALL ASSERTIONS PASSED!");