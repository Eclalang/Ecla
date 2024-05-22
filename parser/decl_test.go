package parser

import (
	"github.com/Eclalang/Ecla/lexer"
	"testing"
)

var f FunctionDecl = FunctionDecl{
	FunctionToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "function",
		Position:  1,
		Line:      1,
	},
	Prototype: FunctionPrototype{
		RightBrace: lexer.Token{
			TokenType: lexer.RBRACE,
			Value:     "}",
			Position:  0,
			Line:      12,
		},
	},
}

func TestFunctionDecl_StartPos(t *testing.T) {
	if f.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestFunctionDecl_EndPos(t *testing.T) {
	if f.EndPos() != 0 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestFunctionDecl_StartLine(t *testing.T) {
	if f.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestFunctionDecl_EndLine(t *testing.T) {
	if f.EndLine() != 12 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestFunctionDecl_declNode(t *testing.T) {
	f.declNode()
}

var s StructDecl = StructDecl{
	StructToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "struct",
		Position:  1,
		Line:      1,
	},
	RightBrace: lexer.Token{
		TokenType: lexer.RBRACE,
		Value:     "}",
		Position:  0,
		Line:      12,
	},
}

func TestStructDecl_StartPos(t *testing.T) {
	if s.StartPos() != 1 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestStructDecl_EndPos(t *testing.T) {
	if s.EndPos() != 0 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestStructDecl_StartLine(t *testing.T) {
	if s.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestStructDecl_EndLine(t *testing.T) {
	if s.EndLine() != 12 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestStructDecl_declNode(t *testing.T) {
	s.declNode()
}

var v VariableDecl = VariableDecl{
	VarToken: lexer.Token{
		TokenType: lexer.TEXT,
		Value:     "var",
		Position:  0,
		Line:      1,
	},
	Value: &BinaryExpr{
		LeftExpr: &Literal{
			Token: lexer.Token{
				TokenType: lexer.INT,
				Value:     "1",
				Position:  12,
				Line:      1,
			},
		},
		Operator: lexer.Token{
			TokenType: lexer.ADD,
			Value:     "+",
			Position:  13,
			Line:      1,
		},
		RightExpr: &Literal{
			Token: lexer.Token{
				TokenType: lexer.INT,
				Value:     "1",
				Position:  14,
				Line:      1,
			},
		},
	},
}

func TestVariableDecl_StartPos(t *testing.T) {
	if v.StartPos() != 0 {
		t.Error("StartPos failed to return the correct value")
	}
}

func TestVariableDecl_EndPos(t *testing.T) {
	if v.EndPos() != 14 {
		t.Error("EndPos failed to return the correct value")
	}
}

func TestVariableDecl_StartLine(t *testing.T) {
	if v.StartLine() != 1 {
		t.Error("StartLine failed to return the correct value")
	}
}

func TestVariableDecl_EndLine(t *testing.T) {
	if v.EndLine() != 1 {
		t.Error("EndLine failed to return the correct value")
	}
}

func TestVariableDecl_declNode(t *testing.T) {
	v.declNode()
}
