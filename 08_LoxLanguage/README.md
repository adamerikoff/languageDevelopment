# Lox: A Dynamically Typed, Interpreted Language

Lox is a simple, dynamically typed programming language designed to help users explore and understand the core concepts of language implementation. Whether you're a curious learner or an aspiring language developer, Lox provides an approachable way to dive into the mechanics of modern programming languages.

## Key Features

- **Dynamic Typing**: 
  Variables don't have fixed data types. Their type is determined at runtime based on the value assigned to them.
  
- **Garbage Collection**: 
  Automatic memory management handles memory allocation and deallocation, freeing the programmer from manual memory management.

- **Simple Syntax**: 
  Lox boasts a clean and concise syntax, making it easy to learn and write.

- **First-Class Functions**: 
  Functions in Lox are first-class citizens, meaning they can be assigned to variables, passed as arguments, and returned from other functions.

- **Classes and Inheritance**: 
  Lox supports object-oriented programming with classes and inheritance, allowing users to build reusable and modular code.

## Language Architecture

Lox's architecture is built around three main components:

1. **Lexer**: 
   The lexer (or scanner) breaks the source code into tokens such as keywords, identifiers, operators, and literals. This is the first step in transforming the source code into a structured format.

2. **Parser**: 
   The parser constructs an Abstract Syntax Tree (AST) from the token stream. The AST represents the syntactic structure of the program in a hierarchical form.

3. **Interpreter**: 
   The interpreter traverses the AST, executing statements and evaluating expressions to bring the program to life.

## Why Learn Lox?

Lox is specifically designed to be approachable for those interested in understanding the building blocks of programming languages. By working with Lox, you'll gain hands-on experience with:

- Lexical analysis and tokenization.
- Parsing and constructing abstract syntax trees.
- Writing interpreters for executing code.

Lox serves as an excellent stepping stone for diving deeper into language design and compiler theory.

## Resources

- [Crafting Interpreters](https://craftinginterpreters.com/).
