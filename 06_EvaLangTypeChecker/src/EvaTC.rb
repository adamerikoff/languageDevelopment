class EvaTC
  def is_number(exp)
    exp.is_a?(Numeric)
  end

  def tc(exp)
    if is_number(exp)
      "number"
    else
      "not a number"
    end
  end
end
