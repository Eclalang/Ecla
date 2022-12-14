package parser

const (
	Int        = "int"
	String     = "string"
	Float      = "float"
	Bool       = "bool"
	ArrayStart = "["
)

var (
	Keywords = map[string]interface{}{
		"var":      nil,
		"function": nil,
		"print":    nil,
		"import":   nil,
		"type":     nil,
		"for":      nil,
		"while":    nil,
		"if":       nil,
		"else":     nil,
	}
	VarTypes = map[string]interface{}{
		Int:        nil,
		Float:      nil,
		String:     nil,
		Bool:       nil,
		ArrayStart: nil,
	}
)
