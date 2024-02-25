# Console

Console library is used for dealing with console.

## Index

- [Clear()](#clear)
- [Input() string](#input)
- [InputFloat() (float, error)](#inputfloat)
- [InputInt() (int, error)](#inputint)
- [Print(args ...interface{})](#print)
- [Printf(format string, args ...interface{})](#printf)
- [PrintInColor(color string, args ...interface{})](#printincolor)
- [Println(args ...interface{})](#println)

## Clear
```
function clear()
```
Clear the console

### Example :
```ecla
import "console";

function testClear() {
    console.print("Some text to clear.");
    console.clear();
    console.print("All cleared :)");
}
```

## Input
```
function input() string
```
Takes user input from console

### Example :
```ecla
import "console";

function testInput() {
    console.print("Enter your name : ");
    var input string = console.input();
    console.println("Hello " + input + "!");
}
```

## InputFloat
```
function inputFloat() (float, error)
```
Read float input

### Example :
```ecla
import "console";

function testInputFloat() {
    console.print("Enter a float : ");
    var input float = console.inputFloat();
    console.println("You entered " + input + ".");
}
```

## InputInt
```
function inputInt() (int, error)
```
Read int input

### Example :
```ecla
import "console";

function testInputInt() {
    console.print("Enter an int : ");
    var input int = console.inputInt();
    console.println("You entered " + input + ".");
}
```

## Print
```
function print(args ...interface{})
```
Print args to console without newline

### Example :
```ecla
import "console";

function testPrint() {
    console.print("Hello ");
    console.print("World!");
}
```

## Printf
```
function printf(format string, args ...interface{})
```
Print args as formatted string to console

### Example :
```ecla
import "console";

function testPrintf() {
    console.printf("Hel%s %s!", "lo", "World");
}
```

## PrintInColor
```
function printInColor(color string, args ...interface{})
```
Print args to console with color

### Example :
```ecla
import "console";

function testPrintInColor() {
    console.printInColor("\033[0;31m", "Hello World!");
}
```

## Println
```
function println(args ...interface{})
```
Print args to console

### Example :
```ecla
import "console";

function testPrintln() {
    console.println("Hello World!");
    console.println("Hello", "user", "?");
}
```
