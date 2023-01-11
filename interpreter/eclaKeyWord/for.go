package eclaKeyWord

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

type ForRange struct {
	Scope     []eclaType.Type
	Body      []parser.Node
	Condition parser.Expr
	Post      parser.Stmt
}

type ForI struct {
	Scope    []eclaType.Type
	InitDecl parser.Decl
	Body     []parser.Node
}

type For interface {
	Execute() error
}

func NewFor(f parser.ForStmt) For {
	tokenEmpty := lexer.Token{}
	if f.RangeToken != tokenEmpty {
		//TODO
		return &ForI{Scope: []eclaType.Type{}, InitDecl: f.InitDecl, Body: f.Body}
	} else {
		return &ForRange{Scope: []eclaType.Type{}}
	}
}

func (f *ForRange) Execute() error {
	return nil
}

func (f *ForI) Execute() error {
	return nil
}
