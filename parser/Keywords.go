package parser

const (
	Int    = "int"
	String = "string"
	Float  = "float"
	Bool   = "bool"
)

var (
	Keywords = map[string]interface{}{
		"var":      nil,
		"function": nil,
		"print":    nil,
		"import":   nil,
		"type":     nil,
		"for":      nil,
		"if":       nil,
	}
	VarTypes = map[string]interface{}{
		Int:    nil,
		Float:  nil,
		String: nil,
		Bool:   nil,
	}
)
