const assert = require("assert");
const { Eva } = require("../src/Eva");
const { Environment } = require("../src/Environment");

// Initialize Eva instance with base environment variables.
const eva = new Eva(new Environment({
    true: true,
    false: false,
    nil: null,
}));

// Number and String Evaluations
assert.strictEqual(eva.eval(1), 1);
assert.strictEqual(eva.eval(`"hello"`), "hello");

// Arithmetic Operations
assert.strictEqual(eva.eval(["+", 5, 6]), 11);
assert.strictEqual(eva.eval(["-", 5, 6]), -1);
assert.strictEqual(eva.eval(["*", 5, 6]), 30);
assert.strictEqual(eva.eval(["/", 5, 2]), 2.5);
assert.strictEqual(eva.eval(["+", ["-", 5, 4], 6]), 7);

// Variable Declarations and Lookups
assert.strictEqual(eva.eval(["define", "x", -990]), -990);
assert.strictEqual(eva.eval("x"), -990);
assert.strictEqual(eva.eval(["define", "z", "true"]), true);
assert.strictEqual(eva.eval(["define", "zero", "nil"]), null);

// Block Evaluations and Nested Scopes
assert.strictEqual(
    eva.eval(["begin", 
        ["define", "x", 2],
        ["define", "w", -2],
        ["+", ["*", "x", "w"], -4],
    ]), 
    -8
);

assert.strictEqual(
    eva.eval(["begin", 
        ["define", "x", 2],
        ["begin",
            ["define", "x", 5],
            "x"
        ],
        "x"
    ]), 
    2
);

assert.strictEqual(
    eva.eval(["begin", 
        ["define", "z", 111],
        ["define", "x", 2],
        ["begin",
            ["define", "x", ["+", "z", 10]],
            "x"
        ],
        "x"
    ]), 
    2
);

assert.strictEqual(
    eva.eval(["begin", 
        ["define", "data", '"string_data"'],
        ["begin",
            ["assign", "data", 33]
        ],
        "data"
    ]), 
    33
);

console.log("ALL ASSERTIONS PASSED!");
