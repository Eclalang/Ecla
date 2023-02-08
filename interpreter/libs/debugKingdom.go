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
			"testMap": nil,
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
		tempMap, ok := newArgs[0].(map[interface{}]interface{})
		if !ok {
			return nil
		}
		newMap := make(map[string]string)
		for k, v := range tempMap {
			newMap[k.(string)] = v.(string)
		}
		d.TestMap(newMap)
	}
	return eclaType.Null{}
}

func (d *DebugKingdom) TestMap(test map[string]string) {
	for k, v := range test {
		fmt.Println(k, v)
	}
}
