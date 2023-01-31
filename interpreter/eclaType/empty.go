package eclaType

import "errors"

type Empty struct {
	typ string
}

func NewEmpty(typ string) Empty {
	return Empty{typ}
}

func (e Empty) GetValue() any {
	return e
}

// SetValue returns nil
func (e Empty) SetValue(value any) error {
	return nil
}

// String returns empty
func (e Empty) String() string {
	return "empty"
}

// GetString returns empty
func (e Empty) GetString() String {
	return String("empty")
}

// GetType returns the Empty type
func (e Empty) GetType() string {
	return e.typ
}

// returns error
func (e Empty) GetIndex(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Add(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Div(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Eq(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) NotEq(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) And(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Or(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Not() (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

// returns error
func (e Empty) Append(other Type) (Type, error) {
	return nil, errors.New("cannot get index from bool")
}

func (e Empty) IsNull() bool {
	return false
}
