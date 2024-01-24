# Hash

La librairie Hash permet d'hacher un texte.

## Index

- [Hashmd5(str string) string](#hashmd5)
- [Hashsha1(str string) string](#hashsha1)
- [Hashsha224(str string) string](#hashsha224)
- [Hashsha256(str string) string](#hashsha256)
- [Hashsha384(str string) string](#hashsha384)
- [Hashsha512(str string) string](#hashsha512)
- [Hashsha512_224(str string) string](#hashsha512224)
- [Hashsha512_256(str string) string](#hashsha512256)

## Hashmd5
```
function hashmd5(str string) string
```
Hache un string avec le MD5

### Example :
```ecla
import "console";
import "hash";

function testHashmd5() {
    var str string = "Hello World!";
    var hashed string = hash.hashmd5(str);
    console.println(hashed);
}
```

## Hashsha1
```
function hashsha1(str string) string
```
Hache un string avec sha1

### Example :
```ecla
import "console";
import "hash";

function testHashsha1() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha1(str);
    console.println(hashed);
}
```

## Hashsha224
```
function hashsha224(str string) string
```
Hache un string avec sha224

### Example :
```ecla
import "console";
import "hash";

function testHashsha224() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha224(str);
    console.println(hashed);
}
```

## Hashsha256
```  
function hashsha256(str string) string
```
Hache un string avec sha256

### Example :
```ecla
import "console";
import "hash";

function testHashsha256() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha256(str);
    console.println(hashed);
}
```

## Hashsha384
```
function hashsha384(str string) string
```
Hache un string avec sha384

### Example :
```ecla
import "console";
import "hash";

function testHashsha384() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha384(str);
    console.println(hashed);
}
```

## Hashsha512
```
function hashsha512(str string) string
```
Hache un string avec sha512

### Example :
```ecla
import "console";
import "hash";

function testHashsha512() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha512(str);
    console.println(hashed);
}
```

## Hashsha512_224
```
function hashsha512_224(str string) string
```
Hache un string avec sha512_224

### Example :
```ecla
import "console";
import "hash";

function testHashsha512_224() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha512_224(str);
    console.println(hashed);
}
```

## Hashsha512_256
```
function hashsha512_256(str string) string
```
Hache un string avec sha512_256

### Example :
```ecla
import "console";
import "hash";

function testHashsha512_256() {
    var str string = "Hello World!";
    var hashed string = hash.hashsha512_256(str);
    console.println(hashed);
}
```