package libs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/interpreter/libs/utils"
	"github.com/Eclalang/cast"
)

type Cast struct {
	functionMap map[string]interface{}
}

func NewCast() *Cast {
	return &Cast{
		functionMap: map[string]interface{}{
			"atoi":       nil,
			"floatToInt": nil,
			"intToFloat": nil,
			"parseBool":  nil,
			"parseFloat": nil,
		},
	}
}

func (c *Cast) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := c.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package cast", name))
	}
	switch name {
	case "atoi":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(cast.Atoi(newArgs[0].(string)))}, nil
		}
	case "floatToInt":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return []eclaType.Type{utils.GoToEclaType(cast.FloatToInt(newArgs[0].(float64)))}, nil
		}
	case "intToFloat":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Int {
			return []eclaType.Type{utils.GoToEclaType(cast.IntToFloat(newArgs[0].(int)))}, nil
		}
	case "parseBool":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(cast.ParseBool(newArgs[0].(string)))}, nil
		}
	case "parseFloat":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.Int {
			return []eclaType.Type{utils.GoToEclaType(cast.ParseFloat(newArgs[0].(string), newArgs[1].(int)))}, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package cast", name))
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}
