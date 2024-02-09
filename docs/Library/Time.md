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

### Example :
```ecla
import "console";
import "time";

function testConvertRoman() {
    console.println(time.convertRoman("4"));
}
```

## Date
```
function date(year, month, day, hour, min, sec int) string
```
Returns a string representation of a date

### Example :
```ecla
import "console";
import "time";

function testDate() {
    console.println(time.date(2020, 12, 31, 23, 59, 59));
}
```

## Now
```
function now() string
```
Returns a string representation of the current time

### Example :
```ecla
import "console";
import "time";

function testNow() {
    console.println(time.now());
}
```

## Sleep
```
function sleep(sec int)
```
Pauses the current goroutine for a specified number of seconds

### Example :
```ecla
import "console";
import "time";

function testSleep() {
    console.println("Sleeping for 2 seconds");
    time.sleep(2);
    console.println("Done");
}
```

## Strftime
```
function strftime(format, date string) string
```
Returns a string representation of a date according to a specified format

### Example :
```ecla
import "console";
import "time";

function testStrftime() {
    console.println(time.strftime("%Y-%m-%d %H:%M:%S", "2020-12-31 23:59:59"));
}
```

## Timer
```
function timer(sec int)
```
Waits for a specified number of seconds

### Example :
```ecla
import "console";
import "time";

function testTimer() {
    console.println("Waiting for 2 seconds");
    time.timer(2);
    console.println("Done");
}
```

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