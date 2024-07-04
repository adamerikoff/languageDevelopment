import runTest from "./runTestFunction.js";

function runLambdaTests(aileen) {
    runTest(
        aileen,
        `
    (SECTION
			(FUNCTION onClick (callback)
				(SECTION
					(ASSIGN x 10)
					(ASSIGN y 20)
					(callback (+ x y))))
			(onClick (LAMBDA (data) (* data 10))))
  `,
        300
    );

    runTest(
        aileen,
        `
		((LAMBDA (x) (* x x)) 2)
		`,
        4
    );

    runTest(
        aileen,
        `
		(SECTION
			(ASSIGN square (LAMBDA (x) (* x x)))
			(square 2))
		`,
        4
    );
}

export default runLambdaTests;