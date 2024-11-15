require_relative "../src/EvaTC"

# Define a helper method to test and print results
def assert_equal(expected, actual, message)
  if expected == actual
    puts "PASS: #{message}"
  else
    puts "!" * 30
    puts "FAIL: #{message}"
    puts "  Expected: #{expected.inspect}"
    puts "  Actual: #{actual.inspect}"
  end
end

# Initialize the EvaTC instance
evaluator = EvaTC.new

# Test cases
puts "Running tests..."

# Test number inputs
assert_equal("number", evaluator.tc(5), "Should return 'number' for integers")
assert_equal("number", evaluator.tc(3.14), "Should return 'number' for floats")

# Test non-number inputs
assert_equal("not a number", evaluator.tc("hello"), "Should return 'not a number' for strings")
assert_equal("not a number", evaluator.tc(nil), "Should return 'not a number' for nil")
assert_equal("not a number", evaluator.tc([]), "Should return 'not a number' for arrays")
assert_equal("not a number", evaluator.tc({}), "Should return 'not a number' for hashes")

# Edge cases
assert_equal("number", evaluator.tc(0), "Should return 'number' for zero")
assert_equal("number", evaluator.tc(-10), "Should return 'number' for negative integers")
assert_equal("number", evaluator.tc(1e5), "Should return 'number' for large numbers")

puts "Tests complete."
