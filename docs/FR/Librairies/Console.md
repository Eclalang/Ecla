# Console

La librairie console est utilisée pour gérer ce qui est lié à la console.

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

## Confirm
```
function confirm(prompt string) bool
```
> Obsolète

Récupère la confirmation de l'utilisateur

### Exemple :
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
Récupère l'entrée de l'utilisateur depuis la console

### Exemple :
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
Lit les floats en entrée

### Exemple :
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
Lit les int en entrée

### Exemple :
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
Écrit les args formatté en string dans la console

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
    console.printInColor("red", "Hello World!");
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

## ProgressBar
```
function progressBar(percent int)
```
> Obsolète

Affiche une barre de progression

### Exemple :
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