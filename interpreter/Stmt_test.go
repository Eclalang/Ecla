package interpreter

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/lexer"
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

	var errCheck = false
	env.ErrorHandle.HookExit(
		func(int) {
			errCheck = true
		})

	if getPointerToSelectorExpr(parser.SelectorExpr{
		Expr: parser.Literal{
			Type:  "VAR",
			Value: "a",
		},
		Sel: parser.Literal{
			Type:  "STRING",
			Value: "a",
		},
	}, env, nil); !errCheck {
		t.Error("Expected a, got nil")
	}

	env = NewEnv()
	env.SetCode("struct A{a : []int;}a := A{[1,2,3]};")
	env.Execute()

	tokens := lexer.Lexer("a.a[1];")
	pars := parser.Parser{Tokens: tokens, ErrorHandler: env.ErrorHandle}
	pas := pars.Parse()
	token := pas.ParseTree.Operations[0]

	switch token.(type) {
	case parser.SelectorExpr:
		if v, _ := env.GetVar("a"); getPointerToSelectorExpr(token.(parser.SelectorExpr), env, nil) != func() *eclaType.Type {
			v, _ := (*v.Value.(*eclaType.Struct).Fields["a"]).GetIndex(eclaType.Int(1))
			return v
		}() {
			t.Error("Expected a, got nil")
		}
	default:
		t.Error("Expected a, got nil")
	}

	env = NewEnv()
	env.SetCode("struct A{a : map[int]int;}a := A{{1:1,2:2,3:3}};")
	env.Execute()

	tokens = lexer.Lexer("a.a[1];")
	pars = parser.Parser{Tokens: tokens, ErrorHandler: env.ErrorHandle}

	pas = pars.Parse()
	token = pas.ParseTree.Operations[0]

	switch token.(type) {
	case parser.SelectorExpr:
		if v, _ := env.GetVar("a"); getPointerToSelectorExpr(token.(parser.SelectorExpr), env, nil) != func() *eclaType.Type {
			v, _ := (*v.Value.(*eclaType.Struct).Fields["a"]).GetIndex(eclaType.Int(1))
			return v
		}() {
			t.Error("Expected a, got nil")
		}
	default:
		t.Error("Expected a, got nil")
	}

}
