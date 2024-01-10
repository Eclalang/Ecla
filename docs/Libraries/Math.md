# Math

Math library implements mathematical functions.

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
Returns the absolute value of nb

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
Returns arc cosine of x

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
Returns arc hyperbolic cosine of x

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
Returns arc sine of x

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
Returns arc hyperbolic sine of x

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
Returns arc hyperbolic tangent of x

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
Returns the cubic root of nb

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
Returns the least integer that it's bigger or equal to x

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
Returns adjacent / hypotenuse

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
Returns hyperbolic cosine of x

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
Returns deg converted to radians

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
Returns e^nb (exponential)

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
Returns the factorial number of nb

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
Returns the greatest integer which is smaller or equal to nb

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
Returns the naturel logarithm of nb

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
Returns base-10 logarithm of nb

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
Returns the larger number between a and b

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
Returns the smaller number between a and b

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
Returns remainder of a / b

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
Returns Pi (3.14...)

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
Returns x^y

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
Returns rad converted to degrees

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
Returns random number between min and max

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
Returns nb rounded to the nearest int

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
Returns opposite / hypotenuse

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
Returns hyperbolic sine of x

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
Returns the square root of nb

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
Returns opposite / adjacent

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
Returns hyperbolic tangent of x

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
Returns the integer value of nb

### Example :
```ecla
import "console";
import "math";

function testTrunc() {
    console.println(math.trunc(1.5));
}
```