# Token struct to store type and value of each token
Token = Struct.new(:type, :value)

class Tokenizer
  # Define token types with their corresponding regex patterns
  TOKEN_TYPES = [
    [:def, /\bdef\b/],            # Keyword 'def'
    [:end, /\bend\b/],            # Keyword 'end'
    [:identifier, /\b[a-zA-Z]+\b/], # Identifiers (letters only)
    [:integer, /\b[0-9]+\b/],     # Integers (one or more digits)
    [:oparen, /\(/],              # Open parenthesis '('
    [:cparen, /\)/],              # Close parenthesis ')'
    [:comma, /,/],                # Comma ','
  ]

  # Initialize with the code to be tokenized
  def initialize(code)
    @code = code
  end

  # Main method to tokenize the entire code
  def tokenize
    tokens = []
    until @code.empty?
      tokens << tokenize_one_token      # Add one token at a time
      @code = @code.strip               # Strip whitespace after each token
    end
    tokens                              # Return the list of tokens
  end

  # Tokenize a single token by matching it with each defined token type
  def tokenize_one_token
    TOKEN_TYPES.each do |type, re|
      re = /\A(#{re})/                  # Ensure pattern matches start of string
      if @code =~ re                    # Check if the pattern matches
        value = $1                      # Capture the matched value
        @code = @code[value.length..-1] # Remove the matched portion from @code
        return Token.new(type, value)   # Return a new token with type and value
      end
    end

    # Raise an error if no token type matches the current part of the code
    raise RuntimeError.new("Couldn't match the token on #{@code.inspect}!")
  end
end
