def number_tests
  [
    { input: 5, expected: "number", description: "Should return 'number' for integers" },
    { input: 3.14, expected: "number", description: "Should return 'number' for floats" },
    { input: 0, expected: "number", description: "Should return 'number' for zero" },
    { input: -10, expected: "number", description: "Should return 'number' for negative integers" },
    { input: 1e5, expected: "number", description: "Should return 'number' for large numbers" }
  ]
end
