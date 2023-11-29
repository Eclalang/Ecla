# Strings

Strings library implements function to alter string.

## Index

- [Contains(str string, substr string) bool](#contains)
- [ContainsAny(str string, chars string) bool](#containsany)
- [Count(str string, substr string) int](#count)
- [Cut(str string, substr string) string](#cut)
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

## ContainsAny
```
function containsAny(str string, chars string) bool
```
Returns true if the string contains any of the characters

## Count
```
function count(str string, substr string) int
```
Returns the number of non-overlapping instances of substr in str

## Cut
```
function cut(str string, substr string) string
```
Returns a string before and after the separator, and a bool if it's found or not

## HasPrefix
```
function hasPrefix(str string, prefix string) bool
```
Returns true if the string starts by the prefix

## HasSuffix
```
function hasSuffix(str string, suffix string) bool
```
Returns true if the string ends by the suffix

## IndexOf
```
function indexOf(str string, substr string) int
```
Returns the index of the first instance of substr in str, or -1 if not found

## Join
```
function join(elems []string, sep string) string
```
Returns a concatenated string from an array of string separated by sep

## Replace
```
function replace(str string, old string, new string, nb int) string
```
Returns a string with the first instance of old replaced by new

## ReplaceAll
```
function replaceAll(str string, old string, new string) string
```
Returns a string with all instances of old replaced by new

## Split
```
function split(str string, sep string) []string
```
Returns an array of the substrings between the separator, or an array only containing str if it doesn't contain sep

## SplitAfter
```
function splitAfter(str string, sep string) []string
```
Returns an array of the substrings after the separator, or an array only containing str if it doesn't contain sep

## SplitAfterN
```
function splitAfterN(str string, sep string, nb int) []string
```
Returns an array of the substrings after the separator, or an array only containing str if it doesn't contain sep. The count determines the number of substrings to return

## SplitN
```
function splitN(str string, sep string, nb int) []string
```
Returns an array of the substrings between the separator, or an array only containing str if it doesn't contain sep. The count determines the number of substrings to return

## ToLower
```
function toLower(str string) string
```
Returns a string with all characters in lowercase

## ToUpper
```
function toUpper(str string) string
```
Returns a string with all characters in uppercase

## Trim
```
function trim(str string, cutset string) string
```
Returns a string with all cut characters removed