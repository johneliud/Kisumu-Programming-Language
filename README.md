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
- var age = 1;
```

### Declaring String Types
```bash
- var name = "Kisumu";
```

### Declaring Boolean Types
```bash
- var isTrue = true;
```


### Declaring Functions
```bash
func(a, b) { a + b; };
```

### Printing Result
```bash
print()
```

### Writing Source Code

The program currently offers a file `source-code.ksm` where the user can write their source for interpretation by the language.

**NOTE**: Every source code line written on the file should end with a `semi-colon`.

## Running The Program

Once you have written your source code on the `source-code.ksm`, open the terminal and type the command:
```bash
go run .
```

This command helps in evaluating what is present in your source code.

## Examples

As standard with every other programming language, let's start with having a program that prints `Hello World!` on the standard output.
1. Declare a variable `message` and assign the string `Hello World!` to it.
    ```bash
    var message = "Hello World!";
    ```

2. Print variable `message`
    ```bash
    print(message);
    ```

3. Run the program
    ```bash
    go run .
    ```

