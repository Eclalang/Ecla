# Console

Console library is used for dealing with console.

## Index

- [Clear()](#clear)
- [Confirm(prompt string) bool](#confirm)
- [Input(args ...interface{})](#input)
- [InputFloat(prompt string)](#inputfloat)
- [InputInt(prompt string)](#inputint)
- [Print(args ...interface{})](#print)
- [Printf(format string, args ...interface{})](#printf)
- [PrintInColor(color string, args ...interface{})](#printincolor)
- [Println(args ...interface{})](#println)
- [ProgressBar(percent int)](#progressbar)

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

## Confirm
```
function confirm(prompt string) bool
```
> deprecated

Get user confirmation

### Example :
```ecla
import "console";

function testConfirm() {
    var check bool = console.confirm("Do you want to continue :");
    if (check) {
        console.println("You have confirmed.");
    } else {
        console.println("You have not confirmed.");
    }
}
```

## Input
```
function input(args ...interface{})
```
Takes user input from console

### Example :
```ecla
import "console";

function testInput() {
    var input string = console.input("Enter your name : ");
    console.println("Hello " + input + "!");
}
```

## InputFloat
```
function inputFloat(prompt string)
```
Read float input

### Example :
```ecla
import "console";

function testInputFloat() {
    var input float = console.inputFloat("Enter a float : ");
    console.println("You entered " + input + ".");
}
```

## InputInt
```
function inputInt(prompt string)
```
Read int input

### Example :
```ecla
import "console";

function testInputInt() {
    var input int = console.inputInt("Enter an int : ");
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
    console.printInColor("red", "Hello World!");
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

## ProgressBar
```
function progressBar(percent int)
```
> deprecated
  
Display a progress bar

### Example :
```ecla
import "console";
import "time";

function testProgressBar() {
    for (var i int = 0, i <= 100, i+= 10) {
        console.progressBar(i);
        time.sleep(1);
        console.clear();
    }
}
```