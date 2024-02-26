# Console

La librairie console est utilisée pour gérer ce qui est lié à la console.

## Index

- [Clear()](#clear)
- [Input() string](#input)
- [InputFloat() float](#inputfloat)
- [InputInt() int](#inputint)
- [Print(args ...interface{})](#print)
- [Printf(format string, args ...interface{})](#printf)
- [PrintInColor(color string, args ...interface{})](#printincolor)
- [Println(args ...interface{})](#println)

## Clear
```
function clear()
```
Efface le contenu de la console

### Exemple :
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
function input(args ...interface{})
```
Récupère l'entrée de l'utilisateur depuis la console

### Exemple :
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
function inputFloat() float
```
Lit les floats en entrée

### Exemple :
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
function inputInt() int
```
Lit les int en entrée

### Exemple :
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
Écrit les args dans la console sans nouvelle ligne

### Exemple :
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
Écrit les args formatés en string dans la console

### Exemple :
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
Écrit les args dans la console avec la couleur

### Exemple :
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
Écrit les args dans la console

### Exemple :
```ecla
import "console";

function testPrintln() {
    console.println("Hello World!");
    console.println("Hello", "user", "?");
}
```
