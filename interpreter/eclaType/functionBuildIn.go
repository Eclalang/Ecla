package eclaType

import (
	"errors"
	"github.com/Eclalang/Ecla/interpreter/utils"
	"strconv"
)

type FunctionBuildIn struct {
	Name string
	f    func([]Type) ([]Type, error)
}

func (f *FunctionBuildIn) Call(args []Type) ([]Type, error) {
	return f.f(args)
}

// implement Type interface

func (f *FunctionBuildIn) GetValue() any {
	return f
}

func (f *FunctionBuildIn) SetValue(value any) error {
	return errors.New("cannot set value of function")
}

func (f *FunctionBuildIn) String() string {
	return "function"
}

func (f *FunctionBuildIn) GetString() String {
	return "function"
}

func (f *FunctionBuildIn) GetType() string {
	//TODO Change this
	return "function()"
}

func (f *FunctionBuildIn) GetIndex(number Type) (*Type, error) {
	return nil, errors.New("cannot get index of function")
}

func (f *FunctionBuildIn) Add(other Type) (Type, error) {
	return nil, errors.New("cannot add function")
}

func (f *FunctionBuildIn) Sub(other Type) (Type, error) {
	return nil, errors.New("cannot subtract function")
}

func (f *FunctionBuildIn) Mul(other Type) (Type, error) {
	return nil, errors.New("cannot multiply function")
}

func (f *FunctionBuildIn) Div(other Type) (Type, error) {
	return nil, errors.New("cannot divide function")
}

func (f *FunctionBuildIn) Mod(other Type) (Type, error) {
	return nil, errors.New("cannot mod function")
}

func (f *FunctionBuildIn) DivEc(other Type) (Type, error) {
	return nil, errors.New("cannot divide ec by function")
}

func (f *FunctionBuildIn) Eq(other Type) (Type, error) {
	return nil, errors.New("cannot eq function")
}

func (f *FunctionBuildIn) NotEq(other Type) (Type, error) {
	return nil, errors.New("cannot notEq function")
}

func (f *FunctionBuildIn) And(other Type) (Type, error) {
	return nil, errors.New("cannot and function")
}

func (f *FunctionBuildIn) Or(other Type) (Type, error) {
	return nil, errors.New("cannot or function")
}

func (f *FunctionBuildIn) Not() (Type, error) {
	return nil, errors.New("cannot not function")
}

func (f *FunctionBuildIn) Xor(other Type) (Type, error) {
	return nil, errors.New("cannot xor function")
}

func (f *FunctionBuildIn) Gt(other Type) (Type, error) {
	return nil, errors.New("cannot gt function")
}

func (f *FunctionBuildIn) GtEq(other Type) (Type, error) {
	return nil, errors.New("cannot gtEq function")
}

func (f *FunctionBuildIn) Lw(other Type) (Type, error) {
	return nil, errors.New("cannot lw function")
}

func (f *FunctionBuildIn) LwEq(other Type) (Type, error) {
	return nil, errors.New("cannot lwEq function")
}

func (f *FunctionBuildIn) Append(other Type) (Type, error) {
	return nil, errors.New("cannot append function")
}

func (f *FunctionBuildIn) IsNull() bool {
	return false
}

func (f *FunctionBuildIn) GetSize() int {
	return utils.Sizeof(f)
}

func (f *FunctionBuildIn) Len() (int, error) {
	return 0, errors.New("cannot get len of function")
}

func NewTypeOf() *FunctionBuildIn {
	return &FunctionBuildIn{
		Name: "typeOf",
		f: func(args []Type) ([]Type, error) {
			if len(args) != 1 {
				return nil, errors.New("typeOf function takes exactly one argument")
			}
			return []Type{String(args[0].GetType())}, nil
		},
	}
}

func NewSizeOf() *FunctionBuildIn {
	return &FunctionBuildIn{
		Name: "sizeOf",
		f: func(args []Type) ([]Type, error) {
			if len(args) != 1 {
				return nil, errors.New("sizeOf function takes exactly one argument")
			}
			return []Type{NewInt(strconv.Itoa(args[0].GetSize()))}, nil
		},
	}
}

func NewLen() *FunctionBuildIn {
	return &FunctionBuildIn{
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

func NewAppend() *FunctionBuildIn {
	return &FunctionBuildIn{
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
