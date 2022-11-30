package parser

var (
	Keywords = map[string]interface{}{
		"var":      nil,
		"function": nil,
		"print":    nil,
		"import":   nil,
	}
	VarTypes = map[string]interface{}{
		"int":    nil,
		"float":  nil,
		"string": nil,
		"bool":   nil,
	}
)
