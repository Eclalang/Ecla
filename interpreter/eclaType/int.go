package eclaType

import (
	"errors"
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
	default:
		return nil, errors.New("cannot add " + string(other.GetString()) + " to int")
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
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from int")
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
		return i % other.(Int), nil
	case Char:
		return i % Int(other.(Char)), nil
	default:
		return nil, errors.New("cannot mod " + string(other.GetString()) + " by int")
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
	default:
		return nil, errors.New("cannot multiply " + string(other.GetString()) + " by int")
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
		if Int(other.(Char)) == 0 {
			return nil, errors.New("Cannot divide by zero")
		}
		return Float(i) / Float(other.(Char)), nil
	case Float:
		if other.(Float) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return Float(i) / other.(Float), nil
	default:
		return nil, errors.New("cannot divide " + string(other.GetString()) + " by int")
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
		return (i - i%other.(Int)) / other.(Int), nil
	case Char:
		return (i - i%Int(other.(Char))) / Int(other.(Char)), nil
	case Float:
		return nil, errors.New("cannot divide ec by float")
	default:
		return nil, errors.New("cannot divide " + string(other.GetString()) + " by int")
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
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
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
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
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
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
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
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
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
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
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
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// And returns errors
func (i Int) And(other Type) (Type, error) {
	return nil, errors.New("cannot and int")
}

// Or returns errors
func (i Int) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or int")
}

// Not returns errors
func (i Int) Not() (Type, error) {
	return nil, errors.New("cannot opposite int")
}

// Append returns errors
func (i Int) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append int")
}

func (i Int) IsNull() bool {
	return false
}
