package eclaType

import (
	"errors"
	"fmt"
)

// NewList creates ?
/* func NewList(value string) List {
	result, _ := strconv.Atoi(value)
	return Int(result)
} */

type List []Type

// GetValue returns the value of the list
func (l List) GetValue() any {
	return l
}

// GetString returns the string of list
func (l List) GetString() String {
	return String(fmt.Sprint(l))
}

// GetType returns the type List
func (l List) GetType() String {
	return "List"
}

// Add adds two Type objects  compatible with List
func (l List) Add(other Type) (Type, error) {
	switch other.(type) {
	case List:
		return append(l, other.(List)...), nil
	case String:
		return l.GetString() + other.GetString(), nil
	}
	return nil, fmt.Errorf("cannot add %s to list", other.GetString())
}

// Sub returns errors because you cannot subtract lists
func (l List) Sub(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot subtract from list")
}

// Mod returns errors because you cannot mod lists
func (l List) Mod(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot mod list")
}

// Mul if other is Int , return n * List
func (l List) Mul(other Type) (Type, error) {
	switch other.(type) {
	case Int:
		result := List{}
		for i := 0; i < int(other.(Int)); i++ {
			result = append(result, l...)
		}
		return result, nil
	}
	return nil, fmt.Errorf("cannot multiply list by %s", other.GetString())
}

// Div returns errors because you cannot divide lists
func (l List) Div(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot divide list")
}

// DivEc returns error because you cannot div ec lists
func (l List) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by list")
}

// Eq returns true if two Type objects are equal
func (l List) Eq(other Type) (Type, error) {
	switch other.(type) {
	case List:
		if len(l) != len(other.(List)) {
			return Bool(false), nil
		}
		for i, v := range l {
			if v != other.(List)[i] {
				return Bool(false), nil
			}
		}
		return Bool(true), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// NotEq returns true if two Type objects are not equal
func (l List) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case List:
		if len(l) != len(other.(List)) {
			return Bool(true), nil
		}
		for i, v := range l {
			if v != other.(List)[i] {
				return Bool(true), nil
			}
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// Gt returns true if the first Type object is greater than the second
func (l List) Gt(other Type) (Type, error) {
	switch other.(type) {
	case List:
		if len(l) > len(other.(List)) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// GtEq returns true if the first Type object is greater than or equal the second
func (l List) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case List:
		if len(l) >= len(other.(List)) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// Lw returns true if the first Type object is lower than the second
func (l List) Lw(other Type) (Type, error) {
	switch other.(type) {
	case List:
		if len(l) < len(other.(List)) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// LwEq returns true if the first Type object is lower than or equal the second
func (l List) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case List:
		if len(l) <= len(other.(List)) {
			return Bool(true), nil
		}
		return Bool(false), nil
	}
	return nil, errors.New(string("cannot compare list to " + other.GetString()))
}

// And returns errors
func (l List) And(other Type) (Type, error) {
	return nil, errors.New("cannot and list")
}

// Or returns errors
func (l List) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or list")
}

// Not returns errors
func (l List) Not(other Type) (Type, error) {
	return nil, errors.New("cannot opposite list")
}
