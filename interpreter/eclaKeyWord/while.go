package eclaKeyWord

import "github.com/tot0p/Ecla/parser"

type While struct {
	Condition parser.Expr
	Body      []parser.Node
}

func NewWhile(condition parser.Expr, body []parser.Node) While {
	return While{condition, body}
}
