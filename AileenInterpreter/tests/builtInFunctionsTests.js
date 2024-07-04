import assert from "assert";

import runTest from "./runTestFunction.js";

function runBuiltInFunctionsTests(aileen) {
    runTest(aileen, `(+ 1 5)`, 6);
    runTest(aileen, `(+ (+ 2 3) 5)`, 10);
    runTest(aileen, `(+ (* 2 3) 5)`, 11);

    // Comparison:

    runTest(aileen, `(> 1 5)`, false);
    runTest(aileen, `(< 1 5)`, true);

    runTest(aileen, `(>= 5 5)`, true);
    runTest(aileen, `(<= 5 5)`, true);
    runTest(aileen, `(=  5 5)`, true);
    runTest(aileen, `(*  5 5)`, 25);

    assert.strictEqual(aileen.evaluateExpression(
            ["SECTION",
                ["ASSIGN", "result", 0],
                ["--", "result"],
                "result",
            ]),
        -1
    );

    assert.strictEqual(aileen.evaluateExpression(
            ["SECTION",
                ["ASSIGN", "result", 122],
                ["--", "result"],
                "result",
            ]),
        121
    );



    assert.strictEqual(aileen.evaluateExpression(
            ["SECTION",
                ["ASSIGN", "result", 0],
                ["++", "result"],
                "result",
            ]),
        1
    );

    assert.strictEqual(aileen.evaluateExpression(
            ["SECTION",
                ["ASSIGN", "result", 122],
                ["++", "result"],
                "result",
            ]),
        123
    );
}

export default runBuiltInFunctionsTests;