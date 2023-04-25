package libs

import (
	"github.com/Eclalang/json"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
)

type Json struct {
	functionMap map[string]interface{}
}

func NewJson() *Json {
	return &Json{
		functionMap: map[string]interface{}{
			"marshal":   nil,
			"unmarshal": nil,
		},
	}
}

func (j *Json) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := j.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "marshal":
		return utils.GoToEclaType(json.Marshal(newArgs[0]))
	case "unmarshal":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(json.Unmarshal(newArgs[0].(string)))
		}
		// TODO : Error
	default:
		return nil
	}

	return eclaType.Null{}
}
