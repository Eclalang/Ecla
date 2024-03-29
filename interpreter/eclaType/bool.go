package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"strconv"
)

// NewBool creates a new Bool
func NewBool(value string) (Bool, error) {
	result, err := strconv.ParseBool(value)
	if err != nil {
		return false, errors.New("cannot convert " + value + " to bool")
	}
	return Bool(result), nil
}

type Bool bool

// GetValue returns the value of the bool
func (b Bool) GetValue() any {
	return b
}

// SetValue
func (b Bool) SetValue(value any) error {
	return errors.New("cannot set value to bool")
}

func (b Bool) String() string {
	return strconv.FormatBool(bool(b))
}

// GetString returns the string representation of the bool
func (b Bool) GetString() String {
	if b {
		return "true"
	}
	return "false"
}

// GetType returns the type Bool
func (b Bool) GetType() string {
	return "bool"
}

// returns error
func (b Bool) GetIndex(other Type) (*Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (b Bool) Add(other Type) (Type, error) {
	if other.GetType() == "string" {
		return b.GetString() + other.GetString(), nil
	}
	return nil, errors.New("cannot add " + b.String() + " and " + other.String())
}

// returns error
func (b Bool) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract " + other.String() + " from " + b.String())
}

// returns error
func (b Bool) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot get remainder of  " + b.String() + " by " + other.String())
}

// returns error"cannot add " + b.String() + " and " + other.String()
func (b Bool) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply " + b.String() + " by " + other.String())
}

// returns error
func (b Bool) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide " + b.String() + " by " + other.String())
}

// returns error
func (b Bool) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get quotient of " + b.String() + " by " + other.String())
}

// Eq returns true if two Type objects are equal
func (b Bool) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if (b == Bool(true) && other.GetValue() != Int(0)) || (b == Bool(false) && other.GetValue() == Int(0)) {
			return Bool(true), nil
		} else {
			return Bool(false), nil
		}
	case Char:
		if (b == Bool(true) && other.GetValue() != Char(0)) || (b == Bool(false) && other.GetValue() == Char(0)) {
			return Bool(true), nil
		} else {
			return Bool(false), nil
		}
	case Float:
		if (b == Bool(true) && other.GetValue() != Float(0)) || (b == Bool(false) && other.GetValue() == Float(0)) {
			return Bool(true), nil
		} else {
			return Bool(false), nil
		}
	case *Any:
		return b.Eq(other.(*Any).Value)
	case Bool:
		return Bool(b == other.GetValue()), nil
	default:
		return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
	}
}

// NotEq returns true if two Type objects are not equal
func (b Bool) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if (b == Bool(true) && other.GetValue() != Int(0)) || (b == Bool(false) && other.GetValue() == Int(0)) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if (b == Bool(true) && other.GetValue() != Char(0)) || (b == Bool(false) && other.GetValue() == Char(0)) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if (b == Bool(true) && other.GetValue() != Float(0)) || (b == Bool(false) && other.GetValue() == Float(0)) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		return Bool(b != other.GetValue()), nil
	case *Any:
		return b.NotEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
	}
}

// returns error
func (b Bool) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
}

// returns error
func (b Bool) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
}

// returns error
func (b Bool) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
}

// returns error
func (b Bool) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
}

// And returns true if both Types are true
func (b Bool) And(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if b == Bool(false) || other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if b == Bool(false) || other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if b == Bool(false) || other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if b == Bool(true) && other.GetValue() == Bool(true) {
			return Bool(true), nil
		} else {
			return Bool(false), nil
		}
	case *Any:
		return b.And(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
	}
}

// Or returns true if either Type is true
func (b Bool) Or(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if b == Bool(false) && other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if b == Bool(false) && other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if b == Bool(false) && other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if b == Bool(true) || other.GetValue() == Bool(true) {
			return Bool(true), nil
		} else {
			return Bool(false), nil
		}
	case *Any:
		return b.Or(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
	}
}

// Not returns the opposite of the bool
func (b Bool) Not() (Type, error) {
	return !b, nil
}

// Xor returns true if only one of the Types is true
func (b Bool) Xor(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if b == Bool(true) && other.GetValue() != Int(0) {
			return Bool(false), nil
		} else if b == Bool(false) && other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if b == Bool(true) && other.GetValue() != Char(0) {
			return Bool(false), nil
		} else if b == Bool(false) && other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if b == Bool(true) && other.GetValue() != Float(0) {
			return Bool(false), nil
		} else if b == Bool(false) && other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if b == Bool(true) && other.GetValue() == Bool(true) {
			return Bool(false), nil
		} else if b == Bool(false) && other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return b.Xor(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + b.String() + " and " + other.String())
	}
}

// Append returns errors
func (b Bool) Append(other Type) (Type, error) {
	return nil, errors.New("cannot add " + other.String() + " to " + b.String())
}

func (b Bool) IsNull() bool {
	return false
}

func (b Bool) GetSize() int {
	return utils.Sizeof(b)
}

func (b Bool) Len() (int, error) {
	return -1, errors.New("cannot get length of bool")
}
