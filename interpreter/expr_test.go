package interpreter

import (
	"github.com/Eclalang/Ecla/lexer"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func Test_RunTree(t *testing.T) {
	env := NewEnv()

	bus := RunTree(parser.Literal{Type: lexer.INT, Value: "0"}, env)

	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	env = NewEnv()

	bus = RunTree(parser.BinaryExpr{
		LeftExpr:  parser.Literal{Type: lexer.INT, Value: "0"},
		RightExpr: parser.Literal{Type: lexer.INT, Value: "0"},
		Operator: lexer.Token{
			TokenType: lexer.ADD,
		},
	}, env)

	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	env = NewEnv()

	bus = RunTree(parser.UnaryExpr{
		RightExpr: parser.Literal{
			Type:  lexer.INT,
			Value: "0",
		},
		Operator: lexer.Token{
			TokenType: lexer.INC,
		},
	}, env)
	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	env = NewEnv()

	bus = RunTree(parser.ParenExpr{
		Expression: parser.Literal{
			Type:  lexer.INT,
			Value: "0",
		},
	}, env)
	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	env = NewEnv()

	bus = RunTree(parser.VariableDecl{
		Name: "test",
		Type: parser.Int,
		Value: parser.Literal{
			Type:  lexer.INT,
			Value: "0",
		},
	}, env)

	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	bus = RunTree(parser.VariableAssignStmt{
		Names: []parser.Expr{
			parser.Literal{
				Type:  "VAR",
				Value: "test",
			},
		},
		Values: []parser.Expr{
			parser.Literal{
				Type:  lexer.INT,
				Value: "0",
			},
		},
		Operator: parser.ASSIGN,
	}, env)

	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	bus = RunTree(parser.WhileStmt{
		Cond: parser.BinaryExpr{
			LeftExpr: parser.Literal{
				Type:  lexer.INT,
				Value: "0",
			},
			RightExpr: parser.Literal{
				Type:  "VAR",
				Value: "test",
			},
			Operator: lexer.Token{
				TokenType: lexer.EQUAL,
			},
		},
		Body: []parser.Node{
			parser.VariableAssignStmt{
				Names: []parser.Expr{
					parser.Literal{
						Type:  "VAR",
						Value: "test",
					},
				},
				Values: []parser.Expr{
					parser.Literal{
						Type:  lexer.INT,
						Value: "1",
					},
				},
				Operator: parser.ASSIGN,
			},
		},
	}, env)

	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

	bus = RunTree(parser.ForStmt{
		InitDecl: parser.VariableDecl{
			Name: "i",
			Type: parser.Int,
			Value: parser.Literal{
				Type:  lexer.INT,
				Value: "0",
			},
		},
		CondExpr: parser.BinaryExpr{
			LeftExpr: parser.Literal{
				Type:  "VAR",
				Value: "i",
			},
			RightExpr: parser.Literal{
				Type:  lexer.INT,
				Value: "10",
			},
			Operator: lexer.Token{
				TokenType: lexer.LEQ,
			},
		},
		PostAssignStmt: parser.VariableAssignStmt{
			Names: []parser.Expr{
				parser.Literal{
					Type:  "VAR",
					Value: "i",
				},
			},
			Operator: "++",
			Values: []parser.Expr{
				nil,
			},
		},
		Body: []parser.Node{
			parser.VariableAssignStmt{
				Names: []parser.Expr{
					parser.Literal{
						Type:  "VAR",
						Value: "test",
					},
				},
				Values: []parser.Expr{
					parser.Literal{
						Type:  lexer.INT,
						Value: "1",
					},
				},
				Operator: parser.ASSIGN,
			},
		},
	}, env)

	if bus == nil {
		t.Error("Expected bus to be non-nil")
	}

}
