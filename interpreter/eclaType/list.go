package eclaType

import (
	"errors"
	"fmt"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"github.com/Eclalang/Ecla/parser"
)

func NewList(t string) (Type, error) {
	return &List{[]Type{}, t}, nil
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
		// Check type of list
		for _, v := range v.([]Type) {
			if v.GetType() != l.Typ[2:] {
				return errors.New(fmt.Sprintf("cannot set value %s to list of type %s", v.GetType(), l.Typ))
			}
		}
		l.Value = v.([]Type)
		return nil
	case *List:
		if l.GetValueType() != v.(*List).GetValueType() {
			return errors.New(fmt.Sprintf("cannot set value %s to list of type %s", v.(*List).Typ, l.Typ))
		}
		t := v.(*List)
		*l = *t
		return nil
	}
	return fmt.Errorf("cannot set value %s of list", v)
}

func (l *List) String() string {
	var s string
	for index, v := range l.Value {
		if index == len(l.Value)-1 {
			s += string(v.GetString())
		} else {
			s += string(v.GetString()) + ", "
		}
	}
	return "[" + s + "]"
}

// GetString returns the string of list
func (l *List) GetString() String {
	return String(fmt.Sprint(l))
}

// GetType returns the type List
func (l *List) GetType() string {
	return l.Typ
}

func (l *List) SetType(other string) {
	l.Typ = other
}

func (l *List) GetIndex(index Type) (*Type, error) {

	if index.GetType() == "int" {
		ind := int(index.GetValue().(Int))
		if ind >= len(l.Value) || ind < 0 {
			return nil, errors.New("index out of range")
		}
		return &l.Value[ind], nil
	}
	return nil, errors.New("index must be an int")

}

// Add adds two Type objects  compatible with List
func (l *List) Add(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case *List:
		if l.Typ == other.(*List).Typ {
			return &List{append(l.Value, other.(*List).Value...), l.Typ}, nil
		}
		if l.GetValueType() == parser.Int && other.(*List).GetValueType() == parser.Char {
			tmpList := l
			for _, elem := range other.(*List).Value {
				tmpList.Value = append(tmpList.Value, elem.(Char).GetValueAsInt())
			}
			return tmpList, nil
		}
		if l.GetValueType() == parser.Char && other.(*List).GetValueType() == parser.Int {
			tmpList := l
			for _, elem := range other.(*List).Value {
				i := int(elem.(Int))
				var err error = nil
				c, err := NewChar(string(rune(i)))
				if err != nil {
					return nil, err
				}
				tmpList.Value = append(tmpList.Value, c)
			}
			return tmpList, nil
		}
		return nil, errors.New("cannot add list of " + l.GetValueType() + " with list of " + other.(*List).GetValueType())
	case String:
		return l.GetString() + other.GetString(), nil
	case *Any:
		return l.Add(other.(*Any).Value)
	}
	return nil, fmt.Errorf("cannot add %s to list", other.GetString())
}

// Sub returns errors because you cannot subtract lists
func (l *List) Sub(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot subtract " + other.String() + " from list")
}

// Mod returns errors because you cannot mod lists
func (l *List) Mod(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot get remainder of list by " + other.String())
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
	case *Any:
		return l.Mul(other.(*Any).Value)
	}
	return nil, fmt.Errorf("cannot multiply list by %s", other.GetString())
}

// Div returns errors because you cannot divide lists
func (l *List) Div(other Type) (Type, error) {
	return nil, fmt.Errorf("cannot divide list by " + other.String())
}

// DivEc returns error because you cannot div ec lists
func (l *List) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot get quotient of list by " + other.String())
}

// Eq returns true if two Type objects are equal
func (l *List) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare list of " + l.GetValueType() + " with list of " + other.(*List).GetValueType())
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
	case *Any:
		return l.Eq(other.(*Any).Value)
	}
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// NotEq returns true if two Type objects are not equal
func (l *List) NotEq(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare list of " + l.GetValueType() + " with list of " + other.(*List).GetValueType())
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
	case *Any:
		return l.NotEq(other.(*Any).Value)
	}
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// Gt returns true if the first Type object is greater than the second
func (l *List) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare list of " + l.GetValueType() + " with list of " + other.(*List).GetValueType())
		}
		if len(l.Value) > len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	case *Any:
		return l.Gt(other.(*Any).Value)
	}
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// GtEq returns true if the first Type object is greater than or equal the second
func (l *List) GtEq(other Type) (Type, error) {

	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare list of " + l.GetValueType() + " with list of " + other.(*List).GetValueType())
		}
		if len(l.Value) >= len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	case *Any:
		return l.GtEq(other.(*Any).Value)
	}
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// Lw returns true if the first Type object is lower than the second
func (l *List) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare list of " + l.GetValueType() + " with list of " + other.(*List).GetValueType())
		}
		if len(l.Value) < len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	case *Any:
		return l.Lw(other.(*Any).Value)
	}
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// LwEq returns true if the first Type object is lower than or equal the second
func (l *List) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if l.Typ != other.(*List).Typ {
			return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.(*List).GetValueType())
		}
		if len(l.Value) <= len(other.(*List).Value) {
			return Bool(true), nil
		}
		return Bool(false), nil
	case *Any:
		return l.LwEq(other.(*Any).Value)
	}
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with list of " + other.GetType())
}

// And returns errors
func (l *List) And(other Type) (Type, error) {
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// Or returns errors
func (l *List) Or(other Type) (Type, error) {
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// Not returns errors
func (l *List) Not() (Type, error) {
	return nil, errors.New("cannot \"not\" list")
}

// Xor returns errors
func (l *List) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot compare list of " + l.GetValueType() + " with " + other.GetType())
}

// Append to list
func (l *List) Append(other Type) (Type, error) {
	switch other.(type) {
	case *List:
		if other.(*List).Typ == l.GetValueType() {
			l.Value = append(l.Value, other.(*List))
			return l, nil
		}
		if l.GetValueType() == parser.Int && other.(*List).GetValueType() == parser.Char ||
			l.GetValueType() == parser.Char && other.(*List).GetType() == parser.Int {
			return l.Add(other)
		}
	case *Any:
		return l.Append(other.(*Any).Value)
	default:
		if other.GetType() == l.GetValueType() {
			l.Value = append(l.Value, other)
			return l, nil
		}
	}
	return nil, errors.New("cannot append " + other.GetType() + " to list of " + l.GetValueType())
}

func (l *List) IsNull() bool {
	return false
}

func (l *List) GetValueType() string {
	return l.Typ[2:]
}

// utils Functions for lists trainmen

func CheckTypeOfList(l *List, t string) bool {
	for _, v := range l.Value {
		if v.GetType() != l.Typ {
			return false
		}
	}
	return true
}

func IsList(t string) bool {
	// via []int or []string [][]int ,string int map[string]int []map[string]int
	if !(len(t) <= 2) {
		if t[0] == '[' && t[1] == ']' {
			return true
		}
	}
	return false
}

func (l *List) GetSize() int {
	return utils.Sizeof(l)
}

func (l *List) Len() (int, error) {
	return len(l.Value), nil
}
