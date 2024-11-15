def string_tests
  [
    { input: '"hello"', expected: "string", description: "Should return 'string' for double-quoted strings" },
    { input: '"world"', expected: "string", description: "Should return 'string' for double-quoted strings" },
    { input: '"hello world!"', expected: "string", description: "Should return 'string' for double-quoted strings with punctuation" },

  ]
end
