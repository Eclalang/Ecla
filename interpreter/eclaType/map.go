package eclaType

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/utils"
)

type Map struct {
	Keys   []Type
	Values []Type
	Typ    string
	TypKey string
	TypVal string
}

func NewMap() *Map {
	return &Map{[]Type{}, []Type{}, "", "", ""}
}

func (m *Map) SetAutoType() error {
	var typ string
	if len(m.Values) == 0 {
		typ = "empty"
	} else {
		m.TypKey = m.Keys[0].GetType()
		m.TypVal = m.Values[0].GetType()
		for _, v := range m.Values {
			if v.GetType() != m.TypVal {
				return errors.New("cannot create a map with different value types")
			}
		}
		for _, v := range m.Keys {
			if v.GetType() != m.TypKey {
				return errors.New("cannot create a map with different key types")
			}
		}
		typ = "map[" + m.TypKey + "]" + m.TypVal
	}

	m.Typ = typ
	return nil
}

func (m *Map) GetValue() any {
	return m
}

func (m *Map) SetValue(v any) error {
	switch v.(type) {
	case *Map:
		t := v.(*Map)
		*m = *t
		return nil
	case *Any:
		return m.SetValue(v.(*Any).Value)
	default:
		return errors.New(fmt.Sprintf("cannot set value %s to map of type %s", v, m.Typ))
	}
}

func (m *Map) String() string {
	var s string
	for index, v := range m.Values {
		if index == len(m.Values)-1 {
			s += m.Keys[index].GetString().String() + ": " + v.GetString().String()
		} else {
			s += m.Keys[index].GetString().String() + ": " + v.GetString().String() + ", "
		}
	}
	return "{" + s + "}"
}

// GetString returns the string of map
func (m *Map) GetString() String {
	return String(fmt.Sprint(m))
}

// GetType returns the type Struct
func (m *Map) GetType() string {
	return m.Typ
}

func (m *Map) SetType(t string) {
	m.Typ = t
	m.TypKey, m.TypVal = GetTypeOfKeyAndValue(t)
}

func (m *Map) Set(key Type, value Type) {
	if key.GetType() != m.TypKey && m.TypKey != "string" {
		panic("cannot set " + key.String() + " as key of type " + m.GetKeyTypes())
	}
	if value.GetType() != m.TypVal && m.TypVal != "string" {
		panic("cannot set " + key.String() + " as value of type " + m.GetValueTypes())
	}
	for index, k := range m.Keys {
		if k.GetString().String() == key.GetString().String() {
			m.Values[index] = value
			return
		}
	}
	m.Keys = append(m.Keys, key)
	m.Values = append(m.Values, value)
}

func (m *Map) AddKey(key Type) error {
	if key.GetType() != m.TypKey && m.TypKey != "string" {
		return errors.New("cannot set " + key.String() + " as key of type " + m.GetKeyTypes())
	}
	m.Keys = append(m.Keys, key)
	m.Values = append(m.Values, NewNullType(m.TypVal))
	return nil
}

func (m *Map) Get(key Type) (Type, bool) {
	for index, k := range m.Keys {
		if k.GetString().String() == key.GetString().String() {
			return m.Values[index], true
		}
	}
	return nil, false
}

func (m *Map) GetIndex(index Type) (*Type, error) {
	if index.GetType() != m.TypKey && m.TypKey != "string" {
		return nil, errors.New("cannot set " + index.String() + " as index of type " + m.GetKeyTypes())
	}
	for i, k := range m.Keys {
		if k.GetString().String() == index.GetString().String() {
			return &m.Values[i], nil
		}
	}
	return nil, errors.New("index not found")
}

func (m *Map) GetKey(value Type) (Type, bool) {
	for index, v := range m.Values {
		if v.GetString().String() == value.GetString().String() {
			return m.Keys[index], true
		}
	}
	return nil, false
}

func (m *Map) Add(value Type) (Type, error) {
	switch value.(type) {
	case *Var:
		value = value.(*Var).Value
	}
	switch value.(type) {
	case *Map:
		if value.GetType() != m.Typ && m.TypKey != "string" && m.TypVal != "string" {
			return nil, errors.New("cannot add map with " + value.String())
		}

		for index, v := range value.(*Map).Keys {
			m.Set(v, value.(*Map).Values[index])
		}
		return m, nil
	case String:
		return m.GetString().Add(value)
	case *Any:
		return m.Add(value.(*Any).Value)
	}
	return nil, errors.New("cannot add map with " + value.String())
}

func (m *Map) Delete(key Type) {
	for index, k := range m.Keys {
		if k.GetString().String() == key.GetString().String() {
			m.Keys = append(m.Keys[:index], m.Keys[index+1:]...)
			m.Values = append(m.Values[:index], m.Values[index+1:]...)
			return
		}
	}
}

func (m *Map) Sub(value Type) (Type, error) {
	switch value.(type) {
	case *Var:
		value = value.(*Var).Value
	}
	switch value.(type) {
	case *Map:
		if value.GetType() != m.Typ && m.TypKey != "string" && m.TypVal != "string" {
			return nil, errors.New("cannot subtract " + value.String() + " from map")
		}
		for _, v := range value.(*Map).Keys {
			value.(*Map).Delete(v)
		}
	case *Any:
		return m.Sub(value.(*Any).Value)
	}
	return m, nil
}

func (m *Map) Mul(value Type) (Type, error) {
	return nil, errors.New("cannot multiply " + value.String() + " with map")
}

func (m *Map) Div(value Type) (Type, error) {
	return nil, errors.New("cannot divide map by " + value.String())
}

func (m *Map) Mod(value Type) (Type, error) {
	return nil, errors.New("cannot get remainder of map by " + value.String())
}

func (m *Map) DivEc(value Type) (Type, error) {
	return nil, errors.New("cannot get quotient of map by " + value.String())
}

// TODO add case var ?
func (m *Map) Eq(value Type) (Type, error) {
	switch value.(type) {
	case *Var:
		value = value.(*Var).Value
	}
	switch value.(type) {
	case *Map:
		if len(m.Keys) != len(value.(*Map).Keys) {
			return Bool(false), nil
		}
		for index, _ := range m.Keys {
			if v, err := m.Values[index].Eq(value.(*Map).Values[index]); err == nil && v.(Bool) == false {
				return Bool(false), nil
			}
		}
		return Bool(true), nil
	case *Any:
		return m.Eq(value.(*Any).Value)
	default:
		return Bool(false), errors.New("cannot compare map of type " + m.GetType() + " with " + value.String())
	}
}

func (m *Map) NotEq(value Type) (Type, error) {
	v, err := m.Eq(value)
	if err != nil {
		return nil, err
	}
	return !v.(Bool), nil
}

func (m *Map) And(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) Or(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) Not() (Type, error) {
	return nil, errors.New("cannot \"not\" map")
}

func (m *Map) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) Append(other Type) (Type, error) {
	return nil, errors.New("cannot compare map with " + other.String())
}

func (m *Map) IsNull() bool {
	return false
}

func IsMap(typ string) bool {
	// exemple of type good : map[string]int
	// exemple of good type : map[[]int]int
	// exemple of bad type : []map[int]int
	if len(typ) < 4 {
		return false
	}
	if typ[:4] != "map[" {
		return false
	}
	return true
}

func GetTypeOfKeyAndValue(v string) (string, string) {
	val := ""
	count := 0
	for i := 4; i < len(v); i++ {
		if v[i] == '[' {
			count++
		}
		if v[i] == ']' {
			count--
		}
		if count == 0 && v[i] == ']' {
			val = v[i+1:]
			break
		}
	}
	return v[4 : len(v)-len(val)-1], val
}

func (m *Map) GetKeyTypes() string {
	return m.TypKey
}

func (m *Map) GetValueTypes() string {
	return m.TypVal
}

func (m *Map) GetSize() int {
	return utils.Sizeof(m)
}

func (m *Map) Len() (int, error) {
	return len(m.Keys), nil
}
