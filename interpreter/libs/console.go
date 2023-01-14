package libs

import (
	"github.com/Eclalang/console"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
)

type Console struct {
}

func (c *Console) Call(name string, args []eclaType.Type) eclaType.Type {
	switch name {
	case "print":
		newargs := make([]any, len(args))
		for k, arg := range args {
			newargs[k] = utils.EclaTypeToGo(arg)
		}
		console.Print(newargs...)
	}
	return eclaType.Empty{}
}
