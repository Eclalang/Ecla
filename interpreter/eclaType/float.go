package eclaType

import (
	"errors"
	"fmt"
	"strconv"
)

// NewFloat creates a new Float
func NewFloat(value string) Float {
	result, _ := strconv.Atoi(value)
	return Float(result)
}

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

// DivEc returns error because you cannot div ec float
func (f Float) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by float")
}

// Eq returns true if two Type objects are equal
func (f Float) Eq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(f == Float(other.(Int))), nil
	case Float:
		return Bool(f == other.(Float)), nil
	default:
		return nil, errors.New("cannot compare float to " + string(other.GetString()))
	}
}

// NotEq returns true if two Type objects are not equal
func (f Float) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(f != Float(other.(Int))), nil
	case Float:
		return Bool(f != other.(Float)), nil
	default:
		return nil, errors.New("cannot compare float to " + string(other.GetString()))
	}
}

// Gt returns true if the float is greater than the other
func (f Float) Gt(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(f > Float(other.(Int))), nil
	case Float:
		return Bool(f > other.(Float)), nil
	default:
		return nil, errors.New("cannot compare float to " + string(other.GetString()))
	}
}

// GtEq returns true if the float is greater than or equal to the other
func (f Float) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(f >= Float(other.(Int))), nil
	case Float:
		return Bool(f >= other.(Float)), nil
	default:
		return nil, errors.New("cannot compare float to " + string(other.GetString()))
	}
}

// Lw returns true if the float is lower than the other
func (f Float) Lw(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(f < Float(other.(Int))), nil
	case Float:
		return Bool(f < other.(Float)), nil
	default:
		return nil, errors.New("cannot compare float to " + string(other.GetString()))
	}
}

// LwEq returns true if the float is lower than or equal to the other
func (f Float) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(f <= Float(other.(Int))), nil
	case Float:
		return Bool(f <= other.(Float)), nil
	default:
		return nil, errors.New("cannot compare float to " + string(other.GetString()))
	}
}

// And returns error
func (f Float) And(other Type) (Type, error) {
	return nil, errors.New("cannot and float")
}
