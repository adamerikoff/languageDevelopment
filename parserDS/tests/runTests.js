import Parser from "../src/Parser.js";

const parser = new Parser();

const sourceCode = "+1.2314";

const ast = parser.parse(sourceCode);

console.log(JSON.stringify(ast, null, 2));