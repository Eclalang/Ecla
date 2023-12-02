# Regex

Regex library implements regular expression search.

## Index

- [Find(regex string, str string) string](#find)
- [FindAll(regex string, str string) []string](#findall)
- [FindAllIndex(regex string, str string) [][]int](#findallindex)
- [FindIndex(regex string, str string) []int](#findindex)
- [Match(regex string, str string) bool](#match)
- [ReplaceAll(regex string, str string, new string) string](#replaceall)

## Find
```
function find(regex string, str string) string
```
Returns the first match of the regex in the string

### Example :
```ecla
import "console";
import "regex";

function testFind() {
    var str string = "abc";
    var match string = regex.find("a", str);
    console.println(match);
}
```

## FindAll
```
function findAll(regex string, str string) []string
```
Returns all matches of the regex in the string

### Example :
```ecla
import "console";
import "regex";

function testFindAll() {
    var str string = "abc abc";
    var matches []string = regex.findAll("a", str);
    console.println(matches);
}
```

## FindAllIndex
```
function findAllIndex(regex string, str string) [][]int
```
Returns the indexes of all matches of the regex in the string

### Example :
```ecla
import "console";
import "regex";

function testFindAllIndex() {
    var str string = "abc abc";
    var matches [][]int = regex.findAllIndex("a", str);
    console.println(matches);
}
```

## FindIndex
```
function findIndex(regex string, str string) []int
```
Returns the first and last index of the first match of the regex in the string

### Example :
```ecla
import "console";
import "regex";

function testFindIndex() {
    var str string = "abc abc";
    var matches []int = regex.findIndex("a", str);
    console.println(matches);
}
```

## Match
```
function match(regex string, str string) bool
```
Returns true if the regex matches the string

### Example :
```ecla
import "console";
import "regex";

function testMatch() {
    var str string = "abc";
    var match bool = regex.match("a", str);
    console.println(match);
}
```

## ReplaceAll
```
function replaceAll(regex string, str string, new string) string
```
Replaces all matches of the regex in the string with the new string

### Example :
```ecla
import "console";
import "regex";

function testReplaceAll() {
    var str string = "abc abc";
    var replaced string = regex.replaceAll("a", str, "b");
    console.println(replaced);
}
```