package eclaType

// Type with all methods.
// actually (+,-,*,/,%)
type Type interface {
	GetValue() any
	GetString() string
	//ADD : +
	ADD(other Type) (Type, error)
	//SUB : -
	SUB(other Type) (Type, error)
	//MUL : *
	MUL(other Type) (Type, error)
	//DIV : /
	DIV(other Type) (Type, error)
	//MOD : %
	MOD(other Type) (Type, error)
}
