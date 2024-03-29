package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"strconv"
)

// NewInt creates a new Int
func NewInt(value string) Int {
	result, _ := strconv.Atoi(value)
	return Int(result)
}

type Int int

// GetValue returns the value of the int
func (i Int) GetValue() any {
	return i
}

// SetValue
func (i Int) SetValue(value any) error {
	return errors.New("cannot set value to int")
}

func (i Int) String() string {
	return strconv.Itoa(int(i))
}

// GetString returns the string representation of the int
func (i Int) GetString() String {
	return String(strconv.Itoa(int(i)))
}

// GetType returns the type Int
func (i Int) GetType() string {
	return "int"
}

// returns error
func (i Int) GetIndex(other Type) (*Type, error) {
	return nil, errors.New("cannot get index from int")
}

// Add adds two Type objects compatible with Int
func (i Int) Add(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return i + other.(Int), nil
	case Char:
		return i + Int(other.(Char)), nil
	case Float:
		return Float(i) + other.(Float), nil
	case String:
		return i.GetString() + other.GetString(), nil
	case *Any:
		return i.Add(other.(*Any).Value)
	default:
		return nil, errors.New("cannot add " + i.String() + " and " + other.String())
	}
}

// Sub subtracts two Type objects compatible with Int
func (i Int) Sub(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return i - other.(Int), nil
	case Char:
		return i - Int(other.(Char)), nil
	case Float:
		return Float(i) - other.(Float), nil
	case *Any:
		return i.Sub(other.(*Any).Value)
	default:
		return nil, errors.New("cannot subtract " + other.String() + " from " + i.String())
	}
}

// Mod returns the remainder of the division of two Type objects compatible with Int
func (i Int) Mod(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if other.(Int) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return i % other.(Int), nil
	case Char:
		if other.(Char) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return i % Int(other.(Char)), nil
	case *Any:
		return i.Mod(other.(*Any).Value)
	default:
		return nil, errors.New("cannot get remainder of " + i.String() + " by " + other.String())
	}
}

// Mul multiplies two Type objects compatible with Int
func (i Int) Mul(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return i * other.(Int), nil
	case Char:
		return i * Int(other.(Char)), nil
	case Float:
		return Float(i) * other.(Float), nil
	case String:
		result := String("")
		for j := 0; j < int(i); j++ {
			result += other.GetString()
		}
		return result, nil
	case *Any:
		return i.Mul(other.(*Any).Value)
	default:
		return nil, errors.New("cannot multiply " + i.String() + " and " + other.String())
	}
}

// Div divides two Type objects compatible with Int
func (i Int) Div(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if other.(Int) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return Float(i) / Float(other.(Int)), nil
	case Char:
		if other.(Char) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return Float(i) / Float(other.(Char)), nil
	case Float:
		if other.(Float) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return Float(i) / other.(Float), nil
	case *Any:
		return i.Div(other.(*Any).Value)
	default:
		return nil, errors.New("cannot divide " + i.String() + " by " + other.String())
	}
}

// DivEc divides two Type objects compatible with Int
func (i Int) DivEc(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if other.(Int) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return (i - i%other.(Int)) / other.(Int), nil
	case Char:
		if other.(Char) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return (i - i%Int(other.(Char))) / Int(other.(Char)), nil
	case *Any:
		return i.DivEc(other.(*Any).Value)
	default:
		return nil, errors.New("cannot get quotient of " + i.String() + " by " + other.String())
	}
}

// Eq returns true if two Type objects are equal
func (i Int) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(i == other.(Int)), nil
	case Char:
		return Bool(i == Int(other.(Char))), nil
	case Float:
		return Bool(Float(i) == other.(Float)), nil
	case *Any:
		return i.Eq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// NotEq returns true if two Type objects are not equal
func (i Int) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(i != other.(Int)), nil
	case Char:
		return Bool(i != Int(other.(Char))), nil
	case Float:
		return Bool(Float(i) != other.(Float)), nil
	case *Any:
		return i.NotEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// Gt returns true if the first Type object is greater than the second
func (i Int) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(i > other.(Int)), nil
	case Char:
		return Bool(i > Int(other.(Char))), nil
	case Float:
		return Bool(Float(i) > other.(Float)), nil
	case *Any:
		return i.Gt(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (i Int) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(i >= other.(Int)), nil
	case Char:
		return Bool(i >= Int(other.(Char))), nil
	case Float:
		return Bool(Float(i) >= other.(Float)), nil
	case *Any:
		return i.GtEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// Lw returns true if the first Type object is lower than the second
func (i Int) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(i < other.(Int)), nil
	case Char:
		return Bool(i < Int(other.(Char))), nil
	case Float:
		return Bool(Float(i) < other.(Float)), nil
	case *Any:
		return i.Lw(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (i Int) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(i <= other.(Int)), nil
	case Char:
		return Bool(i <= Int(other.(Char))), nil
	case Float:
		return Bool(Float(i) <= other.(Float)), nil
	case *Any:
		return i.LwEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// And returns errors
func (i Int) And(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if i == Int(0) || other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if i == Int(0) || other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if i == Int(0) || other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if i == Int(0) || other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return i.And(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// Or returns an error
func (i Int) Or(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if i == Int(0) && other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if i == Int(0) && other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if i == Int(0) && other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if i == Int(0) && other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return i.Or(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// Not returns an error
func (i Int) Not() (Type, error) {
	switch i.GetValue() {
	case Int(0):
		return Int(1), nil
	default:
		return Int(0), nil
	}
}

// Xor returns 1 if only one of the Type objects is true
func (i Int) Xor(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if i == Int(0) && other.GetValue() == Int(0) {
			return Bool(false), nil
		} else if i != Int(0) && other.GetValue() != Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if i == Int(0) && other.GetValue() == Char(0) {
			return Bool(false), nil
		} else if i != Int(0) && other.GetValue() != Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if i == Int(0) && other.GetValue() == Float(0) {
			return Bool(false), nil
		} else if i != Int(0) && other.GetValue() != Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if i == Int(0) && other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else if i != Int(0) && other.GetValue() == Bool(true) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return i.Xor(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + i.String() + " and " + other.String())
	}
}

// Append returns errors
func (i Int) Append(other Type) (Type, error) {
	return nil, errors.New("cannot add " + other.String() + " to " + i.String())
}

func (i Int) IsNull() bool {
	return false
}

func (i Int) GetSize() int {
	return utils.Sizeof(i)
}

func (i Int) Len() (int, error) {
	return -1, errors.New("cannot get length of int")
}
