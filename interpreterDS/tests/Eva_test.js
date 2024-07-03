import assert from 'assert';
import Eva from '../Eva.js';

import yyparse from '../evaParser.js';


const eva = new Eva();

function test_string(eva, code, expected) {
    const expression = yyparse.parse(code);
    assert.strictEqual(eva.eval(expression), expected);
}

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
assert.strictEqual(eva.eval(['assign', 'k', ['/', 20, 10]]), 2);
assert.strictEqual(eva.eval('k'), 2);
//assert.strictEqual(eva.eval(['assign', 'z', '"universe"']), 'universe');

assert.strictEqual(eva.eval(
    ['section',
        ['assign', 'r', 11],
        ['assign', 'w', 5],
        ['+', ['-', ['*', 'r', 'w'],5],60],
    ]),
    110);
assert.strictEqual(eva.eval(
        ['section',
            ['assign', 'r', 11],
            ['section',
                ['assign', 'r', 5],
                'r'
            ],
            'r'
        ]),
    11);

assert.strictEqual(eva.eval(
        ['section',
            ['assign', 'x', 10],
            ['section',
                ['assign', 'x', 30],
                'x'
            ],
            'x'
        ]),
    10);
assert.strictEqual(eva.eval(
        ['section',
            ['assign', 'value', 10],
            ['assign', 'result', ['section',
                ['assign', 'x', ['+', 'value', 10]],
                    'x'
                ]],
            'result'
        ]),
    20);
assert.strictEqual(eva.eval(
        ['section',
            ['assign', 'data', 22],
            ['section',
                ['reassign', 'data', 333],
            ],
            'data'
        ]),
    333);
assert.strictEqual(eva.eval(
    ['section',
        ['assign', 'x', 10],
        ['assign', 'y', 0],
        ['assume', ['>', 'x', 10],
            ['reassign', 'y', 20],
            ['reassign', 'y', 30],
        ],
        'y'
    ]),
    30);
assert.strictEqual(eva.eval(
        ['section',
            ['assign', 'counter', 0],
            ['assign', 'result', 0],
            ['loop', ['<', 'counter', 10],
                ['section',
                    ['reassign', 'result', ['+', 'result', 1]],
                    ['reassign', 'counter', ['+', 'counter', 1]],
                ],
            ],
            'result'
        ]),
    10);
test_string(eva, '(+ 1 2)', 3);  // Ensure the list expression is complete
test_string(eva,'(section(assign x 10)(assign y 30)(+ (* x 10) y))', 130);
console.log('All assertions passed!')