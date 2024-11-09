// Token specification for the tokenizer, containing regular expressions and corresponding token types.
// If a token type is `null`, it will be ignored (e.g., whitespace, comments).
const Spec = [
    // Whitespace and comments (ignored tokens)
    [/^\s+/, null],                // Match whitespace
    [/^\/\/.*/, null],             // Match single-line comments
    [/^\/\*[\s\S]*?\*\//, null],   // Match multi-line comments

    // Symbols
    [/^;/, ";"],                   // Semicolon
    [/^\{/, "{"],                  // Opening brace
    [/^\}/, "}"],                  // Closing brace
    [/^\(/, "("],                  // Opening parenthesis
    [/^\)/, ")"],                  // Closing parenthesis
    [/^,/, ","],                   // Comma
    [/^\./, "."],                   // Comma
    [/^\[/, "["],                   // Comma
    [/^\]/, "]"],                   // Comma

    // Keywords
    [/^\blet\b/, "let"],           // `let` keyword for variable declaration
    [/^\bif\b/, "if"],             // `if` keyword for conditionals
    [/^\belse\b/, "else"],         // `else` keyword for alternative conditional branch
    [/^\bfalse\b/, "false"],       // Boolean `false`
    [/^\btrue\b/, "true"],         // Boolean `true`
    [/^\bnull\b/, "null"],         // Null value
    [/^\bwhile\b/, "while"],       // while 
    [/^\bdo\b/, "do"],             // do
    [/^\bfor\b/, "for"],           // for
    [/^\bdef\b/, "def"],           // for
    [/^\breturn\b/, "return"],           // for

    // Literals
    [/^\d+/, "NUMBER"],            // Numeric literal
    [/^"[^"]*"/, "STRING"],        // String literal (double-quoted)

    // Identifiers
    [/^\w+/, "IDENTIFIER"],        // Identifier for variables, functions, etc.

    // Operators
    [/^[=!]=/, "EQUALITY_OPERATOR"],       // Equality operators (==, !=)
    [/^=/, "SIMPLE_ASSIGN"],               // Simple assignment operator (=)
    [/^[\*\/\+\-]=/, "COMPLEX_ASSIGN"],    // Complex assignment operators (*=, /=, +=, -=)
    [/^[+\-]/, "ADDITIVE_OPERATOR"],       // Additive operators (+, -)
    [/^[*\/]/, "MULTIPLICATIVE_OPERATOR"], // Multiplicative operators (*, /)
    [/^[><]=?/, "RELATIONAL_OPERATOR"],    // Relational operators (<, <=, >, >=)
    [/^&&/, "LOGICAL_AND"],                // Logical AND operator (&&)
    [/^\|\|/, "LOGICAL_OR"],               // Logical OR operator (||)
    [/^!/, "LOGICAL_NOT"],                 // Logical NOT operator (!)
];

module.exports = { 
    Spec,
};
