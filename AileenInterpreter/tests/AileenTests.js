import assert from "assert";
import fs from "fs";

import Aileen from "../Aileen.js";
import yyaprse from "../aileenParser.js";

function test(ail, code, expected) {
    const exp = yyaprse.parse(`(SECTION ${code})`);
    assert.strictEqual(ail.evaluateGlobal(exp), expected);
}

const aileen = new Aileen();

assert.strictEqual(aileen.evaluateExpression(["+", 1, 5]), 6);
assert.strictEqual(aileen.evaluateExpression(["+", ["+", 3, 2], 5]), 10);
assert.strictEqual(aileen.evaluateExpression(["-", 1, 5]), -4);
assert.strictEqual(aileen.evaluateExpression(["+", ["*", 3, 2], 5]), 11);
assert.strictEqual(aileen.evaluateExpression(["/", 20, 5]), 4);
assert.strictEqual(aileen.evaluateExpression(["modulus", 21, 5]), 1);

assert.strictEqual(aileen._isNumber(123), true);
assert.strictEqual(aileen._isNumber(true), false);

assert.strictEqual(aileen.evaluateExpression([">", 5, 1]), true);
assert.strictEqual(aileen.evaluateExpression([">", 1, 5]), false);

assert.strictEqual(aileen.evaluateExpression([">=", 5, 1]), true);
assert.strictEqual(aileen.evaluateExpression([">=", 5, 5]), true);
assert.strictEqual(aileen.evaluateExpression([">=", 1, 5]), false);

assert.strictEqual(aileen.evaluateExpression(["<", 5, 1]), false);
assert.strictEqual(aileen.evaluateExpression(["<", 1, 5]), true);

assert.strictEqual(aileen.evaluateExpression(["<=", 1, 5]), true);
assert.strictEqual(aileen.evaluateExpression(["<=", 5, 5]), true);
assert.strictEqual(aileen.evaluateExpression(["<=", 5, 1]), false);

assert.strictEqual(aileen.evaluateExpression(["=", 5, 5]), true);
assert.strictEqual(aileen.evaluateExpression(["=", 5, 1]), false);

assert.strictEqual(aileen.evaluateExpression(["OR", true, true]), true);
assert.strictEqual(aileen.evaluateExpression(["OR", true, false]), true);
assert.strictEqual(aileen.evaluateExpression(["OR", false, true]), true);
assert.strictEqual(aileen.evaluateExpression(["OR", false, false]), false);

assert.strictEqual(aileen.evaluateExpression(["AND", true, true]), true);
assert.strictEqual(aileen.evaluateExpression(["AND", true, false]), false);
assert.strictEqual(aileen.evaluateExpression(["AND", false, true]), false);
assert.strictEqual(aileen.evaluateExpression(["AND", false, false]), false);

assert.strictEqual(aileen.evaluateExpression(["NOT", true]), false);
assert.strictEqual(aileen.evaluateExpression(["NOT", false]), true);

assert.strictEqual(aileen.evaluateExpression(1), 1);
assert.strictEqual(aileen.evaluateExpression(1), 1);
assert.strictEqual(aileen.evaluateExpression('"hello"'), "hello");

assert.strictEqual(aileen.evaluateExpression(["ASSIGN", "x", 10]), 10);
assert.strictEqual(aileen.evaluateExpression("x"), 10);

assert.strictEqual(aileen.evaluateExpression(["ASSIGN", "y", 100]), 100);
assert.strictEqual(aileen.evaluateExpression("y"), 100);

assert.strictEqual(aileen.evaluateExpression("VERSION"), "0.2");
assert.strictEqual(aileen.evaluateExpression(["ASSIGN", "isUser", "TRUE"]), true);

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

test(
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
            ["ASSIGN", "x", 10],
            ["ASSIGN", "y", 0],
            ["CONDITION", [">", "x", 10],
                ["REASSIGN", "y", 20],
                ["REASSIGN", "y", 30],
            ],
            "y"
        ]),
    30);

assert.strictEqual(aileen.evaluateExpression(
        ["SECTION",
            ["ASSIGN", "x", 15],
            ["ASSIGN", "y", 0],
            ["CONDITION", [">", "x", 10],
                ["REASSIGN", "y", 20],
                ["REASSIGN", "y", 30],
            ],
            "y"
        ]),
    20);

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

test(
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

test(
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

test(
    aileen,
    `
		((LAMBDA (x) (* x x)) 2)
		`,
    4
);

test(
    aileen,
    `
		(SECTION
			(ASSIGN square (LAMBDA (x) (* x x)))
			(square 2))
		`,
    4
);

test(aileen,
    `
    (SECTION
      (ASSIGN value 10)
      (ASSIGN result (SECTION
                    (ASSIGN x (+ value 10))
                    x))
      result)
  `,
    20);

test(aileen,
    `
    (SECTION
      (ASSIGN data 10)
      (SECTION (REASSIGN data 100))
      data)
  `,
    100);

test(aileen,
    `
    (SECTION
      (ASSIGN x -10.0)
      (ASSIGN y 20)
      (+ (* x 10) y))
  `,
    -80);

test(
    aileen,
    `
    (CLASS Point ZERO
      (SECTION
        (FUNCTION constructor (THIS x y)
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

test(
    aileen,
    `
    (CLASS Point3D Point
      (SECTION
        (FUNCTION constructor (THIS x y z)
          (SECTION
            ((PROPERTY (PARENT Point3D) constructor) THIS x y)
            (REASSIGN (PROPERTY this z) z)))
        (FUNCTION calc (this)
          (+ ((PROPERTY (PARENT Point3D) calc) THIS)
              (PROPERTY THIS z)))))
    (ASSIGN p (NEW Point3D 10 20 30))
    ((PROPERTY p calc) p)
    `,
    60
);

test(
    aileen,
    `
    (SECTION
      (FUNCTION square (x)
        (* x x))
      (square 2))
  `,
    4
);

test(
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

test(
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

test(
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

test(
    aileen,
    `
    (SECTION
      (FUNCTION factorial (x)
        (if (= x 1)
          1
          (* x (factorial (- x 1)))))
      (factorial 5))
  `,
    120
);

test(
    aileen,
    `
    (LOAD Math)
    ((PROPERTY Math abs) (- 10))
    `,
    10
);

test(
    aileen,
    `
    (ASSIGN abs (PROPERTY Math abs))
    (abs (- 10))
    `,
    10
);

test(
    aileen,
    `
    (PROPERTY Math MAX_VALUE)
    `,
    1000
);

test(
    aileen,
    `
    (LOAD (abs MAX_VALUE) Math)
    (* (PROPERTY Math MAX_VALUE) ((PROPERTY Math abs) (- 10)))
    `,
    10000
);

test(
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

test(
    aileen,
    `
    (ASSIGN abs (PROPERTY Math abs))
    (abs (- 10))
    `,
    10
);

test(
    aileen,
    `
    (PROPERTY Math MAX_VALUE)
    `,
    1000
);

console.log("All assertions passed!");