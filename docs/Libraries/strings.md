# STRINGS

***
##  Documentation.
### Strings is a library that allows you to interact with the operating system.
###

***
## Index.

* [fonction Contains(str string, substr string) bool](#fonction-contains)
* [fonction ContainsAny(str string, chars string) bool](#fonction-containsany)
* [fonction Count(str string, substr string) int](#fonction-count)
* [fonction Cutset(str string, cutset string) string](#fonction-cutset)
* [fonction HasPrefix(str string, prefix string) bool](#fonction-hasprefix)
* [fonction HasSuffix(str string, suffix string) bool](#fonction-hassuffix)
* [fonction IndexOf(str string, substr string) int](#fonction-indexof)
* [fonction Join(strs []string, sep string) string](#fonction-join)
* [fonction Replace(str string, old string, new string) string](#fonction-replace)
* [fonction ReplaceAll(str string, old string, new string) string](#fonction-replaceall)
* [fonction Split(str string, sep string) []string](#fonction-split)
* [fonction SplitAfter(str string, sep string) []string](#fonction-splitafter)
* [fonction SplitAfterN(str string, sep string, n int) []string](#fonction-splitaftern)
* [fonction SplitN(str string, sep string, n int) []string](#fonction-splitn)
* [fonction ToLower(str string) string](#fonction-tolower)
* [fonction ToUpper(str string) string](#fonction-toupper)
* [fonction Trim(str string, cutset string) string](#fonction-trim)
##
### Fonction Contains
```
function contains(str string, substr string) bool
```
Returns true if the string contains the substring
### Fonction ContainsAny
```
function containsAny(str string, chars string) bool
```
Returns true if the string contains any of the characters
### Fonction Count
```
function count(str string, substr string) int
```
Returns the number of non-overlapping instances of substr in str
### Fonction Cutset
```
function cutset(str string, cutset string) string
```
Returns a string before and after the separator, and a bool if it's found or not
### Fonction HasPrefix
```
function hasPrefix(str string, prefix string) bool
```
Returns true if the string starts by the prefix
### Fonction HasSuffix
```
function hasSuffix(str string, suffix string) bool
```
Returns true if the string ends by the suffix
### Fonction IndexOf
```
function indexOf(str string, substr string) int
```
Returns the index of the first instance of substr in str, or -1 if not found
### Fonction Join
```
function join(strs []string, sep string) string
```
Returns a concatenated string from an array of string separated by sep
### Fonction Replace
```
function replace(str string, old string, new string) string
```
Returns a string with the first instance of old replaced by new
### Fonction ReplaceAll
```
function replaceAll(str string, old string, new string) string
```
Returns a string with all instances of old replaced by new
### Fonction Split
```
function split(str string, sep string) []string
```
Returns an array of the substrings between the separator, or an array only containing str if it doesn't contains sep
### Fonction SplitAfter
```
function splitAfter(str string, sep string) []string
```
Returns an array of the substrings after the separator, or an array only containing str if it doesn't contains sep
### Fonction SplitAfterN
```
function splitAfterN(str string, sep string, n int) []string
```
Returns an array of the substrings after the separator, or an array only containing str if it doesn't contains sep. The count determines the number of substrings to return
### Fonction SplitN
```
function splitN(str string, sep string, n int) []string
```
Returns an array of the substrings between the separator, or an array only containing str if it doesn't contains sep. The count determines the number of substrings to return
### Fonction ToLower
```
function toLower(str string) string
```
Returns a string with all characters in lowercase
### Fonction ToUpper
```
function toUpper(str string) string
```
Returns a string with all characters in uppercase
### Fonction Trim
```
function trim(str string, cutset string) string
```
Returns a string with all cut characters removed
##