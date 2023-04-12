package eclaKeyWord

import "github.com/tot0p/Ecla/parser"

// While is the while statement.
type While struct {
	Condition parser.Expr
	Body      []parser.Node
}

// NewWhile returns a new while statement.
func NewWhile(condition parser.Expr, body []parser.Node) While {
	return While{condition, body}
}
