package parser

import "github.com/tot0p/Ecla/lexer"

// Literal is a struct that defines a literal value for all types
type Literal struct {
	Token lexer.Token
	Type  string
	Value string
}

func (l Literal) StartPos() int {
	return l.Token.Position
}

func (l Literal) EndPos() int {
	return l.Token.Position
}
func (l Literal) precedence() int {
	return TokenPrecedence(l.Token)
}

func (l Literal) String() string {
	//TODO: Implement method
	return ""
}

func (l Literal) exprNode() {}

// BinaryExpr is a struct that defines a binary operation between two expressions
type BinaryExpr struct {
	LeftExpr  Expr
	Operator  lexer.Token
	RightExpr Expr
}

func (b BinaryExpr) StartPos() int {
	return b.LeftExpr.StartPos()
}

func (b BinaryExpr) EndPos() int {
	return b.RightExpr.EndPos()
}

func (b BinaryExpr) precedence() int {
	return TokenPrecedence(b.Operator)
}

func (b BinaryExpr) String() string {
	//TODO: Implement method
	return ""
}

func (b BinaryExpr) exprNode() {}

// UnaryExpr is a struct that defines a unary operation on an expression
type UnaryExpr struct {
	Operator  lexer.Token
	RightExpr Expr
}

func (u UnaryExpr) StartPos() int {
	return u.Operator.Position
}

func (u UnaryExpr) EndPos() int {
	return u.RightExpr.EndPos()
}

func (u UnaryExpr) precedence() int {
	return TokenPrecedence(u.Operator)
}

func (u UnaryExpr) String() string {
	//TODO: Implement method
	return ""
}

func (u UnaryExpr) exprNode() {}

// ParenExpr is a struct that defines a parenthesized expression
type ParenExpr struct {
	Lparen     lexer.Token
	Expression Expr
	Rparen     lexer.Token
}

func (p ParenExpr) StartPos() int {
	return p.Lparen.Position
}

func (p ParenExpr) EndPos() int {
	return p.Rparen.Position
}

func (p ParenExpr) precedence() int {
	return HighestPrecedence
}

func (p ParenExpr) String() string {
	//TODO: Implement method
	return ""
}

func (p ParenExpr) exprNode() {}

type PrintStmt struct {
	PrintToken lexer.Token
	Lparen     lexer.Token
	Rparen     lexer.Token
	Expression Expr
}

func (p PrintStmt) StartPos() int {
	return p.PrintToken.Position
}

func (p PrintStmt) EndPos() int {
	return p.Rparen.Position
}

func (p PrintStmt) String() string {
	//TODO: Implement method
	return ""
}

func (p PrintStmt) stmtNode() {}
