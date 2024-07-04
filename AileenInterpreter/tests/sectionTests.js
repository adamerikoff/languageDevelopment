import assert from 'assert';

import runTest from "./runTestFunction.js";

function runSectionTests(aileen) {

    assert.strictEqual(
        aileen.evaluateExpression([
            "SECTION",
            ["ASSIGN", "x", 10],
            ["ASSIGN", "y", 20],
            ["+", ["*", "x", "y"], 30],
        ]),
        230
    );

    assert.strictEqual(
        aileen.evaluateExpression([
            "SECTION",
            ["ASSIGN", "x", 10],
            ["SECTION", ["ASSIGN", "x", 20], "x"],
            "x",
        ]),
        10
    );

    runTest(aileen,
        `
    (SECTION
      (ASSIGN value 10)
      (ASSIGN result (SECTION
                    (ASSIGN x (+ value 10))
                    x))
      result)
  `,
        20);

    runTest(aileen,
        `
    (SECTION
      (ASSIGN data 10)
      (SECTION (REASSIGN data 100))
      data)
  `,
        100);

    runTest(aileen,
        `
    (SECTION
      (ASSIGN x -10.0)
      (ASSIGN y 20)
      (+ (* x 10) y))
  `,
        -80);
}

export default runSectionTests;