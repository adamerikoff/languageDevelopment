class Type
  attr_reader :name

  def initialize(name)
    @name = name
  end

  def equals(other)
    @name == other.name
  end

  def self.number
    @number ||= new("number")
  end

  def self.string
    @string ||= new("string")
  end
end
