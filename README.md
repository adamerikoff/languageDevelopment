# ponGo Programming Language

<img src="https://github.com/adamerikoff/ponGo/blob/main/pongo.jpg" align="center" alt="Size Limit logo by Anton Lovchikov" width="100">

**ponGo** repository contains an interpreter for the "ponGo" programming language, as described in [Write an Interpreter in Go](https://interpreterbook.com).

## Getting Started

To start using ponGo, follow these steps:

1. **Install Go**: Make sure you have Go installed on your system. You can download and install it from the [official Go website](https://golang.org/).

2. **Clone the Repository**: Clone this repository to your local machine using the following command:
   ```
   git clone https://github.com/yourusername/ponGo.git
   ```
3. **Run ponGo**: Navigate to the cloned repository and use the `ponGo` command followed by the path to your ponGo source file to run it:
   ```
   cd ponGo
   go run .
   ```

4. **Experiment**: Start writing and running ponGo code! Experiment with different language features, try out example programs, and have fun exploring the capabilities of ponGo.

## Syntax Example
```
// Define a factorial function
let factorial = fn(x) {
    if (x == 0) {
        return 1;
    } else {
        return x * factorial(x - 1);
    }
};

// Define a variable and assign the result of the factorial function
let result = factorial(5);

// Print the result
puts(result);  // Output should be 120

```
