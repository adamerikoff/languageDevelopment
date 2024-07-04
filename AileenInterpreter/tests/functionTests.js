import runTest from "./runTestFunction.js";

function runFunctionTests(aileen) {

    runTest(
        aileen,
        `
    (SECTION
      (FUNCTION square (x)
        (* x x))
      (square 2))
  `,
        4
    );

    runTest(
        aileen,
        `
    (SECTION
      (ASSIGN x 10)
      (FUNCTION foo () x)
      (FUNCTION bar ()
        (SECTION
          (ASSIGN x 20)
          (+ (foo) x)))
      (bar))
  `,
        30
    );

    runTest(
        aileen,
        `
    (SECTION
      (FUNCTION calc (x y)
        (SECTION
          (ASSIGN z 30)
          (+ (* x y) z)
        ))
      (calc 10 20))
  `,
        230
    );

    runTest(
        aileen,
        `
    (SECTION
      (ASSIGN value 100)
      (FUNCTION calc (x y)
        (SECTION
          (ASSIGN z (+ x y))

          (FUNCTION inner (foo)
            (+ (+ foo z) value))

          inner

        ))
      (ASSIGN fn (calc 10 20))
      (fn 30))
  `,
        160
    );

    runTest(
        aileen,
        `
    (SECTION
      (FUNCTION factorial (x)
        (CONDITION (= x 1)
          1
          (* x (factorial (- x 1)))))
      (factorial 5))
  `,
        120
    );
}

export default runFunctionTests;