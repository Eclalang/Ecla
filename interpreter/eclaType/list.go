package eclaType

import (
	"errors"
	"fmt"
	"strings"
)

func NewList(t string) (Type, error) {
	if !IsList(t) {
		return nil, errors.New("not a list")
	}

	return &List{[]Type{}, t[2:]}, nil
}

type List struct {
	Value []Type
	Typ   string
}

// GetValue returns the value of the list
func (l *List) GetValue() any {
	return l
}

// SetValue
func (l *List) SetValue(v any) error {
	switch v.(type) {
	case []Type:
		if l.Typ == v.([]Type)[0].GetType() {
			l.Value = v.([]Type)
			return nil
		}
	case *List:

		t := v.(*List)
		*l = *t
		return nil
	}
	return errors.New("cannot set value of list")
}

func (l *List) String() string {
	var s string
	for _, v := range l.Value {
		s += string(v.GetString()) + ", "
	}
	return "[" + s + "]"
}

// GetString returns the string of list
func (l *List) GetString() String {
	return String(fmt.Sprint(l))
}

// GetType returns the type List
func (l *List) GetType() string {
	return "list"
}

// Add adds two Type objects  compatible with List
func (l *List) Add(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ == other.(*List).Typ {
			return &List{append(l.Value, other.(*List).Value...), l.Typ}, nil
		}
		return nil, errors.New("cannot add lists of different types")
	case String:
		return l.GetString() + other.GetString(), nil
	}
	return nil, fmt.Errorf("cannot add %s to list", other.GetString())
}

// Sub returns errors because you cannot subtract lists
func (l *List) Sub(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot subtract from list")
}

// Mod returns errors because you cannot mod lists
func (l *List) Mod(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot mod list")
}

// Mul if other is Int , return n * List
func (l *List) Mul(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		result := List{[]Type{}, l.Typ}
		for i := 0; i < int(other.(Int)); i++ {
			result.Value = append(result.Value, l.Value...)
		}
		return &result, nil
	}
	return nil, fmt.Errorf("cannot multiply list by %s", other.GetString())
}

// Div returns errors because you cannot divide lists
func (l *List) Div(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot divide list")
}

// DivEc returns error because you cannot div ec lists
func (l *List) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by list")
}

// Eq returns true if two Type objects are equal
func (l *List) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare lists of different types")
		}
		if len(l.Value) != len(other.(*List).Value) {
			return Bool(false), nil
		}
		for i, v := range l.Value {
			if v != other.(*List).Value[i] {
				return Bool(false), nil
			}
		}
		return Bool(true), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// NotEq returns true if two Type objects are not equal
func (l *List) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare lists of different types")
		}
		if len(l.Value) != len(other.(*List).Value) {
			return Bool(true), nil
		}
		for i, v := range l.Value {
			if v != other.(*List).Value[i] {
				return Bool(true), nil
			}
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// Gt returns true if the first Type object is greater than the second
func (l *List) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare lists of different types")
		}
		if len(l.Value) > len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// GtEq returns true if the first Type object is greater than or equal the second
func (l *List) GtEq(other Type) (Type, error) {

	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare lists of different types")
		}
		if len(l.Value) >= len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// Lw returns true if the first Type object is lower than the second
func (l *List) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare lists of different types")
		}
		if len(l.Value) < len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// LwEq returns true if the first Type object is lower than or equal the second
func (l *List) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare lists of different types")
		}
		if len(l.Value) <= len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// And returns errors
func (l *List) And(other Type) (Type, error) {
	return nil, errors.New("cannot and list")
}

// Or returns errors
func (l *List) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or list")
}

// Not returns errors
func (l *List) Not() (Type, error) {
	return nil, errors.New("cannot opposite list")
}

func IsList(t string) bool {
	return strings.Contains(t, "[")
}

// append to list
func (l *List) Append(other Type) (Type, error) {
	if other.GetType() == "list" {
		if l.Typ == other.(*List).Typ {
			l.Value = append(l.Value, other.(*List).Value...)
			return l, nil
		}
	} else {
		if l.Typ == other.GetType() {
			l.Value = append(l.Value, other)
			return l, nil
		}
	}
	return nil, errors.New("cannot append to list")
}

// utils Functions for lists trainmen

func CheckTypeOfList(l *List, t string) bool {
	for _, v := range l.Value {
		if v.GetType() != t {
			return false
		}
	}
	return true
}
