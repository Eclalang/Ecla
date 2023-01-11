package eclaKeyWord

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/parser"
)

type ForRange struct {
	Scope []eclaType.Type
	Body  []parser.Node
}

type ForI struct {
	Scope     []eclaType.Type
	Body      []parser.Node
	Condition parser.Expr
	Post      parser.Stmt
}

func NewForRange(scope []eclaType.Type, body []parser.Node) ForRange {
	return ForRange{scope, body}
}

func NewForI(scope []eclaType.Type, body []parser.Node, condition parser.Expr, post parser.Stmt) ForI {
	return ForI{scope, body, condition, post}
}
