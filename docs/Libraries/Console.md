# Console

Console library is used for dealing with console.

## Index

- [Clear()](#clear)
- [Confirm(prompt string)](#confirm)
- [Input(args ...interface{})](#input)
- [InputFloat(prompt string)](#inputfloat)
- [InputInt(prompt string)](#inputint)
- [Print(args ...interface{})](#print)
- [Printf(format string, args ...interface{})](#printf)
- [PrintInColor(color string, args ...interface{})](#printincolor)
- [Println(args ...interface{})](#println)
- [ProgressBar(percent int)](#progressbar)

## Clear
```
function clear()
```
Clear the console

## Confirm
```
function confirm(prompt string)
```
Get user confirmation

## Input
```
function input(args ...interface{})
```
Takes user input from console

## InputFloat
```
function inputFloat(prompt string)
```
Read float input

## InputInt
```
function inputInt(prompt string)
```
Read int input

## Print
```
function print(args ...interface{})
```
Print args to console without newline

## Printf
```
function printf(format string, args ...interface{})
```
Print args as formatted string to console

## PrintInColor
```
function printInColor(color string, args ...interface{})
```
Print args to console with color

## Println
```
function println(args ...interface{})
```
Print args to console

## ProgressBar
```
function progressBar(percent int)
```
Display a progress bar