package eclaKeyWord

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/parser"
)

type For struct {
	Condition any
	Body      any
	Scope     []eclaType.Type
}

func NewFor(parser.ForStmt) *For {
	return &For{}
}

func (f *For) Execute() error {
	return nil
}
