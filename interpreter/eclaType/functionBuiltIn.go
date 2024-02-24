package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"strconv"
)

type FunctionBuiltIn struct {
	Name string
	f    func([]Type) ([]Type, error)
}

func (f *FunctionBuiltIn) Call(args []Type) ([]Type, error) {
	return f.f(args)
}

// implement Type interface

func (f *FunctionBuiltIn) GetValue() any {
	return f
}

func (f *FunctionBuiltIn) SetValue(value any) error {
	return errors.New("cannot set value of function")
}

func (f *FunctionBuiltIn) String() string {
	return "function"
}

func (f *FunctionBuiltIn) GetString() String {
	return "function"
}

func (f *FunctionBuiltIn) GetType() string {
	//TODO Change this
	return "function()"
}

func (f *FunctionBuiltIn) GetIndex(number Type) (*Type, error) {
	return nil, errors.New("cannot get index of function")
}

func (f *FunctionBuiltIn) Add(other Type) (Type, error) {
	return nil, errors.New("cannot add function")
}

func (f *FunctionBuiltIn) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract function")
}

func (f *FunctionBuiltIn) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply function")
}

func (f *FunctionBuiltIn) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide function")
}

func (f *FunctionBuiltIn) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod function")
}

func (f *FunctionBuiltIn) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by function")
}

func (f *FunctionBuiltIn) Eq(other Type) (Type, error) {
	return nil, errors.New("cannot eq function")
}

func (f *FunctionBuiltIn) NotEq(other Type) (Type, error) {
	return nil, errors.New("cannot notEq function")
}

func (f *FunctionBuiltIn) And(other Type) (Type, error) {
	return nil, errors.New("cannot and function")
}

func (f *FunctionBuiltIn) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or function")
}

func (f *FunctionBuiltIn) Not() (Type, error) {
	return nil, errors.New("cannot not function")
}

func (f *FunctionBuiltIn) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor function")
}

func (f *FunctionBuiltIn) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot gt function")
}

func (f *FunctionBuiltIn) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot gtEq function")
}

func (f *FunctionBuiltIn) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot lw function")
}

func (f *FunctionBuiltIn) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot lwEq function")
}

func (f *FunctionBuiltIn) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append function")
}

func (f *FunctionBuiltIn) IsNull() bool {
	return false
}

func (f *FunctionBuiltIn) GetSize() int {
	return utils.Sizeof(f)
}

func (f *FunctionBuiltIn) Len() (int, error) {
	return 0, errors.New("cannot get len of function")
}

func NewTypeOf() *FunctionBuiltIn {
	return &FunctionBuiltIn{
		Name: "typeOf",
		f: func(args []Type) ([]Type, error) {
			if len(args) != 1 {
				return nil, errors.New("typeOf function takes exactly one argument")
			}
			return []Type{String(args[0].GetType())}, nil
		},
	}
}

func NewSizeOf() *FunctionBuiltIn {
	return &FunctionBuiltIn{
		Name: "sizeOf",
		f: func(args []Type) ([]Type, error) {
			if len(args) != 1 {
				return nil, errors.New("sizeOf function takes exactly one argument")
			}
			return []Type{NewInt(strconv.Itoa(args[0].GetSize()))}, nil
		},
	}
}

func NewLen() *FunctionBuiltIn {
	return &FunctionBuiltIn{
		Name: "len",
		f: func(args []Type) ([]Type, error) {
			if len(args) != 1 {
				return nil, errors.New("len function takes exactly one argument")
			}
			l, err := args[0].Len()
			return []Type{NewInt(strconv.Itoa(l))}, err
		},
	}
}

func NewAppend() *FunctionBuiltIn {
	return &FunctionBuiltIn{
		Name: "append",
		f: func(args []Type) ([]Type, error) {
			if len(args) < 2 {
				return nil, errors.New("append function takes exactly two arguments")
			}
			var r Type = args[0]
			var err error
			for i := 1; i < len(args); i++ {
				r, err = r.Append(args[i])
				if err != nil {
					return nil, err
				}
			}
			return []Type{r}, err
		},
	}
}
