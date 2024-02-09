# Encoding

La librairie Encoding est utilisée pour convertir de la data à partir du byte-level et en byte-level ou de représentation textuelle.

## Index

- [AsciiToString(intArr []int) string ](#asciitostring)
- [DecodeBase64(str string) []int ](#decodebase64)
- [DecodeGob(intArr []int) string ](#decodegob)
- [DecodeHex(str string) []int](#decodehex)
- [EncodeBase64(intArr []int) string](#encodebase64)
- [EncodeGob(str string) []int](#encodegob)
- [EncodeHex(intArr []int) string](#encodehex)
- [StringToAscii(str string) []int](#stringtoascii)

## AsciiToString
```
function asciiToString(intArr []int) string
```
Convertit l'ascii en string

### Exemple :
```ecla
import "encoding";
import "console";

function testAsciiToString() {
    var intArr []int = [65, 66, 67];
    var str string = encoding.asciiToString(intArr);
    console.println(str);
}
```

## DecodeBase64
```
function decodeBase64(str string) []int
```
Décode la base64 en ascii

### Exemple :
```ecla
import "encoding";
import "console";

function testDecodeBase64() {
    var str string = "QUJD";
    var intArr []int = encoding.decodeBase64(str);
    console.println(intArr);
}
```

## DecodeGob
```
function decodeGob(intArr []int) string
```
Décode gob en string

### Exemple :
```ecla
import "encoding";
import "console";

function testDecodeGob() {
    var intArr []int = [71, 111, 98, 32, 105, 115, 32, 97, 119, 101, 115, 111, 109, 101, 33];
    var str string = encoding.decodeGob(intArr);
    console.println(str);
}
```

## DecodeHex
```
function decodeHex(str string) []int
```
Décode l'hexadécimal en ascii

### Exemple :
```ecla
import "encoding";
import "console";

function testDecodeHex() {
    var str string = "414243";
    var intArr []int = encoding.decodeHex(str);
    console.println(intArr);
}
```

## EncodeBase64
```
function encodeBase64(intArr []int) string
```
Encode l'ascii en base64

### Exemple :
```ecla
import "encoding";
import "console";

function testEncodeBase64() {
    var intArr []int = [65, 66, 67];
    var str string = encoding.encodeBase64(intArr);
    console.println(str);
}
```

## EncodeGob
```
function encodeGob(str string) []int
```
Encode un string en gob

### Exemple :
```ecla
import "encoding";
import "console";

function testEncodeGob() {
    var str string = "Gob is awesome!";
    var intArr []int = encoding.encodeGob(str);
    console.println(intArr);
}
```

## EncodeHex
```
function encodeHex(intArr []int) string
```
Encode l'ascii en hexadécimal

### Exemple :
```ecla
import "encoding";
import "console";

function testEncodeHex() {
    var intArr []int = [65, 66, 67];
    var str string = encoding.encodeHex(intArr);
    console.println(str);
}
```

## StringToAscii
```
function stringToAscii(str string) []int
```
Convertit un string en ascii

### Exemple :
```ecla
import "encoding";
import "console";

function testStringToAscii() {
    var str string = "ABC";
    var intArr []int = encoding.stringToAscii(str);
    console.println(intArr);
}
```