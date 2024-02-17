package eclaType

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/parser"
)

type Struct struct {
	Fields     map[string]Type
	Typ        string
	Definition *eclaDecl.StructDecl
}

func NewStruct(def *eclaDecl.StructDecl) *Struct {
	return &Struct{map[string]Type{}, def.Name, def}
}

func (s *Struct) AddField(index int, val Type) {
	s.Fields[s.Definition.Order[index]] = val
}

func (s *Struct) GetValue() any {
	return s
}

func (s *Struct) SetValue(v any) error {
	switch v.(type) {
	case *Struct:
		t := v.(*Struct)
		*s = *t
		return nil
	case *Any:
		return s.SetValue(v.(*Any).Value)
	default:
		return errors.New("cannot set value of struct")
	}
}

func (s *Struct) String() string {
	var str string
	for _, fieldName := range s.Definition.Order {
		str += s.Fields[fieldName].String() + ", "
	}
	return s.Typ + "{" + str[:len(str)-2] + "}"
}

// GetString returns the string of map
func (s *Struct) GetString() String {
	return String(fmt.Sprint(s))
}

// GetType returns the type Struct
func (s *Struct) GetType() string {
	return "struct " + s.Typ
}

func (s *Struct) SetType(t string) {
	s.Typ = t
}

func (s *Struct) Set(fieldName string, FieldValue Type) error {
	switch FieldValue.(type) {
	case *Var:
		FieldValue = FieldValue.(*Var).Value
	}
	switch FieldValue.(type) {
	case *Any:
		FieldValue = FieldValue.(*Any).Value
	}

	if _, ok := s.Fields[fieldName]; !ok {
		return errors.New("fieldName does not exist")
	}
	if s.Fields[fieldName].GetType() == parser.Any {
		return s.Fields[fieldName].(*Any).SetAny(FieldValue)
	}
	if s.Fields[fieldName].GetType() != FieldValue.GetType() {
		return errors.New("fieldValue is of type " + FieldValue.GetType() + ", expected " + s.Fields[fieldName].GetType())
	}

	s.Fields[fieldName] = FieldValue
	return nil
}

func (s *Struct) Get(fieldName string) (Type, error) {
	if _, ok := s.Fields[fieldName]; !ok {
		return nil, errors.New("fieldName does not exist")
	}
	return s.Fields[fieldName], nil
}

func (s *Struct) GetIndex(index Type) (*Type, error) {
	if index.GetType() != parser.String {
		return nil, errors.New("index type not match")
	}
	val, err := s.Get(index.String())
	return &val, err
}

func (s *Struct) Add(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case String:
		return s.GetString() + other.GetString(), nil
	case *Any:
		return s.Add(other.(*Any).Value)
	}
	return nil, fmt.Errorf("cannot add %s to list", other.GetString())
}

func (s *Struct) Sub(value Type) (Type, error) {
	return nil, errors.New("Cannot sub struct")
}

func (s *Struct) Mul(value Type) (Type, error) {
	return nil, errors.New("cannot mul struct")
}

func (s *Struct) Div(value Type) (Type, error) {
	return nil, errors.New("cannot div struct")
}

func (s *Struct) Mod(value Type) (Type, error) {
	return nil, errors.New("cannot mod struct")
}

func (s *Struct) DivEc(value Type) (Type, error) {
	return nil, errors.New("cannot divec struct")
}
func (s *Struct) DivMod(value Type) (Type, error) {
	return nil, errors.New("cannot divmod struct")
}

// TODO add case var ?
func (s *Struct) Eq(value Type) (Type, error) {
	switch value.(type) {
	case *Struct:
		if s.Typ != value.(*Struct).Typ {
			return Bool(false), nil
		}
		if len(s.Fields) != len(value.(*Struct).Fields) {
			return Bool(false), nil
		}
		for name, value := range s.Fields {
			vVal, ok := value.(*Struct).Fields[name]
			if !ok || (vVal != value) {
				return Bool(false), nil
			}
		}
		return Bool(true), nil
	case *Any:
		return s.Eq(value.(*Any).Value)
	default:
		return Bool(false), errors.New("cannot compare struct with other type")
	}
}

func (s *Struct) NotEq(value Type) (Type, error) {
	v, err := s.Eq(value)
	if err != nil {
		return nil, err
	}
	return !v.(Bool), nil
}

func (s *Struct) And(other Type) (Type, error) {
	return nil, errors.New("cannot and struct")
}

func (s *Struct) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or struct")
}

func (s *Struct) Not() (Type, error) {
	return nil, errors.New("cannot not struct")
}

func (s *Struct) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor struct")
}

func (s *Struct) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot gt struct")
}

func (s *Struct) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot gtEq struct")
}

func (s *Struct) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot lw struct")
}

func (s *Struct) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot lwEq struct")
}

func (s *Struct) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append struct")
}

func (s *Struct) IsNull() bool {
	return false
}

func IsStruct(typ string) bool {
	if len(typ) < 7 {
		return false
	}
	if typ[:6] != "struct" {
		return false
	}
	return true
}
