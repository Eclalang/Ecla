package eclaType

import (
	"errors"
	"strconv"
)

// NewBool creates a new Bool
func NewBool(value string) Bool {
	result, _ := strconv.ParseBool(value)
	return Bool(result)
}

type Bool bool

// GetValue returns the value of the bool
func (b Bool) GetValue() any {
	return b
}

// GetString returns the string representation of the bool
func (b Bool) GetString() String {
	if b {
		return "true"
	}
	return "false"
}

// returns error
func (b Bool) Add(other Type) (Type, error) {
	return nil, errors.New("cannot add to bool")
}

// returns error
func (b Bool) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract from bool")
}

// returns error
func (b Bool) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod bool")
}

// returns error
func (b Bool) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply bool")
}

// returns error
func (b Bool) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide bool")
}

// returns error
func (b Bool) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec bool")
}

// Eq returns true if two Type objects are equal
func (b Bool) Eq(other Type) (Type, error) {
	return Bool(b == other.GetValue()), nil
}

// NotEq returns true if two Type objects are not equal
func (b Bool) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case Bool:
		return Bool(b != other.GetValue()), nil
	default:
		return nil, errors.New(string("cannot compare bool to " + other.GetString()))
	}
}

// returns error
func (b Bool) Gt(other Type) (Type, error) {
	return nil, errors.New(string("cannot compare bool to " + other.GetString()))
}

// returns error
func (b Bool) GtEq(other Type) (Type, error) {
	return nil, errors.New(string("cannot compare bool to " + other.GetString()))
}

// returns error
func (b Bool) Lw(other Type) (Type, error) {
	return nil, errors.New(string("cannot compare bool to " + other.GetString()))
}

// returns error
func (b Bool) LwEq(other Type) (Type, error) {
	return nil, errors.New(string("cannot compare bool to " + other.GetString()))
}
