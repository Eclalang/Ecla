package libs

import (
	"errors"
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
)

type DebugKingdom struct {
	functionMap map[string]interface{}
}

func NewDebugKingdom() *DebugKingdom {
	return &DebugKingdom{
		functionMap: map[string]interface{}{
			"clear": nil,
		},
	}
}

func (d *DebugKingdom) Call(name string, args []eclaType.Type) ([]eclaType.Type, error) {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := d.functionMap[name]; !ok {
		return nil, errors.New(fmt.Sprintf("Method %s not found in package debugKingdom", name))
	}
	switch name {
	case "clear":
		fmt.Print("\033[H\033[2J")
	}
	return []eclaType.Type{eclaType.Null{}}, nil
}
