const assert = require("assert");
const { Eva } = require("../src/Eva");

const eva = new Eva();

assert.strictEqual(eva.eval(1), 1);
assert.strictEqual(eva.eval(`"hello"`), "hello");

console.log("ALL ASSERTIONS PASSED!");