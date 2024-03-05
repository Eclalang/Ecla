package interpreter

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestRunImportStmt(t *testing.T) {
	env := NewEnv()
	env.Vars = InitBuildIn()
	env.SetCode("import \"console\";")
	env.Execute()
	if env.Libs["console"] == nil {
		t.Error("Expected fmt, got nil")
	}
}

func TestAssignementTypeChecking(t *testing.T) {
	env := NewEnv()
	if AssignementTypeChecking(parser.VariableAssignStmt{
		Names: []parser.Expr{
			parser.Literal{
				Type:  "string",
				Value: "a",
			},
		},
		Values: []parser.Expr{
			parser.Literal{
				Type:  "int",
				Value: "1",
			},
		},
	}, "int", "int", env) {
		t.Error("Expected false, got true")
	}

	if !AssignementTypeChecking(parser.VariableAssignStmt{
		Names: []parser.Expr{
			parser.Literal{
				Type:  "int",
				Value: "1",
			},
		},
		Values: []parser.Expr{
			parser.Literal{
				Type:  "int",
				Value: "1",
			},
		},
	}, "any", "int", env) {
		t.Error("Expected true, got false")
	}

	var e = false
	env.ErrorHandle.HookExit(
		func(int) {
			e = true
		})
	AssignementTypeChecking(parser.VariableAssignStmt{
		Names: []parser.Expr{
			parser.Literal{
				Type:  "int",
				Value: "1",
			},
		},
		Values: []parser.Expr{
			parser.Literal{
				Type:  "int",
				Value: "1",
			},
		},
	}, "int", "string", env)
	if !e {
		t.Error("Expected false, got true")
	}
}

func Test_getPointerToSelectorExpr(t *testing.T) {
	env := NewEnv()

	env.SetCode("struct A{a : int;};a:=A{1};")
	env.Execute()

	if v, _ := env.GetVar("a"); getPointerToSelectorExpr(parser.SelectorExpr{
		Expr: parser.Literal{
			Type:  "VAR",
			Value: "a",
		},
		Sel: parser.Literal{
			Type:  "VAR",
			Value: "a",
		},
	}, env, nil) != v.Value.(*eclaType.Struct).Fields["a"] {
		t.Error("Expected a, got nil")
	}
}
