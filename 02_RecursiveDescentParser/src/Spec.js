const Spec = [
    [/^\s+/, null],
    [/^\/\/.*/, null],
    [/^\/\*[\s\S]*?\*\//, null],
    [/^;/, ";"],
    [/^\{/, "{"],
    [/^\}/, "}"],
    [/^\(/, "("],
    [/^\)/, ")"],
    [/^,/, ","],
    [/^\blet\b/, "let"],
    [/^\bif\b/, "if"],
    [/^\belse\b/, "else"],
    [/^\d+/, "NUMBER"],
    [/^\w+/, "IDENTIFIER"],
    [/^=/, "SIMPLE_ASSIGN"],
    [/^[\*\/\+\-]=/, "COMPEX_ASSIGN"],
    [/^[+\-]/, "ADDITIVE_OPERATOR"],
    [/^[*\/]/, "MULTIPLICATIVE_OPERATOR"],
    [/^[><]=?/, "RELATIONAL_OPERATOR"],
    [/^"[^"]*"/, "STRING"],
];


module.exports = { 
    Spec,
};