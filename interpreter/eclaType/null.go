package eclaType

import "errors"

type Null struct {
	typ string
}

func NewNull() Null {
	return Null{""}
}

func NewNullType(typ string) Null {
	return Null{typ}
}

func (n Null) GetValue() any {
	return nil
}

func (n Null) SetValue(value any) error {
	return nil
}

func (n Null) String() string {
	return "null"
}

func (n Null) GetString() String {
	return "null"
}

func (n Null) GetType() string {
	return n.typ
}

func (n Null) GetIndex(number Type) (*Type, error) {
	return nil, nil
}

func (n Null) Add(other Type) (Type, error) {
	if other.GetType() == "string" {
		return String("null" + string(other.GetString())), nil
	}
	return nil, errors.New("cannot add null to " + other.GetType())
}

func (n Null) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract null")
}

func (n Null) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply null")
}

func (n Null) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide null")
}

func (n Null) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod null")
}

func (n Null) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by null")
}

func (n Null) Eq(other Type) (Type, error) {
	if other.GetType() == "null" {
		return Bool(true), nil
	}
	return Bool(false), nil
}

func (n Null) NotEq(other Type) (Type, error) {
	if other.GetType() == "null" {
		return Bool(false), nil
	}
	return Bool(true), nil
}

func (n Null) And(other Type) (Type, error) {
	return nil, errors.New("cannot and null")
}

func (n Null) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or null")
}

func (n Null) Not() (Type, error) {
	return nil, errors.New("cannot not null")
}

func (n Null) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor null")
}

func (n Null) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot gt null")
}

func (n Null) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot gtEq null")
}

func (n Null) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot lw null")
}

func (n Null) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot lwEq null")
}

func (n Null) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append null")
}

func (n Null) IsNull() bool {
	return true
}
