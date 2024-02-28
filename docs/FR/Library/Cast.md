# Cast

La librairie cast est utilisée pour la conversion de type

## Index

- [Atoi(str string) int](#atoi)
- [FloatToInt(f float) int](#floattoint)
- [IntToFloat(i int) float](#inttofloat)
- [ParseBool(str string) bool](#parsebool)
- [ParseFloat(str string, bitSize int) float](#parsefloat)

## Atoi
```
function atoi(str string) int
```
Convertit un string en int

### Exemple :
```ecla
import "cast";
import "console";

function testAtoi() {
    var str string = "42";
    var x int = cast.atoi(str);
    console.print(x);
}
```

## FloatToInt
```
function floatToInt(x float) int
```
Convertit un float en int

### Exemple :
```ecla
import "cast";
import "console";

function testFloatToInt() {
    var x float = 42.42;
    var y int = cast.floatToInt(x);
    console.print(y);
}
```

## IntToFloat
```
function intToFloat(x int) float
```
Convertit un int en float

### Exemple :
```ecla
import "cast";
import "console";

function testIntToFloat() {
    var x int = 42;
    var y float = cast.intToFloat(x);
    console.print(y);
}
```

## ParseBool
```
function parseBool(str string) bool
```
Convertit un string en bool

### Exemple :
```ecla
import "cast";
import "console";

function testParseBool() {
    var str string = "true";
    var check bool = cast.parseBool(str);
    console.print(check);
}
```

## ParseFloat
```
function parseFloat(str string, x int) float
```
Convertit un string en float

### Exemple :
```ecla
import "cast";
import "console";

function testParseFloat() {
    var str string = "42.42";
    var x int = 64;
    var y float = cast.parseFloat(str, x);
    console.print(y);
}
```