package eclaKeyWord

import (
	"errors"
	"github.com/tot0p/Ecla/interpreter/eclaType"
)

type Var struct {
	Name  string
	Value eclaType.Type
}

func (v *Var) String() string {
	return v.Name + " = " + string(v.Value.GetString())
}

func (v *Var) GetString() eclaType.String {
	return eclaType.String(v.String())
}

func (v *Var) GetValue() any {
	return v.Value.GetValue()
}

func (v *Var) GetType() eclaType.String {
	return "Var"
}

// SetValue sets the value of the variable
func (v *Var) SetValue(value eclaType.Type) error {
	if v.Value.GetType() == value.GetType() {
		v.Value = value
		return nil
	}
	return errors.New("cannot set value of " + v.Name + " to " + string(value.GetString()) + " because it is of type " + string(value.GetType()) + " and not " + string(v.Value.GetType()))
}

// NewVar creates a new variable
func NewVar(name string, value eclaType.Type) *Var {
	return &Var{Name: name, Value: value}
}

// Add adds two Type objects
func (v *Var) Add(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Add(other)
}

// Sub returns errors because you cannot subtract strings
func (v *Var) Sub(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Sub(other)
}

// Mul returns errors because you cannot multiply strings
func (v *Var) Mul(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Mul(other)
}

// Div returns errors because you cannot divide strings
func (v *Var) Div(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Div(other)
}

// Mod returns errors because you cannot mod strings
func (v *Var) Mod(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Mod(other)
}

// DivEc
func (v *Var) DivEc(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.DivEc(other)
}

// Eq
func (v *Var) Eq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Eq(other)
}

// NotEq
func (v *Var) NotEq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.NotEq(other)
}

// Gt
func (v *Var) Gt(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Gt(other)
}

// GtEq
func (v *Var) GtEq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.GtEq(other)
}

// Lw
func (v *Var) Lw(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Lw(other)
}

// LwEq
func (v *Var) LwEq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.LwEq(other)
}

// And
func (v *Var) And(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.And(other)
}

// Or
func (v *Var) Or(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Or(other)
}

// Not
func (v *Var) Not(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Not(other)
}
