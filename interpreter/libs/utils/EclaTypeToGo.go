package utils

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"reflect"
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
	case eclaType.Char:
		return rune(arg.(eclaType.Char))
	case *eclaType.List:
		// TODO : base the type of the array on the type of the ecla list using eclaType.List.GetType()
		arrType := reflect.SliceOf(reflect.TypeOf(EclaTypeToGo(arg.(*eclaType.List).Value[0])))
		arr := reflect.MakeSlice(arrType, len(arg.(*eclaType.List).Value), len(arg.(*eclaType.List).Value))
		for i, val := range arg.(*eclaType.List).Value {
			arr.Index(i).Set(reflect.ValueOf(EclaTypeToGo(val)))
		}
		return arr.Interface()
	case *eclaType.Map:
		var types = make(map[any]any)
		for i := 0; i < len(arg.(*eclaType.Map).Keys); i++ {
			types[EclaTypeToGo(arg.(*eclaType.Map).Keys[i])] = EclaTypeToGo(arg.(*eclaType.Map).Values[i])
		}
		mapType := reflect.MapOf(reflect.TypeOf(types), reflect.TypeOf(types))
		mapVal := reflect.MakeMap(mapType)
		for k, v := range types {
			mapVal.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(v))
		}
		return mapVal.Interface()
	default:
		return nil
	}
}
