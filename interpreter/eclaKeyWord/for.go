package eclaKeyWord

import (
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"github.com/tot0p/Ecla/lexer"
	"github.com/tot0p/Ecla/parser"
)

type ForRange struct {
	Scope      []eclaType.Type
	RangeExpr  parser.Expr
	KeyToken   lexer.Token
	ValueToken lexer.Token
	Body       []parser.Node
}

type ForI struct {
	Scope     []eclaType.Type
	Body      []parser.Node
	Condition parser.Expr
	Post      parser.Stmt
}

func NewForRange(scope []eclaType.Type, rangeExpr parser.Expr, keyToken lexer.Token, valueToken lexer.Token, body []parser.Node) ForRange {
	return ForRange{scope, rangeExpr, keyToken, valueToken, body}
}

func NewForI(scope []eclaType.Type, body []parser.Node, condition parser.Expr, post parser.Stmt) ForI {
	return ForI{scope, body, condition, post}
}
