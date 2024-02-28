package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
	"strconv"
)

func StringAreEscapedChar(value string) bool {
	for i := 0; i < len(value); i++ {
		switch value[i] {
		case '\n', '\t', '\r', '\f', '\b', '\a':
			return true
		}
	}
	return false

}

// NewString creates a new String
func NewString(value string) (String, error) {
	// check if there are any escape characters
	if StringAreEscapedChar(value) {
		return String(""), errors.New("cannot create string with escape characters")
	}
	value = `"` + value + `"`
	value, err := strconv.Unquote(value)
	if err != nil {
		return String(""), err
	}
	return String(value), nil
}

type String string

// GetValue returns the value of the string
func (s String) GetValue() any {
	return s
}

// SetValue
func (s String) SetValue(value any) error {
	return errors.New("cannot set value to String")
}

func (s String) String() string {
	return string(s)
}

// GetString returns the string
func (s String) GetString() String {
	return s
}

// GetType returns the type String
func (s String) GetType() string {
	return "string"
}

// TODO refactor with switch ?
// GetIndex returns a single character
func (s String) GetIndex(other Type) (*Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	if len(other.GetType()) >= 4 {
		if other.GetType()[:3] == parser.Any {
			return s.GetIndex(other.(*Any).Value)
		}
	}
	if other.GetType() == "int" {
		ind := int(other.GetValue().(Int))
		if ind >= len(s) || ind < 0 {
			return nil, errors.New("Index out of range")
		}
		res := Char(s[ind])
		temp := Type(res)
		return &temp, nil
	}
	return nil, errors.New("index must be an integer")
}

// Add adds two Type objects
func (s String) Add(other Type) (Type, error) {
	return s + other.GetString(), nil
}

// Sub returns errors because you cannot subtract strings
func (s String) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract from string")
}

// Mod returns errors because you cannot mod strings
func (s String) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod string")
}

// Mul if other is Int , return n * String
func (s String) Mul(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		result := ""
		for i := 0; i < int(other.(Int)); i++ {
			result += string(s)
		}
		return String(result), nil
	case Char:
		result := ""
		for i := 0; i < int(other.(Char)); i++ {
			result += string(s)
		}
		return String(result), nil
	case *Any:
		return s.Mul(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot multiply string by " + other.GetString()))
	}
}

// Div returns errors because you cannot divide strings
func (s String) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide string")
}

// DivEc returns error because you cannot div ec strings
func (s String) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by string")
}

// Eq returns true if two Type objects are equal
func (s String) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return Bool(s == other.GetString()), nil
	case *Any:
		return s.Eq(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare string by " + other.GetString()))
	}
}

// NotEq returns true if two Type objects are not equal
func (s String) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return Bool(s != other.GetString()), nil
	case *Any:
		return s.NotEq(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare string to " + other.GetString()))
	}
}

// Gt returns true if s is greater than other
func (s String) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return Bool(s > other.GetString()), nil
	case *Any:
		return s.Gt(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare string to " + other.GetString()))
	}
}

// GtEq returns true if s is greater than or equal to other
func (s String) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return Bool(s >= other.GetString()), nil
	case *Any:
		return s.GtEq(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare string to " + other.GetString()))
	}
}

// Lw returns true if s is lower than other
func (s String) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return Bool(s < other.GetString()), nil
	case *Any:
		return s.Lw(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare string to " + other.GetString()))
	}
}

// LwEq returns true if s is lower than or equal to other
func (s String) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return Bool(s <= other.GetString()), nil
	case *Any:
		return s.LwEq(other.(*Any).Value)
	default:
		return nil, errors.New(string("cannot compare string to " + other.GetString()))
	}
}

// And returns errors
func (s String) And(other Type) (Type, error) {
	return nil, errors.New("cannot and string")
}

// Or returns errors
func (s String) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or string")
}

// Not returns errors
func (s String) Not() (Type, error) {
	return nil, errors.New("cannot opposite string")
}

// Xor returns errors
func (s String) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor string")
}

func (s String) Append(other Type) (Type, error) {
	switch other.(type) {
	case String:
		return s + other.GetString(), nil
	case Char:
		return s + other.GetString(), nil
	case *Any:
		return s.Append(other.(*Any).Value)
	}
	return nil, errors.New("cannot append string with " + other.GetType())
}

func (s String) IsNull() bool {
	return false
}

func (s String) GetSize() int {
	return utils.Sizeof(s)
}

func (s String) Len() (int, error) {
	return len(s), nil
}
