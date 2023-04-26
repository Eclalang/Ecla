package libs

import (
	enc "github.com/Eclalang/encoding"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
)

type Encoding struct {
	functionMap map[string]interface{}
}

func NewEncoding() *Encoding {
	return &Encoding{
		functionMap: map[string]interface{}{
			"asciiToString": nil,
			"decodeBase64":  nil,
			"decodeGob":     nil,
			"decodeHex":     nil,
			"encodeBase64":  nil,
			"encodeGob":     nil,
			"encodeHex":     nil,
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
	case "decodeBase64":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(enc.DecodeBase64(newArgs[0].(string)))
		}
	case "decodeGob":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			return utils.GoToEclaType(enc.DecodeGob(newArgs[0].([]int)))
		}
	case "decodeHex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(enc.DecodeHex(newArgs[0].(string)))
		}
	case "encodeBase64":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			return utils.GoToEclaType(enc.EncodeBase64(newArgs[0].([]int)))
		}
	case "encodeGob":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(enc.EncodeGob(newArgs[0].(string)))
		}
	case "encodeHex":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.Slice {
			return utils.GoToEclaType(enc.EncodeHex(newArgs[0].([]int)))
		}
	case "stringToAscii":
		return utils.GoToEclaType(enc.StringToAscii(newArgs[0].(string)))
	default:
		return nil
	}
	return eclaType.Null{}
}
