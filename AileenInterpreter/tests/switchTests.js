import runTest from "./runTestFunction.js";

function runSwitchTests(aileen) {
    runTest(
        aileen,
        `
    (SECTION
      (ASSIGN x 10)
      (SWITCH ((= x 10) 100)
              ((> x 10) 200)
              (ELSE     300)))
  `,
        100
    );
}

export default runSwitchTests;