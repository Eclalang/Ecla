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
	return nil
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
	return nil, errors.New("lib does not support indexing")
}

func (l *Lib) Add(other Type) (Type, error) {
	return nil, errors.New("cannot add lib")
}

func (l *Lib) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot sub lib")
}

func (l *Lib) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot mul lib")
}

func (l *Lib) Div(other Type) (Type, error) {
	return nil, errors.New("cannot div lib")
}

func (l *Lib) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod lib")
}

func (l *Lib) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divEc lib")
}

func (l *Lib) Eq(other Type) (Type, error) {
	return nil, errors.New("cannot eq lib")
}

func (l *Lib) NotEq(other Type) (Type, error) {
	return nil, errors.New("cannot notEq lib")
}

func (l *Lib) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot gt lib")
}

func (l *Lib) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot gtEq lib")
}

func (l *Lib) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot lw lib")
}

func (l *Lib) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot lwEq lib")
}

func (l *Lib) And(other Type) (Type, error) {
	return nil, errors.New("cannot and lib")
}

func (l *Lib) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or lib")
}

func (l *Lib) Not() (Type, error) {
	return nil, errors.New("cannot not lib")
}

func (l *Lib) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor lib")
}

func (l *Lib) IsNull() bool {
	return false
}

func (l *Lib) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append lib")
}

func (l *Lib) GetSize() int {
	return utils.Sizeof(l)
}
