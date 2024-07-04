import fs from "fs";
import yyparse from "./aileenParser.js";
import Aileen from "./Aileen.js";

function aileenGlobal(src, aileen) {
    const expression = yyparse.parse(`(SECTION ${src})`);
    return aileen.evaluateGlobal(expression)
}

function main(argv) {
    const [_node, _path, mode, expression] = argv;

    const aileen = new Aileen();
    // Direct expressions
    if (mode === '-e') {
        return;
    }
    // Source files
    if (mode === '-f') {
        const src = fs.readFileSync(expression, 'utf-8')
        return aileenGlobal(src, aileen)
    }
}

main(process.argv);