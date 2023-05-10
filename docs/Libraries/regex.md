# REGEX

***
##  Documentation.
### Regex is a library that allows you to interact with the operating system.
###

***
## Index.

* [fonction Find(regex string, str string) string](#fonction-find)
* [fonction FindAll(regex string, str string) []string](#fonction-findall)
* [fonction FindAllIndex(regex string, str string) []int](#fonction-findallindex)
* [fonction FindIndex(regex string, str string) []int](#fonction-findindex)
* [fonction Match(regex string, str string) bool](#fonction-match)
* [fonction ReplaceAll(regex string, str string, replace string) string](#fonction-replaceall)
##
### Fonction Find
```
Find(regex string, str string) string
```
Returns the first match of the regex in the string
### Fonction FindAll
```
FindAll(regex string, str string) []string
```
Returns all matches of the regex in the string
### Fonction FindAllIndex
```
FindAllIndex(regex string, str string) []int
```
Returns the indexes of all matches of the regex in the string
### Fonction FindIndex
```
FindIndex(regex string, str string) []int
```
Returns the first and last index of the first match of the regex in the string
### Fonction Match
```
Match(regex string, str string) bool
```
Returns true if the regex matches the string
### Fonction ReplaceAll
```
ReplaceAll(regex string, str string, replace string) string
```
Replaces all matches of the regex in the string with the new string
##