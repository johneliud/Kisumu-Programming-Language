# Kisumu-Programming-Language

This project is a simple toy programming language built in Go programming language and uses a syntax similar to Go programming language.

## About The Language

Kisumu Programming Language is an interpreted language built using [Go](https://go.dev/) and adopts a programming syntax similar to that of Go. The language may be limited to a number of functionalities and data types supported by other programming languages and is therefore not intended for professional adoption. The language currently supports the following data types:

1. [Integer](https://en.wikipedia.org/wiki/Integer_(computer_science))

2. [String](https://en.wikipedia.org/wiki/String_(computer_science))

3. [Boolean](https://en.wikipedia.org/wiki/Boolean_data_type)

## Prerequisites

An additional programming environment, [Go](https://go.dev/doc/install) is required to be installed on the users local environment in order to use the language.

## Installation

1. Clone the repository from the remote location to your local development environment using the below command:
    ```bash
    git clone https://github.com/johneliud/Kisumu-Programming-Language.git
    ```

2. Navigate to the programming language source code using the below command:
    ```bash
    cd Kisumu-Programming-Language
    ```

## Usage

### Declaring Integer Types
```bash
var age = 1;
```

### Declaring String Types
```bash
var name = "Kisumu";
```

### Declaring Boolean Types
```bash
var isTrue = true;
```

### Printing Result
```bash
print()
```

### Writing Source Code

The program currently offers a file `source-code.ksm` where the user can write their source for interpretation by the language.

**NOTE**: Every source code line written on the file should end with a `semi-colon`.

### Running The Program

Once you have written your source code on the `source-code.ksm`, open the terminal and type the command:
```bash
go run .
```

This command helps in evaluating what is present in your source code.

## Examples

1. As standard with every other programming language, let's start with having a program that prints `Hello World!` on the standard output.
    - Declare a variable `message` and assign the string `Hello World!` to it.
        ```bash
        var message = "Hello World!";
        ```

    - Print variable `message`
        ```bash
        print(message);
        ```

    - Run the program
        ```bash
        go run .
        ```

2. Program to find the sum of a given set of numbers.
    - Declare a variable `a` and assign the value `2` to it and variable `b` and assign the value `10` to it.
        ```bash
        var a = 2;
        var b = 10;
        ```

    - Declare a variable `c` and assign it to be the result of `a` and `b`.
        ```bash
        var c = a + b;
        ```

    - Print the value of `c`
        ```bash
        print(c)
        ```

    - Run the program
        ```bash
        go run .
        ```

3. String Concatenation

    - Declare variable `a` and assign the value `Kisumu` to it. Declare another variable `b` and assign the value `Language` to it.
        ```bash
        var a = "Kisumu";
        var b = "Language";
        ```

    - Declare variable `c` which will concatenate `a` and `b` with a space separation.
        ```bash
        var c = a + " " + b;
        ```

    - Print the value of `c`
        ```bash
        print(c)
        ```

    - Run the program.
        ```bash
        go run .
        ```