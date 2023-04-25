package libs

import (
	"github.com/Eclalang/hash"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/interpreter/libs/utils"
	"reflect"
)

type Hash struct {
	functionMap map[string]interface{}
}

func NewHash() *Hash {
	return &Hash{
		functionMap: map[string]interface{}{
			"hashmd5":        nil,
			"hashsha1":       nil,
			"hashsha224":     nil,
			"hashsha256":     nil,
			"hashsha384":     nil,
			"hashsha512":     nil,
			"hashsha512_224": nil,
			"hashsha512_256": nil,
		},
	}
}

func (h *Hash) Call(name string, args []eclaType.Type) eclaType.Type {
	newArgs := make([]any, len(args))
	for k, arg := range args {
		newArgs[k] = utils.EclaTypeToGo(arg)
	}
	if _, ok := h.functionMap[name]; !ok {
		return nil
	}
	switch name {
	case "hashmd5":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashmd5(newArgs[0].(string)))
		}
	case "hashsha1":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha1(newArgs[0].(string)))
		}
	case "hashsha224":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha224(newArgs[0].(string)))
		}
	case "hashsha256":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha256(newArgs[0].(string)))
		}
	case "hashsha384":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha384(newArgs[0].(string)))
		}
	case "hashsha512":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha512(newArgs[0].(string)))
		}
	case "hashsha512_224":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha512_224(newArgs[0].(string)))
		}
	case "hashsha512_256":
		if reflect.TypeOf(newArgs[0]).Kind() == reflect.String {
			return utils.GoToEclaType(hash.Hashsha512_256(newArgs[0].(string)))
		}
	default:
		return nil
	}

	return eclaType.Null{}
}
