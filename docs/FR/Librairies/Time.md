# Time

La librairie Time implémente les fonctions pour manipuler et afficher le temps.

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
Convertit un nombre en chiffre romain

### Exemple :
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
Retourne la représentation de la date en string

### Exemple :
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
Retourne une représentation de l'heure actuelle en string

### Exemple :
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
Met en pause la goroutine actuelle pour une durée spéficié

### Exemple :
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
Retourne une représentation en string de la date sous un format spécifié

### Exemple :
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
Attends une durée spécifié en seconde 

### Exemple :
```ecla
import "console";
import "time";

function testTimer() {
    console.println("Waiting for 2 seconds");
    time.timer(2);
    console.println("Done");
}
```

## Format de conversion supporté :
| Pattern |                    Description                    |
|:-------:|:-------------------------------------------------:|
|   %d    |      Jour du mois en nombre décimal (01-31)       |
|   %H    | Heure (Horaire sur 24h) en nombre décimal (00-23) |
|   %M    |         Minutes en nombre décimal (00-59)         |
|   %m    |          Mois en nombre décimal (01-12)           |
|   %S    |        Secondes en nombres décimal (00-59)        |
|   %Y    |              Années avec les siècles              |
|   %y    |          Années sans les siècles (00-99)          |
|   %%    |                   Un signe '%'                    |