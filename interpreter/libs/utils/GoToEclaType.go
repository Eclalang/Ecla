package utils

import "github.com/tot0p/Ecla/interpreter/eclaType"

func GoToEclaType(arg any) eclaType.Type {
	switch arg.(type) {
	case int:
		return eclaType.Int(arg.(int))
	case float64:
		return eclaType.Float(arg.(float32))
	case string:
		return eclaType.String(arg.(string))
	case bool:
		return eclaType.Bool(arg.(bool))
	case []any:
		var types []eclaType.Type
		for _, val := range arg.([]any) {
			types = append(types, GoToEclaType(val))
		}
		return &eclaType.List{Value: types}
	default:
		return eclaType.Null{}
	}
}
