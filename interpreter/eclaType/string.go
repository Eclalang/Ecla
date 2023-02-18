package eclaType

import (
	"errors"
	"strings"
)

// NewString creates a new String
func NewString(value string) String {
	if strings.Contains(value, "\\n") {
		value = strings.ReplaceAll(value, "\\n", "\n")
	}
	if strings.Contains(value, "\\\\") {
		value = strings.ReplaceAll(value, "\\\\", "\\")
	}
	if strings.Contains(value, "\\a") {
		value = strings.ReplaceAll(value, "\\a", "\a")
	}
	if strings.Contains(value, "\\b") {
		value = strings.ReplaceAll(value, "\\b", "\b")
	}
	if strings.Contains(value, "\\f") {
		value = strings.ReplaceAll(value, "\\f", "\f")
	}
	if strings.Contains(value, "\\r") {
		value = strings.ReplaceAll(value, "\\r", "\r")
	}
	if strings.Contains(value, "\\t") {
		value = strings.ReplaceAll(value, "\\t", "\t")
	}
	if strings.Contains(value, "\\v") {
		value = strings.ReplaceAll(value, "\\v", "\v")
	}
	if strings.Contains(value, "\\\\'") {
		value = strings.ReplaceAll(value, "\\\\'", "\\'")
	}
	if strings.Contains(value, "\\\"") {
		value = strings.ReplaceAll(value, "\\\"", "\"")
	}
	return String(value)
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

// GetIndex returns a single character
func (s String) GetIndex(other Type) (Type, error) {

	if other.GetType() == "int" {
		ind := int(other.GetValue().(Int))
		if ind >= len(s) || ind < 0 {
			return nil, errors.New("Index out of range")
		}
		res := String(string(s)[ind])
		return res, nil
	}
	return nil, errors.New("index must be an integer")
}

// Len returns the length of a string
func (s String) Len() int {
	return len(s)
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

func (s String) Append(other Type) (Type, error) {
	switch other.(type) {
	case String:
		return s + other.GetString(), nil
	}
	return nil, errors.New("cannot append string")
}

func (s String) IsNull() bool {
	return false
}
