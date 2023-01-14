package utils

import "github.com/tot0p/Ecla/interpreter/eclaType"

func EclaTypeToGo(arg eclaType.Type) any {
	switch arg.(type) {
	case eclaType.Int:
		return arg.(eclaType.Int).GetValue()
	case eclaType.Float:
		return arg.(eclaType.Float).GetValue()
	case eclaType.String:
		return arg.(eclaType.String).GetValue()
	case eclaType.Bool:
		return arg.(eclaType.Bool).GetValue()
	case *eclaType.List:
		types := make([]any, len(arg.(*eclaType.List).Value))
		for _, val := range arg.(*eclaType.List).Value {
			types = append(types, EclaTypeToGo(val))
		}
		return types
	default:
		return nil
	}
}
