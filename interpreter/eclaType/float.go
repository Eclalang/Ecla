package eclaType

import (
	"errors"
	"fmt"
)

// NewFloat returns a new Float.
func NewFloat(value float32) Type {
	return &Float{Value: value}
}

type Float struct {
	Value float32
}

// GetValue returns the value of the Float.
func (f *Float) GetValue() any {
	return f.Value
}

// GetString returns the string value of the Float.
func (f *Float) GetString() string {
	return fmt.Sprint(f.Value)
}

// ADD returns the sum of the two Type of Ecla or error.
func (f *Float) ADD(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case float32:
		return &Float{Value: f.Value + vOther.(float32)}, nil
	case int:
		return &Float{Value: f.Value + float32(vOther.(int))}, nil
	default:
		return nil, errors.New("cannot add " + other.GetString() + " to float")
	}
}

// SUB returns the difference of the two Type of Ecla or error.
func (f *Float) SUB(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case float32:
		return &Float{Value: f.Value - vOther.(float32)}, nil
	case int:
		return &Float{Value: f.Value - float32(vOther.(int))}, nil
	default:
		return nil, errors.New("cannot subtract " + other.GetString() + " from float")
	}
}

// MUL returns the product of the two Type of Ecla or error.
func (f *Float) MUL(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case float32:
		return &Float{Value: f.Value * vOther.(float32)}, nil
	case int:
		return &Float{Value: f.Value * float32(vOther.(int))}, nil
	default:
		return nil, errors.New("cannot multiply float by " + other.GetString())
	}
}

// DIV returns the quotient of the two Type of Ecla or error.
func (f *Float) DIV(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case float32:
		return &Float{Value: f.Value / vOther.(float32)}, nil
	case int:
		return &Float{Value: f.Value / float32(vOther.(int))}, nil
	default:
		return nil, errors.New("cannot divide float by " + other.GetString())
	}
}

// MOD returns Modulo of two Type of Ecla or error.
func (f *Float) MOD(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case float32:
		return &Float{Value: float32(int(f.Value) % int(vOther.(float32)))}, nil
	case int:
		return &Float{Value: float32(int(f.Value) % vOther.(int))}, nil
	default:
		return nil, errors.New("cannot mod float by " + other.GetString())
	}
}
