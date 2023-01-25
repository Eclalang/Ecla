package interpreter

import (
	"github.com/tot0p/Ecla/interpreter"
	"github.com/tot0p/Ecla/interpreter/eclaType"
	"testing"
)

func TestScope(t *testing.T) {

	scope := interpreter.NewScopeMain()

	v1, err := eclaType.NewVar("v1", "bool", eclaType.Bool(true))
	if err != nil {
		t.Error(err)
	}
	scope.Set("a", v1)
	scope.GoDeep()
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
