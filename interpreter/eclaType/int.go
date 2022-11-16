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

// GetString returns the string representation of the int
func (i Int) GetString() String {
	return String(strconv.Itoa(int(i)))
}

// Add adds two Type objects compatible with Int
func (i Int) Add(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return i + other.(Int), nil
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
	case Int:
		return i - other.(Int), nil
	case Float:
		return Float(i) - other.(Float), nil
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from int")
	}
}

// Mod returns the remainder of the division of two Type objects compatible with Int
func (i Int) Mod(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return i % other.(Int), nil
	default:
		return nil, errors.New("cannot mod " + string(other.GetString()) + " by int")
	}
}

// Mul multiplies two Type objects compatible with Int
func (i Int) Mul(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return i * other.(Int), nil
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
	case Int:
		return i / other.(Int), nil
	case Float:
		return Float(i) / other.(Float), nil
	default:
		return nil, errors.New("cannot divide " + string(other.GetString()) + " by int")
	}
}

// DivEc divides two Type objects compatible with Int
func (i Int) DivEc(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return (i - i%other.(Int)) / other.(Int), nil
	case Float:
		return nil, errors.New("cannot divide ec by float")
	default:
		return nil, errors.New("cannot divide " + string(other.GetString()) + " by int")
	}
}

// Eq returns true if two Type objects are equal
func (i Int) Eq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(i == other.(Int)), nil
	case Float:
		return Bool(Float(i) == other.(Float)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// NotEq returns true if two Type objects are not equal
func (i Int) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(i != other.(Int)), nil
	case Float:
		return Bool(Float(i) != other.(Float)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// Gt returns true if the first Type object is greater than the second
func (i Int) Gt(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(i > other.(Int)), nil
	case Float:
		return Bool(Float(i) > other.(Float)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (i Int) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(i >= other.(Int)), nil
	case Float:
		return Bool(Float(i) >= other.(Float)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// Lw returns true if the first Type object is lower than the second
func (i Int) Lw(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(i < other.(Int)), nil
	case Float:
		return Bool(Float(i) < other.(Float)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (i Int) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		return Bool(i <= other.(Int)), nil
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
func (i Int) Not(other Type) (Type, error) {
	return nil, errors.New("cannot opposite int")
}
