package libs

import (
	"github.com/Eclalang/os"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
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

func (o *Os) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := o.functionMap[name]; !ok {
		return nil
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
		return utils.GoToEclaType(os.Getegid())
	case "getEnv":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(os.GetEnv(newArgs[0].(string)))
		}
	case "geteuid":
		return utils.GoToEclaType(os.Geteuid())
	case "getgid":
		return utils.GoToEclaType(os.Getgid())
	case "getHostname":
		return utils.GoToEclaType(os.GetHostname())
	case "getpid":
		return utils.GoToEclaType(os.Getpid())
	case "getppid":
		return utils.GoToEclaType(os.Getppid())
	case "getuid":
		return utils.GoToEclaType(os.Getuid())
	case "getUserHomeDir":
		return utils.GoToEclaType(os.GetUserHomeDir())
	case "getwd":
		return utils.GoToEclaType(os.Getwd())
	case "mkdir":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			os.Mkdir(newArgs[0].(string))
		}
	case "readDir":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(os.ReadDir(newArgs[0].(string)))
		}
	case "readFile":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(os.ReadFile(newArgs[0].(string)))
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
		return nil
	}

	return eclaType.Null{}
}
