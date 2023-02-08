package eclaKeyWord

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/parser"
)

type Function struct {
	Name   string
	Args   []parser.FunctionParams
	Body   []parser.Node
	Return []string
}

func NewFunction(name string, args []parser.FunctionParams, body []parser.Node, ret []string) *Function {
	return &Function{
		Name:   name,
		Args:   args,
		Body:   body,
		Return: ret,
	}
}

func (f *Function) String() string {
	return fmt.Sprintf("Function{Name: %s, Args: %s, Body: %s, Return: %s}", f.Name, f.Args, f.Body, f.Return)
}

func (f *Function) TypeAndNumberOfArgsIsCorrect(args []eclaType.Type) (bool, map[string]*eclaType.Var) {
	if len(f.Args) != len(args) {
		return false, nil
	}
	var i int = 0
	var argsType = make(map[string]*eclaType.Var)
	for _, arg := range f.Args {
		paramName := arg.Name
		paramType := arg.Type
		elem := args[i]
		switch elem.(type) {
		case *eclaType.Var:
			elem = elem.(*eclaType.Var).Value
		}
		tp := elem.GetType()
		if tp != paramType {
			return false, nil
		}
		v, err := eclaType.NewVar(paramName, tp, elem)
		if err != nil {
			panic(err)
		}
		argsType[paramName] = v
		i++
	}
	return true, argsType
}

func (f *Function) CheckReturn(ret []eclaType.Type) bool {
	if len(f.Return) != len(ret) {
		return false
	}
	var i int = 0
	for _, r := range f.Return {
		elem := ret[i]
		switch elem.(type) {
		case *eclaType.Var:
			elem = elem.(*eclaType.Var).Value
		}
		tp := elem.GetType()
		if tp != r {
			return false
		}
		i++
	}
	return true
}
