package eclaType

import (
	"errors"
	"fmt"
)

type Float float32

// GetValue returns the value of the float
func (f Float) GetValue() any {
	return f
}

// GetString returns the string representation of the float
func (f Float) GetString() String {
	return String(fmt.Sprint(f))
}

// Add adds two Type objects compatible with Float
func (f Float) Add(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return f + Float(other.(Int)), nil
	case Float:
		return f + other.(Float), nil
	case String:
		return f.GetString() + other.GetString(), nil
	default:
		return nil, errors.New("cannot add " + string(other.GetString()) + " to float")
	}
}

// Sub subtracts two Type objects compatible with Float
func (f Float) Sub(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return f - Float(other.(Int)), nil
	case Float:
		return f - other.(Float), nil
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from float")
	}
}

// Mod returns error
func (f Float) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod float")
}

// Mul multiplies two Type objects compatible with Float
func (f Float) Mul(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return f * Float(other.(Int)), nil
	case Float:
		return f * other.(Float), nil
	default:
		return nil, errors.New("cannot multiply " + string(other.GetString()) + " by float")
	}
}

// Div divides two Type objects compatible with Float
func (f Float) Div(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return f / Float(other.(Int)), nil
	case Float:
		return f / other.(Float), nil
	default:
		return nil, errors.New("cannot divide " + string(other.GetString()) + " by float")
	}
}
