package interpreter

import (
	"github.com/Eclalang/Ecla/interpreter/eclaType"
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
	s := readFile("scope.go")
	if s == "" {
		t.Error("Expected a string, got empty string")
	}

	defer func() {
		if r := recover(); r == nil {
			t.Error("Expected a panic, got nil")
		}
	}()

	readFile("nonexistentfile")

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
