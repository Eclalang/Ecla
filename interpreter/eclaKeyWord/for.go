package eclaKeyWord

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

// ForRange is the for range statement.
type ForRange struct {
	Scope      []eclaType.Type
	RangeExpr  parser.Expr
	KeyToken   lexer.Token
	ValueToken lexer.Token
	Body       []parser.Node
}

// ForI is the for i statement.
type ForI struct {
	Scope     []eclaType.Type
	Body      []parser.Node
	Condition parser.Expr
	Post      parser.Stmt
}

// NewForRange returns a new for range statement.
func NewForRange(scope []eclaType.Type, rangeExpr parser.Expr, keyToken lexer.Token, valueToken lexer.Token, body []parser.Node) ForRange {
	return ForRange{scope, rangeExpr, keyToken, valueToken, body}
}

// NewForI returns a new for i statement.
func NewForI(scope []eclaType.Type, body []parser.Node, condition parser.Expr, post parser.Stmt) ForI {
	return ForI{scope, body, condition, post}
}
