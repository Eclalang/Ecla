package eclaType

import "fmt"

type List []Type

// GetValue returns the value of the list
func (l List) GetValue() any {
	return l
}

// GetString returns the string of list
func (l List) GetString() String {
	return String(fmt.Sprint(l))
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
