package eclaType

import (
	"errors"
	"strconv"
)

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
