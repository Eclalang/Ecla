# Strings

La librairie Strings implémente les fonctions pour altérer les strings.

## Index

- [Contains(str string, substr string) bool](#contains)
- [ContainsAny(str string, chars string) bool](#containsany)
- [Count(str string, substr string) int](#count)
- [Cut(str string, substr string) string, string, bool](#cut)
- [HasPrefix(str string, prefix string) bool](#hasprefix)
- [HasSuffix(str string, suffix string) bool](#hassuffix)
- [IndexOf(str string, substr string) int](#indexof)
- [Join(elems []string, sep string) string](#join)
- [Replace(str string, old string, new string, nb int) string](#replace)
- [ReplaceAll(str string, old string, new string) string](#replaceall)
- [Split(str string, sep string) []string](#split)
- [SplitAfter(str string, sep string) []string](#splitafter)
- [SplitAfterN(str string, sep string, nb int) []string](#splitaftern)
- [SplitN(str string, sep string, nb int) []string](#splitn)
- [ToLower(str string) string](#tolower)
- [ToUpper(str string) string](#toupper)
- [Trim(str string, cutset string) string](#trim)

## Contains
```
function contains(str string, substr string) bool
```
Retourne true si le string contient le substring

### Exemple :
```ecla
import "console";
import "strings";

function testContains() {
    var str string = "Hello World";
    var substr string = "World";
    console.println(strings.contains(str, substr));
}
```

## ContainsAny
```
function containsAny(str string, chars string) bool
```
Retourne true si le string contient le caractère chars
### Exemple :
```ecla
import "console";
import "strings";

function testContainsAny() {
    var str string = "Hello World";
    var chars string = "abc";
    console.println(strings.containsAny(str, chars));
}
```

## Count
```
function count(str string, substr string) int
```
Retourne le nombre d'instances non successives de substr dans str

### Exemple :
```ecla
import "console";
import "strings";

function testCount() {
    var str string = "Hello World";
    var substr string = "l";
    console.println(strings.count(str, substr));
}
```

## Cut
```
function cut(str string, substr string) string, string, bool
```

Retourne un string avant et après le séparateur, et un bool s'il est trouvé ou non

### Exemple :
```ecla
import "console";
import "strings";

function testCut() {
    var str string = "Hello World";
    var substr string = "W";
    var before string;
    var after string;
    var found bool;
    before, after, found = strings.cut(str, substr);
    console.println(before);
    console.println(after);
    console.println(found);
}
```

## HasPrefix
```
function hasPrefix(str string, prefix string) bool
```
Retourne true si le string commence par le préfixe

### Exemple :
```ecla
import "console";
import "strings";

function testHasPrefix() {
    var str string = "Hello World";
    var prefix string = "Hello";
    console.println(strings.hasPrefix(str, prefix));
}
```

## HasSuffix
```
function hasSuffix(str string, suffix string) bool
```
Retourne true si le string finit par le suffixe

### Exemple :
```ecla
import "console";
import "strings";

function testHasSuffix() {
    var str string = "Hello World";
    var suffix string = "World";
    console.println(strings.hasSuffix(str, suffix));
}
```

## IndexOf
```
function indexOf(str string, substr string) int
```
Retourne l'index de la première instance du substr dans str, ou -1 s'il n'est pas trouvé

### Exemple :
```ecla
import "console";
import "strings";

function testIndexOf() {
    var str string = "Hello World";
    var substr string = "l";
    console.println(strings.indexOf(str, substr));
}
```

## Join
```
function join(elems []string, sep string) string
```

Retourne un string concaténé d'un array de strings séparé par sep


### Exemple :
```ecla
import "console";
import "strings";

function testJoin() {
    var elems []string;
    elems = append(elems, "Hello");
    elems = append(elems, "World");
    var sep string = " ";
    console.println(strings.join(elems, sep));
}
```

## Replace
```
function replace(str string, old string, new string, nb int) string
```
Retourne un string avec la première instance de old remplacée par new

### Exemple :
```ecla
import "console";
import "strings";

function testReplace() {
    var str string = "Hello World";
    var old string = "l";
    var new string = "z";
    var nb int = 1;
    console.println(strings.replace(str, old, new, nb));
}
```

## ReplaceAll
```
function replaceAll(str string, old string, new string) string
```
Retourne un string avec toutes les instances de old remplacé par new

### Exemple :
```ecla
import "console";
import "strings";

function testReplaceAll() {
    var str string = "Hello World";
    var old string = "l";
    var new string = "z";
    console.println(strings.replaceAll(str, old, new));
}
```

## Split
```
function split(str string, sep string) []string
```
Retourne un array d'une sous-chaine de caractères entre les séparateurs, ou un array qui contient seulement str s'il ne contient pas sep

### Exemple :
```ecla
import "console";
import "strings";

function testSplit() {
    var str string = "Hello World";
    var sep string = " ";
    var array []string;
    array = strings.split(str, sep);
    console.println(array[0]);
    console.println(array[1]);
}
```

## SplitAfter
```
function splitAfter(str string, sep string) []string
```
Retourne un array de sous-chaines de caractères après le séparateur, ou un array qui contient uniquement str s'il ne contient pas sep

### Exemple :
```ecla
import "console";
import "strings";

function testSplitAfter() {
    var str string = "Hello World";
    var sep string = " ";
    var array []string;
    array = strings.splitAfter(str, sep);
    console.println(array[0]);
    console.println(array[1]);
}
```

## SplitAfterN
```
function splitAfterN(str string, sep string, nb int) []string
```
Retourne un array de substring après le séparateur, ou un array qui contient uniquement str s'il ne contient pas sep. Le compte détermine le nombre de substring à retourner

### Exemple :
```ecla
import "console";
import "strings";

function testSplitAfterN() {
    var str string = "Hello World";
    var sep string = " ";
    var nb int = 2;
    var array []string;
    array = strings.splitAfterN(str, sep, nb);
    console.println(array[0]);
    console.println(array[1]);
}
```

## SplitN
```
function splitN(str string, sep string, nb int) []string
```
Retourne un array de substring entre le séparateur, ou un array qui contient seulement str s'il ne contient pas sep. Le compte détermine le nombre de substring à retourner

### Exemple :
```ecla
import "console";
import "strings";

function testSplitN() {
    var str string = "Hello World";
    var sep string = " ";
    var nb int = 2;
    var array []string;
    array = strings.splitN(str, sep, nb);
    console.println(array[0]);
    console.println(array[1]);
}
```

## ToLower
```
function toLower(str string) string
```
Retourne un string avec tous les caractères en minuscule

### Exemple :
```ecla
import "console";
import "strings";

function testToLower() {
    var str string = "Hello World";
    console.println(strings.toLower(str));
}
```

## ToUpper
```
function toUpper(str string) string
```
Retourne un string avec tous les caractères en majuscule

### Exemple :
```ecla
import "console";
import "strings";

function testToUpper() {
    var str string = "Hello World";
    console.println(strings.toUpper(str));
}
```

## Trim
```
function trim(str string, cutset string) string
```
Retourne un string avec tous les caractères coupés retirés

### Exemple :
```ecla
import "console";
import "strings"

function testTrim() {
    var str string = "Hello World";
    var cutset string = "Hd";
    console.println(strings.trim(str, cutset));
}
```