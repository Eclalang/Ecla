# Ecla Programming Language Quickstart Guide

Welcome to Ecla! This guide will help you get started with the basics of the Ecla programming language.

## Installation

Before you can start coding in Ecla, you'll need to install the language on your system. Please follow the [installation guide](https://github.com/Eclalang/Ecla/blob/main/INSTALL.md) for detailed instructions.

## Creating a Simple Hello World Program

Let's dive into creating a simple "Hello, World!" program in Ecla. We'll break down each step:

### Importing Libraries

```ecla
import "console";
```

In Ecla, you can import libraries to extend the functionality of your programs. Here, we import the `console` library to enable basic input and output operations.

### Variable Declaration

```ecla
var hello string = "Hello";
world := "World !";
```

In Ecla, you can declare variables using the `var` keyword followed by the variable name and its type, optionally initializing it with a value. Alternatively, you can use the `:=` shorthand for implicit variable declaration and initialization.

### Concatenating Strings

```ecla
var helloWorld string;
helloWorld = hello + ", " + world;
```

Here, we concatenate the strings "Hello" and "World !" into a new variable `helloWorld` using the `+` operator.

### Creating a Function

```ecla
function printHelloWorld() {
    console.println(helloWorld);
}
```

In Ecla, you can define functions using the `function` keyword followed by the function name and its parameters (if any). Functions encapsulate reusable blocks of code. Here, we define a function `printHelloWorld` that prints the value of `helloWorld` to the console.

### Calling a Function

```ecla
printHelloWorld();
```

To execute the code within a function, you need to call it. Here, we call the `printHelloWorld` function to display "Hello, World !" to the console.

### Full Code

```ecla
import "console";

var hello string = "Hello";
world := "World !";

var helloWorld string;
helloWorld = hello + ", " + world;

function printHelloWorld() {
    console.println(helloWorld);
}

printHelloWorld();
```

Congratulations! You've just created your first program in Ecla.

### Run your Program

To run your program, save the code to a file with a `.ecla` extension (e.g., `hello.ecla`). Then, open a terminal and navigate to the directory where the file is saved. If you've installed Ecla correctly, you can run the program using the following command:

```bash
ecla hello.ecla
```

When you run the program, you should see "Hello, World !" printed to the console.

You're now ready to embark on your journey into the world of programming with Ecla.
If you want to learn more about the language, check out the [official documentation](https://github.com/Eclalang/LearnEcla/blob/main/README.md) 
> the documentation is still under construction, but you can still find some useful information there.

