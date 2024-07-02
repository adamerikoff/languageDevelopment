import assert from 'assert';
import Eva from '../Eva.js';

const eva = new Eva();

assert.strictEqual(eva.eval(1), 1);
assert.strictEqual(eva.eval('"this is a string"'), 'this is a string');

// Addition tests
assert.strictEqual(eva.eval(['+', 1, 6]), 7);
assert.strictEqual(eva.eval(['+', ['+', 1, 6], 6]), 13);

// Subtraction tests
assert.strictEqual(eva.eval(['-', 6, 1]), 5);
assert.strictEqual(eva.eval(['-', ['-', 10, 3], 2]), 5);

// Multiplication tests
assert.strictEqual(eva.eval(['*', 2, 3]), 6);
assert.strictEqual(eva.eval(['*', ['*', 2, 3], 4]), 24);

// Division tests
assert.strictEqual(eva.eval(['/', 6, 2]), 3);
assert.strictEqual(eva.eval(['/', ['/', 12, 3], 2]), 2);

assert.strictEqual(eva.eval(['assign', 'x', 90]), 90);
assert.strictEqual(eva.eval('x'), 90);
assert.strictEqual(eva.eval(['assign', 'x999', 40]), 40);
assert.strictEqual(eva.eval('x999'), 40);
//assert.strictEqual(eva.eval(['assign', 'z', '"universe"']), 'universe');

console.log('All assertions passed!')