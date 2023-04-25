package libs

import (
	"github.com/Eclalang/regex"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
)

type Regex struct {
	functionMap map[string]interface{}
}

func NewRegex() *Regex {
	return &Regex{
		functionMap: map[string]interface{}{
			"find":         nil,
			"findAll":      nil,
			"findAllIndex": nil,
			"findIndex":    nil,
			"match":        nil,
			"replaceAll":   nil,
		},
	}
}

func (r *Regex) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := r.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "find":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(regex.Find(newArgs[0].(string), newArgs[1].(string)))
		}
	case "findAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(regex.FindAll(newArgs[0].(string), newArgs[1].(string)))
		}
	case "findAllIndex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(regex.FindAllIndex(newArgs[0].(string), newArgs[1].(string)))
		}
	case "findIndex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(regex.FindIndex(newArgs[0].(string), newArgs[1].(string)))
		}
	case "match":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(regex.Match(newArgs[0].(string), newArgs[1].(string)))
		}
	case "replaceAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.String {
			return utils.GoToEclaType(regex.ReplaceAll(newArgs[0].(string), newArgs[1].(string), newArgs[2].(string)))
		}
	default:
		return nil
	}

	return eclaType.Null{}
}
