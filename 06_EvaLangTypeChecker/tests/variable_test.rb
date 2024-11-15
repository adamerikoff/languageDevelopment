def variable_test
  [
    # Test: Variable Assignment
    { input: ['var', 'x', 3], expected: Type.number, description: "Should return 'number' for assignment of a number to a variable" },

    # Test: Variable Assignment with a String
    { input: ['var', 'name', '"hello"'], expected: Type.string, description: "Should return 'string' for assignment of a string to a variable" },

    # Test: Variable Reference
    { input: 'x', expected: Type.number, description: "Should return 'number' when referencing a variable of type number" },

    # Test: Variable Reference with String
    { input: 'name', expected: Type.string, description: "Should return 'string' when referencing a variable of type string" },
  ]
end
