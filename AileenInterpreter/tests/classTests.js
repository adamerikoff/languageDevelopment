import runTest from "./runTestFunction.js";

function runClassTests(aileen) {
    runTest(
        aileen,
        `
    (CLASS Point ZERO
      (SECTION
        (FUNCTION INIT (THIS x y)
          (SECTION
            (REASSIGN (PROPERTY THIS x) x)
            (REASSIGN (PROPERTY THIS y) y)))
        (FUNCTION calc (THIS)
          (+ (PROPERTY THIS x) (PROPERTY THIS y)))))
    (ASSIGN p (NEW Point 10 20))
    ((PROPERTY p calc) p)
    `,
        30
    );

    runTest(
        aileen,
        `
    (CLASS Point3D Point
      (SECTION
        (FUNCTION INIT (THIS x y z)
          (SECTION
            ((PROPERTY (PARENT Point3D) INIT) THIS x y)
            (REASSIGN (PROPERTY THIS z) z)))
        (FUNCTION calc (this)
          (+ ((PROPERTY (PARENT Point3D) calc) THIS)
              (PROPERTY THIS z)))))
    (ASSIGN p (NEW Point3D 10 20 30))
    ((PROPERTY p calc) p)
    `,
        60
    );
}

export default runClassTests;