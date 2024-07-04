import assert from "assert";

import yyaprse from "../aileenParser.js";

function runTest(aileen, code, expected) {
    const expression = yyaprse.parse(`(SECTION ${code})`);
    assert.strictEqual(aileen.evaluateGlobal(expression), expected);
}

export default runTest;