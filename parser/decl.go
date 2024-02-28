package parser

import "github.com/Eclalang/Ecla/lexer"

type FunctionDecl struct {
	FunctionToken lexer.Token
	Name          string
	Prototype     FunctionPrototype
	Body          []Node
}

func (f FunctionDecl) StartPos() int {
	return f.FunctionToken.Position
}

func (f FunctionDecl) EndPos() int {
	return f.Prototype.RightBrace.Position
}

func (f FunctionDecl) StartLine() int {
	return f.FunctionToken.Line
}

func (f FunctionDecl) EndLine() int {
	return f.Prototype.RightBrace.Line
}

func (f FunctionDecl) declNode() {}

type StructDecl struct {
	StructToken lexer.Token
	Name        string
	LeftBrace   lexer.Token
	Fields      []StructField
	RightBrace  lexer.Token
}

func (s StructDecl) StartPos() int {
	return s.StructToken.Position
}

func (s StructDecl) EndPos() int {
	return s.RightBrace.Position
}

func (s StructDecl) StartLine() int {
	return s.StructToken.Line
}

func (s StructDecl) EndLine() int {
	return s.RightBrace.Line
}

func (s StructDecl) declNode() {}

type VariableDecl struct {
	VarToken lexer.Token
	Name     string
	Type     string
	Value    Expr
}

func (v VariableDecl) StartPos() int {
	return v.VarToken.Position
}

func (v VariableDecl) EndPos() int {
	return v.Value.EndPos()
}

func (v VariableDecl) StartLine() int {
	return v.VarToken.Line
}

func (v VariableDecl) EndLine() int {
	return v.Value.EndLine()
}

func (v VariableDecl) declNode() {}
