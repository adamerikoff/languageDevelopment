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
    [/^\bfalse\b/, "false"],
    [/^\btrue\b/, "true"],
    [/^\bnull\b/, "null"],
    [/^\d+/, "NUMBER"],
    [/^\w+/, "IDENTIFIER"],
    [/^[=!]=/, "EQUALITY_OPERATOR"],
    [/^=/, "SIMPLE_ASSIGN"],
    [/^[\*\/\+\-]=/, "COMPEX_ASSIGN"],
    [/^[+\-]/, "ADDITIVE_OPERATOR"],
    [/^[*\/]/, "MULTIPLICATIVE_OPERATOR"],
    [/^[><]=?/, "RELATIONAL_OPERATOR"],
    [/^&&/, "LOGICAL_AND"],
    [/^\|\|/, "LOGICAL_OR"],
    [/^"[^"]*"/, "STRING"],
];


module.exports = { 
    Spec,
};