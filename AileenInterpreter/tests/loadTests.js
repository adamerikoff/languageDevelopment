import runTest from "./runTestFunction.js";

function runLoadTest(aileen) {
    runTest(
        aileen,
        `
    (LOAD Math)
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

    runTest(
        aileen,
        `
    (LOAD (abs MAX_VALUE) Math)
    (* (PROPERTY Math MAX_VALUE) ((PROPERTY Math abs) (- 10)))
    `,
        10000
    );
}

export default runLoadTest;