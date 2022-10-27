package eclaType

import "errors"

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
