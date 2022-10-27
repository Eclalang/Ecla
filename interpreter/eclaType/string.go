package eclaType

import (
	"errors"
)

type String string

// GetValue returns the value of the string
func (s String) GetValue() any {
	return s
}

// GetString returns the string
func (s String) GetString() String {
	return s
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
