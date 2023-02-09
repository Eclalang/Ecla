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
		//TODO: Generate the type of the list
		return &eclaType.List{Value: types}
	case map[any]any:
		var keys []eclaType.Type
		var values []eclaType.Type
		for key, val := range arg.(map[any]any) {
			keys = append(keys, GoToEclaType(key))
			values = append(values, GoToEclaType(val))
		}
		//TODO: Generate the type of the keys, values and full types
		return &eclaType.Map{Keys: keys, Values: values}
	default:
		return eclaType.Null{}
	}
}
