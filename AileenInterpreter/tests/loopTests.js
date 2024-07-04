import assert from "assert";

import runTest from "./runTestFunction.js";

function runLoopTests(aileen) {
    runTest(
        aileen,
        `
    (SECTION

      (ASSIGN result 0)

      (LOOPFOR (ASSIGN i 0) (< i 5) (REASSIGN i (+ i 1))
        (REASSIGN result (+ result i)))

      result

    )

  `,
        10
    );

    runTest(
        aileen,
        `
    (SECTION
      (ASSIGN cnt 0)
      (LOOPFOR (ASSIGN x 10)
           (> x 0)
           (-- x)
           (++ cnt))
      cnt)
  `,
        10
    );

  assert.strictEqual(aileen.evaluateExpression(
          ["SECTION",
            ["ASSIGN", "counter", 0],
            ["ASSIGN", "result", 0],
            ["LOOP", ["<", "counter", 10],
              ["SECTION",
                ["REASSIGN", "result", ["+", "result", 1]],
                ["REASSIGN", "counter", ["+", "counter", 1]]
              ]
            ],
            "result"
          ]),
      10);
}

export default runLoopTests;