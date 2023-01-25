package eclaKeyWord

import (
	"fmt"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/parser"
)

type Function struct {
	Name   string
	Args   map[string]string
	Body   []parser.Node
	Return []string
}

func NewFunction(name string, args map[string]string, body []parser.Node, ret []string) *Function {
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
	for name, arg := range f.Args {
		elem := args[i]
		switch elem.(type) {
		case *eclaType.Var:
			elem = elem.(*eclaType.Var).Value
		}
		tp := elem.GetType()
		if tp == "list" {
			switch elem.(type) {
			case *eclaType.List:
				tp = elem.(*eclaType.List).GetFullType()
			}
		}
		if tp != arg {
			return false, nil
		}
		v, err := eclaType.NewVar(name, tp, elem)
		if err != nil {
			panic(err)
		}
		argsType[name] = v
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
		if tp == "list" {
			switch elem.(type) {
			case *eclaType.List:
				tp = elem.(*eclaType.List).GetFullType()
			}
		}
		if tp != r {
			return false
		}
		i++
	}
	return true
}
