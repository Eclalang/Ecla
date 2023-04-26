package libs

import (
	"errors"
	"fmt"
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

func (c *Console) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := c.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package console", name))
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
		return []eclaType.Type{utils.GoToEclaType(console.Input(newArgs...))}, nil
	case "printInColor":
		console.PrintInColor(newArgs[0].(string), newArgs[1:]...)
		// To add later
		//case "printlnInColor":
		//	console.PrintlnInColor(newArgs[0].(string), newArgs[1:]...)
	case "inputInt":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(console.InputInt(newArgs[0].(string)))}, nil
		}
	case "inputFloat":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(console.InputFloat(newArgs[0].(string)))}, nil
		}
	case "confirm":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return []eclaType.Type{utils.GoToEclaType(console.Confirm(newArgs[0].(string)))}, nil
		}
	case "progressBar":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Int {
			console.ProgressBar(newArgs[0].(int))
		}
	case "clear":
		console.Clear()
	default:
		return nil, errors.New(fmt.Sprintf("Method %s not found in package console", name))
	}

	return []eclaType.Type{eclaType.Null{}}, nil
}
