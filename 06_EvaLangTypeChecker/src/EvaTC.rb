class EvaTC
  def is_number?(exp)
    exp.is_a?(Numeric)
  end

  def is_string?(exp)
    if exp.is_a?(String)
      if exp.match(/^"[^"]*"$/)
        return true
      end
    end
    false
  end

  def tc(exp)
    if is_number?(exp)
      "number"
    elsif is_string?(exp)
      "string"
    else
      "not a number"
    end
  end
end
