package interpreter

import (
	"testing"

	"github.com/Eclalang/Ecla/interpreter/eclaType"
)

func TestScope(t *testing.T) {

	scope := NewScopeMain()

	v1, err := eclaType.NewVar("v1", "bool", eclaType.Bool(true))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v1)
	scope.GoDeep(SCOPE_MAIN)
	v2, err := eclaType.NewVar("v1", "int", eclaType.Int(1))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v2)

	v, _ := scope.Get("a")

	if v.GetValue() != eclaType.Int(1) {
		t.Error("Expected 1, got ", v)
	}

	scope.GoUp()

	v, _ = scope.Get("a")

	if v.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", v)
	}

	scope.GoDeep(SCOPE_FUNCTION)

	scope.GoDeep(SCOPE_MAIN)

	scope.GoUp()
}

func TestScopeFunction(t *testing.T) {

	scope := NewScopeMain()

	v1, err := eclaType.NewVar("v1", "bool", eclaType.Bool(true))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v1)
	scope.GoDeep(SCOPE_FUNCTION)
	v2, err := eclaType.NewVar("v1", "int", eclaType.Int(1))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v2)

	v, _ := scope.Get("a")

	if v.GetValue() != eclaType.Int(1) {
		t.Error("Expected 1, got ", v)
	}

	scope.GoUp()

	v, _ = scope.Get("a")

	if v.GetValue() != eclaType.Bool(true) {
		t.Error("Expected true, got ", v)
	}
}

func TestScopeCheckIfVarExistsInCurrentScope(t *testing.T) {
	scope := NewScopeMain()

	v1, err := eclaType.NewVar("v1", "bool", eclaType.Bool(true))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v1)

	if !scope.CheckIfVarExistsInCurrentScope("a") {
		t.Error("Expected true, got false")
	}

	scope.GoDeep(SCOPE_MAIN)

	if scope.CheckIfVarExistsInCurrentScope("a") {
		t.Error("Expected false, got true")
	}
}

func TestScopeCheckIfVarExistsInCurrentScopeFunction(t *testing.T) {
	scope := NewScopeMain()

	v1, err := eclaType.NewVar("v1", "bool", eclaType.Bool(true))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v1)

	if !scope.CheckIfVarExistsInCurrentScope("a") {
		t.Error("Expected true, got false")
	}

	scope.GoDeep(SCOPE_FUNCTION)

	if scope.CheckIfVarExistsInCurrentScope("a") {
		t.Error("Expected false, got true")
	}
}

func TestScope_InFunction(t *testing.T) {
	scope := NewScopeMain()

	if scope.InFunction() {
		t.Error("Expected false, got true")
	}

	scope.GoDeep(SCOPE_FUNCTION)

	if !scope.next.InFunction() {
		t.Error("Expected true, got false")
	}
}

func TestScope_GetNextScope(t *testing.T) {
	scope := NewScopeMain()

	scope.GoDeep(SCOPE_MAIN)

	if scope.GetNextScope() == nil {
		t.Error("Expected not nil, got nil")
	}

	scope.GoUp()
	if scope.GetNextScope() != nil {
		t.Error("Expected nil, got not nil")
	}
}

func TestScope_SetNextScope(t *testing.T) {
	scope := NewScopeMain()

	scope.SetNextScope(NewScopeMain())

	if scope.GetNextScope() == nil {
		t.Error("Expected not nil, got nil")
	}

	if scope.GetNextScope().GetNextScope() != nil {
		t.Error("Expected nil, got not nil")
	}

	scope.SetNextScope(nil)
	if scope.GetNextScope() != nil {
		t.Error("Expected nil, got not nil")
	}
}

func TestScope_GetFunctionScope(t *testing.T) {
	scope := NewScopeMain()

	scope.GoDeep(SCOPE_MAIN)

	if scope.GetFunctionScope() != nil {
		t.Error("Expected nil, got not nil")
	}

	scope.GoDeep(SCOPE_FUNCTION)

	scope.GoDeep(SCOPE_MAIN)

	//get the most deep scope
	for scope.GetNextScope() != nil {
		scope = scope.GetNextScope()
	}

	if scope.GetFunctionScope() == nil {
		t.Error("Expected not nil, got nil")
	}
}

func TestScope_GoDeepWithSpecificScope(t *testing.T) {
	scope := NewScopeMain()

	scope.GoDeepWithSpecificScope(NewScopeMain())

	if scope.next == nil {
		t.Error("Expected not nil, got nil")
	}

	scope.GoDeepWithSpecificScope(NewScopeMain())

	if scope.next.next == nil {
		t.Error("Expected not nil, got nil")
	}
}
