package libs

import (
	"github.com/Eclalang/cast"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
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

func (c *Cast) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := c.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "atoi":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(cast.Atoi(newArgs[0].(string)))
		}
	case "floatToInt":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Float64 {
			return utils.GoToEclaType(cast.FloatToInt(newArgs[0].(float64)))
		}
	case "intToFloat":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Int {
			return utils.GoToEclaType(cast.IntToFloat(newArgs[0].(int)))
		}
	case "parseBool":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(cast.ParseBool(newArgs[0].(string)))
		}
	case "parseFloat":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.Int {
			return utils.GoToEclaType(cast.ParseFloat(newArgs[0].(string), newArgs[1].(int)))
		}
	default:
		return nil
	}

	return eclaType.Null{}
}
