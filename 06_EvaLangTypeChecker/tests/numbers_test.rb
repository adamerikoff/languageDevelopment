def number_tests
  [
    { input: 5, expected: Type.number, description: "Should return Type.number for integers" },
    { input: 3.14, expected: Type.number, description: "Should return Type.number for floats" },
    { input: 0, expected: Type.number, description: "Should return Type.number for zero" },
    { input: -10, expected: Type.number, description: "Should return Type.number for negative integers" },
    { input: 1e5, expected: Type.number, description: "Should return Type.number for large numbers" }
  ]
end
