package eclaType

import (
	"errors"
	"strconv"
)

// NewInt returns a new Int.
func NewInt(value int) Type {
	return &Int{Value: value}
}

type Int struct {
	Value int
}

// GetValue returns the value of the Int.
func (i *Int) GetValue() any {
	return i.Value
}

// GetString returns the string value of the Int .
func (i *Int) GetString() string {
	return strconv.Itoa(i.Value)
}

// ADD returns the sum of the two Type of Ecla or error.
func (i *Int) ADD(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		return &Int{Value: i.Value + vOther.(int)}, nil
	case string:
		return &String{Value: i.GetString() + vOther.(string)}, nil
	default:
		return nil, errors.New("cannot add " + other.GetString() + " to int")
	}
}

// SUB returns the difference of the two Type of Ecla or error.
func (i *Int) SUB(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		return &Int{Value: i.Value - vOther.(int)}, nil
	default:
		return nil, errors.New("cannot subtract " + other.GetString() + " from int")
	}
}

// MUL returns the product of the two Type of Ecla or error.
func (i *Int) MUL(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		return &Int{Value: i.Value * vOther.(int)}, nil
	default:
		return nil, errors.New("cannot multiply int by " + other.GetString())
	}
}

// DIV returns the quotient of the two Type of Ecla or error.
func (i *Int) DIV(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		return &Int{Value: i.Value / vOther.(int)}, nil
	default:
		return nil, errors.New("cannot divide int by " + other.GetString())
	}
}

// MOD returns the remainder of the two Type of Ecla or error.
func (i *Int) MOD(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		return &Int{Value: i.Value % vOther.(int)}, nil
	default:
		return nil, errors.New("cannot mod int by " + other.GetString())
	}
}
