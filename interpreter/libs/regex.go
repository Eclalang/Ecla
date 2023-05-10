package libs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/interpreter/libs/utils"
	"github.com/Eclalang/regex"
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

func (r *Regex) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := r.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package regex", name))
	}
	switch name {
	case "find":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(regex.Find(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "findAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(regex.FindAll(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "findAllIndex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(regex.FindAllIndex(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "findIndex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(regex.FindIndex(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "match":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(regex.Match(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "replaceAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(regex.ReplaceAll(newArgs[0].(string), newArgs[1].(string), newArgs[2].(string)))}, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package regex", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}
