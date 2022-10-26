package eclaType

import (
	"errors"
	"fmt"
)

// NewList returns a new list
func NewList(list []Type) Type {
	return &List{Value: list}
}

type List struct {
	Value []Type
}

// GetValue returns the value of the List.
func (l *List) GetValue() any {
	return l.Value
}

// GetString returns the value of the List.
func (l *List) GetString() string {
	return fmt.Sprint(l.Value)
}

// ADD returns the sum of the two Type of Ecla or error.
func (l *List) ADD(other Type) (Type, error) {
	var lll *List
	vOther := other.GetValue()
	switch vOther.(type) {
	case []Type:
		lll = l
		for _, v := range vOther.([]Type) {
			lll.Value = append(lll.Value, v)
		}
		return lll, nil
	default:
		return nil, errors.New("cannot add " + other.GetString() + " to list")
	}
}

// MUL returns the product of the two Type of Ecla or error.
func (l *List) MUL(other Type) (Type, error) {
	var lll *List
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		for i := 0; i < vOther.(int); i++ {
			for _, v := range l.Value {
				lll.Value = append(lll.Value, v)
			}
		}
		return lll, nil
	default:
		return nil, errors.New("cannot multiply list by " + other.GetString())
	}
}

// SUB returns error.
func (s *List) SUB(other Type) (Type, error) {
	return nil, errors.New("cannot subtract from list")
}

// DIV returns error.
func (s *List) DIV(other Type) (Type, error) {
	return nil, errors.New("cannot divide list")
}

// MOD returns error.
func (s *List) MOD(other Type) (Type, error) {
	return nil, errors.New("cannot mod list")
}
