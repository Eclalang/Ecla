package libs

import (
	"errors"
	"fmt"
	"github.com/Eclalang/math"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
)

type Math struct {
	functionMap map[string]interface{}
}

func NewMath() *Math {
	return &Math{
		functionMap: map[string]interface{}{
			"pi":               nil,
			"cos":              nil,
			"sin":              nil,
			"tan":              nil,
			"ln":               nil,
			"exp":              nil,
			"sqrt":             nil,
			"cbrt":             nil,
			"pow":              nil,
			"fact":             nil,
			"abs":              nil,
			"floor":            nil,
			"ceil":             nil,
			"trunc":            nil,
			"max":              nil,
			"min":              nil,
			"log10":            nil,
			"round":            nil,
			"degreesToRadians": nil,
			"radiansToDegrees": nil,
			"modulo":           nil,
			"random":           nil,
			"acos":             nil,
			"asin":             nil,
			"atan":             nil,
			"cosh":             nil,
			"sinh":             nil,
			"tanh":             nil,
			"acosh":            nil,
			"asinh":            nil,
			"atanh":            nil,
		},
	}
}

func (m *Math) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := m.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package math", name))
	}
	switch name {
	case "pi":
		return []eclaType.Type{utils.GoToEclaType(math.Pi())}, nil
	case "cos":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Cos(newArgs[0].(float64)))}, nil
		}
	case "sin":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Sin(newArgs[0].(float64)))}, nil
		}
	case "tan":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Tan(newArgs[0].(float64)))}, nil
		}
	case "ln":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Ln(newArgs[0].(float64)))}, nil
		}
	case "exp":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Exp(newArgs[0].(float64)))}, nil
		}
	case "sqrt":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Sqrt(newArgs[0].(float64)))}, nil
		}
	case "cbrt":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Cbrt(newArgs[0].(float64)))}, nil
		}
	case "pow":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 && reflect.TypeOf(newArgs[1]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Pow(newArgs[0].(float64), newArgs[1].(float64)))}, nil
		}
	case "fact":
		// TODO: Fix this
	case "abs":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Abs(newArgs[0].(float64)))}, nil
		}
	case "floor":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Floor(newArgs[0].(float64)))}, nil
		}
	case "ceil":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Ceil(newArgs[0].(float64)))}, nil
		}
	case "trunc":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Trunc(newArgs[0].(float64)))}, nil
		}
	case "max":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 && reflect.TypeOf(newArgs[1]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Max(newArgs[0].(float64), newArgs[1].(float64)))}, nil
		}
	case "min":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 && reflect.TypeOf(newArgs[1]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Min(newArgs[0].(float64), newArgs[1].(float64)))}, nil
		}
	case "log10":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Log10(newArgs[0].(float64)))}, nil
		}
	case "round":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Round(newArgs[0].(float64)))}, nil
		}
	case "degreesToRadians":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.DegreesToRadians(newArgs[0].(float64)))}, nil
		}
	case "radiansToDegrees":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.RadiansToDegrees(newArgs[0].(float64)))}, nil
		}
	case "modulo":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 && reflect.TypeOf(newArgs[1]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Modulo(newArgs[0].(float64), newArgs[1].(float64)))}, nil
		}
	case "random":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 && reflect.TypeOf(newArgs[1]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Random(newArgs[0].(float64), newArgs[1].(float64)))}, nil
		}
	case "acos":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Acos(newArgs[0].(float64)))}, nil
		}
	case "asin":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Asin(newArgs[0].(float64)))}, nil
		}
	case "atan":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Atan(newArgs[0].(float64)))}, nil
		}
	case "cosh":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Cosh(newArgs[0].(float64)))}, nil
		}
	case "sinh":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Sinh(newArgs[0].(float64)))}, nil
		}
	case "tanh":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Tanh(newArgs[0].(float64)))}, nil
		}
	case "acosh":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Acosh(newArgs[0].(float64)))}, nil
		}
	case "asinh":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Asinh(newArgs[0].(float64)))}, nil
		}
	case "atanh":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(math.Atanh(newArgs[0].(float64)))}, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package math", name))
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}
