package eclaKeyWord

import (
	"fmt"
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
