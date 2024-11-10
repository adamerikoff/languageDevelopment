const assert = require("assert");
const { Eva } = require("../src/Eva");
const { Environment } = require("../src/Environment");

const eva = new Eva(new Environment({
    true: true,
    false: false,
    nil: null
}));

assert.strictEqual(eva.eval(1), 1);
assert.strictEqual(eva.eval(`"hello"`), "hello");
assert.strictEqual(eva.eval(["+", 5, 6]), 11);
assert.strictEqual(eva.eval(["-", 5, 6]), -1);
assert.strictEqual(eva.eval(["*", 5, 6]), 30);
assert.strictEqual(eva.eval(["/", 5, 2]), 2.5);
assert.strictEqual(eva.eval(["+", ["-", 5, 4], 6]), 7);
assert.strictEqual(eva.eval(["declare", "x", -990]), -990);
assert.strictEqual(eva.eval("x"), -990);
assert.strictEqual(eva.eval(["declare", "z", "true"]), true);
assert.strictEqual(eva.eval(["declare", "zero", "nil"]), null);

console.log("ALL ASSERTIONS PASSED!");