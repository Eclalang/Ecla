package eclaType

import "errors"

// NewChar creates a new Char
func NewChar(value string) Char {
	result := []rune(value)
	return Char(result[0])
}

type Char rune

// GetValue returns the value of the char
func (c Char) GetValue() any {
	return c
}

// SetValue returns an error
func (c Char) SetValue(value any) error {
	return errors.New("cannot set value to char")
}

func (c Char) String() string {
	return string(c)
}

// GetString returns the string representation of the char
func (c Char) GetString() String {
	return String(c)
}

// GetType returns the type Char
func (Char) GetType() string {
	return "char"
}

// GetIndex returns an error
func (Char) GetIndex(index Type) (*Type, error) {
	return nil, errors.New("cannot get index from char")
}

// Add adds two Type objects compatible with Char
func (c Char) Add(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(int(c) + other.GetValue().(int)), nil
	case Char:
		return c + other.(Char), nil
	case String:
		return c.GetString() + other.GetString(), nil
	default:
		return nil, errors.New("cannot add " + string(other.GetString()) + " to char")
	}
}

// Sub subtracts two Type objects compatible with Char
func (c Char) Sub(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(int(c) - other.GetValue().(int)), nil
	case Char:
		return c - other.(Char), nil
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from int")
	}
}

// Mul multiplies two Type objects compatible with Char
func (c Char) Mul(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(int(c) * other.GetValue().(int)), nil
	case Char:
		return c * other.(Char), nil
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from int")
	}
}

// Div divides two Type objects compatible with Char
func (c Char) Div(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(int(c) / other.GetValue().(int)), nil
	case Char:
		return c / other.(Char), nil
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from int")
	}
}

// Mod returns the remainder of the division of two Type objects compatible with Char
func (c Char) Mod(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Int:
		return Char(int(c) % other.GetValue().(int)), nil
	case Char:
		return c % other.(Char), nil
	default:
		return nil, errors.New("cannot subtract " + string(other.GetString()) + " from int")
	}
}

// DivEc divides two Type objects compatible with Char
func (c Char) DivEc(other Type) (Type, error) {
	return c.Div(other)
}

// Eq returns true if two Type objects are equal
func (c Char) Eq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c == other.(Char)), nil
	case Int:
		return Bool(Int(c) == other.(Int)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// NotEq returns true if two Type objects are not equal
func (c Char) NotEq(other Type) (Type, error) {
	b, err := c.Eq(other)
	switch b.(type) {
	case Bool:
		b = !b.(Bool)
	}
	return b, err
}

// Gt returns true if the first Type object is greater than the second
func (c Char) Gt(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c > other.(Char)), nil
	case Int:
		return Bool(Int(c) > other.(Int)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// GtEq returns true if the first Type object is greater than or equal to the second
func (c Char) GtEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c >= other.(Char)), nil
	case Int:
		return Bool(Int(c) >= other.(Int)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// Lw returns true if the first Type object is lower than the second
func (c Char) Lw(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c < other.(Char)), nil
	case Int:
		return Bool(Int(c) < other.(Int)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// LwEq returns true if the first Type object is lower than or equal to the second
func (c Char) LwEq(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return Bool(c <= other.(Char)), nil
	case Int:
		return Bool(Int(c) <= other.(Int)), nil
	default:
		return nil, errors.New("cannot compare " + string(other.GetString()) + " to int")
	}
}

// And returns an error
func (Char) And(other Type) (Type, error) {
	return nil, errors.New("cannot and char")
}

// Or returns an error
func (Char) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or char")
}

// Not returns an error
func (Char) Not() (Type, error) {
	return nil, errors.New("cannot opposite int")
}

func (Char) IsNull() bool {
	return false
}

func (c Char) Append(other Type) (Type, error) {
	switch other.(type) {
	case *Var:
		other = other.(*Var).Value
	}
	switch other.(type) {
	case Char:
		return c.GetString() + other.GetString(), nil
	case String:
		return c.GetString() + other.GetString(), nil
	default:
		return nil, errors.New("cannot add " + string(other.GetString()) + " to char")
	}
}
