package utils

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
)

func EclaTypeToGo(arg eclaType.Type) any {
	switch arg.(type) {
	case eclaType.Int:
		return int(arg.(eclaType.Int))
	case eclaType.Float:
		return float64(arg.(eclaType.Float))
	case eclaType.String:
		return string(arg.(eclaType.String))
	case eclaType.Bool:
		return bool(arg.(eclaType.Bool))
	case *eclaType.List:
		var types []any
		for _, val := range arg.(*eclaType.List).Value {
			types = append(types, EclaTypeToGo(val))
		}
		return types
	default:
		return nil
	}
}
