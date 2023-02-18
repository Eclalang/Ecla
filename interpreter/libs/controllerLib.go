package libs

import "github.com/tot0p/Ecla/interpreter/eclaType"

type Lib interface {
	Call(name string, args []eclaType.Type) eclaType.Type
}

var (
	console_      *Console
	debugKingdom_ *DebugKingdom
	encoding_     *Encoding
)

func init() {
	console_ = NewConsole()
	debugKingdom_ = NewDebugKingdom()
	encoding_ = NewEncoding()
}

func Import(name string) Lib {
	switch name {
	case "console":
		return console_
	case "debugKingdom":
		return debugKingdom_
	case "encoding":
		return encoding_
	default:
		return nil
	}
}
