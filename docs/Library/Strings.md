# Strings

Strings library implements function to alter string.

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
Returns true if the string contains the substring

### Example :
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
Returns true if the string contains any of the characters

### Example :
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
Returns the number of non-overlapping instances of substr in str

### Example :
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
> Not working for the moment

Returns a string before and after the separator, and a bool if it's found or not

### Example :
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
Returns true if the string starts by the prefix

### Example :
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
Returns true if the string ends by the suffix

### Example :
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
Returns the index of the first instance of substr in str, or -1 if not found

### Example :
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
> Not working for the moment

Returns a concatenated string from an array of string separated by sep

### Example :
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
Returns a string with the first instance of old replaced by new

### Example :
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
Returns a string with all instances of old replaced by new

### Example :
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
Returns an array of the substrings between the separator, or an array only containing str if it doesn't contain sep

### Example :
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
Returns an array of the substrings after the separator, or an array only containing str if it doesn't contain sep

### Example :
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
Returns an array of the substrings after the separator, or an array only containing str if it doesn't contain sep. The count determines the number of substrings to return

### Example :
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
Returns an array of the substrings between the separator, or an array only containing str if it doesn't contain sep. The count determines the number of substrings to return

### Example :
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
Returns a string with all characters in lowercase

### Example :
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
Returns a string with all characters in uppercase

### Example :
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
Returns a string with all cut characters removed

### Example :
```ecla
import "console";
import "strings"

function testTrim() {
    var str string = "Hello World";
    var cutset string = "Hd";
    console.println(strings.trim(str, cutset));
}
```