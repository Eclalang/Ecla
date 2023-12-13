package parser

import "github.com/Eclalang/Ecla/lexer"

const (
	Int        = "int"
	String     = "string"
	Char       = "char"
	Float      = "float"
	Bool       = "bool"
	Map        = "map"
	Function   = "function"
	ArrayStart = "["
)

var (
	Keywords = map[string]interface{}{
		"var":        nil,
		"function":   nil,
		"return":     nil,
		"range":      nil,
		"import":     nil,
		"type":       nil,
		"for":        nil,
		"while":      nil,
		"if":         nil,
		"else":       nil,
		"null":       nil,
		lexer.MURLOC: nil,
	}
	VarTypes = map[string]interface{}{
		Int:        nil,
		Float:      nil,
		String:     nil,
		Char:       nil,
		Bool:       nil,
		Map:        nil,
		Function:   nil,
		ArrayStart: nil,
	}
)
