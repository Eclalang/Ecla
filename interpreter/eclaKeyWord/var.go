package eclaKeyWord

import (
	"errors"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/parser"
)

type Var struct {
	Name  string
	Value eclaType.Type
}

func (v *Var) String() string {
	return v.Name + " = " + string(v.Value.GetString())
}

func (v *Var) GetString() eclaType.String {
	return v.Value.GetString()
}

func (v *Var) GetValue() any {
	return v.Value.GetValue()
}

func (v *Var) GetType() string {
	return v.Value.GetType()
}

// SetValue sets the value of the variable
func (v *Var) SetValue(value eclaType.Type) error {
	if v.Value.GetType() == value.GetType() {
		v.Value = value
		return nil
	}
	return errors.New("cannot set value of " + v.Name + " to " + string(value.GetString()) + " because it is of type " + string(value.GetType()) + " and not " + string(v.Value.GetType()))
}

// Add adds two Type objects
func (v *Var) Add(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Add(other)
}

// Sub subtracts two Type objects
func (v *Var) Sub(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Sub(other)
}

// Mul multiplies two Type objects
func (v *Var) Mul(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Mul(other)
}

// Div divides two Type objects
func (v *Var) Div(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Div(other)
}

// Mod modulos two Type objects
func (v *Var) Mod(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Mod(other)
}

// DivEc divides two Type objects
func (v *Var) DivEc(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.DivEc(other)
}

// Eq returns true if the two Type objects are equal
func (v *Var) Eq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Eq(other)
}

// NotEq returns true if the two Type objects are not equal
func (v *Var) NotEq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.NotEq(other)
}

// Gt returns true if the first Type object is greater than the second
func (v *Var) Gt(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Gt(other)
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (v *Var) GtEq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.GtEq(other)
}

// Lw returns true if the first Type object is lower than the second
func (v *Var) Lw(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Lw(other)
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (v *Var) LwEq(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.LwEq(other)
}

// And returns true if the two Type objects are true
func (v *Var) And(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.And(other)
}

// Or returns true if either Type objects is true
func (v *Var) Or(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Or(other)
}

// Not returns the opposite of the Type object
func (v *Var) Not(other eclaType.Type) (eclaType.Type, error) {
	return v.Value.Not(other)
}

func (v *Var) Decrement() {
	var err error
	v.Value, err = v.Value.Sub(eclaType.NewInt("1"))
	if err != nil {
		panic(err)
	}
}

func (v *Var) Increment() {
	var err error
	v.Value, err = v.Value.Add(eclaType.NewInt("1"))
	if err != nil {
		panic(err)
	}
}

// NewVar creates a new variable
func NewVar(name string, Type string, value eclaType.Type) (*Var, error) {
	if Type == parser.String {
		return &Var{
			Name:  name,
			Value: value.GetString(),
		}, nil
	}
	if Type != value.GetType() {
		return nil, errors.New("cannot create variable of type " + Type + " with value of type " + value.GetType())
	}
	return &Var{Name: name, Value: value}, nil
}
