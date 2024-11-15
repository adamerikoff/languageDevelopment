require_relative 'Type'

class EvaTC

  def tc(exp)
    if is_number?(exp)
      return Type.number
    elsif is_string?(exp)
      return Type.string
    elsif is_binary(exp)
      return binary(exp)
    else
      "ERROR"
    end
  end

  def is_number?(exp)
    return exp.is_a?(Numeric)
  end

  def is_string?(exp)
    if exp.is_a?(String)
      if exp.match(/^"[^"]*"$/)
        return true
      end
    end
    return false
  end

  def is_binary(exp)
    if exp[0].is_a?(String)
      if exp[0].match(/[+\-*\/]/)
        return true
      end
    end
    return false
  end

  def binary(exp)
    check_arity(exp, 3)
    t1 = tc(exp[1])
    t2 = tc(exp[2])
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
end
