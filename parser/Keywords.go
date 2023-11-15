package parser

const (
	Int        = "int"
	String     = "string"
	Char       = "char"
	Float      = "float"
	Bool       = "bool"
	Map        = "map"
	ArrayStart = "["
)

var (
	Keywords = map[string]interface{}{
		"var":      nil,
		"function": nil,
		"return":   nil,
		"print":    nil,
		"range":    nil,
		"import":   nil,
		"type":     nil,
		"for":      nil,
		"while":    nil,
		"if":       nil,
		"else":     nil,
		"null":     nil,
	}
	VarTypes = map[string]interface{}{
		Int:        nil,
		Float:      nil,
		String:     nil,
		Char:       nil,
		Bool:       nil,
		Map:        nil,
		ArrayStart: nil,
	}
)
