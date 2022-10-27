package eclaType

type Type interface {
	GetValue() any
	GetString() String
	//Add : +
	Add(other Type) (Type, error)
	//Sub : -
	Sub(other Type) (Type, error)
	//Mul : *
	Mul(other Type) (Type, error)
	//Div : /
	Div(other Type) (Type, error)
	//Mod : %
	Mod(other Type) (Type, error)
}
