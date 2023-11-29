# Regex

Regex library implements regular expression search.

## Index

- [Find(regex string, str string) string](#find)
- [FindAll(regex string, str string) []string](#findall)
- [FindAllIndex(regex string, str string) []int](#findallindex)
- [FindIndex(regex string, str string) []int](#findindex)
- [Match(regex string, str string) bool](#match)
- [ReplaceAll(regex string, str string, new string) string](#replaceall)

## Find
```
function find(regex string, str string) string
```
Returns the first match of the regex in the string

## FindAll
```
function findAll(regex string, str string) []string
```
Returns all matches of the regex in the string

## FindAllIndex
```
function findAllIndex(regex string, str string) []int
```
Returns the indexes of all matches of the regex in the string

## FindIndex
```
function findIndex(regex string, str string) []int
```
Returns the first and last index of the first match of the regex in the string

## Match
```
function match(regex string, str string) bool
```
Returns true if the regex matches the string

## ReplaceAll
```
function replaceAll(regex string, str string, new string) string
```
Replaces all matches of the regex in the string with the new string