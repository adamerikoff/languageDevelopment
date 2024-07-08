import Parser from '../src/Parser';

const parser = new Parser();

const sourceCode = '42';

const ast = parser.parse(sourceCode);

console.log(JSON.stringify(ast, null, 2));