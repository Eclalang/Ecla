package libs

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"math/rand"
)

type DebugKingdom struct {
	functionMap map[string]interface{}
}

func NewDebugKingdom() *DebugKingdom {
	return &DebugKingdom{
		functionMap: map[string]interface{}{
			"testMap":     nil,
			"testMapRet":  nil,
			"testList":    nil,
			"testListRet": nil,
			"randInt":     nil,
			"clear":       nil,
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
	case "testMapRet":
		tempMap, ok := newArgs[0].(map[interface{}]interface{})
		if !ok {
			return nil
		}
		newMap := make(map[string]string)
		for k, v := range tempMap {
			newMap[k.(string)] = v.(string)
		}
		return utils.GoToEclaType(d.TestMapRet(newMap))
	case "testList":
		tempList, ok := newArgs[0].([]interface{})
		if !ok {
			return nil
		}
		newList := make([]string, len(tempList))
		for k, v := range tempList {
			newList[k] = v.(string)
		}
		d.TestList(newList)
	case "testListRet":
		tempList, ok := newArgs[0].([]interface{})
		if !ok {
			return nil
		}
		newList := make([]string, len(tempList))
		for k, v := range tempList {
			newList[k] = v.(string)
		}
		return utils.GoToEclaType(d.TestListRet(newList))
	case "randInt":
		min, ok := newArgs[0].(int)
		if !ok {
			return eclaType.Null{}
		}
		max, ok := newArgs[1].(int)
		if !ok {
			return eclaType.Null{}
		}
		return utils.GoToEclaType(d.RandInt(min, max))
	case "clear":
		fmt.Print("\033[H\033[2J")
	}
	return eclaType.Null{}
}

func (d *DebugKingdom) RandInt(min int, max int) int {
	// generate a random number between min and max
	return rand.Intn(max-min) + min
}

func (d *DebugKingdom) TestMap(test map[string]string) {
	for k, v := range test {
		fmt.Println(k, v)
	}
}

func (d *DebugKingdom) TestMapRet(test map[string]string) map[string]string {
	return test
}

func (d *DebugKingdom) TestList(test []string) {
	for _, v := range test {
		fmt.Println(v)
	}
}

func (d *DebugKingdom) TestListRet(test []string) []string {
	return test
}
