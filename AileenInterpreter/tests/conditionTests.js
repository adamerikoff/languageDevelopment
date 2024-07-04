import assert from "assert";

function runConditionTests(aileen) {
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
}

export default runConditionTests;