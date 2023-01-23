package libs

import (
	"github.com/Eclalang/console"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
)

type Console struct {
}

func (c *Console) Call(name string, args []eclaType.Type) eclaType.Type {
	newargs := make([]any, len(args))
	for k, arg := range args {
		newargs[k] = utils.EclaTypeToGo(arg)
	}
	switch name {
	case "printf":
		// TODO: refactor this line
		console.Printf(newargs[0].(string), newargs[1:]...)
	case "println":
		console.Println(newargs...)
	case "print":
		console.Print(newargs...)
	case "input":
		return utils.GoToEclaType(console.Input(newargs...))
	case "printInColor":
		console.PrintInColor(newargs[0].(string), newargs[1:]...)
		// To add later
		//case "printlnInColor":
		//	console.PrintlnInColor(newargs[0].(string), newargs[1:]...)
	}

	return eclaType.Empty{}
}
