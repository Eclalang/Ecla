package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
)

func NewLib(name string) *Lib {
	return &Lib{Name: name}
}

type Lib struct {
	Name string
}

// Lib Method for interface Type

func (l *Lib) GetValue() any {
	return l
}

func (l *Lib) SetValue(value any) error {
	return errors.New("cannot set value of lib")
}

func (l *Lib) String() string {
	return l.Name
}

func (l *Lib) GetString() String {
	return String(l.Name)
}

func (l *Lib) GetType() string {
	return "lib"
}

func (l *Lib) GetIndex(number Type) (*Type, error) {
	return nil, errors.New("cannot get index from lib")
}

func (l *Lib) Add(other Type) (Type, error) {
	return nil, errors.New("cannot add " + l.String() + " and " + other.String())
}

func (l *Lib) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract " + other.String() + " from " + l.String())
}

func (l *Lib) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply " + l.String() + " by " + other.String())
}

func (l *Lib) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide " + l.String() + " by " + other.String())
}

func (l *Lib) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot get remainder of " + l.String() + " by " + other.String())
}

func (l *Lib) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get quotient of " + l.String() + " by " + other.String())
}

func (l *Lib) Eq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) NotEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) And(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) Or(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) Not() (Type, error) {
	return nil, errors.New("cannot \"not\" a lib")
}

func (l *Lib) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot compare " + l.String() + " and " + other.String())
}

func (l *Lib) IsNull() bool {
	return false
}

func (l *Lib) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append " + other.String() + " to " + l.String())
}

func (l *Lib) GetSize() int {
	return utils.Sizeof(l)
}

func (l *Lib) Len() (int, error) {
	return -1, errors.New("lib does not support len")
}
