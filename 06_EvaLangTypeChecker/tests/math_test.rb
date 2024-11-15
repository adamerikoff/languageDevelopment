def math_test
  [
    # Test: Addition
    { input: ['+', 2, 3], expected: Type.number, description: "Should return 'number' for addition of two numbers" },

    # Test: Subtraction
    { input: ['-', 5, 3], expected: Type.number, description: "Should return 'number' for subtraction of two numbers" },

    # Test: Multiplication
    { input: ['*', 2, 4], expected: Type.number, description: "Should return 'number' for multiplication of two numbers" },

    # Test: Division
    { input: ['/', 10, 2], expected: Type.number, description: "Should return 'number' for division of two numbers" },

  ]
end
