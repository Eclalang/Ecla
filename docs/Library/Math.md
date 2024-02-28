# Math

Math library implements mathematical functions.

## Index.

- [Abs(x float) float](#abs)
- [Acos(x float) float](#acos)
- [Acosh(x float) float](#acosh)
- [Asin(x float) float](#asin)
- [Asinh(x float) float](#asinh)
- [Atan(x float) float](#atan)
- [Atanh(x float) float](#atanh)
- [Cbrt(x float) float](#cbrt)
- [Ceil(x float) float](#ceil)
- [Cos(x float) float](#cos)
- [Cosh(x float) float](#cosh)
- [DegreesToRadians(x float) float](#degreestoradians)
- [Exp(x float) float](#exp)
- [Fact(x float) float](#fact)
- [Floor(x float) float](#floor)
- [Ln(x float) float](#ln)
- [Log10(x float) float](#log10)
- [Max(x float, y float) float](#max)
- [Min(x float, y float) float](#min)
- [Modulo(x float, y float) float](#modulo)
- [Pi() float](#pi)
- [Pow(x int | float, y int | float) float](#pow)
- [RadiansToDegrees(x float) float](#radianstodegrees)
- [Random(x float, y float) float](#random)
- [Round(x float) float](#round)
- [Sin(x float) float](#sin)
- [Sinh(x float) float](#sinh)
- [Sqrt(x float) float](#sqrt)
- [Tan(x float) float](#tan)
- [Tanh(x float) float](#tanh)
- [Trunc(x float) float](#trunc)

## Abs
```
function abs(x float) float
```
Returns the absolute value of x

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
function cbrt(x float) float
```
Returns the cubic root of x

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
function ceil(x float) float
```
Returns the littlest integer that is greater or equal to x

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
function Cos(x float) float
```
Returns adjacent / hypotenuse

### Example :
```ecla
import "console";
import "math";

function testCos() {
    console.println(math.cos(3.0));
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
function degreesToRadians(x float) float
```
Returns x converted to radians

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
function exp(x float) float
```
Returns e^x (exponential)

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
function fact(x float) float
```
Returns the factorial number of x

### Example :
```ecla
import "console";
import "math";

function testFact() {
    console.println(math.fact(5.0));
}
```

## Floor
```
function floor(x float) float
```
Returns the greatest integer which is smaller or equal to x

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
function ln(x float) float
```
Returns the natural logarithm of x

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
function log10(x float) float
```
Returns base-10 logarithm of x

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
function max(x float, y float) float
```
Returns the larger number between x and y

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
function min(x float, y float) float
```
Returns the smaller number between x and y

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
function modulo(x float, y float) float
```
Returns remainder of x / y

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
function pow(x int | float, y int | float) float
```
Returns x^y

### Example :
```ecla
import "console";
import "math";

function testPow() {
    console.println(math.pow(2.0, 3));
}
```

## RadiansToDegrees
```
function radiansToDegrees(x float) float
```
Returns x converted to degrees

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
function random(x float, y float) float
```
Returns random number between x and y

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
function round(x float) float
```
Returns x rounded to the nearest int

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
function sin(x float) float
```
Returns opposite / hypotenuse

### Example :
```ecla
import "console";
import "math";

function testSin() {
    console.println(math.sin(3.0));
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
function sqrt(x float) float
```
Returns the square root of x

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
function tan(x float) float
```
Returns opposite / adjacent

### Example :
```ecla
import "console";
import "math";

function testTan() {
    console.println(math.tan(3.0));
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
function trunc(x float) float
```
Returns the integer value of x

### Example :
```ecla
import "console";
import "math";

function testTrunc() {
    console.println(math.trunc(1.5));
}
```