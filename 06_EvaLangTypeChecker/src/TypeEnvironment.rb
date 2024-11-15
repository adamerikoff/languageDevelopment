class TypeEnvironment
  def initialize(record = {}, parent = nil)
    @record = record
    @parent = parent
  end

  def define(name, type)
    @record[name] = type
    return type
  end

  def lookup(name)
    return @record[name] if @record.key?(name)
    return @parent.lookup(name) if @parent
    raise "Undefined type: #{name}"
  end
end
