package parser

import "github.com/Eclalang/Ecla/lexer"

type AnonymousFunctionCallExpr struct {
	AnonymousFunction AnonymousFunctionExpr
	LeftParen         lexer.Token
	RightParen        lexer.Token
	Args              []Expr
}

func (f AnonymousFunctionCallExpr) StartPos() int {
	return f.AnonymousFunction.StartPos()
}

func (f AnonymousFunctionCallExpr) EndPos() int {
	return f.RightParen.Position
}

func (f AnonymousFunctionCallExpr) StartLine() int {
	return f.AnonymousFunction.StartLine()
}

func (f AnonymousFunctionCallExpr) EndLine() int {
	return f.RightParen.Line
}

func (f AnonymousFunctionCallExpr) precedence() int {
	return HighestPrecedence
}

func (f AnonymousFunctionCallExpr) exprNode() {}

type AnonymousFunctionExpr struct {
	FunctionToken lexer.Token
	Prototype     FunctionPrototype
	Body          []Node
}

func (f AnonymousFunctionExpr) StartPos() int {
	return f.FunctionToken.Position
}

func (f AnonymousFunctionExpr) EndPos() int {
	return f.Prototype.RightBrace.Position
}

func (f AnonymousFunctionExpr) StartLine() int {
	return f.FunctionToken.Line
}

func (f AnonymousFunctionExpr) EndLine() int {
	return f.Prototype.RightBrace.Line
}

func (f AnonymousFunctionExpr) precedence() int {
	return HighestPrecedence
}

func (f AnonymousFunctionExpr) exprNode() {}

type ArrayLiteral struct {
	LBRACKET lexer.Token
	Values   []Expr
	RBRACKET lexer.Token
}

func (p ArrayLiteral) StartPos() int {
	return p.LBRACKET.Position
}

func (p ArrayLiteral) EndPos() int {
	return p.RBRACKET.Position
}

func (p ArrayLiteral) StartLine() int {
	return p.LBRACKET.Line
}

func (p ArrayLiteral) EndLine() int {
	return p.RBRACKET.Line
}

func (p ArrayLiteral) precedence() int {
	return HighestPrecedence
}

func (p ArrayLiteral) exprNode() {}

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

func (b BinaryExpr) StartLine() int {
	return b.LeftExpr.StartLine()
}

func (b BinaryExpr) EndLine() int {
	return b.RightExpr.EndLine()
}

func (b BinaryExpr) precedence() int {
	return TokenPrecedence(b.Operator)
}

func (b BinaryExpr) exprNode() {}

type FunctionCallExpr struct {
	FunctionCallToken lexer.Token
	Name              string
	LeftParen         lexer.Token
	RightParen        lexer.Token
	Args              []Expr
}

func (f FunctionCallExpr) StartPos() int {
	return f.FunctionCallToken.Position
}

func (f FunctionCallExpr) EndPos() int {
	return f.RightParen.Position
}

func (f FunctionCallExpr) StartLine() int {
	return f.FunctionCallToken.Line
}

func (f FunctionCallExpr) EndLine() int {
	return f.RightParen.Line
}

func (f FunctionCallExpr) precedence() int {
	return HighestPrecedence
}

func (f FunctionCallExpr) exprNode() {}

type IndexableAccessExpr struct {
	VariableToken lexer.Token
	VariableName  string
	Indexes       []Expr
	LastBracket   lexer.Token
}

func (a IndexableAccessExpr) StartPos() int {
	return a.VariableToken.Position
}

func (a IndexableAccessExpr) EndPos() int {
	return a.LastBracket.Position
}

func (a IndexableAccessExpr) StartLine() int {
	return a.VariableToken.Line
}

func (a IndexableAccessExpr) EndLine() int {
	return a.LastBracket.Line
}

func (a IndexableAccessExpr) precedence() int {
	return HighestPrecedence
}

func (a IndexableAccessExpr) exprNode() {}

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

func (l Literal) StartLine() int {
	return l.Token.Line
}

func (l Literal) EndLine() int {
	return l.Token.Line
}

func (l Literal) precedence() int {
	return TokenPrecedence(l.Token)
}

func (l Literal) exprNode() {}

type MapLiteral struct {
	LBRACE lexer.Token
	Keys   []Expr
	Values []Expr
	RBRACE lexer.Token
}

func (m MapLiteral) StartPos() int {
	return m.LBRACE.Position
}

func (m MapLiteral) EndPos() int {
	return m.RBRACE.Position
}

func (m MapLiteral) StartLine() int {
	return m.LBRACE.Line
}

func (m MapLiteral) EndLine() int {
	return m.RBRACE.Line
}

func (m MapLiteral) precedence() int {
	return HighestPrecedence
}

func (m MapLiteral) exprNode() {}

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

func (p ParenExpr) StartLine() int {
	return p.Lparen.Line
}

func (p ParenExpr) EndLine() int {
	return p.Rparen.Line
}

func (p ParenExpr) precedence() int {
	return HighestPrecedence
}

func (p ParenExpr) exprNode() {}

type SelectorExpr struct {
	Field lexer.Token
	Expr  Expr
	Sel   Expr
}

func (s SelectorExpr) StartPos() int {
	return s.Field.Position
}

func (s SelectorExpr) EndPos() int {
	return s.Sel.EndPos()
}

func (s SelectorExpr) StartLine() int {
	return s.Field.Line
}

func (s SelectorExpr) EndLine() int {
	return s.Sel.EndLine()
}

func (s SelectorExpr) precedence() int {
	return HighestPrecedence
}

func (s SelectorExpr) exprNode() {}

type StructInstantiationExpr struct {
	StructNameToken lexer.Token
	Name            string
	LeftBrace       lexer.Token
	RightBrace      lexer.Token
	Args            []Expr
}

func (s StructInstantiationExpr) StartPos() int {
	return s.StructNameToken.Position
}

func (s StructInstantiationExpr) EndPos() int {
	return s.RightBrace.Position
}

func (s StructInstantiationExpr) StartLine() int {
	return s.StructNameToken.Line
}

func (s StructInstantiationExpr) EndLine() int {
	return s.RightBrace.Line
}

func (s StructInstantiationExpr) precedence() int {
	return HighestPrecedence
}

func (s StructInstantiationExpr) exprNode() {}

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

func (u UnaryExpr) StartLine() int {
	return u.Operator.Line
}

func (u UnaryExpr) EndLine() int {
	return u.RightExpr.EndLine()
}

func (u UnaryExpr) precedence() int {
	return TokenPrecedence(u.Operator)
}

func (u UnaryExpr) exprNode() {}
