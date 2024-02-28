package eclaType

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"strconv"
)

// NewFloat creates a new Float
func NewFloat(value string) Float {
	result, error := strconv.ParseFloat(value, 32)
	if error != nil {
		panic(error)
	}
	return Float(result)
}

type Float float32

// GetValue returns the value of the float
func (f Float) GetValue() any {
	return f
}

// SetValue
func (f Float) SetValue(value any) error {
	return errors.New("cannot set value to float")
}

func (f Float) String() string {
	return fmt.Sprintf("%g", f)
}

// GetString returns the string representation of the float
func (f Float) GetString() String {
	return String(f.String())
}

// GetType returns the type Float
func (f Float) GetType() string {
	return "float"
}

// returns error
func (f Float) GetIndex(other Type) (*Type, error) {
	return nil, errors.New("cannot get index from float")
}

// Add adds two Type objects compatible with Float
func (f Float) Add(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return f + Float(other.(Int)), nil
	case Float:
		return f + other.(Float), nil
	case String:
		return f.GetString() + other.GetString(), nil
	case Char:
		return f + Float(Int(other.(Char))), nil
	case *Any:
		return f.Add(other.(*Any).Value)
	default:
		return nil, errors.New("cannot add " + f.String() + " and " + other.String())
	}
}

// Sub subtracts two Type objects compatible with Float
func (f Float) Sub(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return f - Float(other.(Int)), nil
	case Char:
		return f - Float(Int(other.(Char))), nil
	case Float:
		return f - other.(Float), nil
	case *Any:
		return f.Sub(other.(*Any).Value)
	default:
		return nil, errors.New("cannot subtract " + other.String() + " from " + f.String())
	}
}

// Mod returns error
func (f Float) Mod(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	case *Any:
		return f.Mod(other.(*Any).Value)
	}
	return nil, errors.New("cannot get remainder of float")
}

// Mul multiplies two Type objects compatible with Float
func (f Float) Mul(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return f * Float(other.(Int)), nil
	case Char:
		return f * Float(Int(other.(Char))), nil
	case Float:
		return f * other.(Float), nil
	case *Any:
		return f.Mul(other.(*Any).Value)
	default:
		return nil, errors.New("cannot multiply " + f.String() + " by " + other.String())
	}
}

// Div divides two Type objects compatible with Float
func (f Float) Div(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return f / Float(other.(Int)), nil
	case Char:
		return f / Float(other.(Char)), nil
	case Float:
		return f / other.(Float), nil
	case *Any:
		return f.Div(other.(*Any).Value)
	default:
		return nil, errors.New("cannot divide " + f.String() + " by " + other.String())
	}
}

// DivEc returns error because you cannot div ec float
func (f Float) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get quotient of float")
}

// Eq returns true if two Type objects are equal
func (f Float) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(f == Float(other.(Int))), nil
	case Char:
		return Bool(f == Float(other.(Char))), nil
	case Float:
		return Bool(f == other.(Float)), nil
	case *Any:
		return f.Eq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
	}
}

// NotEq returns true if two Type objects are not equal
func (f Float) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(f != Float(other.(Int))), nil
	case Char:
		return Bool(f != Float(other.(Char))), nil
	case Float:
		return Bool(f != other.(Float)), nil
	case *Any:
		return f.NotEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
	}
}

// Gt returns true if the float is greater than the other
func (f Float) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(f > Float(other.(Int))), nil
	case Char:
		return Bool(f > Float(other.(Char))), nil
	case Float:
		return Bool(f > other.(Float)), nil
	case *Any:
		return f.Gt(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
	}
}

// GtEq returns true if the float is greater than or equal to the other
func (f Float) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(f >= Float(other.(Int))), nil
	case Char:
		return Bool(f >= Float(other.(Char))), nil
	case Float:
		return Bool(f >= other.(Float)), nil
	case *Any:
		return f.GtEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
	}
}

// Lw returns true if the float is lower than the other
func (f Float) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(f < Float(other.(Int))), nil
	case Char:
		return Bool(f < Float(other.(Char))), nil
	case Float:
		return Bool(f < other.(Float)), nil
	case *Any:
		return f.Lw(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
	}
}

// LwEq returns true if the float is lower than or equal to the other
func (f Float) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Bool(f <= Float(other.(Int))), nil
	case Char:
		return Bool(f <= Float(other.(Char))), nil
	case Float:
		return Bool(f <= other.(Float)), nil
	case *Any:
		return f.LwEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + f.String() + " and " + other.String())
	}
}

// And returns errors
func (f Float) And(other Type) (Type, error) {
	return nil, errors.New("cannot and float")
}

// Or returns errors
func (f Float) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or float")
}

// Not returns errors
func (f Float) Not() (Type, error) {
	return nil, errors.New("cannot opposite float")
}

// Xor
func (f Float) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor float")
}

// Append returns errors
func (f Float) Append(other Type) (Type, error) {
	return nil, errors.New("cannot add " + other.String() + " to " + f.String())
}

func (f Float) IsNull() bool {
	return false
}

func (f Float) GetSize() int {
	return utils.Sizeof(f)
}

func (f Float) Len() (int, error) {
	return 0, errors.New("cannot get length of float")
}
