require "./Nodes.rb"

# Parser class to parse tokens and construct a syntax tree
class Parser
  # Initialize parser with a list of tokens
  def initialize(tokens)
    @tokens = tokens
  end

  # Entry point to parse a function definition (assumes starting with a 'def' keyword)
  def parse
    parse_def
  end

  # Low-level function to consume a token of a specific type from the token list
  # If the next token is not of the expected type, raises an error
  def consume(expected_type)
    token = @tokens.shift                       # Remove and return the next token
    if token.type == expected_type              # Check if it matches the expected type
      token                                     # Return token if types match
    else
      raise RuntimeError.new("Expected token type #{expected_type.inspect}, but got #{token.type.inspect}!")
    end
  end

  # Low-level function to peek at the next token type without consuming it
  def peek(expected_type, offset=0)
    @tokens.fetch(offset).type == expected_type      # Check if the first token matches the expected type
  end

  # Parses a function definition with a name, arguments, and body
  # Expected syntax: def <name>(<args>) <body> end
  def parse_def
    consume(:def)                             # Expect and consume the 'def' keyword
    name = consume(:identifier).value         # Get the function name (identifier)
    arg_names = parse_arg_names               # Parse function arguments
    body = parse_expr                         # Parse the function body
    consume(:end)                             # Expect and consume the 'end' keyword
    DefNode.new(name, arg_names, body)        # Return a DefNode representing the function
  end

  # Parses argument names inside parentheses for a function
  # Expected syntax: (<arg1>, <arg2>, ...)
  def parse_arg_names
    arg_names = []
    consume(:oparen)                          # Expect and consume '('
    if peek(:identifier)                      # Check if there are any arguments
      arg_names << consume(:identifier).value # Consume first argument name
      while peek(:comma)                      # While commas are present, there are more arguments
        consume(:comma)                       # Consume comma
        arg_names << consume(:identifier).value # Consume subsequent argument names
      end
    end
    consume(:cparen)                          # Expect and consume ')'
    arg_names                                 # Return list of argument names
  end

  # Determines if the current expression is an integer or a function call
  def parse_expr
    if peek(:integer)                         # If token is an integer
      parse_int                               # Parse as integer node
    elsif peek(:identifier) && peek(:oparen, 1)
      parse_call                              # Otherwise, parse as a function call
    else
      parse_var_ref
    end
  end

  # Parses an integer value and returns an IntegerNode
  def parse_int
    IntegerNode.new(consume(:integer).value.to_i) # Consume integer and convert to IntegerNode
  end

  # Placeholder method for parsing function calls (not implemented here)
  def parse_call
    name = consume(:identifier).value
    arg_exprs = parse_arg_exprs
    CallNode.new(name, arg_exprs)
  end

  def parse_arg_exprs
    arg_exprs = []
    consume(:oparen)
    if !peek(:cparen)
      arg_exprs << parse_expr
      while peek(:comma)
        consume(:comma)
        arg_exprs << parse_expr
      end
    end
    consume(:cparen)
    arg_exprs
  end

  def parse_var_ref
    VarRefNode.new(consume(:identifier).value)
  end

end
