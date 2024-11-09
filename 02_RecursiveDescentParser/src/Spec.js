const Spec = [
    [/^\s+/, null],
    [/^\/\/.*/, null],
    [/^\/\*[\s\S]*?\*\//, null],
    [/^;/, ";"],
    [/^\{/, "{"],
    [/^\}/, "}"],
    [/^\d+/, "NUMBER"],
    [/^"[^"]*"/, "STRING"],
];


module.exports = { 
    Spec,
};