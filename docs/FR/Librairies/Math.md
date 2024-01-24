# Math

La librairie Math implémente les fonctions mathématiques.

## Index.

- [Abs(nb float) float](#abs)
- [Acos(x float) float](#acos)
- [Acosh(x float) float](#acosh)
- [Asin(x float) float](#asin)
- [Asinh(x float) float](#asinh)
- [Atan(x float) float](#atan)
- [Atanh(x float) float](#atanh)
- [Cbrt(nb float) float](#cbrt)
- [Ceil(nb float) float](#ceil)
- [Cos(adj float, hyp float) float](#cos)
- [Cosh(x float) float](#cosh)
- [DegreesToRadians(deg float) float](#degreestoradians)
- [Exp(nb float) float](#exp)
- [Fact(nb int) int](#fact)
- [Floor(nb float) float](#floor)
- [Ln(nb float) float](#ln)
- [Log10(nb float) float](#log10)
- [Max(a float, b float) float](#max)
- [Min(a float, b float) float](#min)
- [Modulo(a float, b float) float](#modulo)
- [Pi() float](#pi)
- [Pow(x float, y float) float](#pow)
- [RadiansToDegrees(rad float) float](#radianstodegrees)
- [Random(min float, max float) float](#random)
- [Round(nb float) float](#round)
- [Sin(opp float, hyp float) float](#sin)
- [Sinh(x float) float](#sinh)
- [Sqrt(nb float) float](#sqrt)
- [Tan(opp float, adj float) float](#tan)
- [Tanh(x float) float](#tanh)
- [Trunc(nb float) float](#trunc)

## Abs
```
function abs(nb float) float
```
Retourne la valeur absolue de nb

### Example :
```ecla
import "console";
import "math";

function testAbs() {
    console.println(math.abs(-5.0));
}
```

## Acos
```
function acos(x float) float
```
Retourne arc cosinus de x 

### Example :
```ecla
import "console";
import "math";

function testAcos() {
    console.println(math.acos(0.5));
}
```

## Acosh
```
function acosh(x float) float
```
Retourne l'arc cosinus hyperbolique de x

### Example :
```ecla
import "console";
import "math";

function testAcosh() {
    console.println(math.acosh(1.5));
}
```

## Asin
```
function asin(x float) float
```
Retourne l'arc sinus de x

### Example :
```ecla
import "console";
import "math";

function testAsin() {
    console.println(math.asin(0.5));
}
```

## Asinh
```
function asinh(x float) float
```
Retourne l'arc sinus hyperbolique de x

### Example :
```ecla
import "console";
import "math";

function testAsinh() {
    console.println(math.asinh(1.5));
}
```

## Atan
```
function atan(x float) float
```
Returns arc tangent of x
Retourne l'arc tangente de x

### Example :
```ecla
import "console";
import "math";

function testAtan() {
    console.println(math.atan(0.5));
}
```

## Atanh
```
function atanh(x float) float
```
Retourne l'arc tangente hyperbolique de x

### Example :
```ecla
import "console";
import "math";

function testAtanh() {
    console.println(math.atanh(0.5));
}
```

## Cbrt
```
function cbrt(nb float) float
```
Retourne la racine cubique de nb

### Example :
```ecla
import "console";
import "math";

function testCbrt() {
    console.println(math.cbrt(8.0));
}
```

## Ceil
```
function ceil(nb float) float
```
Retourne l'entier le plus petit qui est plus grand ou égal à x

### Example :
```ecla
import "console";
import "math";

function testCeil() {
    console.println(math.ceil(1.5));
}
```

## Cos
```
function Cos(adj float, hyp float) float
```
Retourne l'adjacente / hypoténuse

### Example :
```ecla
import "console";
import "math";

function testCos() {
    console.println(math.cos(3.0, 5.0));
}
```

## Cosh
```
function cosh(x float) float
```
Retourne le cosinus hyperbolique de x

### Example :
```ecla
import "console";
import "math";

function testCosh() {
    console.println(math.cosh(1.5));
}
```

## DegreesToRadians
```
function degreesToRadians(deg float) float
```
Retourne les degrès converti en radians

### Example :
```ecla
import "console";
import "math";

function testDegreesToRadians() {
    console.println(math.degreesToRadians(180.0));
}
```

## Exp
```
function exp(nb float) float
```
Retourne e^nb (exponentielle)

### Example :
```ecla
import "console";
import "math";

function testExp() {
    console.println(math.exp(1.0));
}
```

## Fact
```
function fact(nb int) int
```
Retourne la factorielle de nb

### Example :
```ecla
import "console";
import "math";

function testFact() {
    console.println(math.fact(5));
}
```

## Floor
```
function floor(nb float) float
```
Retourne le plus grand nombre entier qui est plus petit ou égal à nb

### Example :
```ecla
import "console";
import "math";

function testFloor() {
    console.println(math.floor(1.5));
}
```

## Ln
```
function ln(nb float) float
```
Retourne le logarithme naturel de nb


### Example :
```ecla
import "console";
import "math";

function testLn() {
    console.println(math.ln(1.0));
}
```

## Log10
```
function log10(nb float) float
```
Retourne le logarithme en base-10 de nb

### Example :
```ecla
import "console";
import "math";

function testLog10() {
    console.println(math.log10(10.0));
}
```

## Max
```
function max(a float, b float) float
```
Retourne le nombre le plus grand entre a et b

### Example :
```ecla
import "console";
import "math";

function testMax() {
    console.println(math.max(1.0, 2.0));
}
```

## Min
```
function min(a float, b float) float
```
Retourne le nombre le plus petit entre a et b

### Example :
```ecla
import "console";
import "math";

function testMin() {
    console.println(math.min(1.0, 2.0));
}
```

## Modulo
```
function modulo(a float, b float) float
```
Retourne le reste de a / b

### Example :
```ecla
import "console";
import "math";

function testModulo() {
    console.println(math.modulo(5.0, 2.0));
}
```

## Pi
```
function Pi() float
```
Retourne Pi (3.14...)

### Example :
```ecla
import "console";
import "math";

function testPi() {
    console.println(math.pi());
}
```

## Pow
```
function pow(x float, y float) float
```
Retourne x^y

### Example :
```ecla
import "console";
import "math";

function testPow() {
    console.println(math.pow(2.0, 3.0));
}
```

## RadiansToDegrees
```
function radiansToDegrees(rad float) float
```
Retourne les radians converti en degrés

### Example :
```ecla
import "console";
import "math";

function testRadiansToDegrees() {
    console.println(math.radiansToDegrees(3.14));
}
```

## Random
```
function random(min float, max float) float
```
Retourne un nombre aléatoire entre min et max

### Example :
```ecla
import "console";
import "math";

function testRandom() {
    console.println(math.random(0.0, 1.0));
}
```

## Round
```
function round(nb float) float
```
Retourne nb arrondi à l'entier le plus proche

### Example :
```ecla
import "console";
import "math";

function testRound() {
    console.println(math.round(1.5));
}
```

## Sin
```
function sin(opp float, hyp float) float
```
Retourne opposé / hypoténuse

### Example :
```ecla
import "console";
import "math";

function testSin() {
    console.println(math.sin(3.0, 5.0));
}
```

## Sinh
```
function sinh(x float) float
```
Retourne le sinus hyperbolique de x

### Example :
```ecla
import "console";
import "math";

function testSinh() {
    console.println(math.sinh(1.5));
}
```

## Sqrt
```
function sqrt(nb float) float
```
Retourne la racine carré de nb

### Example :
```ecla
import "console";
import "math";

function testSqrt() {
    console.println(math.sqrt(4.0));
}
```

## Tan
```
function tan(opp float, adj float) float
```
Retourne opposé / adjacent

### Example :
```ecla
import "console";
import "math";

function testTan() {
    console.println(math.tan(3.0, 5.0));
}
```

## Tanh
```
function tanh(x float) float
```
Retourne la tangente hyperbolique de x

### Example :
```ecla
import "console";
import "math";

function testTanh() {
    console.println(math.tanh(1.5));
}
```

## Trunc
```
function trunc(nb float) float
```
Retourne la valeur entière de nb

### Example :
```ecla
import "console";
import "math";

function testTrunc() {
    console.println(math.trunc(1.5));
}
```