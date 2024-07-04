import runTest from "./runTestFunction.js";

function runModuleTests(aileen) {
    runTest(
        aileen,
        `
    (MODULE Math
      (SECTION
        (FUNCTION abs (value)
          (CONDITION (< value 0)
              (- value)
              value))
        (FUNCTION square (x)
          (* x x))
        (ASSIGN MAX_VALUE 1000)
      )
    )
    ((PROPERTY Math abs) (- 10))
    `,
        10
    );

    runTest(
        aileen,
        `
    (ASSIGN abs (PROPERTY Math abs))
    (abs (- 10))
    `,
        10
    );

    runTest(
        aileen,
        `
    (PROPERTY Math MAX_VALUE)
    `,
        1000
    );
}

export default runModuleTests;