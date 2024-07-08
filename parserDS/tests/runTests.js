import Parser from "../src/Parser.js";

const parser = new Parser();

const sourceCode = "'asfaf'";

const ast = parser.parse(sourceCode);

console.log(JSON.stringify(ast, null, 2));