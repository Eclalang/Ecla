package libs

import (
	"github.com/Eclalang/console"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
)

type Console struct {
	functionMap map[string]interface{}
}

func NewConsole() *Console {
	return &Console{
		functionMap: map[string]interface{}{
			"printf":       nil,
			"println":      nil,
			"print":        nil,
			"input":        nil,
			"printInColor": nil,
			"inputInt":     nil,
			"inputFloat":   nil,
			"confirm":      nil,
			"progressBar":  nil,
			"clear":        nil,
		},
	}
}

func (c *Console) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := c.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "printf":
		// TODO: refactor this line
		console.Printf(newArgs[0].(string), newArgs[1:]...)
	case "println":
		console.Println(newArgs...)
	case "print":
		console.Print(newArgs...)
	case "input":
		return utils.GoToEclaType(console.Input(newArgs...))
	case "printInColor":
		console.PrintInColor(newArgs[0].(string), newArgs[1:]...)
		// To add later
		//case "printlnInColor":
		//	console.PrintlnInColor(newArgs[0].(string), newArgs[1:]...)
	case "inputInt":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(console.InputInt(newArgs[0].(string)))
		}
	case "inputFloat":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(console.InputFloat(newArgs[0].(string)))
		}
	case "confirm":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(console.Confirm(newArgs[0].(string)))
		}
	case "progressBar":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Int {
			console.ProgressBar(newArgs[0].(int))
		}
	case "clear":
		console.Clear()
	default:
		return nil
	}

	return eclaType.Null{}
}
