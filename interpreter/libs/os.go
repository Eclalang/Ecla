package libs

import (
	"errors"
	"fmt"
	"reflect"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/interpreter/libs/utils"
	"github.com/Eclalang/os"
)

type Os struct {
	functionMap map[string]interface{}
}

func NewOs() *Os {
	return &Os{
		functionMap: map[string]interface{}{
			"chown":          nil,
			"clearEnv":       nil,
			"create":         nil,
			"getegid":        nil,
			"getEnv":         nil,
			"geteuid":        nil,
			"getgid":         nil,
			"getHostname":    nil,
			"getpid":         nil,
			"getppid":        nil,
			"getuid":         nil,
			"getUserHomeDir": nil,
			"getwd":          nil,
			"mkdir":          nil,
			"readDir":        nil,
			"readFile":       nil,
			"remove":         nil,
			"removeAll":      nil,
			"setEnv":         nil,
			"setEnvByFile":   nil,
			"unsetEnv":       nil,
			"writeFile":      nil,
		},
	}
}

func (o *Os) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := o.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package os", name))
	}
	switch name {
	case "chown":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.Int && reflect.TypeOf(newArgs[2]).Kind() == reflect.Int {
			os.Chown(newArgs[0].(string), newArgs[1].(int), newArgs[2].(int))
		}
	case "clearEnv":
		os.ClearEnv()
	case "create":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.Create(newArgs[0].(string))
		}
	case "getegid":
		return []eclaType.Type{utils.GoToEclaType(os.Getegid())}, nil
	case "getEnv":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(os.GetEnv(newArgs[0].(string)))}, nil
		}
	case "geteuid":
		return []eclaType.Type{utils.GoToEclaType(os.Geteuid())}, nil
	case "getgid":
		return []eclaType.Type{utils.GoToEclaType(os.Getgid())}, nil
	case "getHostname":
		return []eclaType.Type{utils.GoToEclaType(os.GetHostname())}, nil
	case "getpid":
		return []eclaType.Type{utils.GoToEclaType(os.Getpid())}, nil
	case "getppid":
		return []eclaType.Type{utils.GoToEclaType(os.Getppid())}, nil
	case "getuid":
		return []eclaType.Type{utils.GoToEclaType(os.Getuid())}, nil
	case "getUserHomeDir":
		return []eclaType.Type{utils.GoToEclaType(os.GetUserHomeDir())}, nil
	case "getwd":
		return []eclaType.Type{utils.GoToEclaType(os.Getwd())}, nil
	case "mkdir":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.Mkdir(newArgs[0].(string))
		}
	case "readDir":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(os.ReadDir(newArgs[0].(string)))}, nil
		}
	case "readFile":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(os.ReadFile(newArgs[0].(string)))}, nil
		}
	case "remove":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.Remove(newArgs[0].(string))
		}
	case "removeAll":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.RemoveAll(newArgs[0].(string))
		}
	case "setEnv":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			os.SetEnv(newArgs[0].(string), newArgs[1].(string))
		}
	case "setEnvByFile":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.SetEnvByFile(newArgs[0].(string))
		}
	case "unsetEnv":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.UnsetEnv(newArgs[0].(string))
		}
	case "writeFile":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String && reflect.TypeOf(newArgs[1]).Kind() == reflect.String {
			os.WriteFile(newArgs[0].(string), newArgs[1].(string))
		}
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package os", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}
