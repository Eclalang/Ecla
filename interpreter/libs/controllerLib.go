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
	json_         *Json         // Json is the json lib.
	os_           *Os           // Os is the os lib.
	hash_         *Hash         // Hash is the hash lib.
	regex_        *Regex        // Regex is the regex lib.
	math_         *Math         // Math is the math lib.
	strings_      *Strings      // Strings is the strings lib.
)

// InitLibs initializes the libs.
func init() {
	console_ = NewConsole()
	debugKingdom_ = NewDebugKingdom()
	encoding_ = NewEncoding()
	json_ = NewJson()
	os_ = NewOs()
	hash_ = NewHash()
	regex_ = NewRegex()
	math_ = NewMath()
	strings_ = NewStrings()
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
	case "json":
		return json_
	case "os":
		return os_
	case "hash":
		return hash_
	case "regex":
		return regex_
	case "math":
		return math_
	case "strings":
		return strings_
	default:
		return nil
	}
}
