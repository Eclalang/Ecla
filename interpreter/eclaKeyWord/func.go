package eclaKeyWord

import (
	"errors"
	"github.com/tot0p/Ecla/interpreter/eclaType"
)

type Function struct {
	Name  string
	Args  []eclaType.Type
	Scope []eclaType.Type
	Code  any
}

func (f *Function) GetType() string {
	return "Function"
}

func (f *Function) NewFunction(name string, args []eclaType.Type, scope []eclaType.Type, code any) *Function {
	return &Function{Name: name, Args: args, Scope: scope, Code: code}
}

func (f *Function) Call(args []eclaType.Type) error {
	//TODO EXECUTE CODE
	if len(args) < len(f.Args) {
		return errors.New("cannot call function " + f.Name + " too few arguments passed into")
	}
	if len(args) > len(f.Args) {
		return errors.New("cannot call function " + f.Name + " too many arguments passed into")
	}
	return nil
}
