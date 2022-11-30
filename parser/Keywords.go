package parser

var (
	Keywords = map[string]interface{}{
		"var":      nil,
		"function": nil,
		"print":    nil,
		"import":   nil,
	}
	Int      = "int"
	Float    = "float"
	String   = "string"
	Bool     = "bool"
	VarTypes = map[string]interface{}{
		Int:    nil,
		Float:  nil,
		String: nil,
		Bool:   nil,
	}
)
