package parser

const (
	// varible type names

	Int        = "int"
	String     = "string"
	Char       = "char"
	Float      = "float"
	Bool       = "bool"
	Map        = "map"
	Function   = "function"
	ArrayStart = "["
	Any        = "any"

	// keywords
	Var    = "var"
	Return = "return"
	Range  = "range"
	Import = "import"
	For    = "for"
	While  = "while"
	If     = "if"
	Else   = "else"
	Null   = "null"
	Struct = "struct"
	Murloc = "mgrlmgrl"

	// built-in functions
	TypeOf = "typeOf"
	Eval   = "eval"
	Len    = "len"
	SizeOf = "sizeOf"
	Append = "append"
)

var (
	Keywords = map[string]interface{}{
		Var:      nil,
		Function: nil,
		Return:   nil,
		Range:    nil,
		Import:   nil,
		For:      nil,
		While:    nil,
		If:       nil,
		Else:     nil,
		Null:     nil,
		Any:      nil,
		Struct:   nil,
		Murloc:   nil,
		TypeOf:   nil,
		Eval:     nil,
		Len:      nil,
		SizeOf:   nil,
		Append:   nil,
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
		Any:        nil,
	}
	DefaultVarTypes = map[string]interface{}{
		Int:        nil,
		Float:      nil,
		String:     nil,
		Char:       nil,
		Bool:       nil,
		Map:        nil,
		Function:   nil,
		ArrayStart: nil,
		Any:        nil,
	}
)
