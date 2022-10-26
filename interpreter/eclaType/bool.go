package eclaType

import "errors"

// NewBool returns a new Bool.
func NewBool(value bool) Type {
	return &Bool{Value: value}
}

type Bool struct {
	Value bool
}

// GetValue returns the value of the Bool.
func (b *Bool) GetValue() any {
	return b.Value
}

// GetString returns the string value of the Bool.
func (b *Bool) GetString() string {
	if b.Value {
		return "true"
	}
	return "false"
}

// ADD returns error.
func (b *Bool) ADD(other Type) (Type, error) {
	return nil, errors.New("cannot add to bool")
}

// SUB returns error.
func (b *Bool) SUB(other Type) (Type, error) {
	return nil, errors.New("cannot subtract from bool")
}

// MUL returns error.
func (b *Bool) MUL(other Type) (Type, error) {
	return nil, errors.New("cannot multiply bool")
}

// DIV returns error.
func (b *Bool) DIV(other Type) (Type, error) {
	return nil, errors.New("cannot divide bool")
}

// MOD returns error.
func (b *Bool) MOD(other Type) (Type, error) {
	return nil, errors.New("cannot mod bool")
}
