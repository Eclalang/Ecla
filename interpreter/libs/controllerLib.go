package libs

import "github.com/tot0p/Ecla/interpreter/eclaType"

type Lib interface {
	Call(name string, args []eclaType.Type) eclaType.Type
}

func Import(name string) Lib {
	switch name {
	case "console":
		return new(Console)
	default:
		return nil
	}
}
