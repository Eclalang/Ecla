# Encoding

Encoding library is used to convert data to and from byte-level and textual representation.

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
Converts ascii to string

### Example :
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
Decodes base64 to ascii

### Example :
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
Decodes gob to string

### Example :
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
Decodes hex to ascii

### Example :
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
Encodes ascii to base64

### Example :
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
Encodes string to gob

### Example :
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
Encodes ascii to hex

### Example :
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
Converts string to ascii

### Example :
```ecla
import "encoding";
import "console";

function testStringToAscii() {
    var str string = "ABC";
    var intArr []int = encoding.stringToAscii(str);
    console.println(intArr);
}
```