package interpreter

import (
	"github.com/Eclalang/Ecla/interpreter/eclaDecl"
	"github.com/Eclalang/Ecla/interpreter/eclaType"
	"github.com/Eclalang/Ecla/parser"
	"testing"
)

func TestEnv_NewEnv(t *testing.T) {
	env := NewEnv()
	if env == nil {
		t.Error("Expected a new environment, got nil")
	}
}

func TestEnv_InitBuildIn(t *testing.T) {
	env := NewEnv()
	env.Vars = InitBuildIn()
	if env.Vars == nil {
		t.Error("Expected a new scope, got nil")
	}

	if v, ok := env.Vars.Get("typeOf"); v == nil || !ok {
		t.Error("Expected a new typeOf, got nil")
	}

	if v, ok := env.Vars.Get("sizeOf"); v == nil || !ok {
		t.Error("Expected a new sizeOf, got nil")
	}

	if v, ok := env.Vars.Get("len"); v == nil || !ok {
		t.Error("Expected a new len, got nil")
	}

	if v, ok := env.Vars.Get("append"); v == nil || !ok {
		t.Error("Expected a new append, got nil")
	}

}

func TestEnv_SetFile(t *testing.T) {
	env := NewEnv()
	env.SetFile("test")
	if env.File != "test" {
		t.Error("Expected test, got ", env.File)
	}
}

func TestEnv_SetCode(t *testing.T) {
	env := NewEnv()
	env.SetCode("test")
	if env.Code != "test" {
		t.Error("Expected test, got ", env.Code)
	}
}

func TestEnv_NewTemporaryEnv(t *testing.T) {
	env := NewTemporaryEnv(nil)
	if env == nil {
		t.Error("Expected a new environment, got nil")
	}
}

func TestEnv_String(t *testing.T) {
	env := NewEnv()
	if env.String() == "" {
		t.Error("Expected a string, got empty string")
	}
}

func TestEnv_ReadFile(t *testing.T) {
	s, _ := readFile("scope.go")
	if s == "" {
		t.Error("Expected a string, got empty string")
	}

	_, err := readFile("nonexistentfile")
	if err == nil {
		t.Error("Expected an error, got nil")
	}

}

func TestEnv_GetVar(t *testing.T) {
	env := NewEnv()
	v, _ := eclaType.NewVar("v1", "int", eclaType.Int(1))
	env.Vars.Set("a", v)
	if v, ok := env.GetVar("a"); v == nil || !ok {
		t.Error("Expected a new variable, got nil")
	}
}

func TestEnv_SetVar(t *testing.T) {
	env := NewEnv()
	v, _ := eclaType.NewVar("v1", "int", eclaType.Int(1))
	env.SetVar("a", v)
	if v, ok := env.GetVar("a"); v == nil || !ok {
		t.Error("Expected a new variable, got nil")
	}
}

func TestEnv_CheckIfVarExistsInCurrentScope(t *testing.T) {
	env := NewEnv()
	v, _ := eclaType.NewVar("v1", "int", eclaType.Int(1))
	env.Vars.Set("a", v)
	if !env.CheckIfVarExistsInCurrentScope("a") {
		t.Error("Expected true, got false")
	}
}

func TestEnv_NewScope(t *testing.T) {
	env := NewEnv()
	env.NewScope(SCOPE_MAIN)
	if env.Vars == nil {
		t.Error("Expected a new scope, got nil")
	}
	if env.Vars.next == nil {
		t.Error("Expected a new scope, got nil")
	}
}

func TestEnv_SetScope(t *testing.T) {
	env := NewEnv()
	env.NewScope(SCOPE_MAIN)
	env.SetScope(env.Vars.next)
	if env.Vars.next != env.Vars.next.next {
		t.Error("Expected a new scope, got nil")
	}
}

func TestEnv_EndScope(t *testing.T) {
	env := NewEnv()
	env.NewScope(SCOPE_MAIN)
	env.EndScope()
	if env.Vars.next != nil {
		t.Error("Expected a new scope, got nil")
	}
}

func TestEnv_SetFunction(t *testing.T) {
	env := NewEnv()
	f := eclaType.NewFunction("test", nil, nil, nil)
	err := env.SetFunction("test", f)
	if err != nil {
		t.Error("Expected nil, got ", err)
	}
	if v, ok := env.Vars.Get("test"); v == nil || !ok {
		t.Error("Expected a new function, got nil")
	}
}

func TestEnv_Execute(t *testing.T) {
	env := NewEnv()
	env.SetCode("import \"console\";\nconsole.println(\"Hello, World!\");\n")
	env.Execute()
}

func TestEnv_ExecuteFile(t *testing.T) {
	env := NewEnv()
	env.SetFile("../DEMO/AllTests.ecla")
	env.Execute()
}

func TestEnv_ExecuteMetrics(t *testing.T) {
	env := NewEnv()
	env.SetCode("import \"console\";\nconsole.println(\"Hello, World!\");\n")
	env.ExecuteMetrics()
}

func TestEnv_ExecuteMetricsFile(t *testing.T) {
	env := NewEnv()
	env.SetFile("../DEMO/AllTests.ecla")
	env.ExecuteMetrics()
}

func TestEnv_AddTypeDecl(t *testing.T) {
	typ := eclaDecl.NewStructDecl(parser.StructDecl{
		Name:   "test",
		Fields: make([]parser.StructField, 0),
	})

	env := NewEnv()

	env.AddTypeDecl(typ)
	if v, ok := env.GetTypeDecl("test"); v == nil || !ok {
		t.Error("Expected a new type, got nil")
	}

}

func TestEnv_ConvertToLib(t *testing.T) {
	env := NewEnv()
	l := env.ConvertToLib(env)
	if l == nil {
		t.Error("Expected a new lib, got nil")
	}
}

func TestEnv_AddFunctionExecuted(t *testing.T) {
	env := NewEnv()
	f := eclaType.NewFunction("test", nil, nil, nil)
	env.AddFunctionExecuted(f)
	if len(env.ExecutedFunc) != 1 {
		t.Error("Expected 1, got ", len(env.ExecutedFunc))
	}
}

func TestEnv_Load(t *testing.T) {
	env := NewEnv()
	env.SetFile("../DEMO/AllTests.ecla")
	env.Load()
}

func TestEnv_Import(t *testing.T) {
	env := NewEnv()
	env.Import(parser.ImportStmt{
		ModulePath: "../DEMO/AllTests.ecla",
	})

}

func TestEnvLib_Call(t *testing.T) {
	env := NewEnv()
	env.Import(parser.ImportStmt{
		ModulePath: "console",
	})

	if env.Libs["console"] == nil {
		t.Error("Expected a new lib, got nil")
	}

	_, err := env.Libs["console"].Call("println", []eclaType.Type{eclaType.String("Hello, World!")})
	if err != nil {
		t.Error("Expected nil, got ", err)
	}

	_, err = env.Libs["console"].Call("noexistfunction", []eclaType.Type{eclaType.String("Hello, World!")})
	if err == nil {
		t.Error("Expected an error, got nil")
	}
}
