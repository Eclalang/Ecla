package eclaType

import "errors"

// NewString returns a new String.
func NewString(value string) Type {
	return &String{Value: value}
}

type String struct {
	Value string
}

// GetValue returns the value of the String.
func (s *String) GetValue() any {
	return s.Value
}

// GetString returns the value of the String.
func (s *String) GetString() string {
	return s.Value
}

// ADD returns the concatenation of the two Type of Ecla.
func (s *String) ADD(other Type) (Type, error) {
	return &String{Value: s.Value + other.GetString()}, nil
}

// SUB returns error.
func (s *String) SUB(other Type) (Type, error) {
	return nil, errors.New("cannot subtract from string")
}

// MUL returns n * String if other is int else return error .
func (s *String) MUL(other Type) (Type, error) {
	vOther := other.GetValue()
	switch vOther.(type) {
	case int:
		result := s.Value
		for i := 0; i < vOther.(int); i++ {
			result += s.Value
		}
		return &String{Value: result}, nil
	default:
		return nil, errors.New("cannot multiply string by " + other.GetString())
	}

}

// DIV returns error.
func (s *String) DIV(other Type) (Type, error) {
	return nil, errors.New("cannot divide string")
}

// MOD returns error.
func (s *String) MOD(other Type) (Type, error) {
	return nil, errors.New("cannot mod string")
}
