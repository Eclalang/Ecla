# Regex

La librairie Regex implémente les expressions de recherche régulière

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
Retourne la première correspondance de regex dans le string

### Exemple :
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
Retourne toutes les correspondances de regex dans le string

### Exemple :
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
Retourne l'index de toutes les correspondances de regex dans le string

### Exemple :
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
Retourne le premier et dernier index de la première correspondance de regex dans le string

### Exemple :
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
Retourne true si le regex correspond au string

### Exemple :
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
Remplace toutes les correspondances de regex dans le string avec le nouveau string

### Exemple :
```ecla
import "console";
import "regex";

function testReplaceAll() {
    var str string = "abc abc";
    var replaced string = regex.replaceAll("a", str, "b");
    console.println(replaced);
}
```