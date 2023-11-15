# TIME

***
##  Documentation.
### Time is a library that allows you to interact with the operating system.
###

***
## Index.

* [fonction Date() string](#fonction-date)
* [fonction Now() string](#fonction-now)
* [fonction Sleep(seconds int)](#fonction-sleep)
* [fonction Strftime(format string, t time.Time) string](#fonction-strftime)
* [fonction Timer(seconds int)](#fonction-timer)

##
### Fonction Date
```
function Date() string
```
Returns a string representation of a date
### Fonction Now
```
function Now() string
```
Returns a string representation of the current time
### Fonction Sleep
```
function Sleep(seconds int)
```
Pauses the current goroutine for a specified number of seconds
### Fonction Strftime
```
function Strftime(format string, t time.Time) string
```
Returns a string representation of a date according to a specified format
### Fonction Timer
```
function Timer(seconds int)
```
Waits for a specified number of seconds
##
# Supported Conversion Formats :
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