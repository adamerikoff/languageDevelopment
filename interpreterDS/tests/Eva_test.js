import assert from 'assert';
import Eva from '../Eva.js';

const eva = new Eva();

assert.strictEqual(eva.eval(1), 1);
assert.strictEqual(eva.eval('"this is a string"'), 'this is a string');

assert.strictEqual(eva.eval(['+', 1, 6]), 7);

console.log('All assertions passed!')