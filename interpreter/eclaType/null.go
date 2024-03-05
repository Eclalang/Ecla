package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
)

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
	return errors.New("cannot get index from null")
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
	return nil, errors.New("cannot add " + n.String() + " and " + other.String())
}

func (n Null) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract " + other.String() + " from " + n.String())
}

func (n Null) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply " + n.String() + " by " + other.String())
}

func (n Null) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide " + n.String() + " by " + other.String())
}

func (n Null) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot get remainder of " + n.String() + " by " + other.String())
}

func (n Null) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get quotient of " + n.String() + " by " + other.String())
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
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) Or(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) Not() (Type, error) {
	return nil, errors.New("cannot \"not\" null")
}

func (n Null) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) Append(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + n.String() + " and " + other.String())
}

func (n Null) IsNull() bool {
	return true
}

func (n Null) GetSize() int {
	return utils.Sizeof(n)
}

func (n Null) Len() (int, error) {
	return -1, errors.New("cannot get length of null")
}
