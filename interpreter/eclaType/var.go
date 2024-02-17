package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/parser"
)

type Var struct {
	Name  string
	Value Type
}

func (v *Var) String() string {
	return v.Name + " = " + string(v.Value.GetString())
}

func (v *Var) GetString() String {
	return v.Value.GetString()
}

func (v *Var) GetValue() any {
	return v.Value.GetValue()
}

// SetValue TODO: Refactor this method
func (v *Var) SetValue(value any) error {
	return errors.New("this should not be called")
}

func (v *Var) GetType() string {
	return v.Value.GetType()
}

func (v *Var) GetIndex(i Type) (*Type, error) {
	return v.Value.GetIndex(i)
}

// SetVar sets the value of the variable
func (v *Var) SetVar(value Type) error {
	switch value.(type) {
	case *Var:
		v.Value = value.(*Var).Value
		return nil
	}
	typ2 := v.Value.GetType()
	if value.IsNull() {
		v.Value = NewNullType(typ2)
		return nil
	}
	typ := value.GetType()
	if typ2 == typ {
		v.Value = value
		return nil
	}
	return errors.New("cannot set value of " + v.Name + " to " + string(value.GetString()) + " because it is of type " + string(typ) + " and not " + string(typ2))
}

// Add adds two Type objects
func (v *Var) Add(other Type) (Type, error) {
	return v.Value.Add(other)
}

// Sub subtracts two Type objects
func (v *Var) Sub(other Type) (Type, error) {

	return v.Value.Sub(other)
}

// Mul multiplies two Type objects
func (v *Var) Mul(other Type) (Type, error) {

	return v.Value.Mul(other)
}

// Div divides two Type objects
func (v *Var) Div(other Type) (Type, error) {

	return v.Value.Div(other)
}

// Mod modulos two Type objects
func (v *Var) Mod(other Type) (Type, error) {

	return v.Value.Mod(other)
}

// DivEc divides two Type objects
func (v *Var) DivEc(other Type) (Type, error) {

	return v.Value.DivEc(other)
}

// Eq returns true if the two Type objects are equal
func (v *Var) Eq(other Type) (Type, error) {

	return v.Value.Eq(other)
}

// NotEq returns true if the two Type objects are not equal
func (v *Var) NotEq(other Type) (Type, error) {

	return v.Value.NotEq(other)
}

// Gt returns true if the first Type object is greater than the second
func (v *Var) Gt(other Type) (Type, error) {

	return v.Value.Gt(other)
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (v *Var) GtEq(other Type) (Type, error) {

	return v.Value.GtEq(other)
}

// Lw returns true if the first Type object is lower than the second
func (v *Var) Lw(other Type) (Type, error) {

	return v.Value.Lw(other)
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (v *Var) LwEq(other Type) (Type, error) {

	return v.Value.LwEq(other)
}

// And returns true if the two Type objects are true
func (v *Var) And(other Type) (Type, error) {

	return v.Value.And(other)
}

// Or returns true if either Type objects is true
func (v *Var) Or(other Type) (Type, error) {

	return v.Value.Or(other)
}

// Not returns the opposite of the Type object
func (v *Var) Not() (Type, error) {
	return v.Value.Not()
}

// Xor returns true if only one of the Type objects is true
func (v *Var) Xor(other Type) (Type, error) {
	return v.Value.Xor(other)
}

func (v *Var) Decrement() {
	var err error
	v.Value, err = v.Value.Sub(NewInt("1"))
	if err != nil {
		panic(err)
	}
}

func (v *Var) Increment() {
	var err error
	v.Value, err = v.Value.Add(NewInt("1"))
	if err != nil {
		panic(err)
	}
}

func (v *Var) Append(other Type) (Type, error) {

	return v.Value.Append(other)
}

func (v *Var) IsNull() bool {
	return v.Value.IsNull()
}

func (v *Var) IsFunction() bool {
	if len(v.Value.GetType()) >= 9 {
		if v.Value.GetType()[:8] == "function" {
			return true
		}
	}
	if len(v.Value.GetType()) >= 13 {
		if v.Value.GetType()[4:12] == "function" {
			return true
		}
	}
	return false
}

func (v *Var) IsAny() bool {
	if len(v.Value.GetType()) >= 4 {
		return v.Value.GetType()[:3] == parser.Any
	}
	return false
}

func (v *Var) GetFunction() *Function {
	switch v.Value.(type) {
	case *Function:
		return v.Value.(*Function)
	case *Any:
		switch v.Value.(*Any).Value.(type) {
		case *Function:
			return v.Value.(*Any).Value.(*Function)
		}
	}
	return nil
}

// NewVar creates a new variable
func NewVar(name string, Type string, value Type) (*Var, error) {
	if Type == parser.String {
		return &Var{
			Name:  name,
			Value: value.GetString(),
		}, nil
	}
	if Type == parser.Float && value.GetType() == parser.Int {
		return &Var{
			Name:  name,
			Value: NewFloat(value.String()),
		}, nil

	}
	if Type == parser.Any {
		val := value
		if value.GetType() != parser.Any {
			val = NewAny(value)
		}
		return &Var{
			Name:  name,
			Value: val,
		}, nil
	}

	if Type == "" {
		Type = value.GetType()
	} else if Type != value.GetType() && !value.IsNull() {
		return nil, errors.New("cannot create variable of type " + Type + " with value of type " + value.GetType())
	}
	if value.IsNull() {
		value = NewNullType(Type)
	}

	switch value.(type) {
	case *Var:
		value = value.(*Var).Value
	}
	return &Var{Name: name, Value: value}, nil
}

func NewVarEmpty(name string, Type string) (*Var, error) {
	return &Var{Name: name, Value: NewNullType(Type)}, nil
}
