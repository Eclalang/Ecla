package utils

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

// EclaTypeToGo converts an eclaType to a go type.
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
	case *eclaType.Map:
		var types = make(map[any]any)
		for i := 0; i < len(arg.(*eclaType.Map).Keys); i++ {
			types[EclaTypeToGo(arg.(*eclaType.Map).Keys[i])] = EclaTypeToGo(arg.(*eclaType.Map).Values[i])
		}
		return types
	default:
		return nil
	}
}
