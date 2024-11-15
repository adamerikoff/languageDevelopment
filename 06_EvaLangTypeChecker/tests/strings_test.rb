def string_tests
  [
    { input: '"hello"', expected: Type.string, description: "Should return Type.string for double-quoted strings" },
    { input: '"world"', expected: Type.string, description: "Should return Type.string for double-quoted strings" },
    { input: '"hello world!"', expected: Type.string, description: "Should return Type.string for double-quoted strings with punctuation" }
  ]
end
