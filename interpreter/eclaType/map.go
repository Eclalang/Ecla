package eclaType

import (
	"errors"
	"fmt"
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
				return errors.New("you cannot have a map with different value type")
			}
		}
		for _, v := range m.Keys {
			if v.GetType() != m.TypKey {
				return errors.New("you cannot have a map with different key type")
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
	default:
		return errors.New("cannot set value of map")
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

// GetType returns the type Map
func (m *Map) GetType() string {
	return m.Typ
}

func (m *Map) SetType(t string) {
	m.Typ = t
	m.TypKey, m.TypVal = GetTypeOfKeyAndValue(t)
}

func (m *Map) Set(key Type, value Type) {
	if key.GetType() != m.TypKey && m.TypKey != "string" {
		panic("key type not match")
	}
	if value.GetType() != m.TypVal && m.TypVal != "string" {
		panic("value type not match")
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
		return nil, errors.New("index type not match")
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
			return nil, errors.New("cannot add map with " + value.GetType())
		}

		for index, v := range value.(*Map).Keys {
			m.Set(v, value.(*Map).Values[index])
		}

		return m, nil
	}
	return nil, errors.New("cannot add map with " + value.GetType())
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
			return nil, errors.New("cannot add map with " + value.GetType())
		}
		for _, v := range value.(*Map).Keys {
			value.(*Map).Delete(v)
		}
	}
	return m, nil
}

func (m *Map) Mul(value Type) (Type, error) {
	return nil, errors.New("cannot mul map")
}

func (m *Map) Div(value Type) (Type, error) {
	return nil, errors.New("cannot div map")
}

func (m *Map) Mod(value Type) (Type, error) {
	return nil, errors.New("cannot mod map")
}

func (m *Map) DivEc(value Type) (Type, error) {
	return nil, errors.New("cannot divec map")
}
func (m *Map) DivMod(value Type) (Type, error) {
	return nil, errors.New("cannot divmod map")
}

func (m *Map) Eq(value Type) (Type, error) {
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
	default:
		return Bool(false), errors.New("cannot compare map with other type")
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
	return nil, errors.New("cannot and null")
}

func (m *Map) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or null")
}

func (m *Map) Not() (Type, error) {
	return nil, errors.New("cannot not null")
}

func (m *Map) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot gt null")
}

func (m *Map) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot gtEq null")
}

func (m *Map) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot lw null")
}

func (m *Map) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot lwEq null")
}

func (m *Map) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append null")
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
