# Time

Time library implements functions to manipulate and display time.

## Index

- [ConvertRoman(str string) string](#convertroman)
- [Date(year, month, day, hour, min, sec int) string](#date)
- [Now() string](#now)
- [Sleep(sec int)](#sleep)
- [Strftime(format, date string) string](#strftime)
- [Timer(sec int)](#timer)

## ConvertRoman
```
function convertRoman(str string) string
```
Converts a number to a roman number

## Date
```
function date(year, month, day, hour, min, sec int) string
```
Returns a string representation of a date

## Now
```
function now() string
```
Returns a string representation of the current time

## Sleep
```
function sleep(sec int)
```
Pauses the current goroutine for a specified number of seconds

## Strftime
```
function strftime(format, date string) string
```
Returns a string representation of a date according to a specified format

## Timer
```
function timer(sec int)
```
Waits for a specified number of seconds

## Supported Conversion Formats :
| Pattern |                   Description                    |
|:-------:|:------------------------------------------------:|
|   %d    |   Day of the month as a decimal number (01-31)   |
|   %H    | Hour (24-hour clock) as a decimal number (00-23) |
|   %M    |        Minute as a decimal number (00-59)        |
|   %m    |        Month as a decimal number (01-12)         |
|   %S    |        Second as a decimal number (00-59)        |
|   %Y    |                Year with century                 |
|   %y    |           Year without century (00-99)           |
|   %%    |                    A '%' sign                    |