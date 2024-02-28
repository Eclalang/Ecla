package eclaType

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"strconv"
)

// NewChar creates a new Char
func NewChar(value string) (Char, error) {
	if len(value) == 0 {
		return Char(0), nil
	}
	value = `'` + value + `'`
	value, err := strconv.Unquote(value)
	if err != nil {
		return 0, err
	}
	result := []rune(value)
	if len(result) > 1 {
		return 0, errors.New(fmt.Sprint(value, " is not a char"))
	}
	return Char(result[0]), nil
}

type Char rune

// GetValue returns the value of the char
func (c Char) GetValue() any {
	return c
}

// SetValue returns an error
func (c Char) SetValue(value any) error {
	return errors.New("cannot set value to char")
}

func (c Char) String() string {
	return string(c)
}

// GetString returns the string representation of the char
func (c Char) GetString() String {
	return String(c)
}

// GetType returns the type Char
func (Char) GetType() string {
	return "char"
}

// GetIndex returns an error
func (Char) GetIndex(index Type) (*Type, error) {
	return nil, errors.New("cannot get index from char")
}

// Add adds two Type objects compatible with Char
func (c Char) Add(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(Int(c) + other.GetValue().(Int)), nil
	case Char:
		return c + other.(Char), nil
	case String:
		return c.GetString() + other.GetString(), nil
	case *Any:
		return c.Add(other.(*Any).Value)
	default:
		return nil, errors.New("cannot add " + string(other.GetString()) + " to char")
	}
}

// Sub subtracts two Type objects compatible with Char
func (c Char) Sub(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(Int(c) - other.GetValue().(Int)), nil
	case Char:
		return c - other.(Char), nil
	case *Any:
		return c.Sub(other.(*Any).Value)
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from char")
	}
}

// Mul multiplies two Type objects compatible with Char
func (c Char) Mul(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(Int(c) * other.GetValue().(Int)), nil
	case Char:
		return c * other.(Char), nil
	case *Any:
		return c.Mul(other.(*Any).Value)
	default:
		return nil, errors.New("cannot multiply " + string(other.GetString()) + " with char")
	}
}

// Div divides two Type objects compatible with Char
func (c Char) Div(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if other.(Int) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return Char(Int(c) / other.GetValue().(Int)), nil
	case Char:
		if other.(Char) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return c / other.(Char), nil
	case *Any:
		return c.Div(other.(*Any).Value)
	default:
		return nil, errors.New("cannot divide char by " + string(other.GetString()))
	}
}

// Mod returns the remainder of the division of two Type objects compatible with Char
func (c Char) Mod(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if other.(Int) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return Char(Int(c) % other.GetValue().(Int)), nil
	case Char:
		if other.(Char) == 0 {
			return nil, errors.New("cannot divide by zero")
		}
		return c % other.(Char), nil
	case *Any:
		return c.Mod(other.(*Any).Value)
	default:
		return nil, errors.New("cannot get modulo of char by " + string(other.GetString()))
	}
}

// DivEc divides two Type objects compatible with Char
func (c Char) DivEc(other Type) (Type, error) {
	return c.Div(other)
}

// Eq returns true if two Type objects are equal
func (c Char) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c == other.(Char)), nil
	case Int:
		return Bool(Int(c) == other.(Int)), nil
	case *Any:
		return c.Eq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to char")
	}
}

// NotEq returns true if two Type objects are not equal
func (c Char) NotEq(other Type) (Type, error) {
	b, err := c.Eq(other)
	switch b.(type) {
	case Bool:
		b = !b.(Bool)
	}
	return b, err
}

// Gt returns true if the first Type object is greater than the second
func (c Char) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c > other.(Char)), nil
	case Int:
		return Bool(Int(c) > other.(Int)), nil
	case *Any:
		return c.Gt(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to char")
	}
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (c Char) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c >= other.(Char)), nil
	case Int:
		return Bool(Int(c) >= other.(Int)), nil
	case *Any:
		return c.GtEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to char")
	}
}

// Lw returns true if the first Type object is lower than the second
func (c Char) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c < other.(Char)), nil
	case Int:
		return Bool(Int(c) < other.(Int)), nil
	case *Any:
		return c.Lw(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to char")
	}
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (c Char) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c <= other.(Char)), nil
	case Int:
		return Bool(Int(c) <= other.(Int)), nil
	case *Any:
		return c.LwEq(other.(*Any).Value)
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to char")
	}
}

// And returns an error
func (c Char) And(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if c == Char(0) || other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if c == Char(0) || other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if c == Char(0) || other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if c == Char(0) || other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return c.And(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare char to " + other.GetString()))
	}
}

// Or returns an error
func (c Char) Or(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if c == Char(0) && other.GetValue() == Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if c == Char(0) && other.GetValue() == Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if c == Char(0) && other.GetValue() == Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if c == Char(0) && other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return c.Or(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare char to " + other.GetString()))
	}
}

// Not returns an error
func (c Char) Not() (Type, error) {
	switch c.GetValue() {
	case Char(0):
		return Bool(true), nil
	default:
		return Bool(false), nil
	}
}

func (c Char) Xor(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		if c == Char(0) && other.GetValue() == Int(0) {
			return Bool(false), nil
		} else if c != Char(0) && other.GetValue() != Int(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Float:
		if c == Char(0) && other.GetValue() == Float(0) {
			return Bool(false), nil
		} else if c != Char(0) && other.GetValue() != Float(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Char:
		if c == Char(0) && other.GetValue() == Char(0) {
			return Bool(false), nil
		} else if c != Char(0) && other.GetValue() != Char(0) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case Bool:
		if c == Char(0) && other.GetValue() == Bool(false) {
			return Bool(false), nil
		} else if c != Char(0) && other.GetValue() != Bool(false) {
			return Bool(false), nil
		} else {
			return Bool(true), nil
		}
	case *Any:
		return c.Xor(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare char to " + other.GetString()))
	}
}

func (Char) IsNull() bool {
	return false
}

func (c Char) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append to char")
}

func (c Char) GetSize() int {
	return utils.Sizeof(c)
}

func (c Char) Len() (int, error) {
	return -1, errors.New("cannot get length of char")
}

func (c Char) GetValueAsInt() Int {
	return NewInt(strconv.Itoa(int(c)))
}
