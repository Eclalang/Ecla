package libs

import "github.com/tot0p/Ecla/interpreter/eclaType"

type Lib interface {
	Call(name string, args []eclaType.Type) eclaType.Type
}

var (
	console_      *Console
	debugKingdom_ *DebugKingdom
)

func init() {
	console_ = NewConsole()
	debugKingdom_ = NewDebugKingdom()
}

func Import(name string) Lib {
	switch name {
	case "console":
		return console_
	case "debugKingdom":
		return debugKingdom_
	default:
		return nil
	}
}
