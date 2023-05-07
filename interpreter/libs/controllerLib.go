package libs

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
)

// Lib is the interface of a lib.
type Lib interface {
	Call(name string, args []eclaType.Type) ([]eclaType.Type, error)
}

var (
	allLibs map[string]func() Lib
)

// InitLibs initializes the libs.
func init() {

	allLibs = map[string]func() Lib{
		"console": func() Lib {
			return NewConsole()
		},
		"debugKingdom": func() Lib {
			return NewDebugKingdom()
		},
		"encoding": func() Lib {
			return NewEncoding()
		},
		"json": func() Lib {
			return NewJson()
		},
		"os": func() Lib {
			return NewOs()
		},
		"hash": func() Lib {
			return NewHash()
		},
		"regex": func() Lib {
			return NewRegex()
		},
		"math": func() Lib {
			return NewMath()
		},
		"strings": func() Lib {
			return NewStrings()
		},
		"cast": func() Lib {
			return NewCast()
		},
		"time": func() Lib {
			return NewTime()
		},
	}

}

// Import imports the lib with the given name.
func Import(name string) Lib {
	if lib, ok := allLibs[name]; ok {
		return lib()
	}
	return nil
}
