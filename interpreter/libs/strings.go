package libs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/interpreter/libs/utils"
	"github.com/Eclalang/strings"
)

type Strings struct {
	functionMap map[string]interface{}
}

func NewStrings() *Strings {
	return &Strings{
		functionMap: map[string]interface{}{
			"contains":    nil,
			"containsAny": nil,
			"count":       nil,
			"cut":         nil,
			"hasPrefix":   nil,
			"hasSuffix":   nil,
			"indexOf":     nil,
			"join":        nil,
			"replace":     nil,
			"replaceAll":  nil,
			"split":       nil,
			"splitAfter":  nil,
			"splitAfterN": nil,
			"splitN":      nil,
			"toLower":     nil,
			"toUpper":     nil,
			"trim":        nil,
		},
	}
}

func (s *Strings) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := s.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package strings", name))
	}
	switch name {
	case "contains":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.Contains(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "containsAny":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.ContainsAny(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "count":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.Count(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "cut":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.Cut(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "hasPrefix":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.HasPrefix(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "hasSuffix":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.HasSuffix(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "indexOf":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.IndexOf(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "join":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.Join(newArgs[0].([]string), newArgs[1].(string)))}, nil
		}
	case "replace":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.String && reflect.TypeOf(newArgs[3]).Kind() == reflect.Int {
			return []eclaType.Type{utils.GoToEclaType(strings.Replace(newArgs[0].(string), newArgs[1].(string), newArgs[2].(string), newArgs[3].(int)))}, nil
		}
	case "replaceAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.ReplaceAll(newArgs[0].(string), newArgs[1].(string), newArgs[2].(string)))}, nil
		}
	case "split":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.Split(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "splitAfter":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.SplitAfter(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	case "splitAfterN":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int {
			return []eclaType.Type{utils.GoToEclaType(strings.SplitAfterN(newArgs[0].(string), newArgs[1].(string), newArgs[2].(int)))}, nil
		}
	case "splitN":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int {
			return []eclaType.Type{utils.GoToEclaType(strings.SplitN(newArgs[0].(string), newArgs[1].(string), newArgs[2].(int)))}, nil
		}
	case "toLower":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.ToLower(newArgs[0].(string)))}, nil
		}
	case "toUpper":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.ToUpper(newArgs[0].(string)))}, nil
		}
	case "trim":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(strings.Trim(newArgs[0].(string), newArgs[1].(string)))}, nil
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package strings", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}
