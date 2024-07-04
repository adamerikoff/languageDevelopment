import runBooleanTests from "./booleanTests.js";
import runBuiltInFunctionsTests from "./builtInFunctionsTests.js";
import runClassTests from "./classTests.js";
import runConditionTests from "./conditionTests.js";
import runFunctionTests from "./functionTests.js";
import runLambdaTests from "./lambdaTests.js";
import runLoadTest from "./loadTests.js";
import runLoopTests from "./loopTests.js";
import runSectionTests from "./sectionTests.js";
import runSwitchTests from "./switchTests.js";
import runVariableTests from "./variableTests.js";

import Aileen from "../Aileen.js";

const aileen = new Aileen();
runVariableTests(aileen);
console.log("Variable Tests: All assertions passed!");

runBooleanTests(aileen);
console.log("Boolean Tests: All assertions passed!");

runBuiltInFunctionsTests(aileen);
console.log("Built-In Functions Tests: All assertions passed!");

runConditionTests(aileen);
console.log("Condition Tests: All assertions passed!");

runFunctionTests(aileen);
console.log("Function Tests: All assertions passed!");

runLambdaTests(aileen);
console.log("Lambda Tests: All assertions passed!");

runLoopTests(aileen);
console.log("Loop Tests: All assertions passed!");

runSectionTests(aileen);
console.log("Section Tests: All assertions passed!");

runSwitchTests(aileen);
console.log("Switch Tests: All assertions passed!");

runLoadTest(aileen);
console.log("Load Tests: All assertions passed!");

runClassTests(aileen);
console.log("Class Tests: All assertions passed!");

console.log("All tests completed!");