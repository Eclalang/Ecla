package libs

import (
	"github.com/Eclalang/strings"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
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

func (s *Strings) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := s.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "contains":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.Contains(newArgs[0].(string), newArgs[1].(string)))
		}
	case "containsAny":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.ContainsAny(newArgs[0].(string), newArgs[1].(string)))
		}
	case "count":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.Count(newArgs[0].(string), newArgs[1].(string)))
		}
	case "cut":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.Cut(newArgs[0].(string), newArgs[1].(string)))
		}
	case "hasPrefix":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.HasPrefix(newArgs[0].(string), newArgs[1].(string)))
		}
	case "hasSuffix":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.HasSuffix(newArgs[0].(string), newArgs[1].(string)))
		}
	case "indexOf":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.IndexOf(newArgs[0].(string), newArgs[1].(string)))
		}
	case "join":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.Join(newArgs[0].([]string), newArgs[1].(string)))
		}
	case "replace":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.String && reflect.TypeOf(newArgs[3]).Kind() == reflect.Int {
			return utils.GoToEclaType(strings.Replace(newArgs[0].(string), newArgs[1].(string), newArgs[2].(string), newArgs[3].(int)))
		}
	case "replaceAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.ReplaceAll(newArgs[0].(string), newArgs[1].(string), newArgs[2].(string)))
		}
	case "split":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.Split(newArgs[0].(string), newArgs[1].(string)))
		}
	case "splitAfter":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.SplitAfter(newArgs[0].(string), newArgs[1].(string)))
		}
	case "splitAfterN":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int {
			return utils.GoToEclaType(strings.SplitAfterN(newArgs[0].(string), newArgs[1].(string), newArgs[2].(int)))
		}
	case "splitN":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int {
			return utils.GoToEclaType(strings.SplitN(newArgs[0].(string), newArgs[1].(string), newArgs[2].(int)))
		}
	case "toLower":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.ToLower(newArgs[0].(string)))
		}
	case "toUpper":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.ToUpper(newArgs[0].(string)))
		}
	case "trim":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			return utils.GoToEclaType(strings.Trim(newArgs[0].(string), newArgs[1].(string)))
		}
	default:
		return nil
	}

	return eclaType.Null{}
}
