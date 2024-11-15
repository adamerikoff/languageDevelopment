require_relative '../src/EvaTC'
require_relative 'numbers_test'
require_relative 'strings_test'
require_relative 'math_test'

def run_all_tests
  evaluator = EvaTC.new

  # Run all test categories
  puts "Running Number Tests..."
  run_tests(evaluator, number_tests)

  puts "\nRunning String Tests..."
  run_tests(evaluator, string_tests)

  puts "\nRunning Math Tests..."
  run_tests(evaluator, math_test)

  puts "\nAll tests complete."
end

def assert_equal(expected, actual, message)
  if expected == actual
    puts "PASS: #{message}"
  else
    puts "FAIL: #{message}"
    puts "  Expected: #{expected.inspect}"
    puts "  Actual: #{actual.inspect}"
  end
end

def run_tests(evaluator, test_cases)
  test_cases.each do |test_case|
    input = test_case[:input]
    expected = test_case[:expected]
    description = test_case[:description]

    result = evaluator.tc(input)
    assert_equal(expected, result, description)
  end
end

run_all_tests
