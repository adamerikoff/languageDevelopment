# ponGo Programming Language

<img src="https://github.com/adamerikoff/ponGo/blob/main/pongo.jpg" align="center" alt="Size Limit logo by Anton Lovchikov" width="100">

**ponGo** is a simple and expressive programming language implemented in Go, designed to be easy to learn and use for both beginners and experienced developers alike.

## Features

- **Pet Project**: ponGo is a personal project aimed at exploring language design and implementation. It's a great way to learn about compilers and interpreters.
- **Simplicity**: ponGo aims to be simple and easy to understand, with a clean syntax inspired by popular programming languages.
- **Expressiveness**: Despite its simplicity, ponGo is capable of expressing a wide range of computational tasks, from basic arithmetic operations to complex control flow and data manipulation.
- **Flexibility**: The language is designed to be flexible, allowing developers to build various types of applications, including command-line tools, web servers, and more.
- **Interpreted**: ponGo is implemented as an interpreter in Go, making it easy to run and experiment with code without the need for compilation.

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
	@ five = 5.
	@ ten = 10.

	@ add = fn(x, y) {
		x + y.
	}.
	@ result = add(five, ten).

	!-/*5.
   	5 < 10 > 5.

	if (5 < 10) {
		<< T.
	} else {
		<< F.
 	}
	10 == 10.
	10 != 9.
   
```

## Documentation

to add

## Contributing

Contributions to ponGo are welcome! Whether you want to report a bug, suggest a feature, or contribute code, please feel free to open an issue or submit a pull request on GitHub.

## License

ponGo is open-source software released under the [MIT License](LICENSE). Feel free to use, modify, and distribute it according to your needs.