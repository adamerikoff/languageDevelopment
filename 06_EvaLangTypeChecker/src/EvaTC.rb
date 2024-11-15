require_relative 'Type'
require_relative 'TypeEnvironment'

class EvaTC
  def initialize
    @global = createGlobal
  end

  def tc(exp, env = @global)
    if is_number?(exp)
      return Type.number
    elsif is_string?(exp)
      return Type.string
    elsif is_binary(exp)
      return binary(exp, env)
    elsif exp[0] == 'var'
      # Variable declaration
      tag, name, value = exp
      value_type = tc(value, env)

      if name.is_a?(Array)
        var_name, type = name
        expected_type = Type.from_string(type)
        expect(value_type, expected_type, value, type)
        return env.define(var_name, expected_type)
      end
      return env.define(name, value_type)
    elsif is_variable_name(exp)
      return env.lookup(exp)
    else
      return "ERROR"
    end
  end

  def is_number?(exp)
    return exp.is_a?(Numeric)
  end

  def is_string?(exp)
    return exp.is_a?(String) && exp.match(/^"[^"]*"$/)
  end

  def is_binary(exp)
    return exp.is_a?(Array) && exp[0].match(/[+\-*\/<>=]/)
  end

  def binary(exp, env)
    check_arity(exp, 3)
    t1 = tc(exp[1], env)
    t2 = tc(exp[2], env)
    return expect(t2, t1, exp[2], exp)
  end

  def check_arity(exp, arity)
    if exp.length != arity
      raise "Operator expects #{arity - 1} operands, #{exp.length - 1} given in #{exp}"
    end
  end

  def expect(actual_type, expected, value, exp)
    if !actual_type.equals(expected)
      raise "Expected #{expected.name}, but got #{actual_type.name} for value #{value} in expression #{exp}"
    end
    return actual_type
  end

  def createGlobal
    return TypeEnvironment.new({
      "ver" => Type.string
    })
  end

  def is_variable_name(exp)
    return exp.is_a?(String) && exp.match(/[a-zA-Z_][a-zA-Z0-9_]*$/)
  end
end
