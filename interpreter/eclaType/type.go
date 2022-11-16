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
	//DivEc : //
	DivEc(other Type) (Type, error)
	//Eq : ==
	Eq(other Type) (Type, error)
	//NotEq : !=
	NotEq(other Type) (Type, error)
	//Gt : >
	Gt(other Type) (Type, error)
	//GtEq : >=
	GtEq(other Type) (Type, error)
	//Lw : <
	Lw(other Type) (Type, error)
	//LwEq : <=
	LwEq(other Type) (Type, error)
	//And : &&
	//And(other Type) (Type, error)
	//Or : ||
	//Or(other Type) (Type, error)
	//Opposite : !
	//Opposite(other Type) (Type, error)
}
