package eclaType

import (
	"errors"
	"fmt"
	"strings"
)

type Dict struct {
	Key       []Type
	Value     []Type
	KeyType   string
	ValueType string
}

// "{type:type}" is a dict
func NewDict(value string) Dict {
	value = strings.Trim(value, "{}")
	types := strings.Split(value, ":")
	return Dict{[]Type{}, []Type{}, types[0], types[1]}
}

// GetValue returns the value of the list
func (d Dict) GetValue() any {
	return d
}

func (d Dict) SetValue(value any) error {
	switch value.(type) {
	case Dict:
		d = value.(Dict)
		return nil
	}
	return errors.New("cannot set value of dict")
}

func (d Dict) String() string {
	var s string
	for _, v := range d.Value {
		s += string(v.GetString()) + ", "
	}
	return "{" + s + "}"
}

// GetString returns the string of list
func (d Dict) GetString() String {
	return String(fmt.Sprint(d))
}

// GetType returns the type List
func (d Dict) GetType() string {
	return "dict"
}

// Add adds two Type objects  compatible with List
func (d Dict) Add(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		return Dict{append(d.Key, other.(Dict).Key...), append(d.Value, other.(Dict).Value...), d.KeyType, d.ValueType}, nil
	case String:
		return d.GetString() + other.GetString(), nil
	}
	return nil, fmt.Errorf("cannot add %s to dict", other.GetString())
}

// Sub returns errors because you cannot subtract lists
func (d Dict) Sub(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot subtract from dict")
}

// Mod returns errors because you cannot mod lists
func (d Dict) Mod(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot mod dict")
}

// Mul if other is Int , return n * List
func (d Dict) Mul(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		result := Dict{}
		for i := 0; i < int(other.(Int)); i++ {
			result.Key = append(result.Key, d.Key...)
			result.Value = append(result.Value, d.Value...)
		}
		return result, nil
	}
	return nil, fmt.Errorf("cannot multiply dict by %s", other.GetString())
}

// Div returns errors because you cannot divide lists
func (d Dict) Div(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot divide dict")
}

// DivEc returns error because you cannot div ec lists
func (d Dict) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by dict")
}

// Eq returns true if two Type objects are equal
func (d Dict) Eq(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		if len(d.Value) != len(other.(Dict).Value) {
			return Bool(false), nil
		}
		for i, v := range d.Value {
			if v != other.(Dict).Value[i] {
				return Bool(false), nil
			}
		}
		return Bool(true), nil
	}
	return nil, errors.New(string("cannot compare dict to " + other.GetString()))
}

// NotEq returns true if two Type objects are not equal
func (d Dict) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		if len(d.Value) != len(other.(Dict).Value) {
			return Bool(true), nil
		}
		for i, v := range d.Value {
			if v != other.(Dict).Value[i] {
				return Bool(true), nil
			}
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare dict to " + other.GetString()))
}

// Gt returns true if the first Type object is greater than the second
func (d Dict) Gt(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		if len(d.Value) > len(other.(Dict).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare dict to " + other.GetString()))
}

// GtEq returns true if the first Type object is greater than or equal the second
func (d Dict) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		if len(d.Value) >= len(other.(Dict).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare dict to " + other.GetString()))
}

// Lw returns true if the first Type object is lower than the second
func (d Dict) Lw(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		if len(d.Value) < len(other.(Dict).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare dict to " + other.GetString()))
}

// LwEq returns true if the first Type object is lower than or equal the second
func (d Dict) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case Dict:
		if len(d.Value) <= len(other.(Dict).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare dict to " + other.GetString()))
}

// And returns errors
func (d Dict) And(other Type) (Type, error) {
	return nil, errors.New("cannot and dict")
}

// Or returns errors
func (d Dict) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or dict")
}

// Not returns errors
func (d Dict) Not() (Type, error) {
	return nil, errors.New("cannot opposite dict")
}

// Append
func (d Dict) Append(other Type) (Type, error) {
	if other.GetType() == "dict" {
		return Dict{append(d.Key, other.(Dict).Key...), append(d.Value, other.(Dict).Value...), d.KeyType, d.ValueType}, nil
	} else {
		if d.ValueType == other.GetType() {
			d.Key = append(d.Key, other.GetString())
			d.Value = append(d.Value, other)
			return d, nil
		}
	}
	return nil, errors.New("cannot append dict")
}
