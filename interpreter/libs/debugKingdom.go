package libs

import (
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
			// TODO use the functionMap for the switch statement in Call
			"testMap": (*DebugKingdom).Test,
		},
	}
}

func (d *DebugKingdom) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := d.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "testMap":
		newMap, ok := newArgs[0].(map[string]string)
		if ok {
			d.Test(newMap)
		}
	}
	return eclaType.Null{}
}

func (d *DebugKingdom) Test(map[string]string) {
	for k, v := range d.functionMap {
		fmt.Println(k, v)
	}
}
