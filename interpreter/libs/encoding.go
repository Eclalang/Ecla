package libs

import (
	enc "github.com/Eclalang/encoding"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
)

type Encoding struct {
	functionMap map[string]interface{}
}

func NewEncoding() *Encoding {
	return &Encoding{
		functionMap: map[string]interface{}{
			"asciiToString": nil,
			"stringToAscii": nil,
		},
	}
}

func (e *Encoding) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := e.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "asciiToString":
		arg1 := newArgs[0].([]interface{})
		var arg2 []int
		for _, val := range arg1 {
			arg2 = append(arg2, val.(int))
		}
		return utils.GoToEclaType(enc.AsciiToString(arg2))
	case "stringToAscii":
		return utils.GoToEclaType(enc.StringToAscii(newArgs[0].(string)))
	}
	return eclaType.Null{}
}
