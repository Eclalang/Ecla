package libs

import "github.com/tot0p/Ecla/interpreter/eclaType"

// Lib is the interface of a lib.
type Lib interface {
	Call(name string, args []eclaType.Type) eclaType.Type
}

var (
	console_      *Console      // Console is the console lib.
	debugKingdom_ *DebugKingdom // DebugKingdom is the debugKingdom lib.
	encoding_     *Encoding     // Encoding is the encoding lib.
)

// InitLibs initializes the libs.
func init() {
	console_ = NewConsole()
	debugKingdom_ = NewDebugKingdom()
	encoding_ = NewEncoding()
}

// Import imports the lib with the given name.
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
